package api

import "github.com/go-chi/chi/v5"

func (api *Api) BindRoutes() {
	api.Router.Use(api.Sessions.LoadAndSave)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.handleSignupUser)
				r.Post("/loginuser", api.handleLoginUser)
				r.With(api.AuthMiddleware).Post("/logoutuser", api.handleLogoutUser)
			})
		})
	})

}
