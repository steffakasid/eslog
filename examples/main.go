package main

import (
	"github.com/spf13/viper"
	"github.com/steffakasid/eslog"
)

func main() {
    viper.SetDefault("log.level", "info")
    viper.SetDefault("log.format", "json")
    // viper.ReadInConfig() // optional

        logLevel, err := eslog.ParseText(viper.GetString("log.level"))

    cfg := eslog.Config{
        Level:  logLevel,
        Format: viper.GetString("log.format"),
    }

    logger, err := eslog.New(cfg)
    if err != nil {
        panic(err)
    }

    logger.Info("service started",
        eslog.Field("version", "v0.1.0"),
    )
}