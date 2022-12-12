package musicdto

type MusicRequest struct {
	Title     string `json:"title" gorm:"type: varchar"`
	Year      string `json:"year" gorm:"type: varchar"`
	Thumbnail string `json:"thumbnail" gorm:"type: varchar"`
	Attache   string `json:"attache" gorm:"type: varchar"`
	ArtisID   int    `json:"artisId" gorm:"type: int"`
}
