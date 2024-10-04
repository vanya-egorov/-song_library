package entity

type Song struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Group  string `json:"group"`
	Title  string `json:"song"`
	Lyrics string `json:"lyrics"`
}
