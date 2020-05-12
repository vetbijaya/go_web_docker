package main

import "testing"

func Test_sum(t *testing.T) {
	if add(5,7) != 12 {
		t.Fail()
	}
	if add(23, 20) !=43{
		t.Fail()
	}
	if add(200, 321) !=521{
		t.Fail()
	}
}
func Test_product(t *testing.T){
	if multiply(12, 10) !=120{
		t.Fail()
	}
	if multiply(12, 7) != 84{
		t.Fail()
	}
	if multiply(11, 10) != 110{
		t.Fail()
	}
}
	