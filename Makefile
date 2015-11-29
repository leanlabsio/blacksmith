IMAGE = leanlabs/blacksmith
VERSION = 0.0.1

all: build_image

blacksmith: $(find $(CURDIR) -name "*.go" -type f)
	@docker run --rm \
		-v $(CURDIR):/src \
		leanlabs/golang-builder

build_image: blacksmith
	@docker build -t $(IMAGE) .

