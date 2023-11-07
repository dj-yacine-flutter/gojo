// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type AnimeMovie struct {
	ID                int64         `json:"id"`
	OriginalTitle     string        `json:"original_title"`
	Aired             time.Time     `json:"aired"`
	ReleaseYear       int32         `json:"release_year"`
	Rating            string        `json:"rating"`
	Duration          time.Duration `json:"duration"`
	PortriatPoster    string        `json:"portriat_poster"`
	PortriatBlurHash  string        `json:"portriat_blur_hash"`
	LandscapePoster   string        `json:"landscape_poster"`
	LandscapeBlurHash string        `json:"landscape_blur_hash"`
	CreatedAt         time.Time     `json:"created_at"`
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

type AnimeMovieResource struct {
	ID         int64 `json:"id"`
	AnimeID    int64 `json:"anime_id"`
	ResourceID int64 `json:"resource_id"`
}

type AnimeMovieServer struct {
	ID        int64     `json:"id"`
	AnimeID   int64     `json:"anime_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeMovieServerDubVideo struct {
	ID        int64     `json:"id"`
	ServerID  int64     `json:"server_id"`
	VideoID   int64     `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeMovieServerSubVideo struct {
	ID        int64     `json:"id"`
	ServerID  int64     `json:"server_id"`
	VideoID   int64     `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeMovieServerTorrent struct {
	ID        int64     `json:"id"`
	ServerID  int64     `json:"server_id"`
	TorrentID int64     `json:"torrent_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeMovieStudio struct {
	ID       int64 `json:"id"`
	AnimeID  int64 `json:"anime_id"`
	StudioID int32 `json:"studio_id"`
}

type AnimeMovieTorrent struct {
	ID          int64     `json:"id"`
	FileName    string    `json:"file_name"`
	LanguageID  int32     `json:"language_id"`
	TorrentHash string    `json:"torrent_hash"`
	TorrentFile string    `json:"torrent_file"`
	Seeds       int32     `json:"seeds"`
	Peers       int32     `json:"peers"`
	Leechers    int32     `json:"leechers"`
	SizeBytes   int64     `json:"size_bytes"`
	Quality     string    `json:"quality"`
	CreatedAt   time.Time `json:"created_at"`
}

type AnimeMovieVideo struct {
	ID         int64     `json:"id"`
	LanguageID int32     `json:"language_id"`
	Autority   string    `json:"autority"`
	Referer    string    `json:"referer"`
	Link       string    `json:"link"`
	Quality    string    `json:"quality"`
	CreatedAt  time.Time `json:"created_at"`
}

type AnimeResource struct {
	ID              int64     `json:"id"`
	TmdbID          int32     `json:"tmdb_id"`
	ImdbID          string    `json:"imdb_id"`
	OfficialWebsite string    `json:"official_website"`
	WikipediaUrl    string    `json:"wikipedia_url"`
	CrunchyrollUrl  string    `json:"crunchyroll_url"`
	SocialMedia     []string  `json:"social_media"`
	CreatedAt       time.Time `json:"created_at"`
}

type AnimeSerie struct {
	ID                int64     `json:"id"`
	OriginalTitle     string    `json:"original_title"`
	Aired             time.Time `json:"aired"`
	ReleaseYear       int32     `json:"release_year"`
	Rating            string    `json:"rating"`
	PortriatPoster    string    `json:"portriat_poster"`
	PortriatBlurHash  string    `json:"portriat_blur_hash"`
	LandscapePoster   string    `json:"landscape_poster"`
	LandscapeBlurHash string    `json:"landscape_blur_hash"`
	CreatedAt         time.Time `json:"created_at"`
}

type AnimeSerieEpisode struct {
	ID            int64     `json:"id"`
	EpisodeNumber int32     `json:"episode_number"`
	SeasonID      int64     `json:"season_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type AnimeSerieEpisodeMeta struct {
	ID        int64     `json:"id"`
	EpisodeID int64     `json:"episode_id"`
	MetaID    int64     `json:"meta_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeSerieEpisodeServer struct {
	ID        int64     `json:"id"`
	EpisodeID int64     `json:"episode_id"`
	ServerID  int64     `json:"server_id"`
	CreatedAt time.Time `json:"created_at"`
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

type AnimeSerieResource struct {
	ID         int64 `json:"id"`
	AnimeID    int64 `json:"anime_id"`
	ResourceID int64 `json:"resource_id"`
}

type AnimeSerieSeason struct {
	ID                int64     `json:"id"`
	AnimeID           int64     `json:"anime_id"`
	SeasonNumber      int32     `json:"season_number"`
	PortriatPoster    string    `json:"portriat_poster"`
	PortriatBlurHash  string    `json:"portriat_blur_hash"`
	LandscapePoster   string    `json:"landscape_poster"`
	LandscapeBlurHash string    `json:"landscape_blur_hash"`
	CreatedAt         time.Time `json:"created_at"`
}

type AnimeSerieSeasonEpisode struct {
	ID        int64     `json:"id"`
	SeasonID  int64     `json:"season_id"`
	EpisodeID int64     `json:"episode_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeSerieSeasonMeta struct {
	ID        int64     `json:"id"`
	SeasonID  int64     `json:"season_id"`
	MetaID    int64     `json:"meta_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeSerieServer struct {
	ID        int64     `json:"id"`
	EpisodeID int64     `json:"episode_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeSerieServerDubVideo struct {
	ID        int64     `json:"id"`
	ServerID  int64     `json:"server_id"`
	VideoID   int64     `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeSerieServerSubVideo struct {
	ID        int64     `json:"id"`
	ServerID  int64     `json:"server_id"`
	VideoID   int64     `json:"video_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeSerieServerTorrent struct {
	ID        int64     `json:"id"`
	ServerID  int64     `json:"server_id"`
	TorrentID int64     `json:"torrent_id"`
	CreatedAt time.Time `json:"created_at"`
}

type AnimeSerieStudio struct {
	ID       int64 `json:"id"`
	AnimeID  int64 `json:"anime_id"`
	StudioID int32 `json:"studio_id"`
}

type AnimeSerieTorrent struct {
	ID          int64     `json:"id"`
	FileName    string    `json:"file_name"`
	LanguageID  int32     `json:"language_id"`
	TorrentHash string    `json:"torrent_hash"`
	TorrentFile string    `json:"torrent_file"`
	Seeds       int32     `json:"seeds"`
	Peers       int32     `json:"peers"`
	Leechers    int32     `json:"leechers"`
	SizeBytes   int64     `json:"size_bytes"`
	Quality     string    `json:"quality"`
	CreatedAt   time.Time `json:"created_at"`
}

type AnimeSerieVideo struct {
	ID         int64     `json:"id"`
	LanguageID int32     `json:"language_id"`
	Autority   string    `json:"autority"`
	Referer    string    `json:"referer"`
	Link       string    `json:"link"`
	Quality    string    `json:"quality"`
	CreatedAt  time.Time `json:"created_at"`
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
