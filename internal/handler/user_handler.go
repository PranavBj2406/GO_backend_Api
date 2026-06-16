package handler

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"

	"github.com/yourusername/go-backend-task/internal/models"
	"github.com/yourusername/go-backend-task/internal/service"
	"github.com/yourusername/go-backend-task/internal/validator"
)

// UserHandler handles HTTP requests for users.
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser handles POST /users.
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, errors.New("invalid request body"))
	}

	if err := validator.ValidateStruct(req); err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, err)
	}

	user, err := h.service.CreateUser(c.Context(), req)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUserByID handles GET /users/:id.
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := parseIDParam(c)
	if err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, err)
	}

	user, err := h.service.GetUserByID(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser handles PUT /users/:id.
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := parseIDParam(c)
	if err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, err)
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, errors.New("invalid request body"))
	}

	if err := validator.ValidateStruct(req); err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, err)
	}

	user, err := h.service.UpdateUser(c.Context(), id, req)
	if err != nil {
		return h.handleServiceError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// DeleteUser handles DELETE /users/:id.
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := parseIDParam(c)
	if err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, err)
	}

	if err := h.service.DeleteUser(c.Context(), id); err != nil {
		return h.handleServiceError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListUsers handles GET /users with pagination.
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page, err := parsePositiveIntQuery(c, "page", 1)
	if err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, err)
	}

	limit, err := parsePositiveIntQuery(c, "limit", 20)
	if err != nil {
		return h.errorResponse(c, fiber.StatusBadRequest, err)
	}

	offset := int32((page - 1) * limit)
	users, err := h.service.ListUsers(c.Context(), int32(limit), offset)
	if err != nil {
		return h.handleServiceError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) handleServiceError(c *fiber.Ctx, err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return h.errorResponse(c, fiber.StatusNotFound, errors.New("resource not found"))
	}

	if isBadRequestError(err) {
		return h.errorResponse(c, fiber.StatusBadRequest, err)
	}

	return h.errorResponse(c, fiber.StatusInternalServerError, errors.New("internal server error"))
}

func (h *UserHandler) errorResponse(c *fiber.Ctx, status int, err error) error {
	return c.Status(status).JSON(fiber.Map{"error": err.Error()})
}

func parseIDParam(c *fiber.Ctx) (int32, error) {
	idParam := c.Params("id")
	if idParam == "" {
		return 0, errors.New("id is required")
	}

	id64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		return 0, errors.New("id must be a valid integer")
	}
	if id64 < 1 {
		return 0, errors.New("id must be a positive integer")
	}
	return int32(id64), nil
}

func parsePositiveIntQuery(c *fiber.Ctx, key string, defaultValue int) (int, error) {
	valueStr := c.Query(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil || value < 1 {
		return 0, errors.New(key + " must be a positive integer")
	}
	return value, nil
}

func isBadRequestError(err error) bool {
	if err == nil {
		return false
	}

	lowered := strings.ToLower(err.Error())
	return strings.Contains(lowered, "dob") || strings.Contains(lowered, "invalid") || strings.Contains(lowered, "required") || strings.Contains(lowered, "must")
}
