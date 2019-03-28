package bussiness

import (
	"star-edge-cloud/edge/store/implemetion/kv"
)

var store = &kv.KVStore{Dir: "data/badger"}
