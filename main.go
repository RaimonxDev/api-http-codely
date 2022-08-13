package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const httpAddr = ":8080"

func main() {
	fmt.Println("Server running on", httpAddr)

	srv := gin.New()
	srv.GET("/hello", healtHandler)

	log.Fatal(srv.Run(httpAddr))

}

func healtHandler(context *gin.Context) {
	context.String(http.StatusOK, "todo bien")
}
