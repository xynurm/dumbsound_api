package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	ArtisRoutes(r)
	MusicRoutes(r)
	TransactionRoutes(r)
}
