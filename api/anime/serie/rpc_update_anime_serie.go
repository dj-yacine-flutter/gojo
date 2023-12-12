package animeSerie

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/aspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *AnimeSerieServer) UpdateAnimeSerie(ctx context.Context, req *aspb.UpdateAnimeSerieRequest) (*aspb.UpdateAnimeSerieResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update anime serie")
	}

	if violations := validateUpdateAnimeSerieRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	arg := db.UpdateAnimeSerieParams{
		ID: req.GetAnimeID(),
		OriginalTitle: pgtype.Text{
			String: req.GetOriginalTitle(),
			Valid:  req.OriginalTitle != nil,
		},
		FirstYear: pgtype.Int4{
			Int32: req.GetFirstYear(),
			Valid: req.FirstYear != nil,
		},
		LastYear: pgtype.Int4{
			Int32: req.GetLastYear(),
			Valid: req.LastYear != nil,
		},
		MalID: pgtype.Int4{
			Int32: req.GetMalID(),
			Valid: req.MalID != nil,
		},
		TvdbID: pgtype.Int4{
			Int32: req.GetTvdbID(),
			Valid: req.TvdbID != nil,
		},
		TmdbID: pgtype.Int4{
			Int32: req.GetTmdbID(),
			Valid: req.LastYear != nil,
		},
		PortriatPoster: pgtype.Text{
			String: req.GetPortriatPoster(),
			Valid:  req.PortriatPoster != nil,
		},
		PortriatBlurHash: pgtype.Text{
			String: req.GetPortriatBlurHash(),
			Valid:  req.PortriatBlurHash != nil,
		},
		LandscapePoster: pgtype.Text{
			String: req.GetLandscapePoster(),
			Valid:  req.LandscapePoster != nil,
		},
		LandscapeBlurHash: pgtype.Text{
			String: req.GetLandscapeBlurHash(),
			Valid:  req.LandscapeBlurHash != nil,
		},
	}

	anime, err := server.gojo.UpdateAnimeSerie(ctx, arg)
	if err != nil {
		return nil, shared.DatabaseError("failed to update anime serie", err)
	}

	res := &aspb.UpdateAnimeSerieResponse{
		AnimeSerie: shared.ConvertAnimeSerie(anime),
	}
	return res, nil
}

func validateUpdateAnimeSerieRequest(req *aspb.UpdateAnimeSerieRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateInt(req.GetAnimeID()); err != nil {
		violations = append(violations, shared.FieldViolation("ID", err))
	}

	if req.OriginalTitle != nil {
		if err := utils.ValidateString(req.GetOriginalTitle(), 2, 500); err != nil {
			violations = append(violations, shared.FieldViolation("originalTitle", err))
		}
	}

	if req.FirstYear != nil {
		if err := utils.ValidateYear(req.GetFirstYear()); err != nil {
			violations = append(violations, shared.FieldViolation("firstYear", err))
		}
	}

	if req.LastYear != nil {
		if err := utils.ValidateYear(req.GetLastYear()); err != nil {
			violations = append(violations, shared.FieldViolation("lastYear", err))
		}
	}

	if req.MalID != nil {
		if err := utils.ValidateYear(req.GetMalID()); err != nil {
			violations = append(violations, shared.FieldViolation("malID", err))
		}
	}

	if req.TvdbID != nil {
		if err := utils.ValidateYear(req.GetTvdbID()); err != nil {
			violations = append(violations, shared.FieldViolation("tvdbID", err))
		}
	}

	if req.TmdbID != nil {
		if err := utils.ValidateYear(req.GetTmdbID()); err != nil {
			violations = append(violations, shared.FieldViolation("tmdbID", err))
		}
	}

	if req.PortriatPoster != nil {
		if err := utils.ValidateImage(req.GetPortriatPoster()); err != nil {
			violations = append(violations, shared.FieldViolation("portriatPoster", err))
		}
	}

	if req.PortriatBlurHash != nil {
		if err := utils.ValidateString(req.GetPortriatBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("portriatBlurHash", err))
		}
	}

	if req.LandscapePoster != nil {
		if err := utils.ValidateImage(req.GetLandscapePoster()); err != nil {
			violations = append(violations, shared.FieldViolation("landscapePoster", err))
		}
	}

	if req.LandscapeBlurHash != nil {
		if err := utils.ValidateString(req.GetLandscapeBlurHash(), 0, 100); err != nil {
			violations = append(violations, shared.FieldViolation("landscapeBlurHash", err))
		}
	}

	return violations
}
