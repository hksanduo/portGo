/**
* Filename: main.go
* Description: the PortGo main entry point
*   It supports tcp/udp protocol layer traffic forwarding, forward/reverse
*   creation of forwarding links, and multi-level cascading use.
* Author: sanduo
* Time: 2021.01.18
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

/**********************************************************************
* @Function: main()
* @Description: the PortGo entry point, parse command-line argument
* @Parameter: nil
* @Return: nil
**********************************************************************/
func main() {
	app := cli.NewApp()
	app.Name = "portGo"
	app.Version = "V0.1"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		&cli.Author{
			Name:  "hksanduo",
			Email: "hksanduo@gmail.com",
		},
	}
	app.Usage = "port forward tools by go"
	app.Commands = []*cli.Command{
		{
			Name:    "proto",
			Aliases: []string{"p"},
			Usage:   "set network proto",
			Action: func(c *cli.Context) error {
				fmt.Println("test: %q", c.Args().Get(0))
				return nil
			},
		},
	}
	app.Flags = []*cli.Flag{
		&cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Error.Println(err)
	}

	// if len(os.Args) != 4 {
	// 	usage()
	// 	return
	// }
	// proto := os.Args[1]
	// sock1 := os.Args[2]
	// sock2 := os.Args[3]

	// // parse and check argument
	// protocol := PORTFORWARD_PROTO_TCP
	// if strings.ToUpper(proto) == "TCP" {
	// 	protocol = PORTFORWARD_PROTO_TCP
	// } else if strings.ToUpper(proto) == "UDP" {
	// 	protocol = PORTFORWARD_PROTO_UDP
	// } else {
	// 	color.Error.Println("unknown protocol [%s]\n", proto)
	// 	return
	// }

	// m1, a1, err := parseSock(sock1)
	// if err != nil {
	// 	color.Error.Println(err)
	// 	return
	// }
	// m2, a2, err := parseSock(sock2)
	// if err != nil {
	// 	color.Error.Println(err)
	// 	return
	// }

	// // launch
	// args := Args{
	// 	Protocol: protocol,
	// 	Method1:  m1,
	// 	Addr1:    a1,
	// 	Method2:  m2,
	// 	Addr2:    a2,
	// }
	// Launch(args)
}

/**********************************************************************
* @Function: parseSock(sock string) (uint8, string, error)
* @Description: parse and check sock string
* @Parameter: sock string, the sock string from command-line
* @Return: (uint8, string, error), the method, address and error
**********************************************************************/
func parseSock(sock string) (uint8, string, error) {
	// split "method" and "address"
	items := strings.SplitN(sock, ":", 2)
	if len(items) != 2 {
		return PORTFORWARD_SOCK_NIL, "",
			errors.New("host format must [method:address:port]")
	}

	method := items[0]
	address := items[1]
	// check the method field
	if strings.ToUpper(method) == "LISTEN" {
		return PORTFORWARD_SOCK_LISTEN, address, nil
	} else if strings.ToUpper(method) == "CONN" {
		return PORTFORWARD_SOCK_CONN, address, nil
	} else {
		errmsg := fmt.Sprintf("unknown method [%s]", method)
		return PORTFORWARD_SOCK_NIL, "", errors.New(errmsg)
	}
}

/**********************************************************************
* @Function: usage()
* @Description: the PortForward usage
* @Parameter: nil
* @Return: nil
**********************************************************************/
func usage() {
	fmt.Println("Usage:")
	fmt.Println("  ./portGo [proto] [sock1] [sock2]")
	fmt.Println("Option:")
	fmt.Println("  proto      the port forward with protocol(tcp/udp)")
	fmt.Println("  sock       format: [method:address:port]")
	fmt.Println("  method     the sock mode(listen/conn)")
	fmt.Println("Example:")
	fmt.Println("  tcp conn:192.168.1.1:3389 conn:192.168.1.10:23333")
	fmt.Println("  udp listen:192.168.1.3:5353 conn:8.8.8.8:53")
	fmt.Println("  tcp listen:[fe80::1%lo0]:8888 conn:[fe80::1%lo0]:7777")
	fmt.Println()
}
