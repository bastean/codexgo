package smodel

type UseCase[Input any, Output any] interface {
	Run(Input) (Output, error)
}
