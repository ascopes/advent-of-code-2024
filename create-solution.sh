#!/usr/bin/env bash
set -o errexit
set -o nounset
[[ -n ${DEBUG+defined} ]] && set -o xtrace

function usage() {
  echo "USAGE: ${BASH_SOURCE[0]} [-h|--help] <day-number>"
  echo "Create a skeleton for an Advent of Code solution."
  echo
  echo "Arguments:"
  echo "    -h | --help    Show this message and exit."
  echo "    <day-number>   The day to generate the solution for."
  echo ""
}

if [[ ${1:--h} =~ ^(--help|-h)$ ]]; then
  usage
  exit 0
elif ! [[ ${1:-} =~ ^[1-9]|1[0-9]|2[0-5]$ ]]; then
  echo "ERROR: Day number must be an integer between 1 and 25."
  usage
  exit 1
fi

cd "$(dirname "${BASH_SOURCE[0]}")"
mkdir day$1
cd day$1

touch input.txt output.txt

ln -s ../build.mk GNUmakefile

cat > README.md <<EOF
# Day $1
## Part 1

TODO

## Part 2

TODO
EOF

cat > go.mod <<EOF
module github.com/ascopes/advent-of-code-2024/day$1

go 1.23
EOF

cat > main.go <<EOF
package main

import "fmt"

func main() {
    fmt.Println("Advent of Code 2024 Day $1:")
    fmt.Println("    Part 1: TODO")
    fmt.Println("    Part 2: TODO")
}
EOF

echo "Created skeleton in $(pwd)..."
make run
ls -Ahl
git add -v "."

echo "Operation complete."
