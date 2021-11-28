package sandbox3

import (
	"testing"
)

func TestStructPropCannotBeNil(t *testing.T) {
	s := specimen{}
	s.add1.street1 = "a"
	//s.add1 = nil 	<-- if you uncomment this line you will get compiler error
}

type specimen struct {
	add1 address
}

type address struct {
	street1 string
	street2 string
	street3 string
}
