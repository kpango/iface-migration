package old2

type Old2 interface {
	Necessary2()
	Unnecessary2()
}

type old struct {
}

func New() Old2 {
	return new(old)
}

func (o *old) Necessary2() {

}

func (o *old) Unnecessary2() {

}
