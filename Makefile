build-dev:
	docker build -f dev.Dockerfile -t getprint-service-media-dev .

build-prod:
	docker build -f prod.Dockerfile -t getprint-service-media-prod .