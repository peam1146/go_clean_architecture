package common_router

type IContext interface {
	JSON(int, interface{})
	Param(key string) (value string, err error)
	Bind(v interface{}) error
}

type IRouter interface {
	Group(path string) IRouter
	Get(path string, h func(c IContext))
	Put(path string, h func(c IContext))
	Patch(path string, h func(c IContext))
	Delete(path string, h func(c IContext))
	Options(path string, h func(c IContext))
}

type IAPP interface {
	IRouter
	Listen(addr string) error
}
