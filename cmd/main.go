package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // <-- Add this line to register the Postgres driver

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval"
	approvalStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/auth/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/communication"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/config"
	cron2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/cron"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event"
	eventsStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/event/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/http/router"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media"
	mediaStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/media/youtube"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry"
	ministryStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach"
	outreachStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/outreach/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/role"
	rolesStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/role/storage"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
	userStorage "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/storage"
	commonDb "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/db"
	commonlogger "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/pkg/commonlibrary/logger"
)

func main() {
	// Load environment variables from .env file (only in development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	logger := commonlogger.New(cfg)

	// Connect to the PostgreSQL database using sqlx.
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		logger.Sugar().Fatalf("failed to connect to database: %v", err)
	}

	if err := commonDb.RunMigrations(db.DB); err != nil {
		logger.Sugar().Fatalf("failed to run migrations: %v", err)
	}

	userRepo := userStorage.NewUserRepository(db)
	userService := user.NewUserService(logger, userRepo, cfg.JwtSecret)

	rolesRepo := rolesStorage.NewRolesRepository(db)
	rolesService := role.NewRoleService(logger, rolesRepo, nil, nil)

	communicationService := communication.NewCommunicationService(
		cfg.TwilioAccountSID,
		cfg.TwilioAuthToken,
		cfg.TwilioNumber,
	)

	approvalRepository := approvalStorage.NewApprovalRepository(db)
	approvalService := approval.NewApprovalService(logger, approvalRepository, userService, rolesService)

	ministryRepo := ministryStorage.NewMinistryRepository(db) // todo: update all repos to use latest sqlboiler by aaron
	ministryService := ministry.NewMinistryService(
		logger,
		ministryRepo,
		communicationService,
		userService,
		approvalService,
	)

	rolesService.SetApprovalService(approvalService)
	rolesService.SetMinistryService(ministryService)

	authRepo := storage.NewAuthRepository(db)
	authService := auth.NewAuthService(logger, cfg.JwtSecret, authRepo, userService, rolesService)

	outreachRepo := outreachStorage.NewOutreachRepository(db)
	outreachService := outreach.NewOutreachService(logger, outreachRepo)

	mediaRepo := mediaStorage.NewMediaRepository(db)
	mediaService := media.NewMediaService(logger, mediaRepo)

	eventsRepo := eventsStorage.NewEventsRepository(db)
	eventsService := event.NewEventsService(logger, eventsRepo)

	ytClient := youtube.NewYouTubeClient(cfg.YoutubeAPIKey, cfg.YoutubeChannelID)
	mediaCronJob := cron2.NewMediaCronJob(logger, ytClient, mediaService)

	mediaCronJob.RunOnce()

	mediaCronJob.Start()

	mux := router.New(
		logger,
		cfg.JwtSecret,
		userService,
		ministryService,
		outreachService,
		mediaService,
		approvalService,
		eventsService,
		authService,
	)

	logger.Sugar().Infof("Server starting on port %s", cfg.Port)

	addr := fmt.Sprintf(":%s", cfg.Port)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
