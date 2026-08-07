#!/usr/bin/env bash
set -euo pipefail

version=""
previous_tag=""
output_path=""

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/bash/get-contract-release-notes.sh --version v3.0.0 [--output release-notes.md] [--previous-tag v2.1.1]

Generates deterministic Markdown release notes from git commits and changed contract files.
USAGE
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    --version|-v)
      version="${2:-}"
      shift 2
      ;;
    --previous-tag)
      previous_tag="${2:-}"
      shift 2
      ;;
    --output|-o)
      output_path="${2:-}"
      shift 2
      ;;
    --help|-h)
      usage
      exit 0
      ;;
    *)
      echo "Unknown argument: $1" >&2
      usage >&2
      exit 1
      ;;
  esac
done

if [[ ! "$version" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.-]+)?$ ]]; then
  echo "Version is required and must look like v3.0.0." >&2
  exit 1
fi

git_safe() {
  git "$@" 2>/dev/null || true
}

if [[ -z "$previous_tag" ]]; then
  if git rev-parse --verify --quiet "refs/tags/$version" >/dev/null; then
    previous_tag="$(git_safe describe --tags --abbrev=0 "$version^")"
  fi

  if [[ -z "$previous_tag" ]]; then
    previous_tag="$(git_safe describe --tags --abbrev=0 HEAD)"
  fi
fi

if [[ -n "$previous_tag" ]]; then
  range="$previous_tag..HEAD"
else
  range="HEAD"
fi

mapfile -t commits < <(git log --no-merges --pretty=format:'- %s (%h)' "$range" 2>/dev/null || true)
mapfile -t changed_files < <(git diff --name-only "$range" 2>/dev/null || true)

breaking=()
features=()
fixes=()
other=()

for commit in "${commits[@]}"; do
  [[ -z "$commit" ]] && continue
  if [[ "$commit" =~ BREAKING\ CHANGE || "$commit" =~ !: ]]; then
    breaking+=("$commit")
  elif [[ "$commit" =~ ^-\ feat(\(.+\))?: ]]; then
    features+=("$commit")
  elif [[ "$commit" =~ ^-\ fix(\(.+\))?: ]]; then
    fixes+=("$commit")
  else
    other+=("$commit")
  fi
done

emit_section() {
  local title="$1"
  shift
  local items=("$@")

  if [[ ${#items[@]} -eq 0 ]]; then
    return
  fi

  printf '## %s\n' "$title"
  printf '%s\n' "${items[@]}"
  printf '\n'
}

generate_notes() {
  printf '# %s\n\n' "$version"

  if [[ -n "$previous_tag" ]]; then
    printf 'Changes since `%s`.\n\n' "$previous_tag"
  else
    printf 'Initial tracked release notes for this contract version.\n\n'
  fi

  emit_section "Breaking Contract Changes" "${breaking[@]}"
  emit_section "Added" "${features[@]}"
  emit_section "Fixed" "${fixes[@]}"
  emit_section "Other Changes" "${other[@]}"

  contract_files=()
  for file in "${changed_files[@]}"; do
    case "$file" in
      pkg/*|go.mod|VERSIONING.md)
        contract_files+=("- \`$file\`")
        ;;
    esac
  done

  if [[ ${#contract_files[@]} -gt 0 ]]; then
    printf '## Contract Files Changed\n'
    printf '%s\n' "${contract_files[@]}" | sort
    printf '\n'
  fi

  cat <<'NOTES'
## Compatibility Notes
- Consumers should upgrade deliberately and run contract serialization/deserialization tests.
- Major versions may include removed fields, renamed JSON keys, enum value changes, or changed primitive shapes.
- Minor versions may add fields, contracts, enum values, or helper types in a backward-compatible way.
- Patch versions should contain documentation, comments, helper fixes, or non-breaking corrections.
NOTES
}

if [[ -n "$output_path" ]]; then
  mkdir -p "$(dirname "$output_path")"
  generate_notes > "$output_path"
else
  generate_notes
fi
