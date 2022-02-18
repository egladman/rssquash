#!/usr/bin/env bash

set -e

INPUT="${INPUT:-examples/feeds.list}"

DESTDIR="${TMPDIR:-$(mktemp -d)}"
OUTPUTDIR="${DESTDIR}/out"
STAGINGDIR="${DESTDIR}/repo"
FEEDPATH="${OUTPUTDIR}/feed.atom"

GITORIGIN="$(git config  --get remote.origin.url)"
GITBRANCH="gh-pages"

mkdir -p "$OUTPUTDIR" "$STAGINGDIR"

go run main.go --source "$INPUT" > "$FEEDPATH"

cp -r "${OUTPUTDIR}"/* "$STAGINGDIR"

pushd "$STAGINGDIR"
git init
git add .
git commit -m "Updated at $(date)"
git remote add origin "${GITORIGIN}"
git push origin master:refs/heads/"${GITBRANCH}" --force
popd

# TODO cleanup
