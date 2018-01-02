package template

import (
	"fmt"
)

type iTemplate interface {
	Step1()
	Step2()
	Step3()
}

type template struct {
	iTemplate
}

func (t *template) Step() {
	t.iTemplate.Step1()
	t.iTemplate.Step2()
	t.iTemplate.Step3()
}

type obj1 struct {
}

func (t *obj1) Step1() {
	fmt.Println("obj1 Step1")
}
func (t *obj1) Step2() {
	fmt.Println("obj1 Step2")
}
func (t *obj1) Step3() {
	fmt.Println("obj1 Step3")
}

type obj2 struct {
	obj1
}

func (t *obj2) Step2() {
	fmt.Println("obj2 Step2")
}
