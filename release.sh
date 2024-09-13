#!/bin/bash

VERSION=$(cat VERSION)
LATEST_VERSION=$(git tag | tail -1)

if [[ "$#" -ne 2 ]]; then
    echo "Usage: $0 [minor|patch] 'release title'"
    exit 1
fi

if [[ -n $(git status --porcelain) ]]; then
    echo "Please commit your changes before releasing."
    exit 1
fi

if [[ $VERSION != $LATEST_VERSION ]]; then
    echo "Make sure that the string in VERSION equals the current latest tag."
    exit 1
fi

#bump version
IFS="." read -r -a VERSION_PARTS <<< "$VERSION"
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

    NEW_VERSION="$MAJOR.$MINOR.$PATCH"

    echo "$NEW_VERSION" > VERSION

    echo "Version bumped to $NEW_VERSION"
}

bump_version("$1")

git tag $VERSION
git push origin $VERSION
COMMIT_MESSAGE=$(git log -1 --pretty=%B)

sed -i "" "/Changelog/ a\\
${VERSION} - ${COMMIT_MESSAGE}
" CHANGELOG.md

echo "Released utils version $VERSION"
