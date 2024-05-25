package logic

import "fmt"

type Foo struct {
	flag bool
}

func NewFoo(flag bool) Foo {
	return Foo{
		flag: flag,
	}
}

func (f Foo) Rise() {
	fmt.Println(f.flag)
}