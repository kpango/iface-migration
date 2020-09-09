package old3

type Old3 interface {
	Necessary3()
	Unnecessary3()
}

type old struct {
}

func New() Old3 {
	return new(old)
}

func (o *old) Necessary3() {

}

func (o *old) Unnecessary3() {

}
