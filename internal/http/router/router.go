package router

import (
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/media"
	ministry "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/outreach"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user"
	mediaService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media"
	ministryService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	outreachService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles"
	userService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	middleware2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/middleware"
)

func New(
	logger *zap.Logger,
	jwtSecret string,
	userService userService.UserService,
	rolesService roles.RolesService,
	minstryService ministryService.MinistryService,
	outreachService outreachService.OutreachService,
	mediaService mediaService.MediaService,
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
			userHandler := user.NewUserHandler(logger, userService, rolesService, minstryService)
			ministryHandler := ministry.NewMinistryHandler(logger, minstryService)
			outreachHandler := outreach.NewOutreachHandler(logger, outreachService)
			mediaHandler := media.NewMediaHandler(logger, mediaService)

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
			r.Route(
				"/outreach", func(r chi.Router) {
					r.Get("/", outreachHandler.All())
				},
			)
			r.Route(
				"/media", func(r chi.Router) {
					r.Get("/", mediaHandler.All())
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
