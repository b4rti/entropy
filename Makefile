
USER := b4rti
REPO := ddvs
TAG := latest

IMAGE := $(USER)/$(REPO):$(TAG)

all: clean image rootfs plugin

.PHONY: clean
clean:
	rm -rf ./plugin

.PHONY: image
image: clean
	docker build --rm -f Dockerfile -t $(IMAGE) .

.PHONY: rootfs
rootfs: image
	mkdir -p ./plugin/rootfs
	$(eval ID = $(shell docker create b4rti/ddvs true))
	docker export $(ID) | tar -x -C ./plugin/rootfs
	docker rm $(ID) && docker rmi --force $(IMAGE)

.PHONY: plugin
plugin: rootfs
	cp ./config.json ./plugin
	(docker plugin remove --force $(IMAGE) || true)
	docker plugin create $(IMAGE) ./plugin
	rm -rf ./plugin

.PHONY: enable
enable: plugin
	docker plugin enable $(IMAGE)

.PHONY: push
push: plugin
	docker login -u $(USER)
	docker plugin push $(IMAGE)
