package handlers

import (
	artisdto "dumbsound/dto/artis"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerArtis struct {
	ArtisRepository repositories.ArtisRepository
}

func HandlerArtis(ArtisRepository repositories.ArtisRepository) *handlerArtis {
	return &handlerArtis{ArtisRepository}
}

func (h *handlerArtis) FindArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	artists, err := h.ArtisRepository.FindArtists()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	// for i, p := range products {
	// 	products[i].Image = os.Getenv("PATH_FILE") + p.Image

	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: artists}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArtis) CreateArtis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["status"]

	if userRole != "1" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := dto.ErrorResult{Code: http.StatusMethodNotAllowed, Message: "User is not an admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	old, _ := strconv.Atoi(r.FormValue("old"))
	request := artisdto.ArtisRequest{
		Old:         old,
		Name:        r.FormValue("name"),
		Type:        r.FormValue("type"),
		StartCareer: r.FormValue("startCareer"),
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	artis := models.Artis{
		Name:        request.Name,
		Old:         request.Old,
		Type:        request.Type,
		StartCareer: request.StartCareer,
	}

	data, err := h.ArtisRepository.CreateArtis(artis)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ = h.ArtisRepository.GetArtis(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}
