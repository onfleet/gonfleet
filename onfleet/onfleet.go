package onfleet

import (
	"github.com/onfleet/gonfleet/config"
	"github.com/onfleet/gonfleet/resource/worker"
	"github.com/onfleet/gonfleet/util/netw"
)

type Onfleet struct {
	Workers worker.Client
}

func New(params config.InitParams) *Onfleet {
	o := Onfleet{}

	timeout := config.DefaultUserTimeout
	if params.UserTimeout != 0 {
		timeout = params.UserTimeout
	}

	c := config.InitConfig(params)
	httpClient := netw.NewHttpClient(timeout)

	o.Workers = worker.Client{
		Config:     c,
		HttpClient: httpClient,
		SubPath:    "/workers",
	}
	return &o
}
