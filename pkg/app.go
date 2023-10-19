package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//THIS FUNCTION WILL RETURN THE COMMAND LINE APPLICATION READY FOR RUN
func Generated() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line application"
	app.Usage =  "Find server ip's and names in internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	 }
	// HERE I DEFINE THE COMMANDS WILL BE SET ON RUN APPLICATION
	//USING THE STRUCT OF URFAVECLI FOR THAT
	app.Commands = []cli.Command{
		{
		 Name: "find-ip",
		 Usage: "Find ip on internet",
		 Flags: flags,
		 //HERE IS THE FUNCTION OF BE INTIALIZED WHEN THE COMMAND FIND WITH FLAG --HOST AND VALUE BE EXECUTED
		 Action: findIps,
		},

		{
			Name: "find-server",
			Usage: "Find name of servers on the internet",
			Flags: flags,
			Action: findServers,
		},
	}
	
	return app
}

//HERE I DECLARE THE FUNCTION TO FIND THE PUBLIC NAMES OF HOST ON INTERNET USING THE PACKAGE NET 
func  findServers(c *cli.Context){
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, serve := range servers {
		fmt.Println(serve.Host)
	}
}
 

//HERE I DECLARE THE FUNCTION TO FIND THE PUBLIC IP's ON INTERNET USING THE PACKAGE NET 
func findIps(c *cli.Context)  {
	host := c.String("host")

	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}