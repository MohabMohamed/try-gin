package main

import (
	"github.com/MohabMohamed/try-gin/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	r := gin.Default()
	r.Run()
}
