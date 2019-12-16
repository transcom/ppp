package services

import (
	"github.com/gofrs/uuid"

	"github.com/transcom/mymove/pkg/models"
)

// MoveTaskOrderFetcher is the service object interface for FetchMoveTaskOrder
//go:generate mockery -name MoveTaskOrderFetcher
type MoveTaskOrderFetcher interface {
	FetchMoveTaskOrder(moveTaskOrderID uuid.UUID) (*models.MoveTaskOrder, error)
}

//MoveTaskOrderStatusUpdater is the service object interface for MoveTaskOrderUpdate
//go:generate mockery -name MoveTaskOrderStatusUpdater
type MoveTaskOrderStatusUpdater interface {
	UpdateMoveTaskOrderStatus(moveTaskOrderID uuid.UUID, isAvailableToPrime bool) (*models.MoveTaskOrder, error)
}