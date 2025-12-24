package services

import (
	"log"
	"net/http"

	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
	"github.com/DavidAfdal/purchasing-systeam/internal/models"
	"github.com/DavidAfdal/purchasing-systeam/internal/repositories"
	"github.com/DavidAfdal/purchasing-systeam/pkg/errors"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type jwtResponse struct {
	Token      string `json:"access_token"`
	Expired_at string `json:"expired_at"`
}

type UserService interface {
	Register(req *dto.RegisterRequest) (*dto.UserResponse, error)
	Login(req *dto.LoginRequest) (*jwtResponse, error)
	Logout(tokenString string) error
}

type userService struct {
	userRepo     repositories.UserRepo
	tokenUseCase token.TokenUseCase
}

func NewUserService(userRepo repositories.UserRepo, tokenUseCase token.TokenUseCase) UserService {
	return &userService{userRepo: userRepo, tokenUseCase: tokenUseCase}
}

func (s *userService) Register(req *dto.RegisterRequest) (*dto.UserResponse, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		log.Printf("failed to hash password for user %s: %v", req.Username, err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "Something went wrong, please try again later")
	}

	user := &models.User{
		Username: req.Username,
		Password: string(hashPassword),
		Role:     "Staff",
	}

	registeredUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			log.Printf("duplicate username attempt: %s", req.Username)
			return nil, errors.NewAppError(http.StatusBadRequest, "username already exists")
		}

		log.Printf("failed to create user %s: %v", req.Username, err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "Something went wrong, please try again later")
	}

	return s.toUserResponse(registeredUser), nil
}

func (s *userService) Login(req *dto.LoginRequest) (*jwtResponse, error) {
	existedUser, err := s.userRepo.FindUserByUsername(req.Username)
	if err != nil {
		log.Printf("login failed for username %s: %v", req.Username, err)
		return nil, errors.NewAppError(http.StatusUnauthorized, "invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(req.Password)); err != nil {
		log.Printf("invalid password attempt for username %s", req.Username)
		return nil, errors.NewAppError(http.StatusUnauthorized, "invalid username or password")
	}

	claims := s.tokenUseCase.CreateClaims(existedUser.ID.String(), existedUser.Username, existedUser.Role)
	accessToken, expiredAt, err := s.tokenUseCase.GenerateAccessToken(claims)
	if err != nil {
		log.Printf("failed to generate access token for user %s: %v", req.Username, err)
		return nil, errors.NewAppError(http.StatusInternalServerError, "Something went wrong, please try again later")
	}

	data := &jwtResponse{
		Token:      accessToken,
		Expired_at: expiredAt.String(),
	}

	return data, nil
}

func (s *userService) Logout(tokenString string) error {
	return s.tokenUseCase.InvalidateToken(tokenString)
}

func (s *userService) toUserResponse(user *models.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Role:     user.Role,
	}
}
