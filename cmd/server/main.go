package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vanya-egorov/song_library/internal/entity"
	"github.com/vanya-egorov/song_library/internal/usecase/song"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Строка подключения к базе данных
	dsn := "host=localhost user=postgres password=yourpassword dbname=song_library port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Миграции
	err = db.AutoMigrate(&entity.Song{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Создание репозитория, use-case и хендлера
	songRepo := postgres.NewSongRepository(db)
	songUC := song.NewSongUseCase(songRepo)
	songHandler := http.NewSongHandler(songUC)

	// Инициализация роутера
	r := mux.NewRouter()

	// Маршрут для получения песен
	r.HandleFunc("/songs", songHandler.GetSongs).Methods("GET")

	// Запуск сервера
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
