package ministry

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals"
	approvalDomain "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/approvals/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/communication"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/domain"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/mappers"
	"github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/ministry/storage"
	domain2 "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/roles/domain"
	usercontracts "github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API/internal/user/contracts"
)

type MinistryService interface {
	All(ctx context.Context) ([]*domain.Ministry, error)
	AssignLeaderToMinistry(ctx context.Context, ministryID string, userID string) error
	SendApplication(ctx context.Context, userID, ministryID string) error
	GetByID(ctx context.Context, ministryID string) (*domain.Ministry, error)
}

type service struct {
	logger               *zap.Logger
	ministryRepo         storage.MinistryRepository
	communicationService communication.CommunicationService
	userService          usercontracts.UserService
	approvalService      approvals.ApprovalService
}

func NewMinistryService(
	logger *zap.Logger,
	ministryRepo storage.MinistryRepository,
	communicationService communication.CommunicationService,
	userService usercontracts.UserService,
	approvalService approvals.ApprovalService,
) MinistryService {
	return &service{
		logger:               logger,
		ministryRepo:         ministryRepo,
		communicationService: communicationService,
		userService:          userService,
		approvalService:      approvalService,
	}
}

func (ms *service) AssignLeaderToMinistry(ctx context.Context, ministryID string, userID string) error {
	err := ms.ministryRepo.AssignLeaderToMinistry(ctx, ministryID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (ms *service) All(ctx context.Context) ([]*domain.Ministry, error) {
	ministryEntities, err := ms.ministryRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all ministries: %s", err)
	}

	var ministries []*domain.Ministry
	for _, entity := range ministryEntities {
		ministries = append(ministries, mappers.ToDomain(entity))
	}

	return ministries, nil
}

func (ms *service) SendApplication(ctx context.Context, userID, ministryID string) error {
	ms.logger.Info("Sending application to Ministry Leader")
	// 1. Fetch Ministry Leader details(phone number & userID) via ministryID
	ministryDetails, err := ms.ministryRepo.GetMinistryByID(ctx, ministryID)
	if err != nil {
		return err
	}

	var leaderID string
	if ministryDetails.LeaderID.Valid {
		leaderID = ministryDetails.LeaderID.String
	}

	leaderDetails, err := ms.userService.GetUserById(ctx, leaderID)
	if err != nil {
		return err
	}

	var requestedRole string

	requestedRole, err = ms.GetRequestedRole(ministryDetails.Name)
	if err != nil {
		return err
	}

	// 2. Create New Approval (Type,requester_id, approver_id,status)
	approval := &approvalDomain.Approval{
		RequesterID:   userID,
		ApproverID:    &leaderDetails.ID,
		RequestedRole: requestedRole,
		Type:          approvalDomain.MinistryApplication,
		Status:        approvalDomain.Pending,
	}

	err = ms.approvalService.CreateNewApproval(ctx, approval)
	if err != nil {
		return err
	}

	var leaderPhoneNumber string
	if leaderDetails.Phone.Valid {
		leaderPhoneNumber = formatUKPhoneNumber(leaderDetails.Phone.String)
	} else {
		return fmt.Errorf(
			"leader(%s) does not have a valid phone number",
			leaderDetails.ID,
		)
	}

	// 4. Construct and send Message to notify ministry leader that an application has been made
	messageText := "You have received a new application for one of your ministries. Login to the website or mobile app for more details."

	// 5. Send SMS
	err = ms.communicationService.SendSMS(leaderPhoneNumber, messageText)
	if err != nil {
		return err
	}

	return nil
}

func (ms *service) GetRequestedRole(ministryName string) (string, error) {
	switch ministryName {
	case domain2.RoleMediaMinistry,
		domain2.RoleChildrensMinistry,
		domain2.RoleCreativeArtsMinistry,
		domain2.RoleMusicMinistry,
		domain2.RoleProductionMinistry,
		domain2.RoleUsheringSecurity:
		return ministryName, nil
	default:
		return "", fmt.Errorf("unknown ministry %s", ministryName)
	}
}

func (ms *service) GetByID(ctx context.Context, ministryID string) (*domain.Ministry, error) {
	ministryEntity, err := ms.ministryRepo.GetMinistryByID(ctx, ministryID)
	if err != nil {
		return nil, err
	}

	ministryDomain := mappers.ToDomain(ministryEntity)

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
