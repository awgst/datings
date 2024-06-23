package usecase

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/awgst/datings/config"
	"github.com/awgst/datings/internal/customerror"
	"github.com/awgst/datings/internal/entity/model"
	authrequest "github.com/awgst/datings/internal/entity/request/auth"
	authresponse "github.com/awgst/datings/internal/entity/response/auth"
	"github.com/awgst/datings/internal/usecase/repo"
	"github.com/awgst/datings/pkg/password"
	"github.com/awgst/datings/pkg/token"
	"github.com/golang-jwt/jwt/v4"
)

type AuthUsecase interface {
	SignUp(req authrequest.SignupRequest) (authresponse.SignupResponse, error)
	Login(req authrequest.LoginRequest) (authresponse.LoginResponse, error)
}

type authUsecase struct {
	cfg        *config.Config
	userFinder repo.UserFinder
	userWriter repo.UserWriter
	token      token.Token
	password   password.Password
}

func NewAuthUsecase(uc authUsecase) AuthUsecase {
	return uc
}

func (u authUsecase) SignUp(req authrequest.SignupRequest) (authresponse.SignupResponse, error) {
	user, err := u.userFinder.FindByEmail(req.Email)
	if err != nil {
		return authresponse.SignupResponse{}, err
	}

	if user.ID != 0 {
		return authresponse.SignupResponse{}, customerror.Error{
			Code: customerror.ErrorCodeAlreadyExists,
			Err:  fmt.Sprintf("user with email %s already exists", req.Email),
		}
	}

	var createdUser = model.User{
		Email:        req.Email,
		PasswordHash: u.password.Hash(req.Password),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}
	if err = u.userWriter.Create(&createdUser); err != nil {
		return authresponse.SignupResponse{}, err
	}

	accessToken, err := u.token.JwtToken(u.cfg.JWT.Secret, jwt.MapClaims{
		"user_id": createdUser.ID,
		"email":   createdUser.Email,
		"premium": "none",
		"exp":     time.Now().Add(time.Minute * time.Duration(u.cfg.JWT.ExpireInMinutes)).Unix(),
	})
	if err != nil {
		return authresponse.SignupResponse{}, err
	}

	return authresponse.SignupResponse{
		Token: accessToken,
	}, nil
}

func (u authUsecase) Login(req authrequest.LoginRequest) (authresponse.LoginResponse, error) {
	user, err := u.userFinder.FindByEmail(req.Email)
	if err != nil {
		return authresponse.LoginResponse{}, err
	}

	if user.ID == 0 {
		return authresponse.LoginResponse{}, customerror.Error{
			Code: customerror.ErrorCodeInvalidCredentials,
			Err:  nil,
		}
	}

	if !u.password.Compare(user.PasswordHash, req.Password) {
		return authresponse.LoginResponse{}, customerror.Error{
			Code: customerror.ErrorCodeInvalidCredentials,
			Err:  nil,
		}
	}

	premium := "none"
	if user.Premium != nil {
		premium = string(user.Premium.Feature)
	}
	accessToken, err := u.token.JwtToken(u.cfg.JWT.Secret, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"premium": premium,
		"exp":     time.Now().Add(time.Minute * time.Duration(u.cfg.JWT.ExpireInMinutes)).Unix(),
	})
	if err != nil {
		return authresponse.LoginResponse{}, err
	}

	return authresponse.LoginResponse{
		Token: accessToken,
	}, nil
}
