package repository

import (
	"context"

	"github.com/yourusername/go-backend-task/db/sqlc/generated"
)

// UserRepository provides database access for user entities.
// It contains no business logic and delegates directly to SQLC-generated queries.
type UserRepository struct {
	queries *generated.Queries
}

// NewUserRepository creates a new UserRepository with the provided SQLC queries.
func NewUserRepository(q *generated.Queries) *UserRepository {
	return &UserRepository{queries: q}
}

// CreateUser inserts a new user record and returns the new ID.
func (r *UserRepository) CreateUser(ctx context.Context, arg generated.CreateUserParams) (int32, error) {
	return r.queries.CreateUser(ctx, arg)
}

// GetUserByID returns a user by its ID.
func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (generated.User, error) {
	return r.queries.GetUserByID(ctx, id)
}

// UpdateUser updates an existing user record.
func (r *UserRepository) UpdateUser(ctx context.Context, arg generated.UpdateUserParams) (generated.User, error) {
	return r.queries.UpdateUser(ctx, arg)
}

// DeleteUser deletes a user record by ID.
func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}

// ListUsers returns a list of users with limit/offset pagination.
func (r *UserRepository) ListUsers(ctx context.Context, arg generated.ListUsersParams) ([]generated.User, error) {
	return r.queries.ListUsers(ctx, arg)
}
