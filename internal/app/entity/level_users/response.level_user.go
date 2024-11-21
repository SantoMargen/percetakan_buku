package levelusers

import "time"

type LevelUser struct {
	ID         int       `json:"id"`
	LevelUser  string    `json:"level_user"`
	Keterangan string    `json:"keterangan"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
