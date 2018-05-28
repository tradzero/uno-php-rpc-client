// Copyright 2018 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2018-05-26 10:48:05
// Version 1.0.0
// Desc main.go
package main

import (
	"context"
	"net"
	"net/http"
	"strconv"

	"github.com/acrazing/uno"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug         = kingpin.Flag("debug", "enable debug mode").Short('d').Bool()
	serveAddress  = kingpin.Flag("address", "the address to listen").Short('a').Default("0.0.0.0:6110").String()
	serveProtocol = kingpin.Flag("protocol", "the protocol to use").Short('p').Default("http").Enum("grpc", "http")
	minValue      = kingpin.Flag("min-value", "the min value to rent").Short('S').Default("100000").Uint32()
	maxValue      = kingpin.Flag("max-value", "the min value to rent").Short('B').Default("1000000").Uint32()
	ttl           = kingpin.Flag("TTL", "the time to live for rent").Short('L').Default("3600s").Duration()
	ttf           = kingpin.Flag("TTF", "the time to freeze for return").Short('F').Default("3600s").Duration()
	volume        = kingpin.Flag("volume", "the pool volume").Short('V').Default("1000").Uint32()

	worker  = uno.Service.Worker
	service = uno.Service
)

func serveGRPC() {
	listener, err := net.Listen("tcp", *serveAddress)
	if err != nil {
		panic(err)
	}
	grpc.EnableTracing = *debug
	server := grpc.NewServer()
	uno.RegisterUnoServer(server, service)
	server.Serve(listener)
}

func serveHTTP() {
	if *debug {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}
	engine := gin.Default()
	engine.POST("/", func(c *gin.Context) {
		no := worker.Rent()
		if no == 0 {
			c.String(http.StatusServiceUnavailable, "the id pool is exhausted")
		} else {
			c.String(http.StatusOK, strconv.Itoa(int(no)))
		}
	})
	engine.PUT("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		} else if int(uint32(id)) != id {
			c.String(http.StatusBadRequest, "invalid id to relet")
		} else {
			ok := worker.Relet(uint32(id))
			if ok {
				c.String(http.StatusOK, "")
			} else {
				c.String(http.StatusBadRequest, "the specified id is not live")
			}
		}
	})
	engine.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		} else if int(uint32(id)) != id {
			c.String(http.StatusBadRequest, "invalid id to relet")
		} else {
			worker.Return(uint32(id))
			c.String(http.StatusOK, "")
		}
	})
	engine.Run(*serveAddress)
}

func main() {
	kingpin.CommandLine.Help = "Start a uno service"
	kingpin.Parse()
	worker.Init(&uno.Option{
		PoolVolume: *volume,
		TTL:        *ttl,
		TTF:        *ttf,
		MinValue:   *minValue,
		MaxValue:   *maxValue,
		Debug:      *debug,
	})
	go worker.Run(context.Background())
	switch *serveProtocol {
	case "grpc":
		serveGRPC()
	case "http":
		serveHTTP()
	}
}
