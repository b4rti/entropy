---
title: "Getting Started"
date: 2018-07-15T21:06:50+02:00
weight: 15
---


## Overview

This section describes how to get started with Entropy Plugin, available for Docker 1.37+  and [Docker Swarm](https://docs.docker.com/engine/swarm/swarm-tutorial/create-swarm/).

## Installation

Docker plugins can be installed with following command on your Node:

```
$ docker plugin install b4rti/Entropy
```
or you can also use go 
```
$ go get -u -d github.com/b4rti/Entropy

$ cd $GOPATH/src/github.com/b4rti/Entropy

$ go run main.go
```


## Try to test

```
$ docker plugin ls

$ docker volume create -d b4rti/Entropy --name volume-name

$ docker run -v volume-name:/mnt --rm -ti ubuntu
```