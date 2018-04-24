package main

import (
	"os"

	"github.com/alpagap/c14-rclone/pkg/commands"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/juju/errors"
)

func main() {
	log.SetHandler(text.New(os.Stderr))
	if err := commands.Root.Parse(); err != nil {
		log.Fatal(errors.ErrorStack(err))
	}
}
