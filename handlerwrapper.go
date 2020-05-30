package framework

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

// HandlerRouter is a wrapper for httprouter to be compatible with http.Handler
type HandlerRouter struct {
	router *httprouter.Router
}

func wrapHandler(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		h.ServeHTTP(w, r)
	}
}

// NewRouter creates a new httprouter.Router for the HandlerRouter router
func NewRouter() *HandlerRouter {
	return &HandlerRouter{httprouter.New()}
}

// GET is a wrapper for the router GET method
func (h *HandlerRouter) GET(route string, handler http.Handler) {
	h.router.GET(route, wrapHandler(handler))
}

// POST is a wrapper for the router POST method
func (h *HandlerRouter) POST(route string, handler http.Handler) {
	h.router.POST(route, wrapHandler(handler))
}

// PUT is a wrapper for the router PUT method
func (h *HandlerRouter) PUT(route string, handler http.Handler) {
	h.router.PUT(route, wrapHandler(handler))
}

// OPTIONS is a wrapper for the router OPTIONS method
func (h *HandlerRouter) OPTIONS(route string, handler http.Handler) {
	h.router.OPTIONS(route, wrapHandler(handler))
}

// DELETE is a wrapper for the router DELETE method
func (h *HandlerRouter) DELETE(route string, handler http.Handler) {
	h.router.DELETE(route, wrapHandler(handler))
}

// GetRouter returns a pointer to the underlying httprouter.Router
func (h *HandlerRouter) GetRouter() *httprouter.Router { return h.router }
