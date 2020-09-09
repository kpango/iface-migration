package main

import (
	"errors"

	"github.com/kpango/iface-migration/old1"
	"github.com/kpango/iface-migration/old2"
	"github.com/kpango/iface-migration/old3"
)

// 新規インターフェースは適切なもののみを
type Aggregated interface {
	Necessary1()
	Necessary2()
	Necessary3()
}

type aggregated struct {
	old1 old1.Old1
	old2 old2.Old2
	old3 old3.Old3
}

// これはコンパイル時にチェックされます
var _ Aggregated = (*aggregated)(nil)

func New() (Aggregated, error) {
	a := &aggregated{
		old1.New(),
		old2.New(),
		old3.New(),
	}

	// まとめるものが多い場合大変だが、テストコードに書くことでより実装をシンプルに保てる
	var i interface{} = a
	_, ok := i.(old1.Old1)
	if ok {
		return nil, errors.New("err unneccesary method exists")
	}
	_, ok = i.(old2.Old2)
	if ok {
		return nil, errors.New("err unneccesary method exists")
	}
	_, ok = i.(old3.Old3)
	if ok {
		return nil, errors.New("err unneccesary method exists")
	}

	return a, nil
}

func (a aggregated) Necessary1() {
	a.old1.Necessary1()
}

func (a aggregated) Necessary2() {
	a.old2.Necessary2()
}

func (a aggregated) Necessary3() {
	a.old3.Necessary3()
}

// コメントアウトすればーエラーではなくなる
func (a aggregated) Unnecessary1() {
	a.old1.Unnecessary1()
}
