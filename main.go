package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var serviceName string

func main() {
	log.SetOutput(os.Stderr)
	serviceName = "silly-demo"
	if os.Getenv("SERVICE_NAME") != "" {
		serviceName = os.Getenv("SERVICE_NAME")
	}

	// Server
	log.Println("Starting server...")
	router := gin.New()
	router.GET("/ping", pingHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router.Run(fmt.Sprintf(":%s", port))
}

func httpErrorBadRequest(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusBadRequest)
}

func httpErrorInternalServerError(err error, ctx *gin.Context) {
	httpError(err, ctx, http.StatusInternalServerError)
}

func httpError(err error, ctx *gin.Context, status int) {
	log.Println(err.Error())
	ctx.String(status, err.Error())
}
