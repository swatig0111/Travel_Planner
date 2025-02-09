package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Trip struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Destination string             `bson:"destination" json:"destination"`
	StartDate   time.Time          `bson:"start_date" json:"start_date"`
	EndDate     time.Time          `bson:"end_date" json:"end_date"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
