package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	ministry "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/render"
	ministryService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	userService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	middleware2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/middleware"
)

func New(
	logger zap.SugaredLogger,
	userService userService.UserService,
	rolesService roles.RolesService,
	minstryService ministryService.MinistryService,
	jwtSecret string,
) http.Handler {
	// Create a new Chi router.
	router := chi.NewRouter()

	// Add middleware.
	router.Use(middleware.Logger)    // logs every request
	router.Use(middleware.Recoverer) // recovers from panics

	// Define the /alive endpoint.
	registerAliveEndpoint(router)
	router.Route(
		"/api/v1", func(r chi.Router) {
			userHandler := user.NewUserHandler(logger, userService, rolesService)
			ministryHandler := ministry.NewMinistryHandler(logger, minstryService)

			r.Route(
				"/user", func(r chi.Router) {
					r.Post("/register", userHandler.Register())
					r.Post("/login", userHandler.Login())

					// Protected endpoints: wrap these with auth middleware.
					r.Group(func(r chi.Router) {
						r.Use(middleware2.AuthMiddleware([]byte(jwtSecret)))
						r.Post("/update-details", userHandler.UpdateDetails())
						r.Delete("/", userHandler.Delete())
					})
				},
			)
			r.Route(
				"/ministry", func(r chi.Router) {
					r.Get("/", ministryHandler.All())
				},
			)
		},
	)

	return router
}

func registerAliveEndpoint(router *chi.Mux) {
	router.Get("/alive", func(w http.ResponseWriter, r *http.Request) {
		// Return a simple status message.
		render.Json(w, http.StatusOK, "API is alive!")
	})
}
