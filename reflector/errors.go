package reflector

import "errors"

type ReflectorError error

var(
	fieldAcquireError ReflectorError = errors.New("error acquiring reflected field")
)