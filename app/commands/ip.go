package commands

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func IP() cli.Command {
	return cli.Command{
		Name:  "ip",
		Usage: "Seek for IPs related to provided host",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "host", Value: "localhost"},
		},
		Action: searchIP,
	}
}

func searchIP(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
