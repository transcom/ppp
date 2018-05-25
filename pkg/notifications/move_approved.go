package notifications

import (
	"fmt"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"go.uber.org/zap"

	"github.com/transcom/mymove/pkg/auth"
	"github.com/transcom/mymove/pkg/models"
)

// MoveApproved has notification content for approved moves
type MoveApproved struct {
	db      *pop.Connection
	logger  *zap.Logger
	moveID  uuid.UUID
	session *auth.Session // TODO - remove this when we move permissions up to handlers and out of models
}

// NewMoveApproved returns a new move approval notification
func NewMoveApproved(db *pop.Connection,
	logger *zap.Logger,
	session *auth.Session,
	moveID uuid.UUID) *MoveApproved {

	return &MoveApproved{
		db:      db,
		logger:  logger,
		moveID:  moveID,
		session: session,
	}
}

func (m MoveApproved) emails() ([]emailContent, error) {
	var emails []emailContent

	move, err := models.FetchMove(m.db, m.session, m.moveID)
	if err != nil {
		return emails, err
	}

	orders, err := models.FetchOrder(m.db, m.session, move.OrdersID)
	if err != nil {
		return emails, err
	}

	serviceMember, err := models.FetchServiceMember(m.db, m.session, orders.ServiceMemberID)
	if err != nil {
		return emails, err
	}

	if serviceMember.PersonalEmail == nil {
		return emails, fmt.Errorf("no email found for service member")
	}

	smEmail := emailContent{
		recipientEmail: *serviceMember.PersonalEmail,
		subject:        "Move Approved",
		htmlBody:       "Congrats!<br>Your move has been approved!",
		textBody:       "Congrats! Your move has been approved!",
	}

	m.logger.Info("Sent move approval email to service member",
		zap.String("sevice member email address", *serviceMember.PersonalEmail))

	// TODO: Send email to trusted contacts when that's supported
	return append(emails, smEmail), nil
}
