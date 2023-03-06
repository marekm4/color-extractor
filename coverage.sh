#!/usr/bin/env bash
set -euo pipefail

# This is a simple script to generate an HTML coverage report,
# and SVG badge for your Go project.
#
# It's meant to be used manually or as a pre-commit hook.
#
# Place it some where in your code tree and execute it.
# If your tests pass, next to the script you'll find
# the coverage.html report and coverage.svg badge.
#
# You can add the badge to your README.md as such:
#  [![Go Coverage](PATH_TO/coverage.svg)](https://raw.githack.com/URL/coverage.html)
#
# Visit https://raw.githack.com/ to find the correct URL.
#
# To have the script run as a pre-commmit hook,
# symlink the script to .git/hooks/pre-commit:
#
#  ln -s PATH_TO/coverage.sh .git/hooks/pre-commit
#
# Or, if you have other pre-commit hooks,
# call it from your main hook.

# Get the script's directory after resolving a possible symlink.
SCRIPT_DIR="$(dirname -- "$(readlink -f "${BASH_SOURCE[0]}")")"

# Get coverage for all packages in the current directory; store next to script.
go test ./... -coverpkg "$(go list)/..." -coverprofile "$SCRIPT_DIR/coverage.out"

# Create an HTML report; store next to script.
go tool cover -html="$SCRIPT_DIR/coverage.out" -o "$SCRIPT_DIR/coverage.html"

# Extract total coverage: the decimal number from the last line of the function report.
COVERAGE=$(go tool cover -func="$SCRIPT_DIR/coverage.out" | tail -1 | grep -Eo '[0-9]+\.[0-9]')

echo "coverage: $COVERAGE% of statements"

date "+%s,$COVERAGE" >> "$SCRIPT_DIR/coverage.log"

# Pick a color for the badge.
if awk "BEGIN {exit !($COVERAGE >= 90)}"; then
	COLOR=brightgreen
elif awk "BEGIN {exit !($COVERAGE >= 80)}"; then
	COLOR=green
elif awk "BEGIN {exit !($COVERAGE >= 70)}"; then
	COLOR=yellowgreen
elif awk "BEGIN {exit !($COVERAGE >= 60)}"; then
	COLOR=yellow
elif awk "BEGIN {exit !($COVERAGE >= 50)}"; then
	COLOR=orange
else
	COLOR=red
fi

# Download the badge; store next to script.
curl -s "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR" > "$SCRIPT_DIR/coverage.svg"

# When running as a pre-commit hook, add the report and badge to the commit.
if [[ -n "${GIT_INDEX_FILE-}" ]]; then
	git add "$SCRIPT_DIR/coverage.html" "$SCRIPT_DIR/coverage.svg"
fi
