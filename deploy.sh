#!/usr/bin/env bash

set -e

if [[ -z "$1" ]]; then
    printf '%s\n' "Feed path not specified"
    exit 1
fi

DESTDIR="${DESTDIR:-build}"
OUTPUTDIR="${DESTDIR}/out"
STAGINGDIR="${DESTDIR}/repo"

GITORIGIN="$(git config --get remote.origin.url)"
GITBRANCH="gh-pages"

rm -rf "${DESTDIR:?}"/*
mkdir -p "$OUTPUTDIR/${RSSQUASH_PREFIX}" "$STAGINGDIR"

go run main.go --source "$1" > "${OUTPUTDIR}/${RSSQUASH_PREFIX}feed.atom"

cp -r "${OUTPUTDIR}"/* "$STAGINGDIR"

pushd "$STAGINGDIR"
git init
git add .
git commit -m "Updated at $(date)"
git remote add origin "${GITORIGIN}"
git push origin master:refs/heads/"${GITBRANCH}" --force
popd

rm -rf "${DESTDIR:?}"/*
