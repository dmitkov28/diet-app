include .env
export 

.PHONY: dev css templates air stop db goose bootstrap deploy ngrok sync

ngrok:
	docker run -it -e NGROK_AUTHTOKEN="${NGROK_TOKEN}" ngrok/ngrok:latest http host.docker.internal:1323

dev:
	make db && make goose && make air & make templates & make css

db:
	docker compose up -d
	sleep 3

goose: db
	cd ./internal/data/schema && GOOSE_DRIVER=turso GOOSE_DBSTRING="${LOCAL_DB_URL}" goose up

bootstrap: goose
	go run ./cmd/admin/main.go create-user --email "${LOCAL_USER}" --password "${LOCAL_PASSWORD}"

air:
	air

templates:
	templ generate --watch

css:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/main.css --watch --minify

deploy:
	cd terraform && terraform apply -var-file="secret.tfvars" --auto-approve && cd .. && make sync

sync:
	@echo "Syncing S3 bucket..."
	aws s3 sync "${S3_URI}" static/ --delete

stop:
	pkill -f "air" || true
	pkill -f "templ" || true
	pkill -f "tailwindcss" || true
	docker compose down