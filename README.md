# Gorgeous George
_He's not pretty, but he packs a punch_

![](/george.jpeg)

When working with Docker containers I find that it's pretty difficult to view and have a comprehensive understanding of what is being hosted multiple Docker Machine instances. Getting a status view of the multiple Daemons and their currently running docker containers and have the basic commands at your disposal from the web browser would be a boon. Enter, Gorgeous George.

![](/ss.png)

## Usage

```bash
hub clone gophergala2016/GorgeousGeorge
cd GorgeousGeorge
go build
./GorgeousGeorge --help

NAME:
   Gorgeous George - Get a beautiful look at all the Docker containers before they knock you out.

USAGE:
   GorgeousGeorge [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port "10001"	Port on which to host Gorgeous George
   --env		Grab environment variables set by Docker Machine
   --machine_path 	Path to Docker Machine directory, Gorgeous George will scan and monitor all available Docker Machines
   --help, -h		show help
   --version, -v	print the version
```


## Features

- [x] Connecting front-end to back via WebSockets
- [x] Basic CLI options setup
- [x] Rendering of Docker containers in browser
- [ ] Connect simultaneously to multiple daemons based on Docker Machine directory
- [ ] Setup goroutine to ping daemons and update dashboard via websockets
- [ ] Allow addition of new Daemon endpoints from browser
- [ ] Allow stop/start/shutdown of images from browser
