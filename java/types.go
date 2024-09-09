package java

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type TypeClass struct {
	name string
}

type TypeString struct {
	index int
	data  string
}

type TypeAttribute struct {
	name  string
	bytes []byte
}

type TypeMethod struct {
	name         string
	access_flags []string
	descriptor   string
	attributes   TypeAttribute
}

type TypeException struct {
	start_pc   int
	end_pc     int
	handler_pc int
	catch_type int
}

type CodeAttribute struct {
	name                   string
	attribute_length       int
	max_stack              int
	max_locals             int
	code_length            int
	code                   []byte
	exception_table_length int
	exception_table        []TypeException
	attributes_count       int
	attributes             []TypeAttribute
}

func (tm *TypeMethod) GetCodeAttribute(c *Class) CodeAttribute {
	if tm.attributes.name != "Code" {
		panic("No code attribute on the method")
	}

	ca := CodeAttribute{}

	code := tm.attributes.bytes
	ca.name = tm.attributes.name
	ca.attribute_length = len(tm.attributes.bytes)

	ca.max_stack = ToInt(NextBytes(code, 2, true))
	ca.max_locals = ToInt(NextBytes(code, 2))
	ca.code_length = ToInt(NextBytes(code, 4))
	ca.code = NextBytes(code, ca.code_length)
	ca.exception_table_length = ToInt(NextBytes(code, 2))
	if ca.exception_table_length != 0 {
		panic("Parsing exception is not implemented")
	}
	ca.attributes_count = ToInt(NextBytes(code, 2))

	ca.attributes = c.parseAttributes(c.attributes_count, code)

	return ca
}

func (tm TypeMethod) String() string {
	s := fmt.Sprintf("  %v()\n    descriptor: %v\n    access: %v\n", tm.name, tm.descriptor, tm.access_flags)
	s += "    attribute " + tm.attributes.name + ":\n"
	for _, str := range strings.Split(hex.Dump(tm.attributes.bytes), "\n") {
		s += "     " + str + "\n"
	}
	return strings.Trim(s, "\n ")
}

func (ts TypeString) String() string {
	return fmt.Sprintf("\"%v\" [%v]", ts.data, ts.index)
}

func (tc TypeClass) String() string {
	return fmt.Sprintf("class %v", tc.name)
}
