#!/bin/bash

VERSION=$(cat VERSION)

if [[ -n $(git status --porcelain) ]]; then
    echo "Please commit your changes before releasing."
    exit 1
fi

#git tag $VERSION
#git push origin $VERSION

COMMIT_MESSAGE=$(git log -1 --pretty=%B)
sed -i '' '/Changelog/ a\
$VERSION - $COMMIT_MESSAGE\n' CHANGELOG.md

echo "Released utils version $VERSION"
