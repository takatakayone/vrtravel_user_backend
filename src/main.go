package main

import (
	"github.com/takatakayone/vrtravel_user_backend/src/infrastructure/router"
	"github.com/takatakayone/vrtravel_user_backend/src/registry"
)

func main() {
	r := registry.NewRegistry()
	router.NewRouter(r.NewAppHandler())
}