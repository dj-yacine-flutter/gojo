package db

import (
	"context"
	"errors"
)

type CreateAnimeSerieEpisodeTxParams struct {
	Episode      CreateAnimeSerieEpisodeParams
	EpisodeMetas []AnimeMetaTxParam
}

type CreateAnimeSerieEpisodeTxResult struct {
	AnimeSerieEpisode      AnimeSerieEpisode
	AnimeSerieEpisodeMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSerieEpisodeTx(ctx context.Context, arg CreateAnimeSerieEpisodeTxParams) (CreateAnimeSerieEpisodeTxResult, error) {
	var result CreateAnimeSerieEpisodeTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		_, err = q.GetAnimeSerieSeason(ctx, arg.Episode.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		Episode, err := q.CreateAnimeSerieEpisode(ctx, arg.Episode)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSerieEpisode = Episode

		if arg.EpisodeMetas != nil {
			var metaArg CreateMetaParams
			var EpisodeMetaArg CreateAnimeSerieEpisodeMetaParams
			result.AnimeSerieEpisodeMetas = make([]AnimeMetaTxResult, len(arg.EpisodeMetas))

			for i, m := range arg.EpisodeMetas {
				metaArg = CreateMetaParams{
					Title:    m.Title,
					Overview: m.Overview,
				}

				meta, err := q.CreateMeta(ctx, metaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				EpisodeMetaArg = CreateAnimeSerieEpisodeMetaParams{
					EpisodeID:  Episode.ID,
					LanguageID: m.LanguageID,
					MetaID:     meta.ID,
				}

				_, err = q.CreateAnimeSerieEpisodeMeta(ctx, EpisodeMetaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				l, err := q.GetLanguage(ctx, m.LanguageID)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSerieEpisodeMetas[i] = AnimeMetaTxResult{
					Meta:     meta,
					Language: l,
				}

			}

		} else {
			return errors.New("create one meta at least")
		}

		return err
	})

	return result, err
}
