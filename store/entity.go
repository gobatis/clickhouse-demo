package store

import "time"

type Log struct {
	Id        int64     `json:"id"`
	Url       string    `json:"url"`
	Input     string    `json:"input"`
	Output    string    `json:"output"`
	Error     string    `json:"error"`
	CreatedAt time.Time `json:"created_at"`
}
