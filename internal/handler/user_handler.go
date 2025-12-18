package handler

import (
	"strconv"

	"age-calculator/internal/middleware"
	"age-calculator/internal/models"
	"age-calculator/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := middleware.Validate.Struct(req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	res, err := h.svc.Create(
	c.Context(),
	req.Name,
	req.DOB.ToTime(),
    )
    if err != nil {
	return c.SendStatus(fiber.StatusInternalServerError)
}


    return c.Status(fiber.StatusCreated).JSON(res)

}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := h.svc.GetByID(c.Context(), id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(user)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	users, err := h.svc.GetUsers(c.Context(), page, limit)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(users)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := middleware.Validate.Struct(req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	res, err := h.svc.Update(
	    c.Context(),
	    id,
	    req.Name,
	    req.DOB.ToTime(),
    )
    if err != nil {
	    return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.JSON(res)
 
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := h.svc.Delete(c.Context(), id); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
