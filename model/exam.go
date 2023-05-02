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
	Takers  []Class        `json:"taker" gorm:"foreignKey:ExamID"`
	UUID    uuid.UUID      `json:"uuid"`
}
