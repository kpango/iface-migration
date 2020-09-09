package main

import (
	"github.com/kpango/iface-migration/old1"
	"github.com/kpango/iface-migration/old2"
	"github.com/kpango/iface-migration/old3"
)

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

func New() (Aggregated, error) {
	a := &aggregated{
		old1.New(),
		old2.New(),
		old3.New(),
	}

	if _, ok :=interface{a}.(old1.Old1){
		return nil, errors.New("err unneccesary method exists")
	}
	if _, ok :=interface{a}.(old2.Old2){
		return nil, errors.New("err unneccesary method exists")
	}
	if _, ok :=interface{a}.(old3.Old3){
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
