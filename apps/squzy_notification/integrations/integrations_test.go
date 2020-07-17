package integrations

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	api "github.com/squzy/squzy_generated/generated/proto/v1"
	"github.com/stretchr/testify/assert"
	"net/http"
	"squzy/apps/squzy_notification/database"
	"testing"
	"time"
)

type mock struct {

}

type mockError struct {

}

func (m mockError) SendRequest(req *http.Request) (int, []byte, error) {
	return 0, nil, errors.New("")
}

func (m mockError) SendRequestTimeout(req *http.Request, timeout time.Duration) (int, []byte, error) {
	panic("implement me")
}

func (m mockError) SendRequestWithStatusCode(req *http.Request, expectedCode int) (int, []byte, error) {
	panic("implement me")
}

func (m mockError) SendRequestTimeoutStatusCode(req *http.Request, timeout time.Duration, expectedCode int) (int, []byte, error) {
	panic("implement me")
}

func (m mockError) CreateRequest(method string, url string, headers *map[string]string, schedulerID string) *http.Request {
	panic("implement me")
}

func (m mock) SendRequest(req *http.Request) (int, []byte, error) {
	return 0, []byte{}, nil
}

func (m mock) SendRequestTimeout(req *http.Request, timeout time.Duration) (int, []byte, error) {
	panic("implement me")
}

func (m mock) SendRequestWithStatusCode(req *http.Request, expectedCode int) (int, []byte, error) {
	panic("implement me")
}

func (m mock) SendRequestTimeoutStatusCode(req *http.Request, timeout time.Duration, expectedCode int) (int, []byte, error) {
	panic("implement me")
}

func (m mock) CreateRequest(method string, url string, headers *map[string]string, schedulerID string) *http.Request {
	panic("implement me")
}

func TestNew(t *testing.T) {
	t.Run("Shuld: not be nil", func(t *testing.T) {
		s := New(nil)
		assert.NotNil(t, s)
	})
}

func TestIntegrations_Webhook(t *testing.T) {
	t.Run("Should: not throw panic", func(t *testing.T) {
		s := New(&mock{})
		assert.NotPanics(t, func() {
			s.Webhook(context.Background(), &api.Incident{
				Histories: []*api.Incident_HistoryItem{
					{
						Timestamp: ptypes.TimestampNow(),
						Status: api.IncidentStatus_INCIDENT_STATUS_CAN_BE_CLOSED,
					},
				},
			}, &database.WebHookConfig{})
		})
	})
	t.Run("Should: not throw panic", func(t *testing.T) {
		s := New(&mock{})
		assert.NotPanics(t, func() {
			s.Webhook(context.Background(), &api.Incident{
				Histories: []*api.Incident_HistoryItem{
					{
						Timestamp: ptypes.TimestampNow(),
						Status: 123,
					},
				},
			}, &database.WebHookConfig{})
		})
	})
	t.Run("Should: not throw panic", func(t *testing.T) {
		s := New(&mockError{})
		assert.NotPanics(t, func() {
			s.Webhook(context.Background(), &api.Incident{
				Histories: []*api.Incident_HistoryItem{
					{
						Timestamp: ptypes.TimestampNow(),
						Status: api.IncidentStatus_INCIDENT_STATUS_CAN_BE_CLOSED,
					},
				},
			}, &database.WebHookConfig{})
		})
	})
}

func TestIntegrations_Slack(t *testing.T) {
	t.Run("Should: not throw panic", func(t *testing.T) {
		s := New(&mock{})
		assert.NotPanics(t, func() {
			s.Slack(context.Background(), &api.Incident{
				Histories: []*api.Incident_HistoryItem{
					{
						Timestamp: ptypes.TimestampNow(),
						Status: api.IncidentStatus_INCIDENT_STATUS_CAN_BE_CLOSED,
					},
				},
			}, &database.SlackConfig{})
		})
	})
}