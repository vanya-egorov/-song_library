package postgres

import (
	"github.com/vanya-egorov/song_library/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Song{})
}
