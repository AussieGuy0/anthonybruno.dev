#!/bin/bash
# Creates a new post with todays date, and set the title to the first argument 
# given when running the script. Also makes a directory in _includes for code 
# samples
set -e

DATE=$(date "+%Y-%m-%d")
TITLE=$1
CLEANED_TITLE=${TITLE// /-}
FILE=$DATE"-$CLEANED_TITLE.md"
POSTPATH="../_posts/$FILE"
CURR_DIR=$(dirname "$0")

if [[ -z $TITLE ]]
then
    echo "Usage: ./makePost.sh <title>"
    exit 1
fi
cd "$CURR_DIR"
cp post.txt "$POSTPATH"
sed -i -e "s/{title}/$TITLE/g" "$POSTPATH"
vim "$POSTPATH"
