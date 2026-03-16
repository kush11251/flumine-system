package utils

import (
    "errors"
)

func ValidateData(data *models.Data) error {
    if data.Timestamp.IsZero() || data.Value == 0 {
        return errors.New("invalid data")
    }
    return nil
}
