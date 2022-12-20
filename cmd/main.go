package main

import (
	"github.com/imrenagi/imrefs/cmd/cli"
	"github.com/rs/zerolog/log"
)

func main() {
	err := cli.NewCommand().Execute()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
