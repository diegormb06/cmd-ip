package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gen() *cli.App {
	app := cli.NewApp()
	app.Name = "CMD IP"
	app.Usage = "Get ips and host names from internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "Get ips from internet urls",
			Flags: flags,
			Action: getIps,
		},
		{
			Name: "hostname",
			Usage: "Return host name form an url",
			Flags: flags,
			Action: getHostName,
		},
	}

	return app
}

func getIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if(erro != nil) {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getHostName(c *cli.Context) {
	host := c.String("host")

	hostNames, erro := net.LookupNS(host)

	if(erro != nil) {
		log.Fatal(erro)
	}

	for _, hostname := range hostNames {
		fmt.Println(hostname.Host)
	}
}
