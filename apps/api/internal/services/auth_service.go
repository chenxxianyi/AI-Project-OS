package services

import (
	"errors"
	"time"

	"github.com/ai-project-os/api/internal/config"
	"github.com/ai-project-os/api/internal/middleware"
	"github.com/ai-project-os/api/internal/models"
	"github.com/ai-project-os/api/internal/repositories"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo *repositories.UserRepository
	cfg  *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{repo: repositories.NewUserRepository(), cfg: cfg}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func (s *AuthService) Register(req RegisterRequest) (*AuthResponse, error) {
	if req.Email == "" || req.Username == "" || req.Password == "" {
		return nil, errors.New("email, username and password are required")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	existing, err := s.repo.FindByEmail(req.Email)
	if err == nil && existing != nil {
		return nil, errors.New("email already registered")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: string(hash),
		Role:         models.RoleUser,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	token, err := middleware.GenerateToken(s.cfg, user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	zap.L().Info("user registered", zap.String("email", user.Email))
	return &AuthResponse{Token: token, User: *user}, nil
}

func (s *AuthService) Login(req LoginRequest) (*AuthResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := middleware.GenerateToken(s.cfg, user.ID, string(user.Role))
	if err != nil {
		return nil, err
	}

	zap.L().Info("user logged in", zap.String("email", user.Email))
	return &AuthResponse{Token: token, User: *user}, nil
}

func (s *AuthService) GetCurrentUser(userID string) (*models.User, error) {
	return s.repo.FindByID(userID)
}

func (s *AuthService) TokenDuration() time.Duration {
	return time.Duration(s.cfg.JWTExpirationHrs) * time.Hour
}
