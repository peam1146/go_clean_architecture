package router

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type FiberCtx struct {
	*fiber.Ctx
}

func (c *FiberCtx) PaginationQueryParam(query interface{}) error {
	return c.QueryParser(query)
}

func NewFiberCtx(c *fiber.Ctx) *FiberCtx {
	return &FiberCtx{c}
}

func (c *FiberCtx) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberCtx) JSON(statusCode int, v interface{}) {
	c.Ctx.Status(statusCode).JSON(v)
}

func (c *FiberCtx) Param(key string) (value string, err error) {
	value = c.Params(key)
	if value == "" {
		return "", errors.New("No param")
	}

	return value, nil
}

func (c *FiberCtx) Method() string {
	return c.Ctx.Method()
}

func (c *FiberCtx) Path() string {
	return c.Ctx.Path()
}

func (c *FiberCtx) StoreValue(k string, v string) {
	c.Locals(k, v)
}

func (c *FiberCtx) Next() {
	c.Ctx.Next()
}
