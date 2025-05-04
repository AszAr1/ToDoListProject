package database

import (
	"errors"
	"fmt"
)

func NotFoundError(op string) error {
	baseMessage := fmt.Sprintf("Operation %s got not found db-error. Check your input data.", op)
	return errors.New(baseMessage)
}

func CreateFailedError(op string, details string) error {
	baseMessage := fmt.Sprintf("Operation %s got create db-error", op)

	if details != "" {
		baseMessage = fmt.Sprintf("%s. Details: %s", baseMessage, details)
	}

	return errors.New(baseMessage)
}

func DeleteFailedError(op string, details string) error {
	baseMessage := fmt.Sprintf("Operation %s got delete db-error", op)

	if details != "" {
		baseMessage = fmt.Sprintf("%s. Details: %s", baseMessage, details)
	}

	return errors.New(baseMessage)
}

func UpdateFailedError(op string, details string) error {
	baseMessage := fmt.Sprintf("Operation %s got update db-error", op)

	if details != "" {
		baseMessage = fmt.Sprintf("%s. Details: %s", baseMessage, details)
	}

	return errors.New(baseMessage)
}

func ReadFailedError(op string, details string) error {
	baseMessage := fmt.Sprintf("Operation %s got read db-error", op)

	if details != "" {
		baseMessage = fmt.Sprintf("%s. Details: %s", baseMessage, details)
	}

	return errors.New(baseMessage)
}
