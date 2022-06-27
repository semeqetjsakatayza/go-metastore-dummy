package metastore

import (
	"crypto/rand"
	"crypto/rsa"
	"time"
)

func (val *metaValue) prepareRSAPrivateKey(keyBits int, maxAcceptableAge time.Duration, currentModifyAt int64) (ok bool, priKey *rsa.PrivateKey, modifyAt int64, err error) {
	modifyBoundAt := time.Now().Unix() - int64(maxAcceptableAge/time.Second)
	if modifyBoundAt < val.modifyAt {
		modifyAt = val.modifyAt
		if currentModifyAt == val.modifyAt {
			return
		}
		if val.rsaPrivateKey != nil {
			ok = true
			priKey = val.rsaPrivateKey
			return
		}
	}
	val.rsaPrivateKey, err = rsa.GenerateKey(rand.Reader, keyBits)
	if nil != err {
		return
	}
	val.updateModifyAt()
	ok = true
	priKey = val.rsaPrivateKey
	modifyAt = val.modifyAt
	return
}
