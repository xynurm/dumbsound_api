package musicdto

import "dumbsound/models"

type MusicResponse struct {
	ID        int          `json:"id"`
	Title     string       `json:"title"`
	Year      string       `json:"year"`
	Thumbnail string       `json:"thumbnail"`
	Attache   string       `json:"attache"`
	Artis     models.Artis `json:"artis"`
}
