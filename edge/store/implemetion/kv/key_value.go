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
}

// Save - 保存信息
func (kv *KVStore) Save(key string, value []byte) error {
	opts := badger.DefaultOptions
	opts.Dir = kv.Dir
	opts.ValueDir = kv.Dir
	lock.Lock()
	defer lock.Unlock()
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	txn := db.NewTransaction(true)
	defer txn.Discard()
	if err := txn.Set([]byte(key), value); err != nil {
		return err
	}
	_ = txn.Commit()

	return nil
}

// Get - 获取
func (kv *KVStore) Get(key string) []byte {
	opts := badger.DefaultOptions
	opts.Dir = kv.Dir
	opts.ValueDir = kv.Dir
	// opts.ReadOnly = true
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	txn := db.NewTransaction(true)
	defer txn.Discard()
	var valCopy []byte
	db.View(func(txn *badger.Txn) error {
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
func (kv *KVStore) Query(cond string) []byte {
	return nil
}

// Delete - 删除数据
func (kv *KVStore) Delete(key string) int {
	return 0
}
