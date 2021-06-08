package redis_orm

import "errors"

type ORMError error

var (
	AddressFormatError ORMError = errors.New("Address has wrong format")
	AddressLengthError ORMError = errors.New("Address has wrong length")
	UnimplementedError ORMError = errors.New("unimplemented")
)