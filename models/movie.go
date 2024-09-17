package models

type Movie struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Title      string   `json:"title"`
	Rating     int      `json:"rating"`
	DirectorID uint     `json:"director_id"`
	Director   Director `gorm:"foreignKey:DirectorID" json:"director"`
}
