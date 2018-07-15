
REPO := b4rti
NAME := ddvs
TAG := latest

IMAGE := $(REPO)/$(NAME):$(TAG)

all:
	make clean image rootfs plugin

clean:
	rm -rf ./plugin

image:
	docker build --rm -f Dockerfile -t $(IMAGE) .

rootfs:
	mkdir -p ./plugin/rootfs
	$(eval ID = $(shell docker create b4rti/ddvs true))
	docker export $(ID) | tar -x -C ./plugin/rootfs
	docker rm $(ID) && docker rmi --force $(IMAGE)

plugin:
	cp ./config.json ./plugin
	(docker plugin remove --force $(IMAGE) || true)
	docker plugin create $(IMAGE) ./plugin
	docker plugin enable $(IMAGE)

