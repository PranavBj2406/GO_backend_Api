package service

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/yourusername/go-backend-task/db/sqlc/generated"
	"github.com/yourusername/go-backend-task/internal/models"
	"github.com/yourusername/go-backend-task/internal/repository"
	"github.com/yourusername/go-backend-task/internal/utils"
)

// UserService contains business logic for users.
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser validates DOB, calls repository to create a user, and returns API DTO.
func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	// Parse DOB
	dobTime, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, errors.New("dob must be in YYYY-MM-DD format")
	}

	// Reject future dates
	today := time.Now()
	if dobTime.After(today) {
		return models.UserResponse{}, errors.New("dob cannot be in the future")
	}

	// Prepare params for repository (pgtype.Date)
	var dob pgtype.Date
	dob.Time = dobTime

	arg := generated.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	}

	id, err := s.repo.CreateUser(ctx, arg)
	if err != nil {
		return models.UserResponse{}, err
	}

	// Fetch created record to return full response
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return s.toDTO(user)
}

// GetUserByID returns a user DTO by ID.
func (s *UserService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}
	return s.toDTO(user)
}

// UpdateUser validates DOB and updates an existing user, returning the DTO.
func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dobTime, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return models.UserResponse{}, errors.New("dob must be in YYYY-MM-DD format")
	}
	today := time.Now()
	if dobTime.After(today) {
		return models.UserResponse{}, errors.New("dob cannot be in the future")
	}

	var dob pgtype.Date
	dob.Time = dobTime

	arg := generated.UpdateUserParams{
		Name: req.Name,
		Dob:  dob,
		ID:   id,
	}

	user, err := s.repo.UpdateUser(ctx, arg)
	if err != nil {
		return models.UserResponse{}, err
	}
	return s.toDTO(user)
}

// DeleteUser removes a user by ID.
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

// ListUsers returns paginated user DTOs.
func (s *UserService) ListUsers(ctx context.Context, limit, offset int32) ([]models.UserResponse, error) {
	arg := generated.ListUsersParams{Limit: limit, Offset: offset}
	users, err := s.repo.ListUsers(ctx, arg)
	if err != nil {
		return nil, err
	}
	resp := make([]models.UserResponse, 0, len(users))
	for _, u := range users {
		dto, err := s.toDTO(u)
		if err != nil {
			return nil, err
		}
		resp = append(resp, dto)
	}
	return resp, nil
}

// toDTO converts a generated.User to models.UserResponse.
func (s *UserService) toDTO(u generated.User) (models.UserResponse, error) {
	// Extract DOB as time.Time
	var dobStr string
	var dobTime time.Time
	if !u.Dob.Time.IsZero() {
		dobTime = u.Dob.Time
		dobStr = dobTime.Format("2006-01-02")
	} else {
		dobStr = ""
	}

	age := utils.CalculateAge(dobTime)

	return models.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		DOB:  dobStr,
		Age:  age,
	}, nil
}
