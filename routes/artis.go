package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func ArtisRoutes(r *mux.Router) {
	artisRepository := repositories.RepositoryArtis(mysql.DB)
	h := handlers.HandlerArtis(artisRepository)

	r.HandleFunc("/artists", h.FindArtists).Methods("GET")
	r.HandleFunc("/artis", middleware.Auth(h.CreateArtis)).Methods("POST")

}
