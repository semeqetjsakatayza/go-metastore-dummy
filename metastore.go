package metastore

import (
	"crypto/rsa"
	"time"
)

// MetaStore keeps meta informations and handles access operations.
type MetaStore struct {
	d map[string]*metaValue
}

// NewMetaStorage create new instance of meta storage.
func NewMetaStorage() (m *MetaStore) {
	m = &MetaStore{
		d: make(map[string]*metaValue),
	}
	return
}

func (m *MetaStore) prepareValue(metaKey string) (v *metaValue) {
	if v = m.d[metaKey]; v == nil {
		v = &metaValue{}
	}
	return
}

// FetchBool get bool value from store.
func (m *MetaStore) FetchBool(metaKey string, defaultValue bool) (value bool, modifyAt int64, err error) {
	v := m.d[metaKey]
	if v == nil {
		return defaultValue, 0, nil
	}
	return v.b, v.modifyAt, nil
}

// StoreBool put bool value into store.
func (m *MetaStore) StoreBool(metaKey string, value bool) (err error) {
	v := m.prepareValue(metaKey)
	v.setBool(value)
	return
}

// FetchInt32 get int32 value from store.
func (m *MetaStore) FetchInt32(metaKey string, defaultValue int32) (value int32, modifyAt int64, err error) {
	v := m.d[metaKey]
	if v == nil {
		return defaultValue, 0, nil
	}
	return int32(v.i64), v.modifyAt, nil
}

// StoreInt32 put int32 value into store.
func (m *MetaStore) StoreInt32(metaKey string, value int32) (err error) {
	v := m.prepareValue(metaKey)
	v.setInt64(int64(value))
	return
}

// FetchInt64 get int64 value from store.
func (m *MetaStore) FetchInt64(metaKey string, defaultValue int64) (value, modifyAt int64, err error) {
	v := m.d[metaKey]
	if v == nil {
		return defaultValue, 0, nil
	}
	return v.i64, v.modifyAt, nil
}

// StoreInt64 put int64 value into store.
func (m *MetaStore) StoreInt64(metaKey string, value int64) (err error) {
	v := m.prepareValue(metaKey)
	v.setInt64(value)
	return
}

// FetchRevision get revision value from store.
// Return 0 if revision record not exists.
func (m *MetaStore) FetchRevision(metaKey string) (revValue int32, modifyAt int64, err error) {
	return m.FetchInt32(metaKey, 0)
}

// StoreRevision save revision record into store.
func (m *MetaStore) StoreRevision(metaKey string, revValue int32) (err error) {
	return m.StoreInt32(metaKey, revValue)
}

// PrepareRSAPrivateKey read RSA private key from storage.
//
// A new private key will be generate if existed key expires.
func (m *MetaStore) PrepareRSAPrivateKey(metaKey string, keyBits int, maxAcceptableAge time.Duration, currentModifyAt int64) (ok bool, priKey *rsa.PrivateKey, modifyAt int64, err error) {
	v := m.prepareValue(metaKey)
	return v.prepareRSAPrivateKey(keyBits, maxAcceptableAge, currentModifyAt)
}
