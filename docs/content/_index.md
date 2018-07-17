+++
title = "Entropy"
date = 2018-07-15T21:04:34+02:00
weight = 5

+++

# Entropy
## Distributed Docker Volume System


Entropy is a scalable Network Filesystem Plugin for Docker Swarm.  It allows  to combine Disk Space on many Swarm Nodes into a single name space. 

Entropy wants to make easy replication of Docker Volumes from one Node to another and stream all the changes without external service dependencies. Entropy keeps Metadata and the Data itself separately. And Keep your Data precisely where the Data is needed to provide fast I/O.

 - [Quick Start Guide](/installation/)
 - [Administration Guide](/administration/)
 - [Troubleshooting](https://github.com/b4rti/Entropy/issues)

## Why we need Entropy

Docker has exelent Drivers, but none for Volumes. The user cannot simply use/reuse the Volumes for new Containers on other Swarm Nodes without hassle.


## Current Status

- WIP / proof of concept

## Inspired by
    
- https://github.com/sapk/docker-volume-rclone
- https://github.com/ContainX/docker-volume-netshare
- https://github.com/vieux/docker-volume-sshfs
- https://github.com/sapk/docker-volume-gvfs
- https://github.com/calavera/docker-volume-glusterfs
- https://github.com/rexray/rexray
- https://github.com/lizardfs/lizardfs
- https://github.com/minio/minio
- https://github.com/openstack/swift

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
