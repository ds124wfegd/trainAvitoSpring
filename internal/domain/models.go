package domain

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type PVZ struct {
	ID               string    `json:"id"`
	RegistrationDate time.Time `json:"registration_date"`
	City             string    `json:"city"`
}

type Reception struct {
	ID       string    `json:"id"`
	DateTime time.Time `json:"date_time"`
	PVZID    string    `json:"pvz_id"`
	Status   string    `json:"status"`
}

type Product struct {
	ID          string    `json:"id"`
	DateTime    time.Time `json:"date_time"`
	Type        string    `json:"type"`
	ReceptionID string    `json:"reception_id"`
}

type PVZWithReceptions struct {
	PVZ        PVZ                     `json:"pvz"`
	Receptions []ReceptionWithProducts `json:"receptions"`
}

type ReceptionWithProducts struct {
	Reception Reception `json:"reception"`
	Products  []Product `json:"products"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
