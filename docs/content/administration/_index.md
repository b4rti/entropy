---
title: "Administration"
date: 2018-07-16T01:11:33+02:00
---

## Plugin Managing

#### Enable Plugin
```
$ docker plugin enable Entropy
$ docker plugin ls
```

#### Disable Plugin
```
$ docker plugin disable -f Entropy
$ docker plugin ls
```





## Volumes

#### Create a Volume

```
$ docker volume create -d b4rti/Entropy --name volume-name
```

#### List Volume

```
$ docker volume ls
```


#### Delete Volume

```
$ docker volume rm volume-name
```

## Run

Docker:
```
$ docker run -v volume-name:/mnt --rm -ti ubuntu
```
Docker Swarm:
```
...
```


```
$ 
```


