package services

import (
	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

type jwtResponse struct {
	Token      string `json:"access_token"`
	Expired_at string `json:"expired_at"`
}

type UserService interface {
	Register(req *dto.RegisterRequest) (*dto.UserResponse, error)
	Login(req *dto.LoginRequest) (*jwtResponse, error)
}

type userService struct {
	userRepo     repositories.UserRepo
	tokenUseCase token.TokenUseCase
}

func NewUserService(userRepo repositories.UserRepo, tokenUseCase token.TokenUseCase) UserService {
	return &userService{userRepo: userRepo, tokenUseCase: tokenUseCase}
}

func (s *userService) Register(req *dto.RegisterRequest) (*dto.UserResponse, error) {

	hashPassowrd, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: string(hashPassowrd),
		Role:     req.Role,
	}

	registerdUser, err := s.userRepo.CreateUser(user)

	return s.toUserResponse(registerdUser), nil
}

func (s *userService) Login(req *dto.LoginRequest) (*jwtResponse, error) {
	existedUser, err := s.userRepo.GetUserByUsername(req.Username)

	data := jwtResponse{}
	data.Token = ""
	data.Expired_at = ""

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	claims := s.tokenUseCase.CreateClaims(existedUser.ID.String(), existedUser.Username, existedUser.Role)

	accessToken, expiredAt, err := s.tokenUseCase.GenerateAccessToken(claims)

	if err != nil {
		return nil, err
	}

	data.Token = accessToken
	data.Expired_at = expiredAt.String()

	return &data, nil
}

func (s *userService) toUserResponse(user *models.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Role:     user.Role,
	}
}
