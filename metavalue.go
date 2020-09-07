package metastore

import (
	"strconv"
	"time"
)

type metaValue struct {
	b    bool
	i64  int64
	text string

	modifyAt int64
}

func (val *metaValue) updateModifyAt() {
	val.modifyAt = time.Now().Unix()
}

func (val *metaValue) setBool(v bool) {
	if v {
		val.b = true
		val.i64 = 1
		val.text = "1"
	} else {
		val.b = false
		val.i64 = 0
		val.text = "0"
	}
	val.updateModifyAt()
}

func (val *metaValue) setInt64(v int64) {
	if v == 0 {
		val.b = false
	} else {
		val.b = true
	}
	val.i64 = v
	val.text = strconv.FormatInt(v, 10)
	val.updateModifyAt()
}

func (val *metaValue) setText(v string) {
	if (v == "") || (v == "0") || (v == "false") {
		val.b = false
	} else {
		val.b = true
	}
	if v64, err := strconv.ParseInt(v, 10, 64); nil != err {
		val.i64 = 0
	} else {
		val.i64 = v64
	}
	val.text = v
	val.updateModifyAt()
}
