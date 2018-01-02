package template

import "testing"

func Test_template(t *testing.T) {
	o := &template{iTemplate: &obj2{}}
	o.Step()
}
