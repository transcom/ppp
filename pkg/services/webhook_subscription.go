package services

import (
	"github.com/transcom/mymove/pkg/models"
)

// WebhookSubscriptionFetcher is the service object interface for FetchWebhookSubscription
//go:generate mockery -name WebhookSubscriptionFetcher
type WebhookSubscriptionFetcher interface {
	FetchWebhookSubscription(filters []QueryFilter) (models.WebhookSubscription, error)
}

//WebhookSubscriptionUpdater is the service object interface for UpdateWebhookSubscription
type WebhookSubscriptionUpdater interface {
	UpdateWebhookSubscription(webhooksubscription *models.WebhookSubscription, eTag string) (*models.WebhookSubscription, error)
}