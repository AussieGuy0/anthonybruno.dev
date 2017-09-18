#!/bin/bash
# Creates a new post with todays date, and set the title to the first argument 
# given when running the script. Also makes a directory in _includes for code 
# samples

DATE=$(date "+%Y-%m-%d")

TITLE=$1

FILE=$DATE"-$1.md"

POSTPATH="../_posts/${FILE// /-}"

cp post.txt "$POSTPATH"

sed -i -e "s/{title}/$TITLE/g" "$POSTPATH"

mkdir "../_includes/$TITLE"
