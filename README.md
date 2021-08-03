# zlogres

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) [![Go Reference](https://pkg.go.dev/badge/github.com/buildingwatsize/zlogres@v0.1.0.svg)](https://pkg.go.dev/github.com/buildingwatsize/zlogres@v0.1.0) [![GitHub issues](https://img.shields.io/github/issues/buildingwatsize/zlogres)](https://github.com/buildingwatsize/zlogres/issues) [![GitHub forks](https://img.shields.io/github/forks/buildingwatsize/zlogres)](https://github.com/buildingwatsize/zlogres/network) [![GitHub stars](https://img.shields.io/github/stars/buildingwatsize/zlogres)](https://github.com/buildingwatsize/zlogres/stargazers)

zlogres is a middleware for Fiber that logging about api took time since request to response.

## Table of Contents

- [zlogres](#zlogres)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Signatures](#signatures)
  - [Examples](#examples)
  - [Config](#config)
  - [Default Config](#default-config)
  - [Dependencies](#dependencies)
  - [Example Usage](#example-usage)

## Installation

```bash
  go get -u github.com/buildingwatsize/zlogres
```

## Signatures

```go
func New(config ...Config) fiber.Handler
```

## Examples

Import the middleware package that is part of the Fiber web framework

```go
import (
  "github.com/gofiber/fiber/v2"
  "github.com/buildingwatsize/zlogres"
)
```

After you initiate your Fiber app, you can use the following possibilities:

```go
// Default
app.Use(zlogres.New())

// this middleware supported the `requestid` middleware
app.Use(requestid.New())
app.Use(zlogres.New())

// Or extend your config for customization
app.Use(requestid.New(requestid.Config{
  ContextKey: "transaction-id",
}))
app.Use(zlogres.New(zlogres.Config{
  RequestIDContextKey: "transaction-id",
}))
```

## Config

```go
// Config defines the config for middleware.
type Config struct {
  // Optional. Default: nil
  Next func(c *fiber.Ctx) bool

  // Optional. Default: "requestid"
  RequestIDContextKey string
}
```

## Default Config

```go
var ConfigDefault = Config{
  Next:                nil,
  RequestIDContextKey: "requestid",
}
```

## Dependencies

- [Zerolog](https://github.com/rs/zerolog)
- [Fiber](https://github.com/gofiber/fiber)

## Example Usage

Please go to [example/main.go](./example/main.go)
Note: Custom usage please focus on `Custom` section
