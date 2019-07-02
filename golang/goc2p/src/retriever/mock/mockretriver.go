package mock

type Retriver struct {
	contents string
}

func (r Retriver) Get(url string) string{
	return r.contents
}
