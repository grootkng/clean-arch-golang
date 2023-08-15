package main

import (
	"github.com/grootkng/clean-arch-golang/config"
	"github.com/grootkng/clean-arch-golang/internal/pkg/api/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(config.GetEnv().API_PORT)
}
