package models

import (
    "encoding/json"
    "time"
)

type Data struct {
    Timestamp time.Time `json:"timestamp"
    Value      int       `json:"value"
}
