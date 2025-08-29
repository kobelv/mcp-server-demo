// Package main 提供基础的服务入口 在这里进行初始化操作
package main

import (
	"flag"
	"log"
	"mcp-server-demo/infrastructure/adapter"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	var transport string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or http)")
	flag.StringVar(&transport, "transport", "stdio", "Transport type (stdio or http)")
	flag.Parse()

	// Only check for "http" since stdio is the default
	if transport == "http" {
		httpServer := server.NewStreamableHTTPServer(adapter.NewMCPServer())
		log.Printf("HTTP server listening on :8080/mcp")
		if err := httpServer.Start(":8080"); err != nil {
			log.Fatalf("Server error: %v", err)
			panic(any(err))
		}
	} else {
		if err := server.ServeStdio(adapter.NewStdioMCPServer()); err != nil {
			log.Fatalf("Server error: %v", err)
			panic(any(err))
		}
	}
}
