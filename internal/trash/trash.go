package trashdb

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Trash struct {
	Key       string    `json:"-"`
	Path      string    `json:"path"`
	TrashedAt time.Time `json:"trashed_at"`
}

func New(p string) *Trash {
	n := time.Now()
	return &Trash{
		Key:       fmt.Sprintf("%d_%s", n.Unix(), uuid.New()),
		Path:      p,
		TrashedAt: n,
	}
}
