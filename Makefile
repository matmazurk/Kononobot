build:
	docker build -f cmd/mock/api/Api.Dockerfile -t api-mock .
	docker build -f ./Scraper.Dockerfile -t scraper .

run: build
	docker compose up 
	docker compose down