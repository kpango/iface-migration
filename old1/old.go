package old1

type Old1 interface {
	Necessary1()
	Unnecessary1()
}

type old struct {
}

func New() Old1 {
	return new(old)
}

func (o *old) Necessary1() {

}

func (o *old) Unnecessary1() {

}
