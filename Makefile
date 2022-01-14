jekyll_version=3.8

all: fetch_til serve

# Goodreads decided to kill the API :(
fetch_goodreads:
	- cd _scripts/goodreads && go run ./*.go

fetch_til:
	- cp ~/Drive/Notes/reference/TIL/* _tils/

new_post:
	- cd _scripts && ./createPost.sh

serve:
	docker run --rm --volume="$(PWD):/srv/jekyll" --volume="$(PWD)/vendor/bundle:/usr/local/bundle" -p 4000:4000 -it jekyll/jekyll:$(jekyll_version) jekyll serve --watch

