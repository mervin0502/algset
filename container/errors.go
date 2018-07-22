package container

import (
	"errors"
)

var (
	ErrTypeAssertion error = errors.New("wrong type assertion.")

	//Tensor
	ErrDimensionNotEqual error = errors.New("the dimension of two sub is not equal.")
	ErrDimensionTooHigh  error = errors.New("the dimension of the tensor is too high")

	ErrLengthNotEqual error = errors.New("the length of two sub  is not equal.")
)
