package main

import "github.com/fsouza/go-dockerclient"

// message sent to us by the javascript client
type message struct {
	Status string `json:"status"`
}

type initData struct {
	Status     string
	Containers []docker.APIContainers
}
