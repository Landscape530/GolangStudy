package calculate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegular(t *testing.T) {
	ast := assert.New(t)
	s := "floy switch -env   =  {{.d_env}} {{.app_name}} {{.version}} {{.node}}"
	a := Regular(s)
	target := "floy switch -env={{.d_env}} {{.app_name}} {{.version}} {{.node}}"
	ast.Equal(a, target)
}
