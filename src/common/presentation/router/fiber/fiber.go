package router

import (
	"github.com/gofiber/fiber/v2"
	common_router "github.com/peam1146/go_clean_architecture/src/common/presentation/router"
)

type FiberRouter struct {
	fiber.Router
}

func NewFiberRouter(path string, app *fiber.App) *FiberRouter {
	r := app.Group(path)
	return &FiberRouter{r}
}

func (r *FiberRouter) Get(path string, h func(common_router.IContext)) {
	r.Router.Get(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) Put(path string, h func(common_router.IContext)) {
	r.Router.Put(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) Post(path string, h func(common_router.IContext)) {
	r.Router.Post(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) Patch(path string, h func(common_router.IContext)) {
	r.Router.Patch(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) Delete(path string, h func(common_router.IContext)) {
	r.Router.Delete(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) Options(path string, h func(common_router.IContext)) {
	r.Router.Options(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) Listen(addr string) error {
	return r.Router.(*fiber.App).Listen(addr)
}

func (r *FiberRouter) Group(path string) common_router.IRouter {
	g := r.Router.Group(path)
	return &FiberRouter{g}
}

func NewFiber(conf ...fiber.Config) common_router.IAPP {
	r := fiber.New(conf...)
	return &FiberRouter{r}
}
