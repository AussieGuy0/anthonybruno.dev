jekyll_version=4.2.2

all: fetch_til serve

fetch_goodreads:
	- cd _scripts/goodreads && go run ./*.go

fetch_strava:
	- cd _scripts/strava && export $$(cat .env | xargs) && go run main.go strava.go jekyll.go

fetch_til:
	- cp ~/Drive/Notes/knowledge/reference/TIL/* _tils/

new_post:
	- cd _scripts && ./createPost.sh

serve:
	docker compose up
