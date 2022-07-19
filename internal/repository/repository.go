package repository

import "github.com/becunningham/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(models.Reservation) error
}
