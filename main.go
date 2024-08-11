package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"users.com/common"
	"users.com/databases"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error

	err = common.LoadConfig()
	if err != nil {
		return err
	}

	err = databases.Database.Init()
	if err != nil {
		return err
	}

	if common.Config.EnableGinConsoleLog {
		f, _ := os.Create("logs/gin.log")

		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()

	return nil
}

func main() {
	m := Main{}

	if m.initServer() != nil {
		return
	}

	defer databases.Database.Close()

	c := controllers.User{}
}
