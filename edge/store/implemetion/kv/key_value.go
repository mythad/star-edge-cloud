package kv

import (
	"log"
	"sync"

	"github.com/dgraph-io/badger"
)

var lock sync.Mutex

// KVStore -- 主要用于服务数据或扩展数据的存储
type KVStore struct {
	Dir string
	db  *badger.DB
}

// Save - 保存信息
func (it *KVStore) Save(key string, value []byte) error {
	txn := it.db.NewTransaction(true)
	defer txn.Discard()
	if err := txn.Set([]byte(key), value); err != nil {
		return err
	}
	_ = txn.Commit()

	return nil
}

// Get - 获取
func (it *KVStore) Get(key string) []byte {
	txn := it.db.NewTransaction(true)
	defer txn.Discard()
	var valCopy []byte
	it.db.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte(key))

		item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)
			return nil
		})
		return nil
	})
	return valCopy
}

// Query - 查询数据
func (it *KVStore) Query(cond string) []byte {
	return nil
}

// Delete - 删除数据
func (it *KVStore) Delete(key string) int {
	return 0
}

// Initialize -
func (it *KVStore) Initialize() {
	opts := badger.DefaultOptions
	opts.Dir = it.Dir
	opts.ValueDir = it.Dir
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	it.db = db
}

// Dispose -
func (it *KVStore) Dispose(key string) {
	it.db.Close()
}
