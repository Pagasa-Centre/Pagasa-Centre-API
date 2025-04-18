package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/approvals"
	events2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/events"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/media"
	ministry "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/ministry"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/outreach"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/api/user"
	approvals2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/events"
	mediaService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media"
	ministryService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	outreachService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach"
	userService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	middleware2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/middleware"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/render"
)

func New(
	logger *zap.Logger,
	jwtSecret string,
	userService userService.UserService,
	ministryService ministryService.MinistryService,
	outreachService outreachService.OutreachService,
	mediaService mediaService.MediaService,
	approvalService approvals2.ApprovalService,
	eventsService events.EventsService,
) http.Handler {
	// Create a new Chi router.
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // 👈 allow frontend origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Add middleware.
	router.Use(middleware.Logger)    // logs every request
	router.Use(middleware.Recoverer) // recovers from panics

	userHandler := user.NewUserHandler(logger, userService)
	ministryHandler := ministry.NewMinistryHandler(logger, ministryService)
	outreachHandler := outreach.NewOutreachHandler(logger, outreachService)
	mediaHandler := media.NewMediaHandler(logger, mediaService)
	approvalsHandler := approvals.NewApprovalHandler(logger, approvalService)
	eventsHandler := events2.NewEventsHandler(logger, eventsService)

	// Define the /alive endpoint.
	registerAliveEndpoint(router)
	router.Route(
		"/api/v1", func(r chi.Router) {
			r.Route(
				"/user", func(r chi.Router) {
					r.Post("/register", userHandler.Register())
					r.Post("/login", userHandler.Login())

					// Protected endpoints: wrap these with auth middleware.
					r.Group(func(r chi.Router) {
						r.Use(middleware2.AuthMiddlewareString([]byte(jwtSecret)))
						r.Post("/update-details", userHandler.UpdateDetails())
						r.Delete("/", userHandler.Delete())
						r.Get("/approvals", approvalsHandler.All())
						r.Post("/approvals/{id}", approvalsHandler.UpdateApprovalStatus())
					})
				},
			)
			r.Route(
				"/ministry", func(r chi.Router) {
					r.Get("/", ministryHandler.All())

					r.Group(func(r chi.Router) {
						r.Use(middleware2.AuthMiddlewareString([]byte(jwtSecret)))
						r.Post("/application", ministryHandler.Apply())
					})
				},
			)
			r.Route(
				"/outreach", func(r chi.Router) {
					r.Get("/", outreachHandler.All())

					r.Route(
						"/events", func(r chi.Router) {
							r.Get("/", eventsHandler.All())
						},
					)
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
