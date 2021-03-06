package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
	"github.com/zyfdegh/go-dockerpty"
)

func prepare() (endpoints []string, cert, key, ca string) {
	fmt.Println("Welcome to linker web console!")

	endpoints = strings.Split(getSwarmEndpoints(), ",")

	ca, cert, key = PATH_CA, PATH_CERT, PATH_KEY
	if _, err := os.Stat(ca); err != nil {
		log.Fatalf("%s not found\n", ca)
	}
	if _, err := os.Stat(cert); err != nil {
		log.Fatalf("%s not found\n", ca)
	}
	if _, err := os.Stat(key); err != nil {
		log.Fatalf("%s not found\n", ca)
	}
	return endpoints, cert, key, ca
}

// connect to any available swarm endpoint
func startRemoteDockerExec(containerId string) {
	endpoints, cert, key, ca := prepare()
	for _, endpoint := range endpoints {
		err := remoteDockerExec(endpoint, cert, key, ca, containerId)
		if err != nil {
			log.Printf("exec container via swarm[%s] error: %v", endpoint, err)
			continue
		}
	}
}

func remoteDockerExec(endpoint, cert, key, ca, containerId string) (err error) {
	fmt.Printf("Connecting to %s, please wait...\n", endpoint)
	client, err := docker.NewTLSClient(endpoint, cert, key, ca)
	if err != nil {
		log.Printf("new tls client error: %v\n", err)
		return
	}

	fmt.Printf("Connecting to container %s, please wait...\n", containerId)

	// create exec
	createOpts := docker.CreateExecOptions{}
	createOpts.AttachStdin = true
	createOpts.AttachStdout = true
	createOpts.AttachStderr = true
	createOpts.Tty = true
	// select shell sequence
	// bash -> sh -> zsh -> fish -> csh -> tcsh -> scsh -> ksh -> rc
	createOpts.Cmd = []string{"sh", "-c", "bash || sh || zsh || fish || csh || tcsh || scsh || ksh || rc"}
	createOpts.Container = containerId

	exec, err := client.CreateExec(createOpts)
	if err != nil {
		log.Printf("create exec error: %v\n", err)
		return
	}

	// start tty
	err = dockerpty.StartExec(client, exec)
	if err != nil {
		log.Printf("start exec error: %v\n", err)
		return
	}
	return
}

func getSwarmEndpoints() string {
	return os.Getenv(ENV_SWARM_ENDPOINTS)
}
