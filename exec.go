package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

func prepare() (endpoint, cert, key, ca string) {
	fmt.Println("Welcome to linker web console!")

	ip := getSwarmIp()
	if len(strings.TrimSpace(ip)) == 0 {
		log.Fatalln("invalid swarm ip")
	}

	port := getSwarmPort()
	if len(strings.TrimSpace(port)) == 0 {
		log.Fatalln("invalid swarm port")
	}

	endpoint = fmt.Sprintf("tcp://%s:%s", ip, port)

	ca, cert, key = ".certs/ca.pem", ".certs/cert.pem", ".certs/key.pem"
	if _, err := os.Stat(ca); err != nil {
		log.Fatalf("%s not found\n", ca)
	}
	if _, err := os.Stat(cert); err != nil {
		log.Fatalf("%s not found\n", ca)
	}
	if _, err := os.Stat(key); err != nil {
		log.Fatalf("%s not found\n", ca)
	}

	return endpoint, cert, key, ca
}

func remoteDockerExec(containerId string) {
	endpoint, cert, key, ca := prepare()
	fmt.Printf("Connecting to %s, please wait...\n", endpoint)
	client, err := docker.NewTLSClient(endpoint, cert, key, ca)
	if err != nil {
		log.Fatalf("new tls client error: %v\n", err)
	}

	fmt.Printf("Connecting to container %s, please wait...\n", containerId)

	// create exec
	createOpts := docker.CreateExecOptions{}
	createOpts.AttachStdin = true
	createOpts.AttachStdout = true
	createOpts.AttachStderr = true
	createOpts.Tty = true
	createOpts.Cmd = []string{"sh"}
	createOpts.Container = containerId

	exec, err := client.CreateExec(createOpts)
	if err != nil {
		log.Fatalf("create exec error: %v\n", err)
	}

	// start exec
	startOpts := docker.StartExecOptions{}
	startOpts.Tty = true
	startOpts.RawTerminal = true
	startOpts.Detach = false
	startOpts.ErrorStream = os.Stderr
	startOpts.InputStream = os.Stdin
	startOpts.OutputStream = os.Stdout

	err = client.StartExec(exec.ID, startOpts)
	if err != nil {
		log.Fatalf("start exec error: %v\n", err)
	}
}

func getSwarmIp() string {
	return getEnv("SWARM_IP")
}

func getSwarmPort() string {
	return getEnv("SWARM_PORT")
}
