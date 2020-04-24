package services

import (
	"fmt"

	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// PreconditionFailedError is the precondition failed error
type PreconditionFailedError struct {
	id uuid.UUID
	error
}

// NewPreconditionFailedError returns an error for a failed precondition
func NewPreconditionFailedError(id uuid.UUID, err error) PreconditionFailedError {
	return PreconditionFailedError{
		id:    id,
		error: err,
	}
}

// Error is the string representation of the precondition failed error
func (e PreconditionFailedError) Error() string {
	return fmt.Sprintf("id: '%s' could not be updated due to the record being stale", e.id.String())
}

//NotFoundError is returned when a given struct is not found
type NotFoundError struct {
	id      uuid.UUID
	message string
}

// NewNotFoundError returns an error for when a struct can not be found
func NewNotFoundError(id uuid.UUID, message string) NotFoundError {
	return NotFoundError{
		id:      id,
		message: message,
	}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("id: %s not found %s", e.id.String(), e.message)
}

//InvalidInputError is returned when an update fails a validation rule
type InvalidInputError struct {
	id               uuid.UUID
	ValidationErrors *validate.Errors
	message          string
	error
}

// NewInvalidInputError returns an error for invalid input
func NewInvalidInputError(id uuid.UUID, err error, validationErrors *validate.Errors, message string) InvalidInputError {
	return InvalidInputError{
		id:               id,
		error:            err,
		ValidationErrors: validationErrors,
		message:          message,
	}
}

func (e InvalidInputError) Error() string {
	if e.message != "" {
		return fmt.Sprintf(e.message)
	}
	return fmt.Sprintf("invalid input for id: %s. %s", e.id.String(), e.ValidationErrors)
}

// CreateObjectError is returned when object creation in the database failed
type CreateObjectError struct {
	objectType       string
	validationErrors *validate.Errors
	message          string
	err              error
}

func (e CreateObjectError) Error() string {
	if e.message != "" {
		return fmt.Sprintf(e.message)
	}
	return fmt.Sprintf("Could not create object of type: %s. %s", e.objectType, e.validationErrors)
}

// NewCreateObjectError returns an error on object creation
func NewCreateObjectError(objectType string, err error, validationErrors *validate.Errors, message string) CreateObjectError {
	return CreateObjectError{
		objectType:       objectType,
		err:              err,
		validationErrors: validationErrors,
		message:          message,
	}
}

// DetailedMsg returns a detailed msg for the logger (not for payload)
func (e CreateObjectError) DetailedMsg() string {
	basicMsg := e.Error()
	if e.validationErrors != nil && e.validationErrors.Count() > 0 {
		return fmt.Sprintf("%s: %s", basicMsg, e.validationErrors)
	} else if e.err != nil {
		return fmt.Sprintf("%s: %s", basicMsg, e.err.Error())
	}
	return basicMsg
}
