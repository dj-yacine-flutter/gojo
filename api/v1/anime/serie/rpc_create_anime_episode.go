package asapiv1

import (
	"context"
	"errors"
	"time"

	shv1 "github.com/dj-yacine-flutter/gojo/api/v1/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *AnimeSerieServer) CreateAnimeEpisode(ctx context.Context, req *aspbv1.CreateAnimeEpisodeRequest) (*aspbv1.CreateAnimeEpisodeResponse, error) {
	authPayload, err := shv1.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shv1.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie episode")
	}

	if violations := validateCreateAnimeEpisodeRequest(req); violations != nil {
		return nil, shv1.InvalidArgumentError(violations)
	}

	var DBEM = make([]db.AnimeMetaTxParam, len(req.EpisodeMetas))
	for i, am := range req.EpisodeMetas {
		DBEM[i] = db.AnimeMetaTxParam{
			LanguageID: am.GetLanguageID(),
			CreateMetaParams: db.CreateMetaParams{
				Title:    am.GetMeta().GetTitle(),
				Overview: am.GetMeta().GetOverview(),
			},
		}
	}

	arg := db.CreateAnimeEpisodeTxParams{
		Episode: db.CreateAnimeEpisodeParams{
			SeasonID:             req.GetEpisode().GetSeasonID(),
			EpisodeNumber:        int32(req.GetEpisode().GetEpisodeNumber()),
			EpisodeOriginalTitle: req.GetEpisode().GetEpisodeOriginalTitle(),
			Aired:                req.GetEpisode().GetAired().AsTime(),
			Rating:               req.GetEpisode().GetRating(),
			Duration:             req.GetEpisode().GetDuration().AsDuration(),
			Thumbnails:           req.GetEpisode().GetThumbnails(),
			ThumbnailsBlurHash:   req.GetEpisode().GetThumbnailsBlurHash(),
		},
		EpisodeMetas: DBEM,
	}

	data, err := server.gojo.CreateAnimeEpisodeTx(ctx, arg)
	if err != nil {
		return nil, shv1.ApiError("failed to create anime serie episode", err)
	}

	var PBSM = make([]*nfpbv1.AnimeMetaResponse, len(data.AnimeEpisodeMetas))

	for i, am := range data.AnimeEpisodeMetas {
		PBSM[i] = &nfpbv1.AnimeMetaResponse{
			Meta:       shv1.ConvertMeta(am.Meta),
			LanguageID: am.LanguageID,
			CreatedAt:  timestamppb.New(am.Meta.CreatedAt),
		}
	}

	res := &aspbv1.CreateAnimeEpisodeResponse{
		Season:       convertAnimeSeason(data.AnimeSeason),
		Episode:      convertAnimeEpisode(data.AnimeEpisode),
		EpisodeMetas: PBSM,
	}
	return res, nil
}

func validateCreateAnimeEpisodeRequest(req *aspbv1.CreateAnimeEpisodeRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if req.GetEpisode() != nil {
		if err := utils.ValidateInt(req.GetEpisode().GetSeasonID()); err != nil {
			violations = append(violations, shv1.FieldViolation("seasonID", err))
		}

		if err := utils.ValidateInt(int64(req.GetEpisode().GetEpisodeNumber())); err != nil {
			violations = append(violations, shv1.FieldViolation("episodeNumber", err))
		}

		if err := utils.ValidateString(req.GetEpisode().GetEpisodeOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shv1.FieldViolation("episodeOriginalTitle", err))
		}

		if err := utils.ValidateDate(req.GetEpisode().GetAired().AsTime().Format(time.DateOnly)); err != nil {
			violations = append(violations, shv1.FieldViolation("aired", err))
		}

		if err := utils.ValidateString(req.GetEpisode().GetRating(), 2, 30); err != nil {
			violations = append(violations, shv1.FieldViolation("rating", err))
		}

		if err := utils.ValidateDuration(req.GetEpisode().GetDuration().AsDuration().String()); err != nil {
			violations = append(violations, shv1.FieldViolation("duration", err))
		}

		if err := utils.ValidateImage(req.GetEpisode().GetThumbnails()); err != nil {
			violations = append(violations, shv1.FieldViolation("thumbnails", err))
		}

		if err := utils.ValidateString(req.GetEpisode().GetThumbnailsBlurHash(), 0, 100); err != nil {
			violations = append(violations, shv1.FieldViolation("thumbnailsBlurHash", err))
		}

	} else {
		violations = append(violations, shv1.FieldViolation("episode", errors.New("episode :you need to send the episode model")))
	}

	if req.EpisodeMetas != nil {
		for _, am := range req.EpisodeMetas {
			if err := utils.ValidateInt(int64(am.GetLanguageID())); err != nil {
				violations = append(violations, shv1.FieldViolation("languageID", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetTitle(), 2, 500); err != nil {
				violations = append(violations, shv1.FieldViolation("title", err))
			}

			if err := utils.ValidateString(am.GetMeta().GetOverview(), 5, 5000); err != nil {
				violations = append(violations, shv1.FieldViolation("overview", err))
			}
		}
	} else {
		violations = append(violations, shv1.FieldViolation("episodeMetas", errors.New("episodeMetas > meta : you need to send at least one of meta model")))
	}

	return violations
}
