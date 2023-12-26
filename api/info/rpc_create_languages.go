package info

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/nfpb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *InfoServer) CreateLanguages(ctx context.Context, req *nfpb.CreateLanguagesRequest) (*nfpb.CreateLanguagesResponse, error) {
	authPayload, err := shared.AuthorizeUser(ctx, server.tokenMaker, []string{utils.AdminRole, utils.RootRoll})
	if err != nil {
		return nil, shared.UnAuthenticatedError(err)
	}

	if authPayload.Role != utils.RootRoll {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create new language")
	}

	if violations := validateCreateLanguageRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	var LanguageParams []db.CreateLanguageParams
	for _, l := range req.GetLanguages() {
		LanguageParams = append(LanguageParams, db.CreateLanguageParams{
			LanguageName: l.LanguageName,
			LanguageCode: l.LanguageCode,
		})
	}

	result, err := server.gojo.CreateLanguagesTx(ctx, db.CreateLanguagesTxParams{
		CreateLanguageParams: LanguageParams,
	})
	if err != nil {
		return nil, shared.ApiError("failed to create new language", err)
	}

	var Languages []*nfpb.LanguageResponse
	for _, l := range result.Languages {
		language := shared.ConvertLanguage(l)
		Languages = append(Languages, language)
	}

	res := &nfpb.CreateLanguagesResponse{
		Languages: Languages,
	}

	return res, nil
}

func validateCreateLanguageRequest(req *nfpb.CreateLanguagesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateLanguage(req.GetLanguages()); err != nil {
		violations = append(violations, shared.FieldViolation("languages", err))
	}

	return violations
}
