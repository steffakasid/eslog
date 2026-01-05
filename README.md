[![Go Test](https://github.com/steffakasid/eslog/actions/workflows/go-test.yml/badge.svg)](https://github.com/steffakasid/eslog/actions/workflows/go-test.yml) [![CodeQL](https://github.com/steffakasid/eslog/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/steffakasid/eslog/actions/workflows/codeql-analysis.yml) [![Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)

# eslog

Lightweight, structured logger for Go projects. It basically adds some functions to slog which I found missing (e.g. log*f functions and a logFatal function known from standard log package.). Slog is wrapped inside eslog so you can do everything which is possible with slog.

This project is intended to not have any dapendencies.


## Install
go get github.com/steffakasid/eslog@latest

## Quick start
```go
package main

import (
    "github.com/spf13/viper"
    "github.com/steffakasid/eslog"
)

func main() {
    viper.SetDefault("log.level", "info")
    viper.SetDefault("log.format", "json")
    // viper.ReadInConfig() // optional

    cfg := eslog.Config{
        Level:  viper.GetString("log.level"),
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
```

## Configuration
- log.level: debug|info|warn|error
- log.format: json|text
- output (optional): file path or stdout/stderr

## Contributing
PRs welcome. Please follow gofmt and golangci-lint rules.

## License
Apache v2