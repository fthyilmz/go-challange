package src

import (
	"math/rand"
	"sync"
	"time"
)

type DB struct {
	lock          sync.RWMutex
	items         map[interface{}]interface{}
	keys          []interface{}
	sliceKeyIndex map[interface{}]int
}

func NewDB() *DB {
	return &DB{
		items:         make(map[interface{}]interface{}),
		sliceKeyIndex: make(map[interface{}]int),
	}
}

func (db *DB) Add(key, item interface{}) {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.items[key] = item

	db.keys = append(db.keys, key)
}

func (db *DB) Update(key, item interface{}) {
	db.lock.Lock()
	defer db.lock.Unlock()

	db.items[key] = item
}

func (db *DB) Get(key interface{}) interface{} {
	db.lock.RLock()
	defer db.lock.RUnlock()

	return db.items[key]
}

func (db *DB) GetFirst() interface{} {
	db.lock.RLock()
	defer db.lock.RUnlock()

	return db.items[db.keys[0]]
}

func (db *DB) GetRandom() interface{} {
	db.lock.RLock()
	defer db.lock.RUnlock()
	rand.Seed(time.Now().UnixNano())

	key := db.keys[rand.Intn(db.Count())]
	return db.items[key]
}

// Count returns len of the queue
func (db *DB) Count() int {
	db.lock.RLock()
	defer db.lock.RUnlock()

	return len(db.items)
}

// IsEmpty returns true if the queue is empty
func (db *DB) IsEmpty() bool {
	return len(db.items) == 0
}
