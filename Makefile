# Docker image name
IMAGE   = leanlabs/blacksmith

# Docker image tag
VERSION = 0.0.1

# Current working directory for all targets
# executed in docker builders
CWD     = /go/src/github.com/vasiliy-t/blacksmith

all: release

blacksmith: $(shell find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GO15VENDOREXPERIMENT=1 \
		-e GOOS=linux \
		-e GOARCH=amd64 \
		-e CGO_ENABLED=0 \
		golang:1.5.3 go build -ldflags '-s' -v -o $@

node_modules/: package.json
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		leanlabs/npm-builder npm install

.PHONY: build_image
build_image: blacksmith
	@docker build -t $(IMAGE) .

.PHONY: release
release: build_image
	@docker login \
		--email=$$DOCKER_HUB_EMAIL \
		--username=$$DOCKER_HUB_LOGIN \
		--password=$$DOCKER_HUB_PASSWORD
	@docker push $(IMAGE):latest

# Development related targets

# Start Redis server
.PHONY: dev_redis
dev_redis:
	@docker inspect -f {{.State.Running}} bs_dev_redis || \
		docker run -d -p 6379:6379 --name bs_dev_redis leanlabs/redis

# Install nodejs modules and start Gulp watcher
.PHONY: dev_watcher
dev_watcher: node_modules/
	@docker inspect -f {{.State.Running}} bs_dev_watcher || \
		docker run -d -v $(CURDIR):$(CWD) -w $(CWD) leanlabs/npm-builder gulp watch

# Start golang server
.PHONY: dev
dev: dev_redis dev_watcher
	-docker rm -f bs_dev
	@docker run -d \
		-p 80:80 \
		--link bs_dev_redis:redis \
		--name bs_dev \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GO15VENDOREXPERIMENT=1 \
		-e REDIS_ADDR=redis:6379 \
		-e GITHUB_CLIENT_ID=$(GITHUB_CLIENT_ID) \
		-e GITHUB_CLIENT_SECRET=$(GITHUB_CLIENT_SECRET) \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.5.3 run -v main.go daemon
