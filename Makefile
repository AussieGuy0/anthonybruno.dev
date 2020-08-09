jekyll_version=3.8

all: fetch_goodreads serve

fetch_goodreads:
	- cd _scripts/goodreads && go run ./*.go

serve:
	docker run --rm --volume="$(PWD):/srv/jekyll" --volume="$(PWD)/vendor/bundle:/usr/local/bundle" -p 4000:4000 -it jekyll/jekyll:$(jekyll_version) jekyll serve --watch

