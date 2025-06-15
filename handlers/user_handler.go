package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sunnysingha911/admin-service/gen/user-service/userpb"
	"github.com/sunnysingha911/admin-service/grpc"

	"context"
)

type UserHandler struct {
	GrpcClient *grpc.UserClient
}

func NewUserHandler(client *grpc.UserClient) *UserHandler {
	return &UserHandler{GrpcClient: client}
}

// Login handler calls user service Login RPC
func (h *UserHandler) Login(c *fiber.Ctx) error {
	req := new(struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	})

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	res, err := h.GrpcClient.Client.Login(context.Background(), &userpb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"token": res.Token,
		"user":  res.User,
	})
}

// GetUser handler calls GetUser RPC
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	user, err := h.GrpcClient.Client.GetUserById(context.Background(), &userpb.GetUserRequest{
		Id: string(id),
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	res, err := h.GrpcClient.Client.GetAllUsers(context.Background(), &userpb.GetAllUserRequest{
		Page:  int32(page),
		Limit: int32(limit),
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"meta":  res.Meta,
		"users": res.Users,
	})
}

// UpdateUser handler calls UpdateUser RPC
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	req := new(struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	})

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	updatedUser, err := h.GrpcClient.Client.UpdateUser(context.Background(), &userpb.UpdateUserRequest{
		Id:    int32(id),
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(updatedUser)
}

// DeleteUser handler calls DeleteUser RPC
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	res, err := h.GrpcClient.Client.DeleteUser(context.Background(), &userpb.DeleteUserRequest{
		Id: int32(id),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": res.Success})
}

// ListUsers handler calls ListUsers RPC
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	res, err := h.GrpcClient.Client.ListUsers(context.Background(), &userpb.ListUsersRequest{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res.Users)
}
