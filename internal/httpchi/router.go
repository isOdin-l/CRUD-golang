package httpchi

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/isOdin/RestApi/internal/handler"
	mw "github.com/isOdin/RestApi/internal/middleware"
)

func NewRouter(md *mw.Middleware, h *handler.Handler) chi.Router {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// /auth/...
	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.Authorization.SignUpHandler)
		r.Post("/sign-in", h.Authorization.SignInHandler)
	})

	// api/...
	r.Route("/api", func(r chi.Router) {
		r.Use(md.JWTAuth)
		r.Route("/lists", func(r chi.Router) { // api/lists/...
			r.Post("/", h.TodoList.CreateList)
			r.Get("/", h.TodoList.GetAllLists)
			r.Get("/{list_id}", h.TodoList.GetListById)
			r.Put("/{list_id}", h.TodoList.UpdateList)
			r.Delete("/{list_id}", h.TodoList.DeleteList)

			r.Route("/{list_id}/items", func(r chi.Router) { // api/lists/items/...
				r.Post("/", h.TodoItem.CreateItem)
			})
			r.Route("/items", func(r chi.Router) { // api/lists/items/...
				r.Get("/", h.TodoItem.GetAllItems)
				r.Get("/{item_id}", h.TodoItem.GetItemById)
				r.Put("/{item_id}", h.TodoItem.UpdateItem)
				r.Delete("/{item_id}", h.TodoItem.DeleteItem)
			})

		})
	})

	return r
}
