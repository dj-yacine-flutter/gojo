// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type AnimeEpisodeMeta struct {
	ID         int64
	EpisodeID  int64
	LanguageID int32
	MetaID     int64
	CreatedAt  time.Time
}

type AnimeEpisodeServer struct {
	ID        int64
	EpisodeID int64
	CreatedAt time.Time
}

type AnimeEpisodeServerDubTorrent struct {
	ID        int64
	ServerID  int64
	TorrentID int64
	CreatedAt time.Time
}

type AnimeEpisodeServerDubVideo struct {
	ID        int64
	ServerID  int64
	VideoID   int64
	CreatedAt time.Time
}

type AnimeEpisodeServerSubTorrent struct {
	ID        int64
	ServerID  int64
	TorrentID int64
	CreatedAt time.Time
}

type AnimeEpisodeServerSubVideo struct {
	ID        int64
	ServerID  int64
	VideoID   int64
	CreatedAt time.Time
}

type AnimeEpisodeTorrent struct {
	ID          int64
	FileName    string
	LanguageID  int32
	TorrentHash string
	TorrentFile string
	Seeds       int32
	Peers       int32
	Leechers    int32
	SizeBytes   int64
	Quality     string
	CreatedAt   time.Time
}

type AnimeEpisodeVideo struct {
	ID         int64
	LanguageID int32
	Authority  string
	Referer    string
	Link       string
	Quality    string
	CreatedAt  time.Time
}

type AnimeImage struct {
	ID              int64
	ImageHost       string
	ImageUrl        string
	ImageThumbnails string
	ImageBlurhash   string
	ImageHeight     int32
	ImageWidth      int32
	CreatedAt       time.Time
}

type AnimeLink struct {
	ID              int64
	OfficialWebsite string
	WikipediaUrl    string
	CrunchyrollUrl  string
	SocialMedia     []string
	CreatedAt       time.Time
}

type AnimeMovie struct {
	ID                int64
	OriginalTitle     string
	Aired             time.Time
	ReleaseYear       int32
	Rating            string
	Duration          time.Duration
	PortraitPoster    string
	PortraitBlurHash  string
	LandscapePoster   string
	LandscapeBlurHash string
	CreatedAt         time.Time
}

type AnimeMovieBackdropImage struct {
	ID        int64
	AnimeID   int64
	ImageID   int64
	CreatedAt time.Time
}

type AnimeMovieGenre struct {
	ID      int64
	AnimeID int64
	GenreID int32
}

type AnimeMovieLink struct {
	ID      int64
	AnimeID int64
	LinkID  int64
}

type AnimeMovieLogoImage struct {
	ID        int64
	AnimeID   int64
	ImageID   int64
	CreatedAt time.Time
}

type AnimeMovieMeta struct {
	ID         int64
	AnimeID    int64
	LanguageID int32
	MetaID     int64
}

type AnimeMovieOfficialTitle struct {
	ID        int64
	AnimeID   int64
	TitleText string
	CreatedAt time.Time
}

type AnimeMovieOtherTitle struct {
	ID        int64
	AnimeID   int64
	TitleText string
	CreatedAt time.Time
}

type AnimeMoviePosterImage struct {
	ID        int64
	AnimeID   int64
	ImageID   int64
	CreatedAt time.Time
}

type AnimeMovieResource struct {
	ID         int64
	AnimeID    int64
	ResourceID int64
}

type AnimeMovieServer struct {
	ID        int64
	AnimeID   int64
	CreatedAt time.Time
}

type AnimeMovieServerDubTorrent struct {
	ID        int64
	ServerID  int64
	TorrentID int64
	CreatedAt time.Time
}

type AnimeMovieServerDubVideo struct {
	ID        int64
	ServerID  int64
	VideoID   int64
	CreatedAt time.Time
}

type AnimeMovieServerSubTorrent struct {
	ID        int64
	ServerID  int64
	TorrentID int64
	CreatedAt time.Time
}

type AnimeMovieServerSubVideo struct {
	ID        int64
	ServerID  int64
	VideoID   int64
	CreatedAt time.Time
}

type AnimeMovieShortTitle struct {
	ID        int64
	AnimeID   int64
	TitleText string
	CreatedAt time.Time
}

type AnimeMovieStudio struct {
	ID       int64
	AnimeID  int64
	StudioID int32
}

type AnimeMovieTag struct {
	ID        int64
	AnimeID   int64
	TagID     int64
	CreatedAt time.Time
}

type AnimeMovieTorrent struct {
	ID          int64
	FileName    string
	LanguageID  int32
	TorrentHash string
	TorrentFile string
	Seeds       int32
	Peers       int32
	Leechers    int32
	SizeBytes   int64
	Quality     string
	CreatedAt   time.Time
}

type AnimeMovieTrailer struct {
	ID        int64
	AnimeID   int64
	TrailerID int64
	CreatedAt time.Time
}

type AnimeMovieTranslationTitle struct {
	ID        int64
	AnimeID   int64
	TitleText string
	CreatedAt time.Time
}

type AnimeMovieVideo struct {
	ID         int64
	LanguageID int32
	Authority  string
	Referer    string
	Link       string
	Quality    string
	CreatedAt  time.Time
}

type AnimeResource struct {
	ID            int64
	TvdbID        int32
	TmdbID        int32
	ImdbID        string
	LivechartID   int32
	AnimePlanetID string
	AnisearchID   int32
	AnidbID       int32
	KitsuID       int32
	MalID         int32
	NotifyMoeID   string
	AnilistID     int32
	CreatedAt     time.Time
}

type AnimeSeasonEpisode struct {
	ID        int64
	SeasonID  int64
	EpisodeID int64
	CreatedAt time.Time
}

type AnimeSeasonGenre struct {
	ID       int64
	SeasonID int64
	GenreID  int32
}

type AnimeSeasonMeta struct {
	ID         int64
	SeasonID   int64
	LanguageID int32
	MetaID     int64
	CreatedAt  time.Time
}

type AnimeSeasonOfficialTitle struct {
	ID        int64
	SeasonID  int64
	TitleText string
	CreatedAt time.Time
}

type AnimeSeasonOtherTitle struct {
	ID        int64
	SeasonID  int64
	TitleText string
	CreatedAt time.Time
}

type AnimeSeasonPosterImage struct {
	ID        int64
	SeasonID  int64
	ImageID   int64
	CreatedAt time.Time
}

type AnimeSeasonResource struct {
	ID         int64
	SeasonID   int64
	ResourceID int64
}

type AnimeSeasonShortTitle struct {
	ID        int64
	SeasonID  int64
	TitleText string
	CreatedAt time.Time
}

type AnimeSeasonStudio struct {
	ID       int64
	SeasonID int64
	StudioID int32
}

type AnimeSeasonTag struct {
	ID        int64
	SeasonID  int64
	TagID     int64
	CreatedAt time.Time
}

type AnimeSeasonTrailer struct {
	ID        int64
	SeasonID  int64
	TrailerID int64
	CreatedAt time.Time
}

type AnimeSeasonTranslationTitle struct {
	ID        int64
	SeasonID  int64
	TitleText string
	CreatedAt time.Time
}

type AnimeSerie struct {
	ID                int64
	OriginalTitle     string
	FirstYear         int32
	LastYear          int32
	MalID             int32
	TvdbID            int32
	TmdbID            int32
	PortraitPoster    string
	PortraitBlurHash  string
	LandscapePoster   string
	LandscapeBlurHash string
	CreatedAt         time.Time
}

type AnimeSerieBackdropImage struct {
	ID        int64
	AnimeID   int64
	ImageID   int64
	CreatedAt time.Time
}

type AnimeSerieEpisode struct {
	ID                   int64
	SeasonID             int64
	EpisodeNumber        int32
	EpisodeOriginalTitle string
	Aired                time.Time
	Rating               string
	Duration             time.Duration
	Thumbnails           string
	ThumbnailsBlurHash   string
	CreatedAt            time.Time
}

type AnimeSerieLink struct {
	ID      int64
	AnimeID int64
	LinkID  int64
}

type AnimeSerieLogoImage struct {
	ID        int64
	AnimeID   int64
	ImageID   int64
	CreatedAt time.Time
}

type AnimeSerieMeta struct {
	ID         int64
	AnimeID    int64
	LanguageID int32
	MetaID     int64
}

type AnimeSeriePosterImage struct {
	ID        int64
	AnimeID   int64
	ImageID   int64
	CreatedAt time.Time
}

type AnimeSerieSeason struct {
	ID                  int64
	AnimeID             int64
	SeasonOriginalTitle string
	ReleaseYear         int32
	Aired               time.Time
	PortraitPoster      string
	PortraitBlurHash    string
	Rating              string
	CreatedAt           time.Time
}

type AnimeSerieTrailer struct {
	ID        int64
	AnimeID   int64
	TrailerID int64
	CreatedAt time.Time
}

type AnimeTag struct {
	ID        int64
	Tag       string
	CreatedAt time.Time
}

type AnimeTrailer struct {
	ID         int64
	IsOfficial bool
	HostName   string
	HostKey    string
	CreatedAt  time.Time
}

type Genre struct {
	ID        int32
	GenreName string
	CreatedAt time.Time
}

type Language struct {
	ID           int32
	LanguageCode string
	LanguageName string
	CreatedAt    time.Time
}

type Meta struct {
	ID        int64
	Title     string
	Overview  string
	CreatedAt time.Time
}

type Session struct {
	ID           uuid.UUID
	Username     string
	RefreshToken string
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
}

type Studio struct {
	ID         int32
	StudioName string
	CreatedAt  time.Time
}

type User struct {
	ID                int64
	Username          string
	Email             string
	HashedPassword    string
	FullName          string
	PasswordChangedAt time.Time
	CreatedAt         time.Time
	IsEmailVerified   bool
	Role              string
}

type VerifyEmail struct {
	ID         int64
	Username   string
	Email      string
	SecretCode string
	IsUsed     bool
	CreatedAt  time.Time
	ExpiredAt  time.Time
}
