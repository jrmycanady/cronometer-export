#!/bin/bash

# Obtaining the build version from args.
if [ "$#" -eq 0 ]; then
echo "No version value was provided. Please 'use build-release.sh #' where # is the version number."
    exit 1
fi
VERSION=$1
echo "version:    ${VERSION}"

rm -rf ./cronometer-export-release-tmp
mkdir ./cronometer-export-release-tmp

mkdir ./cronometer-export-release-tmp/cronometer-export-macosx-amd64-$VERSION
mkdir ./cronometer-export-release-tmp/cronometer-export-freebsd-amd64-$VERSION
mkdir ./cronometer-export-release-tmp/cronometer-export-linux-amd64-$VERSION
mkdir ./cronometer-export-release-tmp/cronometer-export-linux-arm64-$VERSION
mkdir ./cronometer-export-release-tmp/cronometer-export-openbsd-amd64-$VERSION
mkdir ./cronometer-export-release-tmp/cronometer-export-openbsd-arm64-$VERSION
mkdir ./cronometer-export-release-tmp/cronometer-export-windows-amd64-$VERSION

cd ../

env GOOS=darwin GOARCH=amd64 go build -o ./release/cronometer-export-release-tmp/cronometer-export-macosx-amd64-$VERSION/
env GOOS=freebsd GOARCH=amd64 go build -o ./release/cronometer-export-release-tmp/cronometer-export-freebsd-amd64-$VERSION/
env GOOS=linux GOARCH=amd64 go build -o ./release/cronometer-export-release-tmp/cronometer-export-linux-amd64-$VERSION/
env GOOS=linux GOARCH=arm64 go build -o ./release/cronometer-export-release-tmp/cronometer-export-linux-arm64-$VERSION/
env GOOS=openbsd GOARCH=amd64 go build -o ./release/cronometer-export-release-tmp/cronometer-export-openbsd-amd64-$VERSION/
env GOOS=openbsd GOARCH=arm64 go build -o ./release/cronometer-export-release-tmp/cronometer-export-openbsd-arm64-$VERSION/
env GOOS=windows GOARCH=amd64 go build -o ./release/cronometer-export-release-tmp/cronometer-export-windows-amd64-$VERSION/

cd ./release/cronometer-export-release-tmp
zip -r ./cronometer-export-macosx-amd64.zip ./cronometer-export-macosx-amd64-$VERSION
zip -r ./cronometer-export-freebsd-amd64.zip ./cronometer-export-macosx-amd64-$VERSION
zip -r ./cronometer-export-linux-amd64.zip ./cronometer-export-macosx-amd64-$VERSION
zip -r ./cronometer-export-linux-arm64.zip ./cronometer-export-macosx-amd64-$VERSION
zip -r ./cronometer-export-openbsd-amd64.zip ./cronometer-export-macosx-amd64-$VERSION
zip -r ./cronometer-export-openbsd-arm64.zip ./cronometer-export-macosx-amd64-$VERSION
zip -r ./cronometer-export-windows-amd64.zip ./cronometer-export-macosx-amd64-$VERSION
