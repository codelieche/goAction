package webhealth

import "time"

type CheckResponse struct {
	Status  bool      `json:"status"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
