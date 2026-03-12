package main

import (
	"github.com/MuhammadFelda/task-manager/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	routes.Register(r)
	r.Run()
}