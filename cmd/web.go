package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/jianggushi/topstory/controllers"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var Web = cli.Command{
	Name:        "web",
	Usage:       "start web server",
	Description: "",
	Action:      runWeb,
	Flags:       []cli.Flag{},
}

func runWeb(c *cli.Context) error {
	r := gin.Default()
	r.GET("/nodes", controllers.ListNodes)
	r.GET("/nodes/:id", controllers.GetNodeByID)
	r.GET("/nodes/:id/lastitem", controllers.GetLastItem)
	r.GET("/nodes/:id/items", controllers.GetItems)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Panic("failed to run web server")
	}
	return nil
}
