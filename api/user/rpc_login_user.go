package user

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb/uspb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *UserServer) LoginUser(ctx context.Context, req *uspb.LoginUserRequest) (*uspb.LoginUserResponse, error) {
	if violations := validateLoginUserRequest(req); violations != nil {
		return nil, shared.InvalidArgumentError(violations)
	}

	user, err := server.gojo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, shared.ApiError("failed to find user", err)
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "inccorect password : %s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		user.Role,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token : %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		user.Role,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token : %s", err)
	}

	md := shared.ExtractMetadata(ctx)
	arg := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    md.UserAgent,
		ClientIp:     md.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	session, err := server.gojo.CreateSession(ctx, arg)
	if err != nil {
		return nil, shared.ApiError("failed to renew session", err)
	}

	res := &uspb.LoginUserResponse{
		User: &uspb.User{
			Username:          user.Username,
			FullName:          user.FullName,
			Email:             user.Email,
			PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
			CreatedAt:         timestamppb.New(user.CreatedAt),
		},
		SessionID:             session.ID.String(),
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return res, nil
}

func validateLoginUserRequest(req *uspb.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := utils.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, shared.FieldViolation("username", err))
	}

	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, shared.FieldViolation("password", err))
	}

	return violations
}
