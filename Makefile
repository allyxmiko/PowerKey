PROJECT_NAME := PowerKey
DOCKER_TAG := power-key:latest

build:
	docker container rm -f $(PROJECT_NAME) > /dev/null 2>&1 || true
	docker image rm -f $(DOCKER_TAG) > /dev/null 2>&1 || true
	docker build -t $(DOCKER_TAG) .
	docker run -d -p 3000:3000 --name PowerKey -v /www/software/PowerKey/:/app/data $(DOCKER_TAG)