package routes

import (
	"dumbsound/handlers"
	"dumbsound/pkg/middleware"
	"dumbsound/pkg/mysql"
	"dumbsound/repositories"

	"github.com/gorilla/mux"
)

func MusicRoutes(r *mux.Router) {
	musicRepository := repositories.RepositoryMusic(mysql.DB)
	h := handlers.HandlerMusic(musicRepository)

	r.HandleFunc("/music", middleware.Auth(middleware.UploadThumbnail(middleware.UploadMusic(h.CreateMusic)))).Methods("POST")
	r.HandleFunc("/musics", middleware.Auth(h.FindMusics)).Methods("GET")
	r.HandleFunc("/gmusics", h.GetAllMusics).Methods("GET")
}
