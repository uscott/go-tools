package optim

import (
	"gonum.org/v1/gonum/optimize"
)

type NM struct {
	*optimize.NelderMead
}

func (nm NM) Init(dim, task int) int {
	return nm.NelderMead.Init(dim, task)
}

func (nm NM) Run(
	operation chan<- optimize.Task, result <-chan optimize.Task, tasks []optimize.Task) {
	nm.NelderMead.Run(operation, result, tasks)
}

func (nm NM) Status() (optimize.Status, error) {
	status, err := nm.NelderMead.Status()
	return status, err
}

func (nm NM) Uses(has optimize.Available) (uses optimize.Available, err error) {
	return nm.NelderMead.Uses(has)
}
