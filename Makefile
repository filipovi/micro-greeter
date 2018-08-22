# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help all build-proto build-docker

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help 

build-proto: ## Generates the protobuf files   
	docker run --rm -v $(pwd):$(pwd) -w $(pwd) filipovi/docker-protobuf --go_out=. -I. proto/greeter.proto --micro_out=.

build-docker: ## Build the docker image
	docker build -t micro-greeter .

all: build-proto build-docker ## build all the project

