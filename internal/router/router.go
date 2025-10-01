package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/isOdin/RestApi/api/handler"
)

func NewRouter(h *handler.Handler) chi.Router {
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
		r.Route("/lists", func(r chi.Router) { // api/lists/...
			r.Post("/", h.TodoList.CreateList)
			r.Get("/", h.TodoList.GetAllLists)
			r.Get("/{id}", h.TodoList.GetListById)
			r.Put("/{id}", h.TodoList.UpdateList)
			r.Delete("/{id}", h.TodoList.DeleteList)

			r.Route("/items", func(r chi.Router) { // api/lists/items/...
				r.Post("/", h.TodoItem.CreateItem)
				r.Get("/", h.TodoItem.GetAllItems)
				r.Get("/{item_id}", h.TodoItem.GetItemById)
				r.Put("/{item_id}", h.TodoItem.UpdateItem)
				r.Delete("/{item_id}", h.TodoItem.DeleteItem)
			})
		})
	})

	return r
}
