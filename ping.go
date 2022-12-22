package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func pingHandler(ctx *gin.Context) {
	req := resty.New().R().SetHeader("Content-Type", "application/text")
	url := "http://silly-demo:8080"
	log.Printf("Sending a ping to %s", url)
	resp, err := req.Get(url)
	if err != nil {
		httpErrorBadRequest(err, ctx)
		return
	}
	log.Println(resp.String())
	ctx.String(http.StatusOK, resp.String())
}
