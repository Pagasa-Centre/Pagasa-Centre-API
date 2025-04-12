package communication

import (
	"context"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type CommunicationService interface {
	SendSMS(ctx context.Context, to string, message string) error
}

type service struct {
	client     *twilio.RestClient
	fromNumber string
}

func NewCommunicationService(
	TwilioAccountSID string,
	TwilioAuthToken string,
	TwilioFromNumber string,
) CommunicationService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TwilioAccountSID,
		Password: TwilioAuthToken,
	})

	return &service{
		client:     client,
		fromNumber: TwilioFromNumber,
	}
}

func (s *service) SendSMS(ctx context.Context, to string, message string) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(s.fromNumber)
	params.SetBody(message)

	_, err := s.client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("failed to send SMS: %w", err)
	}

	return nil
}
