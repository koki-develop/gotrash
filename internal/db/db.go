package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/koki-develop/gotrash/internal/trash"
	"github.com/koki-develop/gotrash/internal/util"
	"github.com/tidwall/buntdb"
)

const (
	trashDirname = ".gotrash"
	filesDirname = "can"
	dbFilename   = "db"

	shrinkSize = 10 * 1024 * 1024 // 10MB
)

type DB struct {
	trashDir string
	filesDir string
	dbFile   string

	db *buntdb.DB
}

func Open() (*DB, error) {
	h, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	trashDir := filepath.Join(h, trashDirname)
	if err := util.CreateDir(trashDir); err != nil {
		return nil, err
	}

	filesDir := filepath.Join(trashDir, filesDirname)

	dbFile := filepath.Join(trashDir, dbFilename)
	db, err := buntdb.Open(dbFile)
	if err != nil {
		return nil, err
	}

	return &DB{
		trashDir: trashDir,
		filesDir: filesDir,
		dbFile:   dbFile,

		db: db,
	}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Put(ps []string) error {
	if err := util.CreateDir(db.filesDir); err != nil {
		return err
	}

	for _, p := range ps {
		exists, err := util.Exists(p)
		if err != nil {
			return err
		}
		if !exists {
			return fmt.Errorf("%s: no such file or directory", p)
		}

		p, err = filepath.Abs(p)
		if err != nil {
			return err
		}

		t := trash.New(p)
		v, err := json.Marshal(t)
		if err != nil {
			return err
		}

		err = db.db.Update(func(tx *buntdb.Tx) error {
			if _, _, err := tx.Set(t.Key, string(v), nil); err != nil {
				return err
			}

			if err := os.Rename(p, filepath.Join(db.filesDir, t.Key)); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *DB) List() (trash.TrashList, error) {
	var ts trash.TrashList

	err := db.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(k, v string) bool {
			t := trash.MustFromJSON(k, []byte(v))
			ts = append(ts, t)
			return true
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return ts, nil
}

func (db *DB) Restore(is []int, force bool) error {
	maxi := 0
	m := make(map[int]struct{}, len(is))
	for _, i := range is {
		m[i] = struct{}{}
		if maxi < i {
			maxi = i
		}
	}

	var ts trash.TrashList

	err := db.db.View(func(tx *buntdb.Tx) error {
		i := 0
		err := tx.Ascend("", func(k, v string) bool {
			if _, ok := m[i]; ok {
				ts = append(ts, trash.MustFromJSON(k, []byte(v)))
				delete(m, i)
			}

			if i == maxi {
				return false
			}
			i++
			return true
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	if len(m) > 0 {
		is := []int{}
		for i := range m {
			is = append(is, i)
		}
		return fmt.Errorf("%d: no such index item", is)
	}

	for _, t := range ts {
		if !force {
			exists, err := util.Exists(t.Path)
			if err != nil {
				return err
			}
			if exists {
				return fmt.Errorf("%s: already exists", t.Path)
			}
		}

		err = db.db.Update(func(tx *buntdb.Tx) error {
			if _, err := tx.Delete(t.Key); err != nil {
				return err
			}

			if err := os.Rename(filepath.Join(db.filesDir, t.Key), t.Path); err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	if err := db.shrink(false); err != nil {
		return err
	}

	return nil
}

func (db *DB) ClearAll() error {
	err := db.db.Update(func(tx *buntdb.Tx) error {
		if err := tx.DeleteAll(); err != nil {
			return err
		}

		if err := os.RemoveAll(db.filesDir); err != nil {
			return err
		}

		return nil
	})

	if err := db.shrink(true); err != nil {
		return err
	}

	return err
}

func (db *DB) shrink(force bool) error {
	if force {
		if err := db.db.Shrink(); err != nil {
			return err
		}
		return nil
	}

	info, err := os.Stat(db.dbFile)
	if err != nil {
		return err
	}
	if info.Size() > shrinkSize {
		if err := db.db.Shrink(); err != nil {
			return err
		}
	}

	return nil
}
