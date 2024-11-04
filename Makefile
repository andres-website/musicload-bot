build:
	GOOS=linux go build -o app ./cmd/
	docker build -t downloader/telegrambot .
	rm -f app

run:
	docker run --memory="100m" --cpus="1" --rm downloader/telegrambot 