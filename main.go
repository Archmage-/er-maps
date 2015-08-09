package main

import (
	"fmt"
	"github.com/Archmage-/er-maps/core"
	"github.com/Archmage-/er-maps/modules/minemap"
	"github.com/Archmage-/er-maps/server"
	"github.com/Archmage-/er-maps/server/handler"
	"log"
	"net/http"
)

func main() {
	provider := minemapprovider.NewProvider("http://localhost/map0.txt", 300)
	mapgetsector := server.NewMapGetSector(provider)
	fmt.Println("starting")
	setRouteHandler("/minemap/getsector", mapgetsector)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func setRouteHandler(route string, wrap core.IRequestHandler) {
	http.Handle(route, &handler.RContextHandler{InnerHandler: wrap})
}
