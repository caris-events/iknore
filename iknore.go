package main

import (
	"log"
	"os"

	"github.com/caris-events/iknore/endpoint"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start the iknore file server",
				Action: func(*cli.Context) error {
					return nil
				},
			},
			{
				Name:  "key",
				Usage: "generate the auth key",
				Action: func(*cli.Context) error {
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func router() {
	r := gin.Default()
	e := endpoint.New()

	r.POST("/api/create_pre_signed_file", e.CreatePreSignedFile)
	r.POST("/api/finish_pre_signed_file", e.FinishPreSignedFile)
	r.POST("/api/delete_file", e.DeleteFile)

	r.Run()
}
