package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata"
	"os"
)

func main() {

	logger := zerolog.New(os.Stdout)

	ctx := context.Background()

	cfg := &twelvedata.Conf{
		APIKey: "4e0133f255164c499a387977ce017ebc",
	}

	if err := configor.New(&configor.Config{}).Load(cfg); err != nil {
		logger.Err(err).Msg("init config")

		return
	}

	wsCli := twelvedata.NewWS(
		cfg,
		&logger,
		nil,
	)

	go wsCli.Subscribe(ctx, []string{"AAPL", "META"})

	for e := range wsCli.Consume() {
		fmt.Println(e)
	}
}
