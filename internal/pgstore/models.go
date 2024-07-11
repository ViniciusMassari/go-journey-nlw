// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package pgstore

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Activity struct {
	ID       uuid.UUID
	TripID   uuid.UUID
	Title    string
	OccursAt pgtype.Timestamp
}

type Link struct {
	ID     uuid.UUID
	TripID uuid.UUID
	Title  string
	Url    string
}

type Participant struct {
	ID          uuid.UUID
	TripID      uuid.UUID
	Email       string
	IsConfirmed bool
}

type Trip struct {
	ID          uuid.UUID
	Destination string
	OwnerEmail  string
	OwnerName   string
	IsConfirmed bool
	StartsAt    pgtype.Timestamp
	EndsAt      pgtype.Timestamp
}