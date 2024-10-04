package repository

import (
	"github.com/vanya-egorov/song_library/internal/entity"
	"gorm.io/gorm"
)

type SongRepository struct {
	db *gorm.DB
}

func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) GetAllSongs(limit, offset int, filters map[string]interface{}) ([]entity.Song, error) {
	var songs []entity.Song
	query := r.db.Limit(limit).Offset(offset)

	// Применение фильтров
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *SongRepository) AddSong(song *entity.Song) error {
	return r.db.Create(song).Error
}

func (r *SongRepository) UpdateSong(song *entity.Song) error {
	return r.db.Save(song).Error
}

func (r *SongRepository) DeleteSong(id uint) error {
	return r.db.Delete(&entity.Song{}, id).Error
}
