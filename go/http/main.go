package main

import (
	"github.com/gin-gonic/gin"
	"http/router/user"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	user.InitRouter(v1)

	port := "9999"
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
