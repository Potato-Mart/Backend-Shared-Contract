#!/usr/bin/env bash
set -euo pipefail

version=""
input_path=""
output_path=""
previous_tag=""
model="${OPENAI_RELEASE_NOTES_MODEL:-gpt-5}"

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/get-ai-contract-release-notes.sh --version v3.0.0 --input release-notes.md --output release-notes.md

Polishes deterministic release notes using the OpenAI Responses API.
Requires OPENAI_API_KEY.
USAGE
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    --version|-v)
      version="${2:-}"
      shift 2
      ;;
    --input|-i)
      input_path="${2:-}"
      shift 2
      ;;
    --output|-o)
      output_path="${2:-}"
      shift 2
      ;;
    --previous-tag)
      previous_tag="${2:-}"
      shift 2
      ;;
    --model)
      model="${2:-}"
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

if [[ -z "$input_path" || -z "$output_path" ]]; then
  echo "--input and --output are required." >&2
  exit 1
fi

if [[ -z "${OPENAI_API_KEY:-}" ]]; then
  echo "OPENAI_API_KEY is required." >&2
  exit 1
fi

if [[ -z "$previous_tag" ]]; then
  previous_tag="$(git describe --tags --abbrev=0 HEAD 2>/dev/null || true)"
fi

range="HEAD"
if [[ -n "$previous_tag" ]]; then
  range="$previous_tag..HEAD"
fi

python3 - "$version" "$input_path" "$output_path" "$previous_tag" "$range" "$model" <<'PY'
import json
import os
import subprocess
import sys
import urllib.request

version, input_path, output_path, previous_tag, commit_range, model = sys.argv[1:]

def git(*args):
    try:
        return subprocess.check_output(["git", *args], text=True, stderr=subprocess.DEVNULL).strip()
    except subprocess.CalledProcessError:
        return ""

with open(input_path, "r", encoding="utf-8") as f:
    deterministic_notes = f.read()

commits = git("log", "--no-merges", "--pretty=format:%h %s", commit_range)
changed_files = git("diff", "--name-only", commit_range)
diff_stat = git("diff", "--stat", commit_range)

prompt = f"""You are writing release notes for a Go module that is a shared backend contract package.

Version: {version}
Previous tag: {previous_tag}

Use the deterministic notes, commit subjects, changed files, and diff summary below.
Produce Markdown release notes for engineers consuming the contract.

Rules:
- Keep it factual. Do not invent changes that are not supported by the input.
- Call out breaking contract changes prominently.
- Explain compatibility impact in practical terms.
- Mention changed JSON field names, enum wire values, primitive shape changes, and new exported contract types when present.
- Keep the release notes concise but useful.
- Do not include marketing language.

Deterministic notes:
{deterministic_notes}

Commits:
{commits}

Changed files:
{changed_files}

Diff summary:
{diff_stat}
"""

payload = {
    "model": model,
    "instructions": "Write accurate, concise contract release notes for backend engineers.",
    "input": prompt,
    "text": {"verbosity": "medium"},
}

request = urllib.request.Request(
    "https://api.openai.com/v1/responses",
    data=json.dumps(payload).encode("utf-8"),
    headers={
        "Authorization": f"Bearer {os.environ['OPENAI_API_KEY']}",
        "Content-Type": "application/json",
    },
    method="POST",
)

with urllib.request.urlopen(request, timeout=120) as response:
    data = json.loads(response.read().decode("utf-8"))

text = data.get("output_text", "")
if not text:
    chunks = []
    for item in data.get("output", []):
        for content in item.get("content", []):
            if "text" in content:
                chunks.append(content["text"])
    text = "\n".join(chunks).strip()

if not text:
    raise RuntimeError("OpenAI response did not contain release note text.")

directory = os.path.dirname(output_path)
if directory:
    os.makedirs(directory, exist_ok=True)

with open(output_path, "w", encoding="utf-8") as f:
    f.write(text)
    f.write("\n")
PY
