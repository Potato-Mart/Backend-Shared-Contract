#!/usr/bin/env bash
set -euo pipefail

version=""
bump=""
commit_message=""
release_notes_path="release-notes.md"
use_ai_release_notes=false
no_commit=false
no_tag=false
push_release=false
dry_run=false

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/bash/publish-contract-version.sh --version v3.0.0 [--push]
  bash scripts/bash/publish-contract-version.sh --bump major|minor|patch [--push]

Options:
  --version vX.Y.Z              Exact version to release.
  --bump major|minor|patch      Calculate next version from pkg/versioning/version.go.
  --message "..."               Commit message. Defaults to chore(release): VERSION.
  --release-notes-path PATH     Release notes path. Defaults to release-notes.md.
  --use-ai-release-notes        Polish deterministic notes with OpenAI. Requires OPENAI_API_KEY.
  --no-commit                   Update files but do not commit.
  --no-tag                      Do not create an annotated tag.
  --push                        Push commit and tag.
  --dry-run                     Print target version and exit without changing files.
USAGE
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    --version|-v)
      version="${2:-}"
      shift 2
      ;;
    --bump)
      bump="${2:-}"
      shift 2
      ;;
    --message|-m)
      commit_message="${2:-}"
      shift 2
      ;;
    --release-notes-path)
      release_notes_path="${2:-}"
      shift 2
      ;;
    --use-ai-release-notes)
      use_ai_release_notes=true
      shift
      ;;
    --no-commit)
      no_commit=true
      shift
      ;;
    --no-tag)
      no_tag=true
      shift
      ;;
    --push)
      push_release=true
      shift
      ;;
    --dry-run)
      dry_run=true
      shift
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

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"
cd "$repo_root"

git rev-parse --is-inside-work-tree >/dev/null

version_file="pkg/versioning/version.go"
current_version="$(sed -nE 's/.*ModuleVersion[[:space:]]*=[[:space:]]*"([^"]+)".*/\1/p' "$version_file" | head -n 1)"

if [[ -z "$current_version" ]]; then
  echo "Could not read ModuleVersion from $version_file." >&2
  exit 1
fi

bump_version() {
  local current="$1"
  local kind="$2"

  if [[ ! "$current" =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+) ]]; then
    echo "Current version '$current' is not valid semver." >&2
    exit 1
  fi

  local major="${BASH_REMATCH[1]}"
  local minor="${BASH_REMATCH[2]}"
  local patch="${BASH_REMATCH[3]}"

  case "$kind" in
    major)
      major=$((major + 1)); minor=0; patch=0
      ;;
    minor)
      minor=$((minor + 1)); patch=0
      ;;
    patch)
      patch=$((patch + 1))
      ;;
    *)
      echo "--bump must be major, minor, or patch." >&2
      exit 1
      ;;
  esac

  echo "v$major.$minor.$patch"
}

if [[ -z "$version" ]]; then
  if [[ -z "$bump" ]]; then
    echo "Provide --version vX.Y.Z or --bump major|minor|patch." >&2
    exit 1
  fi
  version="$(bump_version "$current_version" "$bump")"
fi

if [[ ! "$version" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.-]+)?$ ]]; then
  echo "Target version '$version' is not valid semver." >&2
  exit 1
fi

if [[ -z "$commit_message" ]]; then
  commit_message="chore(release): $version"
fi

echo "Current contract version: $current_version"
echo "Target contract version:  $version"

if [[ "$dry_run" == true ]]; then
  echo "Dry run: no files, commits, tags, or pushes will be changed."
  exit 0
fi

major_version="${version%%.*}"

python3 - "$version_file" "$version" "$major_version" <<'PY'
import pathlib
import re
import sys

path = pathlib.Path(sys.argv[1])
version = sys.argv[2]
major = sys.argv[3]

text = path.read_text(encoding="utf-8")
text = re.sub(r'ModuleVersion\s*=\s*"v\d+\.\d+\.\d+(?:-[0-9A-Za-z.-]+)?"', f'ModuleVersion = "{version}"', text)
text = re.sub(r'MajorVersion\s*=\s*"v\d+"', f'MajorVersion  = "{major}"', text)
path.write_text(text, encoding="utf-8")
PY

bash scripts/bash/get-contract-release-notes.sh --version "$version" --output "$release_notes_path"

if [[ "$use_ai_release_notes" == true ]]; then
  bash scripts/bash/get-ai-contract-release-notes.sh --version "$version" --input "$release_notes_path" --output "$release_notes_path"
fi

if [[ -n "$(git status --porcelain)" ]]; then
  echo "Detected changes:"
  git status --short

  if [[ "$no_commit" == false ]]; then
    git add -A
    git commit -m "$commit_message"
  else
    echo "Skipping commit because --no-commit was supplied."
  fi
else
  echo "No file changes to commit."
fi

if [[ "$no_tag" == false ]]; then
  if git rev-parse --verify --quiet "refs/tags/$version" >/dev/null; then
    echo "Tag '$version' already exists." >&2
    exit 1
  fi

  git tag -a "$version" -F "$release_notes_path"
fi

if [[ "$push_release" == true ]]; then
  git push
  if [[ "$no_tag" == false ]]; then
    git push origin "$version"
  fi
fi

echo "Release preparation complete for $version."
