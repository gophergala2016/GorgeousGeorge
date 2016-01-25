package main

import "github.com/fsouza/go-dockerclient"

// message sent to us by the javascript client
type Message struct {
	Status string `json:"status"`
}

type InitData struct {
	Status     string
	Containers []docker.APIContainers
}
