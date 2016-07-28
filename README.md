[![Go Report](https://goreportcard.com/badge/github.com/LinkerNetworks/remote-docker-exec)](https://goreportcard.com/report/github.com/LinkerNetworks/remote-docker-exec)
[![Build Status](https://travis-ci.org/LinkerNetworks/remote-docker-exec.svg)](https://travis-ci.org/LinkerNetworks/remote-docker-exec)

# remote-docker-exec
Connect to docker daemon or swarm with TLS. Run command 'sh' in container.

# Precondition
To run this programme, a TLS enabled docker daemon is required with port binded.

TCP port is not enable in docker daemon by default, click [here][1] to see how to bind a port.

To enable TLS(SSL), you need generate cert files by yourself with **openssl**, follow this [article][2] for detail. After that, restart docker daemon with --tls flags to enable TLS, follow docker [docs][3].

# Params

**ContainerId:** name or ID of container. For docker daemon, support containers created by daemon. For swarm, support containers created by all daemons in cluster.

# Version
Support docker 1.11.x with API version 1.23

# LICENSE
MIT

# Related projects
[LinkerNetworks/linkerDcosDockerFile/linkerWebConsole][https://github.com/LinkerNetworks/linkerDcosDockerFile/tree/console/linkerWebConsole]: Dockerfile, golang binaries and config files

[LinkerNetworks/gotty][https://github.com/LinkerNetworks/gotty]: Linker GoTTY is a simple command line tool that turns remote-docker-exec into web applications.

[1]:https://docs.docker.com/engine/reference/commandline/dockerd/#bind-docker-to-another-host-port-or-a-unix-socket
[2]:https://jamielinux.com/docs/openssl-certificate-authority/
[3]:https://docs.docker.com/engine/security/https/
