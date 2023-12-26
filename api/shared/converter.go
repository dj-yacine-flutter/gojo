package shared

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/ampb"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/pb/shpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertGenre(genre db.Genre) *nfpb.Genre {
	return &nfpb.Genre{
		GenreID:   genre.ID,
		GenreName: genre.GenreName,
		CreatedAt: timestamppb.New(genre.CreatedAt),
	}
}

func ConvertStudio(studio db.Studio) *nfpb.Studio {
	return &nfpb.Studio{
		StudioID:   studio.ID,
		StudioName: studio.StudioName,
		CreatedAt:  timestamppb.New(studio.CreatedAt),
	}
}

func ConvertLanguage(language db.Language) *nfpb.LanguageResponse {
	return &nfpb.LanguageResponse{
		LanguageID:   language.ID,
		LanguageCode: language.LanguageCode,
		LanguageName: language.LanguageName,
		CreatedAt:    timestamppb.New(language.CreatedAt),
	}
}

func ConvertMeta(meta db.Meta) *nfpb.MetaResponse {
	return &nfpb.MetaResponse{
		MetaID:   meta.ID,
		Title:    meta.Title,
		Overview: meta.Overview,
	}
}

func ConvertAnimeMovie(a db.AnimeMovie) *ampb.AnimeMovieResponse {
	return &ampb.AnimeMovieResponse{
		ID:                a.ID,
		OriginalTitle:     a.OriginalTitle,
		Aired:             timestamppb.New(a.Aired),
		ReleaseYear:       a.ReleaseYear,
		Rating:            a.Rating,
		Duration:          durationpb.New(a.Duration),
		PortraitPoster:    a.PortraitPoster,
		PortraitBlurHash:  a.PortraitBlurHash,
		LandscapePoster:   a.LandscapePoster,
		LandscapeBlurHash: a.LandscapeBlurHash,
		CreatedAt:         timestamppb.New(a.CreatedAt),
	}
}

func ConvertAnimeSerie(a db.AnimeSerie) *aspb.AnimeSerieResponse {
	return &aspb.AnimeSerieResponse{
		ID:                a.ID,
		OriginalTitle:     a.OriginalTitle,
		PortraitPoster:    a.PortraitPoster,
		PortraitBlurHash:  a.PortraitBlurHash,
		LandscapePoster:   a.LandscapePoster,
		LandscapeBlurHash: a.LandscapeBlurHash,
		FirstYear:         a.FirstYear,
		LastYear:          a.LastYear,
		MalID:             a.MalID,
		TvdbID:            a.TvdbID,
		TmdbID:            a.TmdbID,
		CreatedAt:         timestamppb.New(a.CreatedAt),
	}
}

func ConvertAnimeResource(r db.AnimeResource) *shpb.AnimeResourceResponse {
	return &shpb.AnimeResourceResponse{
		ID:            r.ID,
		TvdbID:        r.TvdbID,
		TmdbID:        r.TmdbID,
		ImdbID:        r.ImdbID,
		LivechartID:   r.LivechartID,
		AnimePlanetID: r.AnimePlanetID,
		AnisearchID:   r.AnisearchID,
		AnidbID:       r.AnidbID,
		KitsuID:       r.KitsuID,
		MalID:         r.MalID,
		NotifyMoeID:   r.NotifyMoeID,
		AnilistID:     r.AnilistID,
		CreatedAt:     timestamppb.New(r.CreatedAt),
	}
}

func ConvertAnimeLink(l db.AnimeLink) *shpb.AnimeLinkResponse {
	return &shpb.AnimeLinkResponse{
		ID:              l.ID,
		OfficialWebsite: l.OfficialWebsite,
		WikipediaUrl:    l.WikipediaUrl,
		CrunchyrollUrl:  l.CrunchyrollUrl,
		SocialMedia:     l.SocialMedia,
		CreatedAt:       timestamppb.New(l.CreatedAt),
	}
}

func ConvertGenres(gg []db.Genre) []*nfpb.Genre {
	if len(gg) > 0 {
		Genres := make([]*nfpb.Genre, len(gg))

		for i, g := range gg {
			Genres[i] = &nfpb.Genre{
				GenreID:   g.ID,
				GenreName: g.GenreName,
				CreatedAt: timestamppb.New(g.CreatedAt),
			}
		}

		return Genres
	} else {
		return nil
	}
}

func ConvertStudios(ss []db.Studio) []*nfpb.Studio {
	if len(ss) > 0 {
		Studios := make([]*nfpb.Studio, len(ss))

		for i, s := range ss {
			Studios[i] = &nfpb.Studio{
				StudioID:   s.ID,
				StudioName: s.StudioName,
				CreatedAt:  timestamppb.New(s.CreatedAt),
			}
		}

		return Studios
	} else {
		return nil
	}
}

func ConvertAnimeSeason(s db.AnimeSerieSeason) *aspb.AnimeSeasonResponse {
	return &aspb.AnimeSeasonResponse{
		ID:                  s.ID,
		AnimeID:             s.AnimeID,
		SeasonOriginalTitle: s.SeasonOriginalTitle,
		ReleaseYear:         s.ReleaseYear,
		Aired:               timestamppb.New(s.Aired),
		Rating:              s.Rating,
		PortraitPoster:      s.PortraitPoster,
		PortraitBlurHash:    s.PortraitBlurHash,
		CreatedAt:           timestamppb.New(s.CreatedAt),
	}
}

func ConvertAnimeEpisode(e db.AnimeSerieEpisode) *aspb.AnimeEpisodeResponse {
	return &aspb.AnimeEpisodeResponse{
		ID:                   e.ID,
		SeasonID:             e.SeasonID,
		EpisodeNumber:        uint32(e.EpisodeNumber),
		EpisodeOriginalTitle: e.EpisodeOriginalTitle,
		Aired:                timestamppb.New(e.Aired),
		Rating:               e.Rating,
		Duration:             durationpb.New(e.Duration),
		Thumbnails:           e.Thumbnails,
		ThumbnailsBlurHash:   e.ThumbnailsBlurHash,
		CreatedAt:            timestamppb.New(e.CreatedAt),
	}
}

func ConvertAnimeMovieVideos(amv []db.AnimeMovieVideo) []*ampb.AnimeMovieVideoResponse {
	if len(amv) > 0 {
		Videos := make([]*ampb.AnimeMovieVideoResponse, len(amv))

		for i, v := range amv {
			Videos[i] = &ampb.AnimeMovieVideoResponse{
				ID:         v.ID,
				LanguageID: v.LanguageID,
				Authority:  v.Authority,
				Referer:    v.Referer,
				Link:       v.Link,
				Quality:    v.Quality,
				CreatedAt:  timestamppb.New(v.CreatedAt),
			}
		}
		return Videos
	} else {
		return nil
	}
}

func ConvertAnimeMovieTorrents(amt []db.AnimeMovieTorrent) []*ampb.AnimeMovieTorrentResponse {
	if len(amt) > 0 {
		Torrents := make([]*ampb.AnimeMovieTorrentResponse, len(amt))

		for i, t := range amt {
			Torrents[i] = &ampb.AnimeMovieTorrentResponse{
				ID:          t.ID,
				LanguageID:  t.LanguageID,
				FileName:    t.FileName,
				TorrentHash: t.TorrentHash,
				TorrentFile: t.TorrentFile,
				Seeds:       t.Seeds,
				Peers:       t.Peers,
				Leechers:    t.Leechers,
				SizeBytes:   t.SizeBytes,
				Quality:     t.Quality,
				CreatedAt:   timestamppb.New(t.CreatedAt),
			}
		}

		return Torrents
	} else {
		return nil
	}
}

func ConvertAnimeSerieVideos(asv []db.AnimeSerieVideo) []*aspb.AnimeSerieVideoResponse {
	if len(asv) > 0 {
		Videos := make([]*aspb.AnimeSerieVideoResponse, len(asv))

		for i, v := range asv {
			Videos[i] = &aspb.AnimeSerieVideoResponse{
				ID:         v.ID,
				LanguageID: v.LanguageID,
				Authority:  v.Authority,
				Referer:    v.Referer,
				Link:       v.Link,
				Quality:    v.Quality,
				CreatedAt:  timestamppb.New(v.CreatedAt),
			}
		}
		return Videos
	} else {
		return nil
	}
}

func ConvertAnimeSerieTorrents(ast []db.AnimeSerieTorrent) []*aspb.AnimeSerieTorrentResponse {
	if len(ast) > 0 {
		Torrents := make([]*aspb.AnimeSerieTorrentResponse, len(ast))

		for i, t := range ast {
			Torrents[i] = &aspb.AnimeSerieTorrentResponse{
				ID:          t.ID,
				LanguageID:  t.LanguageID,
				FileName:    t.FileName,
				TorrentHash: t.TorrentHash,
				TorrentFile: t.TorrentFile,
				Seeds:       t.Seeds,
				Peers:       t.Peers,
				Leechers:    t.Leechers,
				SizeBytes:   t.SizeBytes,
				Quality:     t.Quality,
				CreatedAt:   timestamppb.New(t.CreatedAt),
			}
		}

		return Torrents
	} else {
		return nil
	}
}

func ConvertAnimeEpiosdes(ase []db.AnimeSerieEpisode) []*aspb.AnimeEpisodeResponse {
	if len(ase) > 0 {
		episodes := make([]*aspb.AnimeEpisodeResponse, len(ase))

		for i, e := range ase {
			episodes[i] = &aspb.AnimeEpisodeResponse{
				ID:                   e.ID,
				SeasonID:             e.SeasonID,
				EpisodeNumber:        uint32(e.EpisodeNumber),
				EpisodeOriginalTitle: e.EpisodeOriginalTitle,
				Aired:                timestamppb.New(e.Aired),
				Rating:               e.Rating,
				Duration:             durationpb.New(e.Duration),
				Thumbnails:           e.Thumbnails,
				ThumbnailsBlurHash:   e.ThumbnailsBlurHash,
				CreatedAt:            timestamppb.New(e.CreatedAt),
			}
		}
		return episodes
	} else {
		return nil
	}
}

func ConvertAnimeImages(ai []db.AnimeImage) []*shpb.ImageResponse {
	if len(ai) > 0 {
		images := make([]*shpb.ImageResponse, len(ai))

		for i, g := range ai {
			images[i] = &shpb.ImageResponse{
				ID:         g.ID,
				Host:       g.ImageHost,
				Url:        g.ImageUrl,
				Thumbnails: g.ImageThumbnails,
				Blurhash:   g.ImageBlurhash,
				Height:     uint32(g.ImageHeight),
				Width:      uint32(g.ImageWidth),
				CreatedAt:  timestamppb.New(g.CreatedAt),
			}
		}

		return images
	} else {
		return nil
	}
}

func ConvertAnimeTrailers(at []db.AnimeTrailer) []*shpb.AnimeTrailerResponse {
	if len(at) > 0 {
		trailers := make([]*shpb.AnimeTrailerResponse, len(at))

		for i, t := range at {
			trailers[i] = &shpb.AnimeTrailerResponse{
				ID:         t.ID,
				IsOfficial: t.IsOfficial,
				HostName:   t.HostName,
				HostKey:    t.HostKey,
				CreatedAt:  timestamppb.New(t.CreatedAt),
			}
		}

		return trailers
	} else {
		return nil
	}
}
