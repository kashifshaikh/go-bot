package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/asdine/storm/codec/gob"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/index"
	bolt "go.etcd.io/bbolt"
)

func NewDB(cfg *Config) *storm.DB {
	var err error
	var db *storm.DB
	// if err = os.RemoveAll(cfg.DbFile); err != nil {
	// 	panic(err)
	// }
	if db, err = storm.Open(cfg.DbFile, storm.Codec(gob.Codec)); err != nil {
		panic(err)
	}
	fmt.Println("Successfully opened DB: " + cfg.DbFile)

	return db
}

// type PK struct {
// 	ID uint `json:"id,omitempty" gorm:"primarykey"`
// 	// ID uint `gorm:"primarykey" json:",omitempty"`
// }

// func (pk *PK) BeforeCreate(tx *gorm.DB) (err error) {
// 	pk.ID = uuid.New()
// 	return nil
// }

// Unix Timestamps

type PK uint64

const TIMESTAMP_NOT_SET = -1

type Timestamps struct {
	CreatedAt int64 `json:"-" storm:"index"`
	UpdatedAt int64 `json:"-"`
	DeletedAt int64 `json:"-" storm:"index"`
}

func OpenBolt(path string) *bolt.DB {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type DB struct {
	bolt *bolt.DB
}

func NotFoundError(err error) bool {
	return err == storm.ErrNotFound || err == index.ErrNotFound
}

func (db *DB) Save(bucket string, key string, data *interface{}) error {
	err := db.bolt.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(Key(bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		enc, err := Encode(data)
		if err != nil {
			return fmt.Errorf("could not encode key %s data %v: %s", key, data, err)
		}
		err = b.Put(Key(key), enc)
		return err
	})
	return err
}

func (db *DB) Get(bucket string, key string, data *interface{}) error {

	err := db.bolt.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket(Key(bucket))
		k := []byte(Key(key))
		err = Decode(b.Get(k), data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) List(bucket string, data map[string]interface{}) error {
	err := db.bolt.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(Key(bucket)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			data[string(k)] = Decode
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) ListPrefix(bucket, prefix string) error {
	err := db.bolt.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(Key(bucket)).Cursor()
		p := Key(prefix)
		for k, v := c.Seek(p); bytes.HasPrefix(k, p); k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) ListRange(bucket, start, stop string) error {
	err := db.bolt.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(Key(bucket)).Cursor()
		min := Key(start)
		max := Key(stop)
		for k, v := c.Seek(min); k != nil && bytes.Compare(k, max) <= 0; k, v = c.Next() {
			fmt.Printf("%s: %s\n", k, v)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func Key(key string) []byte {
	return []byte(key)
}

func Encode(data interface{}) ([]byte, error) {
	enc, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func Decode(data []byte, out *interface{}) error {
	err := json.Unmarshal(data, out)
	if err != nil {
		return err
	}
	return nil
}
