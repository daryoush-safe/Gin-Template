package bootstrap

type Di struct {
	Constants *Constants
}

func Run() *Di {
	di := &Di{}
	di.Constants = NewConstants()

	return di
}
