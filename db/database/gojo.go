package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Gojo interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
	AddAnimeMovieInfoTx(ctx context.Context, arg AddAnimeMovieInfoTxParams) (AddAnimeMovieInfoTxResult, error)
	AddAnimeSerieInfoTx(ctx context.Context, arg AddAnimeSerieInfoTxParams) (AddAnimeSerieInfoTxResult, error)
	CreateAnimeMovieMetasTx(ctx context.Context, arg CreateAnimeMovieMetasTxParams) (CreateAnimeMovieMetasTxResult, error)
	CreateAnimeSerieMetasTx(ctx context.Context, arg CreateAnimeSerieMetasTxParams) (CreateAnimeSerieMetasTxResult, error)
	CreateGenresTx(ctx context.Context, arg CreateGenresTxParams) (CreateGenresTxResult, error)
	CreateStudiosTx(ctx context.Context, arg CreateStudiosTxParams) (CreateStudiosTxResult, error)
	CreateLanguagesTx(ctx context.Context, arg CreateLanguagesTxParams) (CreateLanguagesTxResult, error)
	RenewSessionTx(ctx context.Context, arg RenewSessionTxParams) (RenewSessionTxResult, error)
	CreateAnimeMovieTx(ctx context.Context, arg CreateAnimeMovieTxParams) (CreateAnimeMovieTxResult, error)
	CreateAnimeSerieTx(ctx context.Context, arg CreateAnimeSerieTxParams) (CreateAnimeSerieTxResult, error)
	CreateAnimeMovieResourceTx(ctx context.Context, arg CreateAnimeMovieResourceTxParams) (CreateAnimeMovieResourceTxResult, error)
	CreateAnimeMovieLinkTx(ctx context.Context, arg CreateAnimeMovieLinkTxParams) (CreateAnimeMovieLinkTxResult, error)
	CreateAnimeSerieLinkTx(ctx context.Context, arg CreateAnimeSerieLinkTxParams) (CreateAnimeSerieLinkTxResult, error)
	CreateAnimeSeasonResourceTx(ctx context.Context, arg CreateAnimeSeasonResourceTxParams) (CreateAnimeSeasonResourceTxResult, error)
	AddAnimeMovieDataTx(ctx context.Context, arg AddAnimeMovieDataTxParams) (AddAnimeMovieDataTxResult, error)
	CreateAnimeSerieSeasonTx(ctx context.Context, arg CreateAnimeSerieSeasonTxParams) (CreateAnimeSerieSeasonTxResult, error)
	AddAnimeSerieSeasonMetasTx(ctx context.Context, arg AddAnimeSerieSeasonMetasTxParams) (AddAnimeSerieSeasonMetasTxResult, error)
	CreateAnimeSerieEpisodeTx(ctx context.Context, arg CreateAnimeSerieEpisodeTxParams) (CreateAnimeSerieEpisodeTxResult, error)
	AddAnimeSerieEpisodeMetasTx(ctx context.Context, arg AddAnimeSerieEpisodeMetasTxParams) (AddAnimeSerieEpisodeMetasTxResult, error)
	AddAnimeSerieDataTx(ctx context.Context, arg AddAnimeSerieDataTxParams) (AddAnimeSerieDataTxResult, error)
	CreateAnimeMovieImageTx(ctx context.Context, arg CreateAnimeMovieImageTxParams) (CreateAnimeMovieImageTxResult, error)
}

type SQLGojo struct {
	*Queries
	coonPool *pgxpool.Pool
}

func NewGojo(coonPool *pgxpool.Pool) Gojo {
	return &SQLGojo{
		coonPool: coonPool,
		Queries:  New(coonPool),
	}
}
