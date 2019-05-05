package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
	
}

//helper function to get port
func port()string{
	port:=os.Getenv("PORT")
	if len(port) == 0{
		port = 8080
	}
	return ":" + port
}