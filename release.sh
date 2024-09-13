#!/bin/bash

if [[ "$#" -ne 2 ]]; then
    echo "Usage: $0 [minor|patch] 'release title'"
    exit 1
fi

if [[ -n $(git status --porcelain) ]]; then
    echo "Please commit your changes before releasing."
    exit 1
fi

VERSION=$(cat VERSION | sed 's/^v//')
LATEST_VERSION=$(git tag | tail -1 | sed 's/^v//')

if [[ $VERSION != $LATEST_VERSION ]]; then
    echo "Make sure that the string in VERSION equals the current latest tag."
    exit 1
fi

#bump version
IFS="." read -r -a VERSION_PARTS <<< "$VERSION"
MAJOR=${VERSION_PARTS[0]}
MINOR=${VERSION_PARTS[1]}
PATCH=${VERSION_PARTS[2]}

bump_version() {
    case $1 in
        major)
        ((MAJOR++))
        MINOR=0
        PATCH=0
        ;;
        minor)
        ((MINOR++))
        PATCH=0
        ;;
        patch)
        ((PATCH++))
        ;;
        *)
        echo "Invalid bump type! Use 'major', 'minor', or 'patch'."
        exit 1
        ;;
    esac

    NEW_VERSION="v$MAJOR.$MINOR.$PATCH"

    echo "$NEW_VERSION" > VERSION

    echo "Version bumped to $NEW_VERSION"
}

bump_version "$1"

echo "tagging current commit with new version $NEW_VERSION"
git tag $NEW_VERSION

echo "pushing changes to origin"
git push origin $NEW_VERSION

echo "writing new version to the changelog"
sed -i "" "/Changelog/ a\\
${NEW_VERSION} - $2
" CHANGELOG.md

echo "Succesfully released utils $NEW_VERSION"
