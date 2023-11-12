package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/rs/zerolog"
	"github.com/soulgarden/twelvedata"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout)

	cfg := twelvedata.Conf{
		APIKey: "4e0133f255164c499a387977ce017ebc",
	}

	if err := configor.New(&configor.Config{}).Load(&cfg); err != nil {
		logger.Err(err).Msg("init config")

		return
	}

	cli := twelvedata.NewCli(&cfg, twelvedata.NewHTTPCli(&fasthttp.Client{}, &cfg, &logger), &logger)
	resp, creditsLeft, creditsUsed, err := cli.GetEtfs("", "", "", "", true, true)

	fmt.Println(resp, creditsLeft, creditsUsed, err)
}
