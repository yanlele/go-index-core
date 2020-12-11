package app

import (
	"gin-example/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Warn(err.Key, err.Message)
	}
	return
}
