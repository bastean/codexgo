package model

type UseCase[Input any, Output any] interface {
	Run(Input) (Output, error)
}
