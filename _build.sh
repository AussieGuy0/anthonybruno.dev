#!/bin/bash

jekyll_version="3.8"
cd _scripts/goodreads || exit
go run ./*.go
cd ../.. || exit
docker run --rm --volume="$PWD:/srv/jekyll" --volume="$PWD/vendor/bundle:/usr/local/bundle" -p 4000:4000 -it jekyll/jekyll:$jekyll_version jekyll serve --watch
