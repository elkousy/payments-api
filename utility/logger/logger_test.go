//+build !integration

package logger

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_newLogger(t *testing.T) {
	//Assert
	assert.Equal(t, reflect.TypeOf(LogStdOut), reflect.TypeOf(&zap.SugaredLogger{}))
	assert.Equal(t, reflect.TypeOf(LogStdErr), reflect.TypeOf(&zap.SugaredLogger{}))
}
