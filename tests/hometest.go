package tests

import "github.com/revel/revel/testing"

type HomeTest struct {
	testing.TestSuite
}

func (t *HomeTest) Before() {
	println("Set up")
}

func (t *HomeTest) TestHomePageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *HomeTest) After() {
	println("Tear down")
}
