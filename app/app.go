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

	app.Commands = []cli.Command{
		{
			Name:  "findip",
			Usage: "Busca por IPs",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: func(context *cli.Context) {
				host := context.String("host")
				ips, err := net.LookupIP(host)
				if err != nil {
					log.Fatal(err)
				}

				for _, ip := range ips {
					log.Println(ip)
				}
			},
		},
	}

	return app
}
