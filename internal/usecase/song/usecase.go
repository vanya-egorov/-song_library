package song

import (
	"github.com/vanya-egorov/song_library/internal/entity"
	"github.com/vanya-egorov/song_library/internal/repository"
)

type UseCase interface {
	GetSongs(limit, offset int, filters map[string]interface{}) ([]entity.Song, error)
}

type SongUseCase struct {
	repo repository.SongRepository
}

func NewSongUseCase(repo repository.SongRepository) *SongUseCase {
	return &SongUseCase{repo: repo}
}

func (uc *SongUseCase) GetSongs(limit, offset int, filters map[string]interface{}) ([]entity.Song, error) {
	return uc.repo.GetAllSongs(limit, offset, filters)
}
