package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAnimeMovieGenre(t *testing.T) AnimeMovieGenre {
	anime := createRandomAnimeMovie(t)
	genre := createRandomGenre(t)
	arg := CreateAnimeMovieGenreParams{
		AnimeID: anime.ID,
		GenreID: genre.ID,
	}

	animeMovieGenre, err := testGojo.CreateAnimeMovieGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, animeMovieGenre)

	require.Equal(t, arg.AnimeID, animeMovieGenre.AnimeID)
	require.Equal(t, arg.GenreID, animeMovieGenre.GenreID)
	require.NotZero(t, animeMovieGenre.ID)
	require.NotZero(t, animeMovieGenre.GenreID)

	return animeMovieGenre
}

func TestCreateAnimeMovieGenre(t *testing.T) {
	createRandomAnimeMovieGenre(t)
}

func TestGetAnimeMovieGenre(t *testing.T) {
	a := createRandomAnimeMovie(t)
	g := createRandomGenre(t)
	arg := CreateAnimeMovieGenreParams{
		AnimeID: a.ID,
		GenreID: g.ID,
	}

	Genre1, err := testGojo.CreateAnimeMovieGenre(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Genre1)

	Genre2, err := testGojo.GetAnimeMovieGenre(context.Background(), GetAnimeMovieGenreParams{
		AnimeID: a.ID,
		GenreID: Genre1.GenreID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, Genre2)

	require.Equal(t, Genre1.AnimeID, Genre2.AnimeID)
	require.Equal(t, Genre1.GenreID, Genre2.GenreID)
}

func TestDeleteAnimeMovieGenre(t *testing.T) {
	a := createRandomAnimeMovie(t)
	g := createRandomGenre(t)
	arg1 := CreateAnimeMovieGenreParams{
		AnimeID: a.ID,
		GenreID: g.ID,
	}

	genre, err := testGojo.CreateAnimeMovieGenre(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, genre)

	arg2 := DeleteAnimeMovieGenreParams{
		AnimeID: genre.AnimeID,
		GenreID: genre.GenreID,
	}

	err = testGojo.DeleteAnimeMovieGenre(context.Background(), arg2)
	require.NoError(t, err)
}

func TestListAnimeMovieGenres(t *testing.T) {
	a := createRandomAnimeMovie(t)
	require.NotEmpty(t, a)

	for i := 0; i < 3; i++ {
		genre := createRandomGenre(t)
		arg := CreateAnimeMovieGenreParams{
			AnimeID: a.ID,
			GenreID: genre.ID,
		}
		testGojo.CreateAnimeMovieGenre(context.Background(), arg)
	}

	genres, err := testGojo.ListAnimeMovieGenres(context.Background(), a.ID)
	require.NoError(t, err)
	require.NotNil(t, genres)
}
