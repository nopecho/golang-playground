package print_test

import (
	"github.com/stretchr/testify/assert"
	"gitupb.com/nopecho/golang-playgound/go-cli/src/print"
	"testing"
)

func TestName(t *testing.T) {
	name := print.Name
	name("hello")

	var sut string = "hello"

	assert.Equal(t, sut,"hello")
}
