package song_http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/vanya-egorov/song_library/internal/usecase/song"
)

type SongHandler struct {
	useCase song.UseCase // Использование интерфейса вместо конкретной структуры
}

func NewSongHandler(uc song.UseCase) *SongHandler {
	return &SongHandler{useCase: uc}
}

func (h *SongHandler) GetSongs(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10 // значение по умолчанию
	}

	offset, err := strconv.Atoi(query.Get("offset"))
	if err != nil {
		offset = 0 // значение по умолчанию
	}

	filters := map[string]interface{}{}
	if group := query.Get("group"); group != "" {
		filters["group"] = group
	}
	if title := query.Get("title"); title != "" {
		filters["title"] = title
	}

	songs, err := h.useCase.GetSongs(limit, offset, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error retrieving songs: %v", err)
		return
	}

	if len(songs) == 0 {
		http.Error(w, "No songs found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
}
