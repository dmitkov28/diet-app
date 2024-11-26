.PHONY: dev css templates air

air:
	air

templates:
	templ generate --watch

css:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/main.css --watch --minify


stop:
	pkill -f "air" || true
	pkill -f "templ" || true
	pkill -f "tailwindcss" || true