jekyll_version=4.2.2

all: fetch_til serve

fetch_goodreads:
	- cd _scripts/goodreads && go run ./*.go

fetch_til:
	- cp ~/Drive/Notes/knowledge/reference/TIL/* _tils/

new_post:
	- cd _scripts && ./createPost.sh

serve:
	docker run --rm --volume="$(PWD):/srv/jekyll" --volume="$(PWD)/vendor/bundle:/usr/local/bundle" -p 4000:4000 -p 35729:35729 -it jekyll/jekyll:$(jekyll_version) jekyll serve --livereload

