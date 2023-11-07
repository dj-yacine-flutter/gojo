// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAnimeMovie(ctx context.Context, arg CreateAnimeMovieParams) (AnimeMovie, error)
	CreateAnimeMovieGenre(ctx context.Context, arg CreateAnimeMovieGenreParams) (AnimeMovieGenre, error)
	CreateAnimeMovieMeta(ctx context.Context, arg CreateAnimeMovieMetaParams) (AnimeMovieMeta, error)
	CreateAnimeMovieResource(ctx context.Context, arg CreateAnimeMovieResourceParams) (AnimeMovieResource, error)
	CreateAnimeMovieServer(ctx context.Context, animeID int64) (AnimeMovieServer, error)
	CreateAnimeMovieServerDubVideo(ctx context.Context, arg CreateAnimeMovieServerDubVideoParams) (AnimeMovieServerDubVideo, error)
	CreateAnimeMovieServerSubVideo(ctx context.Context, arg CreateAnimeMovieServerSubVideoParams) (AnimeMovieServerSubVideo, error)
	CreateAnimeMovieServerTorrent(ctx context.Context, arg CreateAnimeMovieServerTorrentParams) (AnimeMovieServerTorrent, error)
	CreateAnimeMovieStudio(ctx context.Context, arg CreateAnimeMovieStudioParams) (AnimeMovieStudio, error)
	CreateAnimeMovieTorrent(ctx context.Context, arg CreateAnimeMovieTorrentParams) (AnimeMovieTorrent, error)
	CreateAnimeMovieVideo(ctx context.Context, arg CreateAnimeMovieVideoParams) (AnimeMovieVideo, error)
	CreateAnimeResource(ctx context.Context, arg CreateAnimeResourceParams) (AnimeResource, error)
	CreateAnimeSerie(ctx context.Context, arg CreateAnimeSerieParams) (AnimeSerie, error)
	CreateAnimeSerieEpisode(ctx context.Context, arg CreateAnimeSerieEpisodeParams) (AnimeSerieEpisode, error)
	CreateAnimeSerieEpisodeMeta(ctx context.Context, arg CreateAnimeSerieEpisodeMetaParams) (AnimeSerieEpisodeMeta, error)
	CreateAnimeSerieEpisodeServer(ctx context.Context, arg CreateAnimeSerieEpisodeServerParams) (AnimeSerieEpisodeServer, error)
	CreateAnimeSerieGenre(ctx context.Context, arg CreateAnimeSerieGenreParams) (AnimeSerieGenre, error)
	CreateAnimeSerieMeta(ctx context.Context, arg CreateAnimeSerieMetaParams) (AnimeSerieMeta, error)
	CreateAnimeSerieResource(ctx context.Context, arg CreateAnimeSerieResourceParams) (AnimeSerieResource, error)
	CreateAnimeSerieSeason(ctx context.Context, arg CreateAnimeSerieSeasonParams) (AnimeSerieSeason, error)
	CreateAnimeSerieSeasonEpisode(ctx context.Context, arg CreateAnimeSerieSeasonEpisodeParams) (AnimeSerieSeasonEpisode, error)
	CreateAnimeSerieSeasonMeta(ctx context.Context, arg CreateAnimeSerieSeasonMetaParams) (AnimeSerieSeasonMeta, error)
	CreateAnimeSerieServer(ctx context.Context, episodeID int64) (AnimeSerieServer, error)
	CreateAnimeSerieServerDubVideo(ctx context.Context, arg CreateAnimeSerieServerDubVideoParams) (AnimeSerieServerDubVideo, error)
	CreateAnimeSerieServerSubVideo(ctx context.Context, arg CreateAnimeSerieServerSubVideoParams) (AnimeSerieServerSubVideo, error)
	CreateAnimeSerieServerTorrent(ctx context.Context, arg CreateAnimeSerieServerTorrentParams) (AnimeSerieServerTorrent, error)
	CreateAnimeSerieStudio(ctx context.Context, arg CreateAnimeSerieStudioParams) (AnimeSerieStudio, error)
	CreateAnimeSerieTorrent(ctx context.Context, arg CreateAnimeSerieTorrentParams) (AnimeSerieTorrent, error)
	CreateAnimeSerieVideo(ctx context.Context, arg CreateAnimeSerieVideoParams) (AnimeSerieVideo, error)
	CreateGenre(ctx context.Context, genreName string) (Genre, error)
	CreateLanguage(ctx context.Context, arg CreateLanguageParams) (Language, error)
	CreateMeta(ctx context.Context, arg CreateMetaParams) (Meta, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateStudio(ctx context.Context, studioName string) (Studio, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	DeleteAnimeMovie(ctx context.Context, id int64) error
	DeleteAnimeMovieGenre(ctx context.Context, arg DeleteAnimeMovieGenreParams) error
	DeleteAnimeMovieMeta(ctx context.Context, arg DeleteAnimeMovieMetaParams) error
	DeleteAnimeMovieResource(ctx context.Context, arg DeleteAnimeMovieResourceParams) error
	DeleteAnimeMovieServer(ctx context.Context, id int64) error
	DeleteAnimeMovieServerDubVideo(ctx context.Context, id int64) error
	DeleteAnimeMovieServerSubVideo(ctx context.Context, id int64) error
	DeleteAnimeMovieServerTorrent(ctx context.Context, id int64) error
	DeleteAnimeMovieStudio(ctx context.Context, arg DeleteAnimeMovieStudioParams) error
	DeleteAnimeMovieTorrent(ctx context.Context, id int64) error
	DeleteAnimeMovieVideo(ctx context.Context, id int64) error
	DeleteAnimeResource(ctx context.Context, id int64) error
	DeleteAnimeSerie(ctx context.Context, id int64) error
	DeleteAnimeSerieEpisode(ctx context.Context, id int64) error
	DeleteAnimeSerieEpisodeMeta(ctx context.Context, id int64) error
	DeleteAnimeSerieEpisodeServer(ctx context.Context, id int64) error
	DeleteAnimeSerieGenre(ctx context.Context, arg DeleteAnimeSerieGenreParams) error
	DeleteAnimeSerieMeta(ctx context.Context, arg DeleteAnimeSerieMetaParams) error
	DeleteAnimeSerieResource(ctx context.Context, arg DeleteAnimeSerieResourceParams) error
	DeleteAnimeSerieSeason(ctx context.Context, id int64) error
	DeleteAnimeSerieSeasonEpisode(ctx context.Context, id int64) error
	DeleteAnimeSerieSeasonMeta(ctx context.Context, id int64) error
	DeleteAnimeSerieServer(ctx context.Context, id int64) error
	DeleteAnimeSerieServerDubVideo(ctx context.Context, id int64) error
	DeleteAnimeSerieServerSubVideo(ctx context.Context, id int64) error
	DeleteAnimeSerieServerTorrent(ctx context.Context, id int64) error
	DeleteAnimeSerieStudio(ctx context.Context, arg DeleteAnimeSerieStudioParams) error
	DeleteAnimeSerieTorrent(ctx context.Context, id int64) error
	DeleteAnimeSerieVideo(ctx context.Context, id int64) error
	DeleteGenre(ctx context.Context, id int32) error
	DeleteLanguage(ctx context.Context, id int32) error
	DeleteMeta(ctx context.Context, id int64) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	DeleteStudio(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int64) error
	GetAnimeMovie(ctx context.Context, id int64) (AnimeMovie, error)
	GetAnimeMovieGenre(ctx context.Context, id int64) (AnimeMovieGenre, error)
	GetAnimeMovieMeta(ctx context.Context, arg GetAnimeMovieMetaParams) (int64, error)
	GetAnimeMovieResource(ctx context.Context, id int64) (AnimeMovieResource, error)
	GetAnimeMovieResourceByAnimeID(ctx context.Context, animeID int64) (AnimeMovieResource, error)
	GetAnimeMovieServer(ctx context.Context, id int64) (AnimeMovieServer, error)
	GetAnimeMovieServerDubVideo(ctx context.Context, id int64) (AnimeMovieServerDubVideo, error)
	GetAnimeMovieServerSubVideo(ctx context.Context, id int64) (AnimeMovieServerSubVideo, error)
	GetAnimeMovieServerTorrent(ctx context.Context, id int64) (AnimeMovieServerTorrent, error)
	GetAnimeMovieStudio(ctx context.Context, id int64) (AnimeMovieStudio, error)
	GetAnimeMovieTorrent(ctx context.Context, id int64) (AnimeMovieTorrent, error)
	GetAnimeMovieVideo(ctx context.Context, id int64) (AnimeMovieVideo, error)
	GetAnimeResource(ctx context.Context, id int64) (AnimeResource, error)
	GetAnimeSerie(ctx context.Context, id int64) (AnimeSerie, error)
	GetAnimeSerieEpisode(ctx context.Context, id int64) (AnimeSerieEpisode, error)
	GetAnimeSerieEpisodeMeta(ctx context.Context, id int64) (AnimeSerieEpisodeMeta, error)
	GetAnimeSerieEpisodeServer(ctx context.Context, id int64) (AnimeSerieEpisodeServer, error)
	GetAnimeSerieGenre(ctx context.Context, id int64) (AnimeSerieGenre, error)
	GetAnimeSerieMeta(ctx context.Context, arg GetAnimeSerieMetaParams) (int64, error)
	GetAnimeSerieResource(ctx context.Context, id int64) (AnimeSerieResource, error)
	GetAnimeSerieResourceByAnimeID(ctx context.Context, animeID int64) (AnimeSerieResource, error)
	GetAnimeSerieSeason(ctx context.Context, id int64) (AnimeSerieSeason, error)
	GetAnimeSerieSeasonEpisode(ctx context.Context, id int64) (AnimeSerieSeasonEpisode, error)
	GetAnimeSerieSeasonMeta(ctx context.Context, id int64) (AnimeSerieSeasonMeta, error)
	GetAnimeSerieServer(ctx context.Context, id int64) (AnimeSerieServer, error)
	GetAnimeSerieServerDubVideo(ctx context.Context, id int64) (AnimeSerieServerDubVideo, error)
	GetAnimeSerieServerSubVideo(ctx context.Context, id int64) (AnimeSerieServerSubVideo, error)
	GetAnimeSerieServerTorrent(ctx context.Context, id int64) (AnimeSerieServerTorrent, error)
	GetAnimeSerieStudio(ctx context.Context, id int64) (AnimeSerieStudio, error)
	GetAnimeSerieTorrent(ctx context.Context, id int64) (AnimeSerieTorrent, error)
	GetAnimeSerieVideo(ctx context.Context, id int64) (AnimeSerieVideo, error)
	GetGenre(ctx context.Context, id int32) (Genre, error)
	GetLanguage(ctx context.Context, id int32) (Language, error)
	GetMeta(ctx context.Context, id int64) (Meta, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetStudio(ctx context.Context, id int32) (Studio, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListAnimeMovieGenres(ctx context.Context, arg ListAnimeMovieGenresParams) ([]int32, error)
	ListAnimeMovieMetas(ctx context.Context, arg ListAnimeMovieMetasParams) ([]int64, error)
	ListAnimeMovieResources(ctx context.Context, arg ListAnimeMovieResourcesParams) ([]int64, error)
	ListAnimeMovieServerDubVideos(ctx context.Context, arg ListAnimeMovieServerDubVideosParams) ([]AnimeMovieServerDubVideo, error)
	ListAnimeMovieServerSubVideos(ctx context.Context, arg ListAnimeMovieServerSubVideosParams) ([]AnimeMovieServerSubVideo, error)
	ListAnimeMovieServerTorrents(ctx context.Context, arg ListAnimeMovieServerTorrentsParams) ([]AnimeMovieServerTorrent, error)
	ListAnimeMovieServers(ctx context.Context, arg ListAnimeMovieServersParams) ([]AnimeMovieServer, error)
	ListAnimeMovieStudios(ctx context.Context, arg ListAnimeMovieStudiosParams) ([]int32, error)
	ListAnimeMovieTorrents(ctx context.Context, arg ListAnimeMovieTorrentsParams) ([]AnimeMovieTorrent, error)
	ListAnimeMovieVideos(ctx context.Context, arg ListAnimeMovieVideosParams) ([]AnimeMovieVideo, error)
	ListAnimeMovies(ctx context.Context, arg ListAnimeMoviesParams) ([]AnimeMovie, error)
	ListAnimeSerieEpisodeMetasByEpisode(ctx context.Context, arg ListAnimeSerieEpisodeMetasByEpisodeParams) ([]AnimeSerieEpisodeMeta, error)
	ListAnimeSerieEpisodeServersByEpisode(ctx context.Context, arg ListAnimeSerieEpisodeServersByEpisodeParams) ([]AnimeSerieEpisodeServer, error)
	ListAnimeSerieEpisodesBySeason(ctx context.Context, arg ListAnimeSerieEpisodesBySeasonParams) ([]AnimeSerieEpisode, error)
	ListAnimeSerieGenres(ctx context.Context, arg ListAnimeSerieGenresParams) ([]int32, error)
	ListAnimeSerieMetas(ctx context.Context, arg ListAnimeSerieMetasParams) ([]int64, error)
	ListAnimeSerieResources(ctx context.Context, arg ListAnimeSerieResourcesParams) ([]int64, error)
	ListAnimeSerieSeasonEpisodesBySeason(ctx context.Context, arg ListAnimeSerieSeasonEpisodesBySeasonParams) ([]AnimeSerieSeasonEpisode, error)
	ListAnimeSerieSeasonMetasBySeason(ctx context.Context, arg ListAnimeSerieSeasonMetasBySeasonParams) ([]AnimeSerieSeasonMeta, error)
	ListAnimeSerieSeasonsByAnime(ctx context.Context, arg ListAnimeSerieSeasonsByAnimeParams) ([]AnimeSerieSeason, error)
	ListAnimeSerieServerDubVideos(ctx context.Context, arg ListAnimeSerieServerDubVideosParams) ([]AnimeSerieServerDubVideo, error)
	ListAnimeSerieServerSubVideos(ctx context.Context, arg ListAnimeSerieServerSubVideosParams) ([]AnimeSerieServerSubVideo, error)
	ListAnimeSerieServerTorrents(ctx context.Context, arg ListAnimeSerieServerTorrentsParams) ([]AnimeSerieServerTorrent, error)
	ListAnimeSerieServers(ctx context.Context, arg ListAnimeSerieServersParams) ([]AnimeSerieServer, error)
	ListAnimeSerieStudios(ctx context.Context, arg ListAnimeSerieStudiosParams) ([]int32, error)
	ListAnimeSerieTorrents(ctx context.Context, arg ListAnimeSerieTorrentsParams) ([]AnimeSerieTorrent, error)
	ListAnimeSerieVideos(ctx context.Context, arg ListAnimeSerieVideosParams) ([]AnimeSerieVideo, error)
	ListAnimeSeries(ctx context.Context, arg ListAnimeSeriesParams) ([]AnimeSerie, error)
	ListGenres(ctx context.Context, arg ListGenresParams) ([]Genre, error)
	ListLanguages(ctx context.Context, arg ListLanguagesParams) ([]Language, error)
	ListStudios(ctx context.Context, arg ListStudiosParams) ([]Studio, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	RefreshSessions(ctx context.Context, username string) error
	UpdateAnimeMovie(ctx context.Context, arg UpdateAnimeMovieParams) (AnimeMovie, error)
	UpdateAnimeMovieMeta(ctx context.Context, arg UpdateAnimeMovieMetaParams) (AnimeMovieMeta, error)
	UpdateAnimeMovieServer(ctx context.Context, arg UpdateAnimeMovieServerParams) (AnimeMovieServer, error)
	UpdateAnimeMovieServerDubVideo(ctx context.Context, arg UpdateAnimeMovieServerDubVideoParams) (AnimeMovieServerDubVideo, error)
	UpdateAnimeMovieServerSubVideo(ctx context.Context, arg UpdateAnimeMovieServerSubVideoParams) (AnimeMovieServerSubVideo, error)
	UpdateAnimeMovieServerTorrent(ctx context.Context, arg UpdateAnimeMovieServerTorrentParams) (AnimeMovieServerTorrent, error)
	UpdateAnimeMovieTorrent(ctx context.Context, arg UpdateAnimeMovieTorrentParams) (AnimeMovieTorrent, error)
	UpdateAnimeMovieVideo(ctx context.Context, arg UpdateAnimeMovieVideoParams) (AnimeMovieVideo, error)
	UpdateAnimeResource(ctx context.Context, arg UpdateAnimeResourceParams) (AnimeResource, error)
	UpdateAnimeSerie(ctx context.Context, arg UpdateAnimeSerieParams) (AnimeSerie, error)
	UpdateAnimeSerieEpisode(ctx context.Context, arg UpdateAnimeSerieEpisodeParams) (AnimeSerieEpisode, error)
	UpdateAnimeSerieEpisodeMeta(ctx context.Context, arg UpdateAnimeSerieEpisodeMetaParams) (AnimeSerieEpisodeMeta, error)
	UpdateAnimeSerieEpisodeServer(ctx context.Context, arg UpdateAnimeSerieEpisodeServerParams) (AnimeSerieEpisodeServer, error)
	UpdateAnimeSerieMeta(ctx context.Context, arg UpdateAnimeSerieMetaParams) (AnimeSerieMeta, error)
	UpdateAnimeSerieSeason(ctx context.Context, arg UpdateAnimeSerieSeasonParams) (AnimeSerieSeason, error)
	UpdateAnimeSerieSeasonEpisode(ctx context.Context, arg UpdateAnimeSerieSeasonEpisodeParams) (AnimeSerieSeasonEpisode, error)
	UpdateAnimeSerieSeasonMeta(ctx context.Context, arg UpdateAnimeSerieSeasonMetaParams) (AnimeSerieSeasonMeta, error)
	UpdateAnimeSerieServer(ctx context.Context, arg UpdateAnimeSerieServerParams) (AnimeSerieServer, error)
	UpdateAnimeSerieServerDubVideo(ctx context.Context, arg UpdateAnimeSerieServerDubVideoParams) (AnimeSerieServerDubVideo, error)
	UpdateAnimeSerieServerSubVideo(ctx context.Context, arg UpdateAnimeSerieServerSubVideoParams) (AnimeSerieServerSubVideo, error)
	UpdateAnimeSerieServerTorrent(ctx context.Context, arg UpdateAnimeSerieServerTorrentParams) (AnimeSerieServerTorrent, error)
	UpdateAnimeSerieTorrent(ctx context.Context, arg UpdateAnimeSerieTorrentParams) (AnimeSerieTorrent, error)
	UpdateAnimeSerieVideo(ctx context.Context, arg UpdateAnimeSerieVideoParams) (AnimeSerieVideo, error)
	UpdateGenre(ctx context.Context, arg UpdateGenreParams) (Genre, error)
	UpdateLanguage(ctx context.Context, arg UpdateLanguageParams) (Language, error)
	UpdateMeta(ctx context.Context, arg UpdateMetaParams) (Meta, error)
	UpdateSession(ctx context.Context, arg UpdateSessionParams) (string, error)
	UpdateStudio(ctx context.Context, arg UpdateStudioParams) (Studio, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
