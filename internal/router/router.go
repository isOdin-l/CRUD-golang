package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/isOdin/RestApi/api/handler"
)

func NewRouter(listHandler *handler.List, itemHandler *handler.Item, authHandler *handler.Auth) chi.Router {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	// /auth/...
	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", authHandler.SignUpHandler)
		r.Post("/sign-in", authHandler.SignInHandler)
	})

	// api/...
	r.Route("/api", func(r chi.Router) {
		r.Route("/lists", func(r chi.Router) { // api/lists/...
			r.Post("/", listHandler.CreateList)
			r.Get("/", listHandler.GetAllLists)
			r.Get("/{id}", listHandler.GetListById)
			r.Put("/{id}", listHandler.UpdateList)
			r.Delete("/{id}", listHandler.DeleteList)

			r.Route("/items", func(r chi.Router) { // api/lists/items/...
				r.Post("/", itemHandler.CreateItem)
				r.Get("/", itemHandler.GetAllItems)
				r.Get("/{item_id}", itemHandler.GetItemById)
				r.Put("/{item_id}", itemHandler.UpdateItem)
				r.Delete("/{item_id}", itemHandler.DeleteItem)
			})
		})
	})

	return r
}
