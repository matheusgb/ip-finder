package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "ip-finder"
	app.Usage = "Busca por IPs e nomes de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "findip",
			Usage:  "Busca por IPs",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "findservers",
			Usage:  "Busca por nomes de servidores",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(context *cli.Context) {
	host := context.String("host")
	servers, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		log.Println(server)
	}
}

func searchServers(context *cli.Context) {
	host := context.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		log.Println(server)
	}
}
