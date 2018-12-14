.PHONY: dev buildimage

default:
	@ echo "Available methods"
	@ echo "make dev			- dev server"
	@ echo "make buildimage		- build docker image"

dev:
	go run ./cmd/valuestorage/*.go ./configs/config.dev.json

buildimage:
	echo Building
	docker build --tag "wutchzone/value-storage-service" --file ./build/Dockerfile .
