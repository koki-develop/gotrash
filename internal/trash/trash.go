package trash

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Trash struct {
	Key       string    `json:"-"`
	Path      string    `json:"path"`
	TrashedAt time.Time `json:"trashed_at"`
}

type TrashList []*Trash

func New(p string) *Trash {
	n := time.Now()

	return &Trash{
		Key:       fmt.Sprintf("%d_%s", n.Unix(), uuid.New()),
		Path:      p,
		TrashedAt: n,
	}
}

func MustFromJSON(k string, b []byte) *Trash {
	var t *Trash
	if err := json.Unmarshal(b, &t); err != nil {
		panic(err)
	}

	t.Key = k
	return t
}

// for fuzzy
func (ts TrashList) String(i int) string {
	return ts[i].Path
}

func (ts TrashList) Len() int {
	return len(ts)
}
