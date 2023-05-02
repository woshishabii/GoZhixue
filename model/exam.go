package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"

	"gorm.io/gorm"
)

type Exam struct {
	gorm.Model
	Title   string `json:"title"`
	StartAt time.Time
	EndAt   time.Time
	Content datatypes.JSON `json:"content"`
	Answer  datatypes.JSON `json:"answer"`
	Taker   []Class        `json:"taker"`
	UUID    uuid.UUID      `json:"uuid"`
}
