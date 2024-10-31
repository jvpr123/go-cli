package commands

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func ServerName() cli.Command {
	return cli.Command{
		Name:  "server-name",
		Usage: "Seek for server names related to provided host",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "host", Value: "localhost"},
		},
		Action: searchServerName,
	}
}

func searchServerName(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
