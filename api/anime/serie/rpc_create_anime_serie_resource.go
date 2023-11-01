package animeSerie

import (
	"context"
	"errors"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) CreateAnimeSerieResource(ctx context.Context, req *aspb.CreateAnimeSerieResourceRequest) (*aspb.CreateAnimeSerieResourceResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create anime serie")
	}

	if violations := validateCreateAnimeSerieResourceRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.CreateAnimeSerieResourceTxParams{
		AnimeID: req.GetAnimeID(),
		CreateAnimeResourceParams: db.CreateAnimeResourceParams{
			TmdbID:          req.Resources.GetTMDbID(),
			ImdbID:          req.Resources.GetIMDbID(),
			WikipediaUrl:    req.Resources.GetWikipediaUrl(),
			OfficialWebsite: req.Resources.GetOfficialWebsite(),
			CrunchyrollUrl:  req.Resources.GetCrunchyrollUrl(),
			SocialMedia:     req.Resources.GetSocialMedia(),
		},
	}

	data, err := server.gojo.CreateAnimeSerieResourceTx(ctx, arg)
	if err != nil {
		db.ErrorSQL(err)
		return nil, status.Errorf(codes.Internal, "failed to create anime serie : %s", err)
	}

	res := &aspb.CreateAnimeSerieResourceResponse{
		AnimeSerie: shared.ConvertAnimeSerie(data.AnimeSerie),
		Resources:  shared.ConvertAnimeResource(data.AnimeResource),
	}
	return res, nil
}

func validateCreateAnimeSerieResourceRequest(req *aspb.CreateAnimeSerieResourceRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("ID", err))
	}

	if req.GetResources() != nil {
		if err := utils.ValidateInt(int64(req.GetResources().GetTMDbID())); err != nil {
			violations = append(violations, shared.FieldViolation("TMDbID", err))
		}

		if err := utils.ValidateURL(req.GetResources().GetWikipediaUrl(), "wikipedia"); err != nil {
			violations = append(violations, shared.FieldViolation("wikipediaUrl", err))
		}

		if req.GetResources().GetCrunchyrollUrl() != "" {
			if err := utils.ValidateURL(req.GetResources().GetCrunchyrollUrl(), "crunchyroll"); err != nil {
				violations = append(violations, shared.FieldViolation("crunchyrollUrl", err))
			}
		}

		if req.GetResources().GetOfficialWebsite() != "" {
			if err := utils.ValidateURL(req.GetResources().GetOfficialWebsite(), ""); err != nil {
				violations = append(violations, shared.FieldViolation("officialWebsite", err))
			}
		}

		if req.GetResources().GetIMDbID() != "" {
			if err := utils.ValidateIMDbID(req.GetResources().GetIMDbID()); err != nil {
				violations = append(violations, shared.FieldViolation("IMDbID", err))
			}
		}

		if req.GetResources().GetSocialMedia() != nil {
			for _, s := range req.GetResources().SocialMedia {
				if err := utils.ValidateURL(s, ""); err != nil {
					violations = append(violations, shared.FieldViolation("socialMedia", err))
				}
			}
		}
	} else {
		violations = append(violations, shared.FieldViolation("resources", errors.New("you need to send the resources model")))
	}

	return violations
}
