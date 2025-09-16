package ministry

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval"
	approvalDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approval/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/communication"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/mapper"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/storage"
	userService "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user"
)

type Service interface {
	All(ctx context.Context) ([]*domain.Ministry, error)
	SendApplication(ctx context.Context, userID, ministryID, reason string) error
	GetByID(ctx context.Context, ministryID string) (*domain.Ministry, error)
}

type service struct {
	logger               *zap.Logger
	ministryRepo         storage.MinistryRepository
	communicationService communication.CommunicationService
	userService          userService.Service
	approvalService      approval.Service
}

func NewMinistryService(
	logger *zap.Logger,
	ministryRepo storage.MinistryRepository,
	communicationService communication.CommunicationService,
	userService userService.Service,
	approvalService approval.Service,
) Service {
	return &service{
		logger:               logger,
		ministryRepo:         ministryRepo,
		communicationService: communicationService,
		userService:          userService,
		approvalService:      approvalService,
	}
}

var ErrMinistryNotFound = errors.New("ministry not found")

func (ms *service) All(ctx context.Context) ([]*domain.Ministry, error) {
	ministryEntities, err := ms.ministryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all ministries: %s", err)
	}

	var ministries []*domain.Ministry

	for _, entity := range ministryEntities {
		ministryLeadersDetails, err := ms.ministryRepo.GetMinistryLeaderUsersByMinistryID(ctx, entity.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get ministry leader users: %s", err)
		}

		var ministryLeaderNames []string

		for _, ministryLeader := range ministryLeadersDetails {
			//todo: turn this into a ministry struct method
			name := fmt.Sprintf("%s %s", ministryLeader.FirstName, ministryLeader.LastName)
			ministryLeaderNames = append(ministryLeaderNames, name)
		}

		ministryActivites, err := ms.ministryRepo.GetMinistryActivitiesByMinistryID(ctx, entity.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get ministry activities: %s", err)
		}

		var activities []string
		for _, ministryActivity := range ministryActivites {
			activities = append(activities, ministryActivity.Name)
		}

		ministries = append(ministries, mapper.ToDomain(entity, ministryLeaderNames, activities))
	}

	return ministries, nil
}

func (ms *service) SendApplication(ctx context.Context, userID, ministryID, reason string) error { //todo:create domain object
	ms.logger.With(
		zap.String("userID", userID),
		zap.String("ministryID", ministryID)).Info("Sending application to Ministry Leader")

	ministryDetails, err := ms.ministryRepo.GetMinistryByID(ctx, ministryID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrMinistryNotFound // todo: move to entity layer
		}

		return err
	}

	// todo: split into fetch ministry and then call user service to get user
	ministryLeadersDetails, err := ms.ministryRepo.GetMinistryLeaderUsersByMinistryID(ctx, ministryID)
	if err != nil {
		return err
	}

	var leadersPhoneNumbers []string
	for _, user := range ministryLeadersDetails {
		leadersPhoneNumbers = append(leadersPhoneNumbers, formatUKPhoneNumber(user.Phone.String))
	}

	roleName := fmt.Sprintf("%s Member", ministryDetails.Name)

	err = ms.approvalService.CreateNewApproval(
		ctx,
		&approvalDomain.Approval{
			RequesterID:   userID,
			RequestedRole: roleName,
			Type:          string(approvalDomain.MinistryApplication),
			Reason:        reason,
			Status:        string(approvalDomain.Pending),
		})
	if err != nil {
		return fmt.Errorf("failed to create new approval: %s", err)
	}

	// 4. Construct and send Message to notify ministry leaders that an application has been made
	messageText := "You have received a new application for one of your ministries. Login to the website or mobile app for more details."

	// 5. Send SMS
	for _, phoneNumber := range leadersPhoneNumbers {
		err = ms.communicationService.SendSMS(phoneNumber, messageText)
		ms.logger.Error("Failed to send SMS to ministry leader",
			zap.String("phone", phoneNumber),
			zap.Error(err),
		)
	}

	return nil
}

func (ms *service) GetByID(ctx context.Context, ministryID string) (*domain.Ministry, error) {
	ministryEntity, err := ms.ministryRepo.GetMinistryByID(ctx, ministryID)
	if err != nil {
		return nil, err
	}

	ministryLeadersDetails, err := ms.ministryRepo.GetMinistryLeaderUsersByMinistryID(ctx, ministryEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get ministry leader users: %s", err)
	}

	var ministryLeaderNames []string

	for _, ministryLeader := range ministryLeadersDetails {
		name := fmt.Sprintf("%s %s", ministryLeader.FirstName, ministryLeader.LastName)
		ministryLeaderNames = append(ministryLeaderNames, name)
	}

	ministryActivites, err := ms.ministryRepo.GetMinistryActivitiesByMinistryID(ctx, ministryEntity.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get ministry activities: %s", err)
	}

	var activities []string
	for _, ministryActivity := range ministryActivites {
		activities = append(activities, ministryActivity.Name)
	}

	ministryDomain := mapper.ToDomain(ministryEntity, ministryLeaderNames, activities)

	return ministryDomain, nil
}

func formatUKPhoneNumber(number string) string {
	if len(number) == 0 {
		return number
	}

	if number[0] == '0' {
		return "+44" + number[1:]
	}

	if number[0] != '+' {
		return "+44" + number // fallback just in case it's missing both 0 and +
	}

	return number // already in E.164
}
