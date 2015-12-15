IMAGE   = leanlabs/blacksmith
VERSION = 0.0.1
CWD     = /go/src/github.com/vasiliy-t/blacksmith

all: release

blacksmith: $(shell find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		leanlabs/golang-builder

build_image: blacksmith
	@docker build -t $(IMAGE) .

release: build_image
	@docker login --email=$$DOCKER_HUB_EMAIL --username=$$DOCKER_HUB_LOGIN --password=$$DOCKER_HUB_PASSWORD
	@docker push $(IMAGE):latest

# Development related targets

dev_redis:
	@docker inspect -f {{.State.Running}} bs_dev_redis || \
		docker run -d -p 6379:6379 --name bs_dev_redis leanlabs/redis

dev: dev_redis
	@docker -d \
		--link bs_dev_redis:redis \
		--name bs_dev \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e GO15VENDOREXPERIMENT=1 \
		--entrypoint=/usr/local/go/bin/go \
		golang:1.5.2 run -v main.go --redis-addr redis:6379
