package main

import (
	"github.com/tpodg/go-config"
	"log"
)

type cfg struct {
	Grpc grpcCfg
	Rest restCfg
}

type grpcCfg struct {
	Target string
}

type restCfg struct {
	Target string
}

func (c *cfg) Parse() {
	conf := config.New()
	conf.WithProviders(&config.Env{})
	if err := conf.Parse(c); err != nil {
		log.Fatal("failed to parse config", c)
	}
}
