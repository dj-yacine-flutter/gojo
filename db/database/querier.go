// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateAnimeGenre(ctx context.Context, arg CreateAnimeGenreParams) error
	CreateAnimeMeta(ctx context.Context, arg CreateAnimeMetaParams) error
	CreateAnimeMovie(ctx context.Context, arg CreateAnimeMovieParams) (AnimeMovie, error)
	CreateAnimeStudio(ctx context.Context, arg CreateAnimeStudioParams) error
	CreateGenre(ctx context.Context, genreName string) (Genre, error)
	CreateLanguage(ctx context.Context, arg CreateLanguageParams) (Language, error)
	CreateMeta(ctx context.Context, arg CreateMetaParams) (Meta, error)
	CreateStudio(ctx context.Context, studioName string) (Studio, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAnimeGenre(ctx context.Context, arg DeleteAnimeGenreParams) error
	DeleteAnimeMeta(ctx context.Context, arg DeleteAnimeMetaParams) error
	DeleteAnimeMovie(ctx context.Context, id int64) error
	DeleteAnimeStudio(ctx context.Context, arg DeleteAnimeStudioParams) error
	DeleteGenre(ctx context.Context, id int32) error
	DeleteLanguage(ctx context.Context, id int32) error
	DeleteMeta(ctx context.Context, id int64) error
	DeleteStudio(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int64) error
	GetAnimeMovie(ctx context.Context, id int64) (AnimeMovie, error)
	GetMetaIDByAnimeAndLanguage(ctx context.Context, arg GetMetaIDByAnimeAndLanguageParams) (int64, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	ListAnimeGenres(ctx context.Context, arg ListAnimeGenresParams) ([]pgtype.Int4, error)
	ListAnimeMetas(ctx context.Context, arg ListAnimeMetasParams) ([]int64, error)
	ListAnimeStudios(ctx context.Context, arg ListAnimeStudiosParams) ([]pgtype.Int4, error)
	ListGenres(ctx context.Context, arg ListGenresParams) ([]Genre, error)
	ListLanguages(ctx context.Context, arg ListLanguagesParams) ([]Language, error)
	ListStudios(ctx context.Context, arg ListStudiosParams) ([]Studio, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateAnimeGenre(ctx context.Context, arg UpdateAnimeGenreParams) (AnimeGenre, error)
	UpdateAnimeMeta(ctx context.Context, arg UpdateAnimeMetaParams) (AnimeMeta, error)
	UpdateAnimeStudio(ctx context.Context, arg UpdateAnimeStudioParams) (AnimeStudio, error)
	UpdateGenre(ctx context.Context, arg UpdateGenreParams) (Genre, error)
	UpdateLanguage(ctx context.Context, arg UpdateLanguageParams) (Language, error)
	UpdateMeta(ctx context.Context, arg UpdateMetaParams) error
	UpdateStudio(ctx context.Context, arg UpdateStudioParams) (Studio, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
