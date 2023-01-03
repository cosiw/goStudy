package api

import "go-server/pkg/db"

type APIs struct {
	db *db.DBHandler
}

func NewAPI(h *db.DBHandler) *APIs {
	return &APIs{db: h}
}
