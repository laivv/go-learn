package main

import (
	"server"
)

func main (){
	Server :=	new(server.Server)
	Server.Run(8001)
}

