IMAGE = leanlabs/blacksmith
VERSION = 0.0.1

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
