package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBreakObject(t *testing.T) {
	b := &Break{}
	if b.Type() != BREAK_OBJ {
		t.Errorf("unexpected type. got=%s", b.Type())
	}
	if b.Inspect() != "break" {
		t.Errorf("unexpected inspect string. got=%s", b.Inspect())
	}
}

func TestContinueObject(t *testing.T) {
	c := &Continue{}
	if c.Type() != CONTINUE_OBJ {
		t.Errorf("unexpected type. got=%s", c.Type())
	}
	if c.Inspect() != "continue" {
		t.Errorf("unexpected inspect string. got=%s", c.Inspect())
	}
}
