package errs

import (
	"fmt"
	"strings"
)

type Field struct {
	Key   string
	Value interface{}
}

type AppError struct {
	Code    string
	Message string
	Err     error
	Fields  []Field
}

func (e *AppError) Error() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("code=%s message=%s", e.Code, e.Message))
	for _, f := range e.Fields {
		b.WriteString(fmt.Sprintf(" %s=%v", f.Key, f.Value))
	}
	if e.Err != nil {
		b.WriteString(fmt.Sprintf(" | cause=%v", e.Err))
	}
	return b.String()
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// New creates an AppError without the "original" error, but with any number of fields
func New(code, message string, args ...interface{}) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Fields:  parseFields(args...),
	}
}

// Wrap creates an AppError by wrapping an existing error and adding fields
func Wrap(err error, code, message string, args ...interface{}) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
		Fields:  parseFields(args...),
	}
}

// parseFields is waiting for an even number of args: key, value, key, value, ...
// If odd, the last key drops without value
func parseFields(args ...interface{}) []Field {
	var fs []Field
	for i := 0; i+1 < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		fs = append(fs, Field{Key: key, Value: args[i+1]})
	}
	return fs
}
