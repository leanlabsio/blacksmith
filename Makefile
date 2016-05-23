# Docker image name
IMAGE   = leanlabs/blacksmith

# Docker image tag
VERSION = 0.0.1

# Current working directory for all targets
# executed in docker builders
CWD     = /go/src/github.com/leanlabsio/blacksmith

all: release

# Frontend application tasks

## Install frontend dependencies
node_modules/: package.json
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		node:6.1.0-slim npm run install

## Build frontend application
build: node_modules
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		node:6.1.0-slim npm run build


# Backend applications tasks

## Embed templates
templates/templates.go: $(find $(CURDIR)/templates -name "*.html" ! -name "templates.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		leanlabs/go-bindata-builder \
		$(DEBUG) \
		-pkg=templates -o $@ templates/...

## Embed frontend application
web/web.go: $(find $(CURDIR)/web/ -name "*" ! -name "web.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		leanlabs/go-bindata-builder \
		$(DEBUG) \
		-pkg=web -o $@ web/...

## Compile application
blacksmith: build web/web.go templates/templates.go $(shell find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e CGO_ENABLED=0 \
		golang:1.6-alpine go build -ldflags '-X main.Version=$(VERSION)' -v -o $@

## Build docker image with application
build_image: blacksmith
	@docker build -t $(IMAGE) .

## Publish image to docker hub
release: build_image
	@docker login \
		--username=$$DOCKER_HUB_LOGIN \
		--password=$$DOCKER_HUB_PASSWORD
	@docker push $(IMAGE):latest


# Development related targets

## Start Redis server
dev_redis:
	@docker inspect -f {{.State.Running}} bs_dev_redis || \
		docker run -d -p 6379:6379 --name bs_dev_redis leanlabs/redis

## Install nodejs modules and start Gulp watcher
dev_watcher: node_modules/
	@docker inspect -f {{.State.Running}} bs_dev_watcher || \
		docker run -d \
			--name bs_dev_watcher \
			-v $(CURDIR):$(CWD) \
			-w $(CWD) \
			node:6.1.0-slim npm run watch

dev : DEBUG=-debug

## Start golang server
dev: build web/web.go templates/templates.go dev_watcher dev_redis
	-docker rm -f bs_dev
	@docker run -d \
		-p 80:80 \
		--link bs_dev_redis:redis \
		--name bs_dev \
		-v $(CURDIR):$(CWD) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-w $(CWD) \
		-e REDIS_ADDR=redis:6379 \
		-e GITHUB_CLIENT_ID=$(GITHUB_CLIENT_ID) \
		-e GITHUB_CLIENT_SECRET=$(GITHUB_CLIENT_SECRET) \
		-e DOCKER_HOST=unix:///var/run/docker.sock \
		-e BS_SERVER_HOSTNAME=$(BS_SERVER_HOSTNAME) \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.6-alpine run -v main.go daemon

clean:
	@rm -rf $(CURDIR)/web
	@rm -f $(CURDIR)/templates/templates.go
