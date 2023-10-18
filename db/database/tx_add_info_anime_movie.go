package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type AddInfoAnimeMovieTxParams struct {
	AnimeID   int64
	GenreIDs  []int32
	StudioIDs []int32
}

type AddInfoAnimeMovieTxResult struct {
	AnimeMovie
	AnimeMovieGenres  []AnimeMovieGenre
	AnimeMovieStudios []AnimeMovieStudio
}

func (gojo *SQLGojo) AddInfoAnimeMovieTx(ctx context.Context, arg AddInfoAnimeMovieTxParams) (AddInfoAnimeMovieTxResult, error) {
	var result AddInfoAnimeMovieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		result.AnimeMovie, err = q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.GenreIDs != nil {
			var argGenre CreateAnimeMovieGenreParams
			for _, g := range arg.GenreIDs {
				argGenre = CreateAnimeMovieGenreParams{
					AnimeID: result.AnimeMovie.ID,
					GenreID: pgtype.Int4{
						Int32: g,
						Valid: true,
					},
				}
				ag, err := q.CreateAnimeMovieGenre(ctx, argGenre)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				result.AnimeMovieGenres = append(result.AnimeMovieGenres, ag)
			}
		}

		if arg.StudioIDs != nil {
			var argStudio CreateAnimeMovieStudioParams
			for _, s := range arg.StudioIDs {
				argStudio = CreateAnimeMovieStudioParams{
					AnimeID: result.AnimeMovie.ID,
					StudioID: pgtype.Int4{
						Int32: s,
						Valid: true,
					},
				}

				as, err := q.CreateAnimeMovieStudio(ctx, argStudio)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				result.AnimeMovieStudios = append(result.AnimeMovieStudios, as)
			}
		}

		return err
	})

	return result, err
}
