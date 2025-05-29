include .env
build:
	go build -o out/ .

run:
	$(MAKE) build
	cd out && DB_USER=$(DB_USER) DB_PASSWORD=$(DB_PASSWORD) ./caisse-app

test:
	golangci-lint run
	docker-compose up -d dbtest
	go test -v ./tests/...

clean:
	rm -rf out/*
	
clean-docker:
	docker stop $$(docker ps -a -q) || true
	docker rm $$(docker ps -a -q) || true
	docker image rm $$(docker image ls -q) || true
	docker volume rm $$(docker volume ls -q) || true

deploy:
	$(MAKE) build
	docker push $(USERNAME)/caisse-app:latest

dev-setup:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	echo $(PWD) | docker login -u $(USERNAME) --password-stdin
	go mod tidy
	