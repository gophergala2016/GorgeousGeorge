package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/fsouza/go-dockerclient"
)

func startWebServer(c *cli.Context, client docker.Client) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	mux.HandleFunc("/ps", func(w http.ResponseWriter, req *http.Request) {
		containers, err := client.ListContainers(docker.ListContainersOptions{
			All:  true,
			Size: true,
		})
		if err != nil {
			log.Fatal(err)
		}

		for _, container := range containers {
			fmt.Fprintf(w, container.ID)
		}
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	n.Run(fmt.Sprintf(":%s", c.String("port")))
}

func main() {

	app := cli.NewApp()
	app.Name = "Gorgeous George"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port",
			Value: "10001",
			Usage: "Port on which to host Gorgeous George",
		},
		cli.BoolFlag{
			Name:  "env",
			Usage: "Grab environment variables set by Docker Machine",
		},
		cli.StringFlag{
			Name:  "machine_path",
			Usage: "Path to Docker Machine directory, Gorgeous George will scan and monitor all available Docker Machines",
		},
	}

	app.Usage = "Get a beautiful look at all the Docker containers in your life."
	app.Action = func(c *cli.Context) {

		var client docker.Client
		if c.Bool("env") {
			client, _ := docker.NewClientFromEnv()
			log.Println(client.Version())
		} else if c.String("machine_path") != "" {
			log.Println("Connecting to Machines")
			// Connect to machines here
		}

		startWebServer(c, client)
	}

	app.Run(os.Args)

}
