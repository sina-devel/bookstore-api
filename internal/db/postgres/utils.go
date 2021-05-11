package postgres

import (
	"errors"
	"gorm.io/gorm"
)

func isErrorNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
