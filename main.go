package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"gocalc/handler"
	"gocalc/handler/operations"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/api/add", handler.MathOperationGet(operations.Add{}))
	r.GET("/api/sub", handler.MathOperationGet(operations.Sub{}))
	r.GET("/api/mul", handler.MathOperationGet(operations.Mul{}))
	r.GET("/api/div", handler.MathOperationGet(operations.Div{}))

	portFlagPtr := flag.Int("port", 8080, "Port to run on")
	flag.Parse()
	err := r.Run(fmt.Sprintf(":%d", *portFlagPtr))
	if err != nil {
		log.Fatal(err)
	}
}
