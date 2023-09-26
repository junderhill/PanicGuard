package panicguard

import (
	"errors"
	"fmt"
)

func Run[T interface{}](f func() T) (ret T, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			err = errors.New(fmt.Sprint(rec))
		}
	}()
	return f(), nil
}
