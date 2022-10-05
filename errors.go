package ego

import (
	"errors"
	"fmt"
)

type GqlError interface {
	ToGqlErrorCode() string
}

type HttpError interface {
	ToHttpStatusCode() int
}

func Wrap(err error, content string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", content, err)
}

func Wrapf(err error, content string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := fmt.Sprintf(content, args...)
	return fmt.Errorf("%s: %w", msg, err)
}

func Cause(err error) error {
	if err == nil {
		return nil
	}
	for err != nil {
		_, ok := err.(interface{ Unwrap() error })
		if !ok {
			return err
		}
		err = errors.Unwrap(err)
	}
	return err
}
