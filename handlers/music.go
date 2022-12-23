package handlers

import (
	"context"
	musicdto "dumbsound/dto/music"
	dto "dumbsound/dto/result"
	"dumbsound/models"
	"dumbsound/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerMusic struct {
	MusicRepository repositories.MusicRepository
}

func HandlerMusic(MusicRepository repositories.MusicRepository) *handlerMusic {
	return &handlerMusic{MusicRepository}
}

func (h *handlerMusic) FindMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userSubscribe := userInfo["subscribe"]

	musics, err := h.MusicRepository.FindMusics()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// looping image
	// for i, m := range musics {
	// 		musics[i].Thumbnail = os.Getenv("PATH_FILE") + m.Thumbnail\
	// }

	if userSubscribe == "true" {
		var responseSubscribe []musicdto.MusicSubscribeResponse

		for _, m := range musics {
			responseSubscribe = append(responseSubscribe, convertResponseSubs(m))
		}
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Status: "success", Data: responseSubscribe}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		var responseMusics []musicdto.MusicResponse

		for _, m := range musics {
			responseMusics = append(responseMusics, convertResponseMusic(m))
		}
		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Status: "success", Data: responseMusics}
		json.NewEncoder(w).Encode(response)
		return
	}

}

func (h *handlerMusic) GetAllMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	musics, err := h.MusicRepository.FindMusics()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var responseMusics []musicdto.MusicResponse

	for _, m := range musics {
		responseMusics = append(responseMusics, convertResponseMusic(m))
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: responseMusics}
	json.NewEncoder(w).Encode(response)

}

func (h *handlerMusic) CreateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["status"]

	dataContex := r.Context().Value("dataThumbnail") // Thumbnail (image)
	fileThumbnail := dataContex.(string)

	musicContex := r.Context().Value("dataMusic") // Thumbnail (image)
	fileMusic := musicContex.(string)
	if userRole != "1" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := dto.ErrorResult{Code: http.StatusMethodNotAllowed, Message: "User is not an admin"}
		json.NewEncoder(w).Encode(response)
		return
	}

	artisId, _ := strconv.Atoi(r.FormValue("artisId"))
	request := musicdto.MusicRequest{
		Title:   r.FormValue("title"),
		Year:    r.FormValue("year"),
		ArtisID: artisId,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Declare Context Background, Cloud Name, API Key, API Secret ...
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	respThumbnail, err := cld.Upload.Upload(ctx, fileThumbnail, uploader.UploadParams{Folder: "dumbsound/image"}) //image upload

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	respMusic, err := cld.Upload.Upload(ctx, fileMusic, uploader.UploadParams{Folder: "dumbsound/music"})

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	music := models.Music{
		Title:     request.Title,
		Year:      request.Year,
		Thumbnail: respThumbnail.SecureURL,
		ArtisID:   request.ArtisID,
		Attache:   respMusic.SecureURL,
	}

	data, err := h.MusicRepository.CreateMusic(music)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, _ = h.MusicRepository.GetMusicID(data.ID)

	musicResponse := musicdto.MusicSubscribeResponse{
		ID:        data.ID,
		Title:     data.Title,
		Thumbnail: data.Thumbnail,
		Attache:   data.Attache,
		Artis:     data.Artis,
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: musicResponse}
	json.NewEncoder(w).Encode(response)
}

func convertResponseMusic(u models.Music) musicdto.MusicResponse {
	return musicdto.MusicResponse{
		ID:        u.ID,
		Title:     u.Title,
		Year:      u.Year,
		Thumbnail: u.Thumbnail,
		Artis:     u.Artis,
	}
}

func convertResponseSubs(u models.Music) musicdto.MusicSubscribeResponse {
	return musicdto.MusicSubscribeResponse{
		ID:        u.ID,
		Title:     u.Title,
		Year:      u.Year,
		Attache:   u.Attache,
		Thumbnail: u.Thumbnail,
		Artis:     u.Artis,
	}
}
