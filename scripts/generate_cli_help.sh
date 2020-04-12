#!/usr/bin/env bash
set -euo pipefail

MY_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
TMP_FILE=$(mktemp)
TMP_OUTPUT=$(mktemp)

cd "${MY_DIR}/../" || exit;

echo "\`\`\`" >"${TMP_FILE}"

export MAX_SCREEN_WIDTH=100
export FORCE_SCREEN_WIDTH="true"
export VERBOSE="true"

go build -o . ./cmd/gonotes
./gonotes help | sed 's/^/  /' | sed 's/ *$//' >>"${TMP_FILE}"

echo "\`\`\`" >> "${TMP_FILE}"

MAX_LEN=$(awk '{ print length($0); }' "${TMP_FILE}" | sort -n | tail -1 )

README_FILE="${MY_DIR}/../README.md"

LEAD='^<!-- AUTO_STAR -->$'
TAIL='^<!-- AUTO_END -->'

BEGIN_GEN=$(grep -n "${LEAD}" <"${README_FILE}" | sed 's/\(.*\):.*/\1/g')
END_GEN=$(grep -n "${TAIL}" <"${README_FILE}" | sed 's/\(.*\):.*/\1/g')
cat <(head -n "${BEGIN_GEN}" "${README_FILE}") \
    "${TMP_FILE}" \
    <(tail -n +"${END_GEN}" "${README_FILE}") \
    >"${TMP_OUTPUT}"

mv "${TMP_OUTPUT}" "${README_FILE}"
