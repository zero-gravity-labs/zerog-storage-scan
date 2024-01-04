package main

import (
	"github.com/Conflux-Chain/go-conflux-util/config"
	"github.com/Conflux-Chain/go-conflux-util/log"
	"github.com/zero-gravity-labs/zerog-storage-scan/cmd"
)

func main() {
	config.MustInit("zerog_storage")
	log.MustInitFromViper()
	cmd.Execute()
}
