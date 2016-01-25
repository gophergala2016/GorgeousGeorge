package main

import (
	"encoding/json"
	"fmt"
	"io"
	// "log"

	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"
	"github.com/fsouza/go-dockerclient"
	"github.com/gorilla/websocket"
)

// var client *docker.Client

func startWebServer(c *cli.Context) {

	// t, err := template.New("home").ParseFiles("index.html")
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "index.html")
	})

	mux.HandleFunc("/ws", handleWebsocket)

	mux.HandleFunc("/ps", func(w http.ResponseWriter, req *http.Request) {
		client, err := docker.NewClientFromEnv()
		if err != nil {
			log.Fatal(err)
		}
		containers, err := client.ListContainers(docker.ListContainersOptions{
			All:  true,
			Size: true,
		})

		log.Println(containers)

		log.Println(err)

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

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func validateMessage(data []byte) (message, error) {

	var msg message

	log.Info(string(data[:]))
	log.Infof("%#v %T %#q`", data, data, data)
	if err := json.Unmarshal(data, &msg); err != nil {
		return msg, err
	}

	if msg.Status == "" {
		return msg, fmt.Errorf("Message has no Handle or Text")
	}

	return msg, nil
}

// handleWebsocket connection. Update to
func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	log.Println("Attempting websocket upgrade...")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithField("err", err).Println("Upgrading to websockets")
		http.Error(w, "Error Upgrading to websockets", 400)
		return
	}

	for {
		mt, data, err := ws.ReadMessage()
		ctx := log.Fields{"mt": mt, "data": data, "err": err}
		if err != nil {
			if err == io.EOF {
				log.WithFields(ctx).Info("Websocket closed!")
			} else {
				log.WithFields(ctx).Error("Error reading websocket message")
			}
			break
		}
		switch mt {
		case websocket.TextMessage:
			msg, err := validateMessage(data)
			if err != nil {
				ctx["msg"] = msg
				ctx["err"] = err
				log.WithFields(ctx).Error("Invalid Message")
				break
			}

			switch msg.Status {

			case "init":

				var data initData

				client, err := docker.NewClientFromEnv()
				if err != nil {
					log.Fatal(err)
				}

				data.Containers, err = client.ListContainers(docker.ListContainersOptions{
					All: true,
				})
				if err != nil {
					log.Fatal(err)
				}

				data.Status = "initResponse"

				jsonOut, err := json.Marshal(data)
				if err != nil {
					log.Fatal(err)
				}
				ws.WriteMessage(websocket.TextMessage, jsonOut)
			}
			log.WithFields(ctx).Info(msg)
			// rw.publish(data)
		default:
			log.WithFields(ctx).Warning("Unknown Message!")
		}
	}

	// rr.deRegister(id)

	ws.WriteMessage(websocket.CloseMessage, []byte{})
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

		if c.Bool("env") {
			client, err := docker.NewClientFromEnv()
			if err != nil {
				log.Fatal(err)
			}
			log.Println(client.Version())
		} else if c.String("machine_path") != "" {
			log.Println("Connecting to Machines")
			// Connect to machines here
		}

		startWebServer(c)
	}

	app.Run(os.Args)

}
