// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type AnimeMovie struct {
	ID            int64         `json:"id"`
	OriginalTitle string        `json:"original_title"`
	Aired         time.Time     `json:"aired"`
	ReleaseYear   int32         `json:"release_year"`
	Duration      time.Duration `json:"duration"`
	CreatedAt     time.Time     `json:"created_at"`
}

type AnimeMovieGenre struct {
	ID      int64 `json:"id"`
	AnimeID int64 `json:"anime_id"`
	GenreID int32 `json:"genre_id"`
}

type AnimeMovieMeta struct {
	ID         int64 `json:"id"`
	AnimeID    int64 `json:"anime_id"`
	LanguageID int32 `json:"language_id"`
	MetaID     int64 `json:"meta_id"`
}

type AnimeMovieStudio struct {
	ID       int64 `json:"id"`
	AnimeID  int64 `json:"anime_id"`
	StudioID int32 `json:"studio_id"`
}

type AnimeSerie struct {
	ID            int64         `json:"id"`
	OriginalTitle string        `json:"original_title"`
	Aired         time.Time     `json:"aired"`
	ReleaseYear   int32         `json:"release_year"`
	Duration      time.Duration `json:"duration"`
	CreatedAt     time.Time     `json:"created_at"`
}

type AnimeSerieGenre struct {
	ID      int64 `json:"id"`
	AnimeID int64 `json:"anime_id"`
	GenreID int32 `json:"genre_id"`
}

type AnimeSerieMeta struct {
	ID         int64 `json:"id"`
	AnimeID    int64 `json:"anime_id"`
	LanguageID int32 `json:"language_id"`
	MetaID     int64 `json:"meta_id"`
}

type AnimeSerieStudio struct {
	ID       int64 `json:"id"`
	AnimeID  int64 `json:"anime_id"`
	StudioID int32 `json:"studio_id"`
}

type Genre struct {
	ID        int32     `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
}

type Language struct {
	ID           int32     `json:"id"`
	LanguageCode string    `json:"language_code"`
	LanguageName string    `json:"language_name"`
	CreatedAt    time.Time `json:"created_at"`
}

type Meta struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Overview  string    `json:"overview"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Studio struct {
	ID         int32     `json:"id"`
	StudioName string    `json:"studio_name"`
	CreatedAt  time.Time `json:"created_at"`
}

type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	IsEmailVerified   bool      `json:"is_email_verified"`
	Role              string    `json:"role"`
}

type VerifyEmail struct {
	ID         int64     `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	CreatedAt  time.Time `json:"created_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}
