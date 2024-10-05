package bootstrap

type Di struct {
	Constants *Constants
	Env       *Env
}

func Run() *Di {
	di := &Di{}
	di.Constants = NewConstants()
	di.Env = NewEnvironments()

	return di
}
