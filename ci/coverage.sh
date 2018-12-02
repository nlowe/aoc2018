#!/usr/bin/env bash
set -euo pipefail

SCRIPT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
"${GOPATH}/bin/goveralls" -coverprofile="${SCRIPT_ROOT}/../coverage.out" -service=travis-ci