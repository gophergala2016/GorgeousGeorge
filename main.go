package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/fsouza/go-dockerclient"
)

func main() {

	client, _ := docker.NewClientFromEnv()

	fmt.Println(client.Version())

	app := cli.NewApp()
	app.Name = "Gorgeous George"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port",
			Value: "10001",
			Usage: "Port on which to host Gorgeous George",
		},
	}

	app.Usage = "Get a beautiful look at all the Docker containers in your life."
	app.Action = func(c *cli.Context) {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, "Welcome to the home page!")
		})

		n := negroni.Classic()
		n.UseHandler(mux)

		n.Run(fmt.Sprintf(":%s", c.String("port")))
	}

	app.Run(os.Args)

}
