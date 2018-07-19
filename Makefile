
USER := b4rti
REPO := entropy
TAG := latest

IMAGE := $(USER)/$(REPO):$(TAG)

BUILD := ./build
ROOTFS := $(BUILD)/rootfs
CONFIG := ./plugin/config.json

all: clean image rootfs plugin

.PHONY: clean
clean:
	rm -rf $(BUILD)

.PHONY: image
image: clean
	docker build --rm -t $(IMAGE) .

.PHONY: rootfs
rootfs: image
	mkdir -p $(ROOTFS)
	$(eval ID = $(shell docker create b4rti/entropy true))
	docker export $(ID) | tar -x -C $(ROOTFS)
	docker rm $(ID) && docker rmi --force $(IMAGE)

.PHONY: plugin
plugin: rootfs
	cp $(CONFIG) $(BUILD)
	(docker plugin remove --force $(IMAGE) || true)
	docker plugin create $(IMAGE) $(BUILD)
	rm -rf $(BUILD)

.PHONY: enable
enable: plugin
	docker plugin enable $(IMAGE)

.PHONY: push
push: plugin
	docker login -u $(USER)
	docker plugin push $(IMAGE)

.PHONY: run
run: image
	docker run --rm -ti \
		-v /run/docker.sock:/run/docker.sock \
		-v /var/lib/docker/plugins:/var/lib/docker/plugins \
		-v /etc/entropy:/etc/entropy \
		$(IMAGE)
