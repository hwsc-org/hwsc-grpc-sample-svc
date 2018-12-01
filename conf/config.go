package conf

import (
	"fmt"
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/env"
	"github.com/micro/go-config/source/file"
	"log"
)

var (
	// GRPCHost address and port of gRPC microservice
	GRPCHost Host
)

func init() {
	// Create new config
	conf := config.NewConfig()
	if err := conf.Load(file.NewSource(file.WithPath("conf/json/config.dev.json"))); err != nil {
		// TODO - This is a hacky solution for the unit test, because of a weird path issue with GoLang Unit Test
		if err := conf.Load(file.NewSource(file.WithPath("../conf/json/config.dev.json"))); err != nil {
			log.Printf("[INFO] Failed to initialize configuration file %v\n", err)
			log.Println("[INFO] Reading ENV variables")
			src := env.NewSource(
				env.WithPrefix("hosts"),
			)
			if err := conf.Load(src); err != nil {
				log.Fatalf("[FATAL] Failed to initialize configuration %v\n", err)

			}
		}
	}

	if err := conf.Get("hosts", "grpc").Scan(&GRPCHost); err != nil {
		log.Fatalf("[FATAL] Failed to get configuration %v\n", err)
	}
}

// Host represents a server.
type Host struct {
	Address string `json:"address"`
	Port    string `json:"port"`
	Network string `json:"network"`
}

func (h *Host) String() string {
	return fmt.Sprintf("%s:%s", h.Address, h.Port)
}
