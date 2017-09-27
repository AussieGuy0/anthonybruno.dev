#!/bin/bash
set -eu

function replaceText() {
    sed -i -e "s@$1@$2@g" "$3"
}
URL=$1
TITLE=$(wget -qO- "$URL" |perl -l -0777 -ne 'print $1 if /<title.*?>\s*(.*?)\s*<\/title/si')
DATE=$(date "+%Y-%m-%d")
SUMMARY=$2
TAGS=$3

FILE="$TITLE.md"

POSTPATH="../_links/$FILE"

cp link.md "$POSTPATH"

replaceText "{url}" "$URL" "$POSTPATH"
replaceText "{title}" "$TITLE" "$POSTPATH"
replaceText "{date}" "$DATE" "$POSTPATH"
replaceText "{summary}" "$SUMMARY" "$POSTPATH"
replaceText "{tags}" "$TAGS" "$POSTPATH"

