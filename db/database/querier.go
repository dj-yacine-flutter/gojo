// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAnimeImage(ctx context.Context, arg CreateAnimeImageParams) (AnimeImage, error)
	CreateAnimeLink(ctx context.Context, arg CreateAnimeLinkParams) (AnimeLink, error)
	CreateAnimeMovie(ctx context.Context, arg CreateAnimeMovieParams) (AnimeMovie, error)
	CreateAnimeMovieBackdropImage(ctx context.Context, arg CreateAnimeMovieBackdropImageParams) (AnimeMovieBackdropImage, error)
	CreateAnimeMovieGenre(ctx context.Context, arg CreateAnimeMovieGenreParams) (AnimeMovieGenre, error)
	CreateAnimeMovieLink(ctx context.Context, arg CreateAnimeMovieLinkParams) (AnimeMovieLink, error)
	CreateAnimeMovieLogoImage(ctx context.Context, arg CreateAnimeMovieLogoImageParams) (AnimeMovieLogoImage, error)
	CreateAnimeMovieMeta(ctx context.Context, arg CreateAnimeMovieMetaParams) (AnimeMovieMeta, error)
	CreateAnimeMoviePosterImage(ctx context.Context, arg CreateAnimeMoviePosterImageParams) (AnimeMoviePosterImage, error)
	CreateAnimeMovieResource(ctx context.Context, arg CreateAnimeMovieResourceParams) (AnimeMovieResource, error)
	CreateAnimeMovieServer(ctx context.Context, animeID int64) (AnimeMovieServer, error)
	CreateAnimeMovieServerDubTorrent(ctx context.Context, arg CreateAnimeMovieServerDubTorrentParams) (AnimeMovieServerDubTorrent, error)
	CreateAnimeMovieServerDubVideo(ctx context.Context, arg CreateAnimeMovieServerDubVideoParams) (AnimeMovieServerDubVideo, error)
	CreateAnimeMovieServerSubTorrent(ctx context.Context, arg CreateAnimeMovieServerSubTorrentParams) (AnimeMovieServerSubTorrent, error)
	CreateAnimeMovieServerSubVideo(ctx context.Context, arg CreateAnimeMovieServerSubVideoParams) (AnimeMovieServerSubVideo, error)
	CreateAnimeMovieStudio(ctx context.Context, arg CreateAnimeMovieStudioParams) (AnimeMovieStudio, error)
	CreateAnimeMovieTorrent(ctx context.Context, arg CreateAnimeMovieTorrentParams) (AnimeMovieTorrent, error)
	CreateAnimeMovieVideo(ctx context.Context, arg CreateAnimeMovieVideoParams) (AnimeMovieVideo, error)
	CreateAnimeResource(ctx context.Context, arg CreateAnimeResourceParams) (AnimeResource, error)
	CreateAnimeSeasonResource(ctx context.Context, arg CreateAnimeSeasonResourceParams) (AnimeSeasonResource, error)
	CreateAnimeSerie(ctx context.Context, arg CreateAnimeSerieParams) (AnimeSerie, error)
	CreateAnimeSerieBackdropImage(ctx context.Context, arg CreateAnimeSerieBackdropImageParams) (AnimeSerieBackdropImage, error)
	CreateAnimeSerieEpisode(ctx context.Context, arg CreateAnimeSerieEpisodeParams) (AnimeSerieEpisode, error)
	CreateAnimeSerieEpisodeMeta(ctx context.Context, arg CreateAnimeSerieEpisodeMetaParams) (AnimeSerieEpisodeMeta, error)
	CreateAnimeSerieEpisodeServer(ctx context.Context, arg CreateAnimeSerieEpisodeServerParams) (AnimeSerieEpisodeServer, error)
	CreateAnimeSerieGenre(ctx context.Context, arg CreateAnimeSerieGenreParams) (AnimeSerieGenre, error)
	CreateAnimeSerieLink(ctx context.Context, arg CreateAnimeSerieLinkParams) (AnimeSerieLink, error)
	CreateAnimeSerieLogoImage(ctx context.Context, arg CreateAnimeSerieLogoImageParams) (AnimeSerieLogoImage, error)
	CreateAnimeSerieMeta(ctx context.Context, arg CreateAnimeSerieMetaParams) (AnimeSerieMeta, error)
	CreateAnimeSeriePosterImage(ctx context.Context, arg CreateAnimeSeriePosterImageParams) (AnimeSeriePosterImage, error)
	CreateAnimeSerieSeason(ctx context.Context, arg CreateAnimeSerieSeasonParams) (AnimeSerieSeason, error)
	CreateAnimeSerieSeasonEpisode(ctx context.Context, arg CreateAnimeSerieSeasonEpisodeParams) (AnimeSerieSeasonEpisode, error)
	CreateAnimeSerieSeasonMeta(ctx context.Context, arg CreateAnimeSerieSeasonMetaParams) (AnimeSerieSeasonMeta, error)
	CreateAnimeSerieSeasonPosterImage(ctx context.Context, arg CreateAnimeSerieSeasonPosterImageParams) (AnimeSerieSeasonPosterImage, error)
	CreateAnimeSerieServer(ctx context.Context, episodeID int64) (AnimeSerieServer, error)
	CreateAnimeSerieServerDubTorrent(ctx context.Context, arg CreateAnimeSerieServerDubTorrentParams) (AnimeSerieServerDubTorrent, error)
	CreateAnimeSerieServerDubVideo(ctx context.Context, arg CreateAnimeSerieServerDubVideoParams) (AnimeSerieServerDubVideo, error)
	CreateAnimeSerieServerSubTorrent(ctx context.Context, arg CreateAnimeSerieServerSubTorrentParams) (AnimeSerieServerSubTorrent, error)
	CreateAnimeSerieServerSubVideo(ctx context.Context, arg CreateAnimeSerieServerSubVideoParams) (AnimeSerieServerSubVideo, error)
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
	DeleteAnimeImage(ctx context.Context, id int64) error
	DeleteAnimeLink(ctx context.Context, id int64) error
	DeleteAnimeMovie(ctx context.Context, id int64) error
	DeleteAnimeMovieBackdropImage(ctx context.Context, arg DeleteAnimeMovieBackdropImageParams) error
	DeleteAnimeMovieGenre(ctx context.Context, arg DeleteAnimeMovieGenreParams) error
	DeleteAnimeMovieLink(ctx context.Context, arg DeleteAnimeMovieLinkParams) error
	DeleteAnimeMovieLogoImage(ctx context.Context, arg DeleteAnimeMovieLogoImageParams) error
	DeleteAnimeMovieMeta(ctx context.Context, arg DeleteAnimeMovieMetaParams) error
	DeleteAnimeMoviePosterImage(ctx context.Context, arg DeleteAnimeMoviePosterImageParams) error
	DeleteAnimeMovieResource(ctx context.Context, arg DeleteAnimeMovieResourceParams) error
	DeleteAnimeMovieServer(ctx context.Context, id int64) error
	DeleteAnimeMovieServerDubTorrent(ctx context.Context, id int64) error
	DeleteAnimeMovieServerDubVideo(ctx context.Context, id int64) error
	DeleteAnimeMovieServerSubTorrent(ctx context.Context, id int64) error
	DeleteAnimeMovieServerSubVideo(ctx context.Context, id int64) error
	DeleteAnimeMovieStudio(ctx context.Context, arg DeleteAnimeMovieStudioParams) error
	DeleteAnimeMovieTorrent(ctx context.Context, id int64) error
	DeleteAnimeMovieVideo(ctx context.Context, id int64) error
	DeleteAnimeResource(ctx context.Context, id int64) error
	DeleteAnimeSeasonResource(ctx context.Context, arg DeleteAnimeSeasonResourceParams) error
	DeleteAnimeSerie(ctx context.Context, id int64) error
	DeleteAnimeSerieBackdropImage(ctx context.Context, arg DeleteAnimeSerieBackdropImageParams) error
	DeleteAnimeSerieEpisode(ctx context.Context, id int64) error
	DeleteAnimeSerieEpisodeMeta(ctx context.Context, arg DeleteAnimeSerieEpisodeMetaParams) error
	DeleteAnimeSerieEpisodeServer(ctx context.Context, id int64) error
	DeleteAnimeSerieGenre(ctx context.Context, arg DeleteAnimeSerieGenreParams) error
	DeleteAnimeSerieLink(ctx context.Context, arg DeleteAnimeSerieLinkParams) error
	DeleteAnimeSerieLogoImage(ctx context.Context, arg DeleteAnimeSerieLogoImageParams) error
	DeleteAnimeSerieMeta(ctx context.Context, arg DeleteAnimeSerieMetaParams) error
	DeleteAnimeSeriePosterImage(ctx context.Context, arg DeleteAnimeSeriePosterImageParams) error
	DeleteAnimeSerieSeason(ctx context.Context, id int64) error
	DeleteAnimeSerieSeasonEpisode(ctx context.Context, id int64) error
	DeleteAnimeSerieSeasonMeta(ctx context.Context, arg DeleteAnimeSerieSeasonMetaParams) error
	DeleteAnimeSerieSeasonPosterImage(ctx context.Context, arg DeleteAnimeSerieSeasonPosterImageParams) error
	DeleteAnimeSerieServer(ctx context.Context, id int64) error
	DeleteAnimeSerieServerDubTorrent(ctx context.Context, id int64) error
	DeleteAnimeSerieServerDubVideo(ctx context.Context, id int64) error
	DeleteAnimeSerieServerSubTorrent(ctx context.Context, id int64) error
	DeleteAnimeSerieServerSubVideo(ctx context.Context, id int64) error
	DeleteAnimeSerieStudio(ctx context.Context, arg DeleteAnimeSerieStudioParams) error
	DeleteAnimeSerieTorrent(ctx context.Context, id int64) error
	DeleteAnimeSerieVideo(ctx context.Context, id int64) error
	DeleteGenre(ctx context.Context, id int32) error
	DeleteLanguage(ctx context.Context, id int32) error
	DeleteMeta(ctx context.Context, id int64) error
	DeleteSession(ctx context.Context, id uuid.UUID) error
	DeleteStudio(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int64) error
	GetAnimeImage(ctx context.Context, id int64) (AnimeImage, error)
	GetAnimeLink(ctx context.Context, id int64) (AnimeLink, error)
	GetAnimeMovie(ctx context.Context, id int64) (AnimeMovie, error)
	GetAnimeMovieBackdropImage(ctx context.Context, id int64) (AnimeMovieBackdropImage, error)
	GetAnimeMovieBackdropImageByAnimeID(ctx context.Context, animeID int64) (AnimeMovieBackdropImage, error)
	GetAnimeMovieGenre(ctx context.Context, arg GetAnimeMovieGenreParams) (AnimeMovieGenre, error)
	GetAnimeMovieLink(ctx context.Context, animeID int64) (AnimeMovieLink, error)
	GetAnimeMovieLogoImage(ctx context.Context, id int64) (AnimeMovieLogoImage, error)
	GetAnimeMovieLogoImageByAnimeID(ctx context.Context, animeID int64) (AnimeMovieLogoImage, error)
	GetAnimeMovieMeta(ctx context.Context, arg GetAnimeMovieMetaParams) (int64, error)
	GetAnimeMovieMetaByID(ctx context.Context, id int64) (AnimeMovieMeta, error)
	GetAnimeMoviePosterImage(ctx context.Context, id int64) (AnimeMoviePosterImage, error)
	GetAnimeMoviePosterImageByAnimeID(ctx context.Context, animeID int64) (AnimeMoviePosterImage, error)
	GetAnimeMovieResource(ctx context.Context, animeID int64) (AnimeMovieResource, error)
	GetAnimeMovieServer(ctx context.Context, id int64) (AnimeMovieServer, error)
	GetAnimeMovieServerDubTorrent(ctx context.Context, id int64) (AnimeMovieServerDubTorrent, error)
	GetAnimeMovieServerDubVideo(ctx context.Context, id int64) (AnimeMovieServerDubVideo, error)
	GetAnimeMovieServerSubTorrent(ctx context.Context, id int64) (AnimeMovieServerSubTorrent, error)
	GetAnimeMovieServerSubVideo(ctx context.Context, id int64) (AnimeMovieServerSubVideo, error)
	GetAnimeMovieStudio(ctx context.Context, arg GetAnimeMovieStudioParams) (AnimeMovieStudio, error)
	GetAnimeMovieTorrent(ctx context.Context, id int64) (AnimeMovieTorrent, error)
	GetAnimeMovieVideo(ctx context.Context, id int64) (AnimeMovieVideo, error)
	GetAnimeResource(ctx context.Context, id int64) (AnimeResource, error)
	GetAnimeSeasonResource(ctx context.Context, seasonID int64) (AnimeSeasonResource, error)
	GetAnimeSerie(ctx context.Context, id int64) (AnimeSerie, error)
	GetAnimeSerieBackdropImage(ctx context.Context, id int64) (AnimeSerieBackdropImage, error)
	GetAnimeSerieBackdropImageByAnimeID(ctx context.Context, animeID int64) (AnimeSerieBackdropImage, error)
	GetAnimeSerieEpisodeByEpisodeID(ctx context.Context, id int64) (AnimeSerieEpisode, error)
	GetAnimeSerieEpisodeMeta(ctx context.Context, arg GetAnimeSerieEpisodeMetaParams) (AnimeSerieEpisodeMeta, error)
	GetAnimeSerieEpisodeServer(ctx context.Context, id int64) (AnimeSerieEpisodeServer, error)
	GetAnimeSerieGenre(ctx context.Context, arg GetAnimeSerieGenreParams) (AnimeSerieGenre, error)
	GetAnimeSerieLink(ctx context.Context, animeID int64) (AnimeSerieLink, error)
	GetAnimeSerieLogoImage(ctx context.Context, id int64) (AnimeSerieLogoImage, error)
	GetAnimeSerieLogoImageByAnimeID(ctx context.Context, animeID int64) (AnimeSerieLogoImage, error)
	GetAnimeSerieMeta(ctx context.Context, arg GetAnimeSerieMetaParams) (int64, error)
	GetAnimeSerieMetaByID(ctx context.Context, id int64) (AnimeSerieMeta, error)
	GetAnimeSeriePosterImage(ctx context.Context, id int64) (AnimeSeriePosterImage, error)
	GetAnimeSeriePosterImageByAnimeID(ctx context.Context, animeID int64) (AnimeSeriePosterImage, error)
	GetAnimeSerieSeason(ctx context.Context, id int64) (AnimeSerieSeason, error)
	GetAnimeSerieSeasonEpisode(ctx context.Context, id int64) (AnimeSerieSeasonEpisode, error)
	GetAnimeSerieSeasonMeta(ctx context.Context, arg GetAnimeSerieSeasonMetaParams) (AnimeSerieSeasonMeta, error)
	GetAnimeSerieSeasonPosterImage(ctx context.Context, id int64) (AnimeSerieSeasonPosterImage, error)
	GetAnimeSerieSeasonPosterImageByAnimeID(ctx context.Context, seasonID int64) (AnimeSerieSeasonPosterImage, error)
	GetAnimeSerieServer(ctx context.Context, id int64) (AnimeSerieServer, error)
	GetAnimeSerieServerDubTorrent(ctx context.Context, id int64) (AnimeSerieServerDubTorrent, error)
	GetAnimeSerieServerDubVideo(ctx context.Context, id int64) (AnimeSerieServerDubVideo, error)
	GetAnimeSerieServerSubTorrent(ctx context.Context, id int64) (AnimeSerieServerSubTorrent, error)
	GetAnimeSerieServerSubVideo(ctx context.Context, id int64) (AnimeSerieServerSubVideo, error)
	GetAnimeSerieStudio(ctx context.Context, arg GetAnimeSerieStudioParams) (AnimeSerieStudio, error)
	GetAnimeSerieTorrent(ctx context.Context, id int64) (AnimeSerieTorrent, error)
	GetAnimeSerieVideo(ctx context.Context, id int64) (AnimeSerieVideo, error)
	GetGenre(ctx context.Context, id int32) (Genre, error)
	GetLanguage(ctx context.Context, id int32) (Language, error)
	GetMeta(ctx context.Context, id int64) (Meta, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetStudio(ctx context.Context, id int32) (Studio, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListAnimeMovieBackdropImages(ctx context.Context, arg ListAnimeMovieBackdropImagesParams) ([]int64, error)
	ListAnimeMovieGenres(ctx context.Context, animeID int64) ([]int32, error)
	ListAnimeMovieLogoImages(ctx context.Context, arg ListAnimeMovieLogoImagesParams) ([]int64, error)
	ListAnimeMovieMetas(ctx context.Context, animeID int64) ([]int64, error)
	ListAnimeMoviePosterImages(ctx context.Context, arg ListAnimeMoviePosterImagesParams) ([]int64, error)
	ListAnimeMovieServerDubTorrents(ctx context.Context, serverID int64) ([]AnimeMovieServerDubTorrent, error)
	ListAnimeMovieServerDubVideos(ctx context.Context, serverID int64) ([]AnimeMovieServerDubVideo, error)
	ListAnimeMovieServerSubTorrents(ctx context.Context, serverID int64) ([]AnimeMovieServerSubTorrent, error)
	ListAnimeMovieServerSubVideos(ctx context.Context, serverID int64) ([]AnimeMovieServerSubVideo, error)
	ListAnimeMovieServers(ctx context.Context, arg ListAnimeMovieServersParams) ([]AnimeMovieServer, error)
	ListAnimeMovieStudios(ctx context.Context, animeID int64) ([]int32, error)
	ListAnimeMovieTorrents(ctx context.Context, arg ListAnimeMovieTorrentsParams) ([]AnimeMovieTorrent, error)
	ListAnimeMovieVideos(ctx context.Context, arg ListAnimeMovieVideosParams) ([]AnimeMovieVideo, error)
	ListAnimeMovies(ctx context.Context, arg ListAnimeMoviesParams) ([]AnimeMovie, error)
	ListAnimeSerieBackdropImages(ctx context.Context, arg ListAnimeSerieBackdropImagesParams) ([]int64, error)
	ListAnimeSerieEpisodeMetasByEpisode(ctx context.Context, arg ListAnimeSerieEpisodeMetasByEpisodeParams) ([]AnimeSerieEpisodeMeta, error)
	ListAnimeSerieEpisodeServersByEpisode(ctx context.Context, arg ListAnimeSerieEpisodeServersByEpisodeParams) ([]AnimeSerieEpisodeServer, error)
	ListAnimeSerieEpisodesBySeasonID(ctx context.Context, arg ListAnimeSerieEpisodesBySeasonIDParams) ([]AnimeSerieEpisode, error)
	ListAnimeSerieGenres(ctx context.Context, animeID int64) ([]int32, error)
	ListAnimeSerieLogoImages(ctx context.Context, arg ListAnimeSerieLogoImagesParams) ([]int64, error)
	ListAnimeSerieMetas(ctx context.Context, animeID int64) ([]int64, error)
	ListAnimeSeriePosterImages(ctx context.Context, arg ListAnimeSeriePosterImagesParams) ([]int64, error)
	ListAnimeSerieSeasonEpisodesBySeason(ctx context.Context, arg ListAnimeSerieSeasonEpisodesBySeasonParams) ([]AnimeSerieSeasonEpisode, error)
	ListAnimeSerieSeasonMetasBySeason(ctx context.Context, arg ListAnimeSerieSeasonMetasBySeasonParams) ([]AnimeSerieSeasonMeta, error)
	ListAnimeSerieSeasonPosterImages(ctx context.Context, arg ListAnimeSerieSeasonPosterImagesParams) ([]int64, error)
	ListAnimeSerieSeasonsByAnimeID(ctx context.Context, arg ListAnimeSerieSeasonsByAnimeIDParams) ([]AnimeSerieSeason, error)
	ListAnimeSerieServerDubTorrents(ctx context.Context, serverID int64) ([]AnimeSerieServerDubTorrent, error)
	ListAnimeSerieServerDubVideos(ctx context.Context, serverID int64) ([]AnimeSerieServerDubVideo, error)
	ListAnimeSerieServerSubTorrents(ctx context.Context, serverID int64) ([]AnimeSerieServerSubTorrent, error)
	ListAnimeSerieServerSubVideos(ctx context.Context, serverID int64) ([]AnimeSerieServerSubVideo, error)
	ListAnimeSerieServers(ctx context.Context, arg ListAnimeSerieServersParams) ([]AnimeSerieServer, error)
	ListAnimeSerieStudios(ctx context.Context, animeID int64) ([]int32, error)
	ListAnimeSerieTorrents(ctx context.Context, arg ListAnimeSerieTorrentsParams) ([]AnimeSerieTorrent, error)
	ListAnimeSerieVideos(ctx context.Context, arg ListAnimeSerieVideosParams) ([]AnimeSerieVideo, error)
	ListAnimeSeries(ctx context.Context, arg ListAnimeSeriesParams) ([]AnimeSerie, error)
	ListGenres(ctx context.Context, arg ListGenresParams) ([]Genre, error)
	ListLanguages(ctx context.Context, arg ListLanguagesParams) ([]Language, error)
	ListStudios(ctx context.Context, arg ListStudiosParams) ([]Studio, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	RefreshSessions(ctx context.Context, username string) error
	UpdateAnimeImage(ctx context.Context, arg UpdateAnimeImageParams) (AnimeImage, error)
	UpdateAnimeLink(ctx context.Context, arg UpdateAnimeLinkParams) (AnimeLink, error)
	UpdateAnimeMovie(ctx context.Context, arg UpdateAnimeMovieParams) (AnimeMovie, error)
	UpdateAnimeMovieMeta(ctx context.Context, arg UpdateAnimeMovieMetaParams) (AnimeMovieMeta, error)
	UpdateAnimeMovieServer(ctx context.Context, arg UpdateAnimeMovieServerParams) (AnimeMovieServer, error)
	UpdateAnimeMovieServerDubTorrent(ctx context.Context, arg UpdateAnimeMovieServerDubTorrentParams) (AnimeMovieServerDubTorrent, error)
	UpdateAnimeMovieServerDubVideo(ctx context.Context, arg UpdateAnimeMovieServerDubVideoParams) (AnimeMovieServerDubVideo, error)
	UpdateAnimeMovieServerSubTorrent(ctx context.Context, arg UpdateAnimeMovieServerSubTorrentParams) (AnimeMovieServerSubTorrent, error)
	UpdateAnimeMovieServerSubVideo(ctx context.Context, arg UpdateAnimeMovieServerSubVideoParams) (AnimeMovieServerSubVideo, error)
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
	UpdateAnimeSerieServerDubTorrent(ctx context.Context, arg UpdateAnimeSerieServerDubTorrentParams) (AnimeSerieServerDubTorrent, error)
	UpdateAnimeSerieServerDubVideo(ctx context.Context, arg UpdateAnimeSerieServerDubVideoParams) (AnimeSerieServerDubVideo, error)
	UpdateAnimeSerieServerSubTorrent(ctx context.Context, arg UpdateAnimeSerieServerSubTorrentParams) (AnimeSerieServerSubTorrent, error)
	UpdateAnimeSerieServerSubVideo(ctx context.Context, arg UpdateAnimeSerieServerSubVideoParams) (AnimeSerieServerSubVideo, error)
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
