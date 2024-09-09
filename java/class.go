package java

import (
	"fmt"
	"log"
)

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type Class struct {
	magic               int
	major               int
	minor               int
	constant_pool_count int
	constant_pool       []CpInfo
	access_flags        int
	this_class          int
	super_class         int
	interfaces_count    int
	interfaces          []Interface
	fields_count        int
	fields              []FieldInfo
	methods_count       int
	methods             []MethodInfo
	attributes_count    int
	attributes          []AttributeInfo
}

func (c *Class) ParseFromBytes(bytes []byte) {
	c.magic = ToInt(NextBytes(bytes, 4, true))
	c.minor = ToInt(NextBytes(bytes, 2))
	c.major = ToInt(NextBytes(bytes, 2))
	c.constant_pool_count = ToInt(NextBytes(bytes, 2))

	c.constant_pool = make([]CpInfo, 0, c.constant_pool_count)
	for i := 0; i < c.constant_pool_count-1; i++ {
		tag := NextBytes(bytes, 1)

		cp_info := make(CpInfo)
		cp_info["tag"] = tag

		switch ToInt(cp_info["tag"]) {
		case CONSTANT_Class:
			cp_info["name_index"] = NextBytes(bytes, 2)
		case CONSTANT_Fieldref:
			cp_info["class_index"] = NextBytes(bytes, 2)
			cp_info["name_and_type_index"] = NextBytes(bytes, 2)
		case CONSTANT_Methodref:
			cp_info["class_index"] = NextBytes(bytes, 2)
			cp_info["name_and_type_index"] = NextBytes(bytes, 2)
		case CONSTANT_InterfaceMethodref:
			panic("Constant type not implemented yet")
		case CONSTANT_String:
			cp_info["string_index"] = NextBytes(bytes, 2)
		case CONSTANT_Integer:
			panic("Constant type not implemented yet")
		case CONSTANT_Float:
			panic("Constant type not implemented yet")
		case CONSTANT_Long:
			panic("Constant type not implemented yet")
		case CONSTANT_Double:
			panic("Constant type not implemented yet")
		case CONSTANT_NameAndType:
			cp_info["name_index"] = NextBytes(bytes, 2)
			cp_info["descriptor_index"] = NextBytes(bytes, 2)
		case CONSTANT_Utf8:
			cp_info["length"] = NextBytes(bytes, 2)
			cp_info["bytes"] = NextBytes(bytes, ToInt(cp_info["length"]))
		case CONSTANT_MethodHandle:
			panic("Constant type not implemented yet")
		case CONSTANT_MethodType:
			panic("Constant type not implemented yet")
		case CONSTANT_InvokeDynamic:
			panic("Constant type not implemented yet")
		default:
			log.Fatal("ERROR: Could not parse constant_pool tag")
		}
		c.constant_pool = append(c.constant_pool, cp_info)
	}

	c.access_flags = ToInt(NextBytes(bytes, 2))
	c.this_class = ToInt(NextBytes(bytes, 2))
	c.super_class = ToInt(NextBytes(bytes, 2))
	c.interfaces_count = ToInt(NextBytes(bytes, 2))
	if c.interfaces_count != 0 {
		panic("Parsing interfaces not implemented yet")
	}
	c.fields_count = ToInt(NextBytes(bytes, 2))
	if c.fields_count != 0 {
		panic("Parsing interfaces not implemented yet")
	}
	c.methods_count = ToInt(NextBytes(bytes, 2))
	for i := 0; i < c.methods_count; i++ {
		method := make(MethodInfo)
		method["access_flags"] = NextBytes(bytes, 2)
		method["name_index"] = NextBytes(bytes, 2)
		method["descriptor_index"] = NextBytes(bytes, 2)
		method["attributes_count"] = NextBytes(bytes, 2)
		for i := 0; i < ToInt(method["attributes_count"]); i++ {
			method["attributes_info-attribute_name_index"] = NextBytes(bytes, 2)
			attribute_length := NextBytes(bytes, 4)
			method["attributes_info-attribute_length"] = attribute_length
			method["attributes_info-info"] = NextBytes(bytes, ToInt(attribute_length))
		}

		c.methods = append(c.methods, method)
	}
	c.attributes_count = ToInt(NextBytes(bytes, 2))
	for i := 0; i < c.attributes_count; i++ {
		attribute := make(AttributeInfo)
		attribute["attribute_name_index"] = NextBytes(bytes, 2)
		attribute_length := NextBytes(bytes, 4)
		attribute["attribute_length"] = attribute_length
		attribute["info"] = NextBytes(bytes, ToInt(attribute_length))

		c.attributes = append(c.attributes, attribute)
	}
}

func (c Class) String() string {
	s := fmt.Sprintf("magic = 0x%x\n", c.magic)
	s += fmt.Sprintf("version = %v.%v\n", c.major, c.minor)
	s += fmt.Sprintf("const_pool_count = %v\n", c.constant_pool_count)
	// s += fmt.Sprintf("const_pool: %v\n", c.constant_pool_count)
	// for index, constant := range c.constant_pool {
	// 	s += fmt.Sprintf("%v: %v", index+1, constant)
	// }
	s += fmt.Sprintf("access_flags = 0b%b\n", c.access_flags)
	s += fmt.Sprintf("this_class = %v\n", c.this_class)
	s += fmt.Sprintf("super_class = %v\n", c.super_class)
	s += fmt.Sprintf("interfaces_count = %v\n", c.interfaces_count)
	s += fmt.Sprintf("interfaces = %v\n", c.interfaces)
	s += fmt.Sprintf("fields_count = %v\n", c.fields_count)
	s += fmt.Sprintf("fields = %v\n", c.fields)
	s += fmt.Sprintf("methods_count = %v\n", c.methods_count)
	s += fmt.Sprintf("methods = %v\n", c.methods)
	s += fmt.Sprintf("attributes_count = %v\n", c.attributes_count)
	s += fmt.Sprintf("attributes = %v\n", c.attributes)

	return s[:len(s)-1]
}

func (c *Class) expectPoolIndexType(i int, t int) {
	v := ToInt(c.constant_pool[i-1]["tag"])

	if v != t {
		panic(fmt.Sprintf("ERROR: Type assertion failed: Expected %v but got %v at index %v", t, v, i))
	}
}

func (c *Class) filterTypeFromPool(t int) ([]CpInfo, []int) {
	collected := make([]CpInfo, 0)
	indexes := make([]int, 0)

	for index, constant := range c.constant_pool {
		if ToInt(constant["tag"]) == t {
			collected = append(collected, constant)
			indexes = append(indexes, index)
		}
	}

	return collected, indexes
}

func (c *Class) Classes() []TypeClass {
	classes := make([]TypeClass, 0)
	raw_classes, _ := c.filterTypeFromPool(CONSTANT_Class)

	for _, raw_class := range raw_classes {
		name_index := ToInt(raw_class["name_index"])
		c.expectPoolIndexType(name_index, CONSTANT_Utf8)
		raw_utf8 := c.constant_pool[name_index-1]
		name := string(raw_utf8["bytes"])

		classes = append(classes, TypeClass{
			name: name,
		})
	}

	return classes
}

func (c *Class) Flags() []string {
	s := make([]string, 0)
	for _, v := range CLASSFILE_ACCESS_FLAGS {
		if c.access_flags&v.value != 0 {
			s = append(s, v.name)
		}
	}
	return s
}

func (c *Class) Strings() []TypeString {
	strings := make([]TypeString, 0)
	raw_strings, indexes := c.filterTypeFromPool(CONSTANT_String)

	for i, raw_string := range raw_strings {
		data_index := ToInt(raw_string["string_index"])
		c.expectPoolIndexType(data_index, CONSTANT_Utf8)
		raw_utf8 := c.constant_pool[data_index-1]
		data := string(raw_utf8["bytes"])

		strings = append(strings, TypeString{
			index: indexes[i],
			data:  data,
		})
	}

	return strings
}

type MainMethod struct {
	TypeMethod
	class *Class
}

type Element map[string][]byte

func (e Element) String() string {
	str := ""
	for k, v := range e {
		switch k {
		case "type", "string":
			str += fmt.Sprintf("  %v: %v", k, string(v))
		case "number":
			str += fmt.Sprintf("  %v: %v", k, ToInt(v))
		default:
			panic(fmt.Sprintf("No print implementation for %v", string(k)))
		}
	}
	return str
}

type Stack []Element

func (s *Stack) Push(key string, value []byte) {
	e := make(Element)
	e[key] = []byte(value)
	*s = append(*s, e)
}

func (s *Stack) Peek() Element {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() Element {
	e := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return e
}

func (s Stack) String() string {
	str := ""
	for i, v := range s {
		str += fmt.Sprintf("[%v] %v\n", i, v)
	}
	return str
}

func (mm *MainMethod) Execute() {
	code := mm.GetCodeAttribute(mm.class).code
	// fmt.Printf("%+#v\n", code)

	/*
		0xb2 0x00 0x07	-> load System.out
		0x12 0x13		-> ldc "Hello, World"
		0xb6 0x00 0x0f 	-> invokevirtual
		0xb1			-> return
	*/

	stack := make(Stack, 0)

	__nbi = 0
	for {
		opcode := NextBytes(code, 1)[0]
		// fmt.Printf("Opcode: %02x\n", opcode)
		switch opcode {
		case 0xb2: // getstatic
			index := ToInt(NextBytes(code, 2))

			static := mm.class.constant_pool[index-1]
			class_index := ToInt(static["class_index"])
			class := mm.class.constant_pool[class_index-1]
			name_and_type_index := ToInt(static["name_and_type_index"])
			name_and_type := mm.class.constant_pool[name_and_type_index-1]

			class_name_index := ToInt(class["name_index"])
			class_name := string(mm.class.constant_pool[class_name_index-1]["bytes"])

			type_name := string(mm.class.constant_pool[ToInt(name_and_type["name_index"])-1]["bytes"])

			// fmt.Printf("getstatic %v.%v\n", class_name, type_name)

			if class_name == "java/lang/System" && type_name == "out" {
				stack.Push("type", []byte("PrintStream"))
			} else {
				panic("Class and type not implemented yet")
			}

		case 0x12: // ldc
			index := ToInt(NextBytes(code, 1))

			static := mm.class.constant_pool[index-1]

			switch ToInt(static["tag"]) {
			case CONSTANT_String:
				value := string(mm.class.constant_pool[ToInt(static["string_index"])-1]["bytes"])
				stack.Push("string", []byte(value))
				// fmt.Printf("ldc \"%v\"\n", value)
			default:
				panic(fmt.Sprintf("Parsing type %v is not implemented yet\n", ToInt(static["tag"])))
			}
		case 0xb6: // invokevirtual
			index := ToInt(NextBytes(code, 2))
			static := mm.class.constant_pool[index-1]
			// fmt.Println("invokevirtual", index)
			class_index := ToInt(static["class_index"])
			name_and_type_index := ToInt(static["name_and_type_index"])

			class := mm.class.constant_pool[class_index-1]
			name_and_type := mm.class.constant_pool[name_and_type_index-1]

			class_name := string(mm.class.constant_pool[ToInt(class["name_index"])-1]["bytes"])
			type_name := string(mm.class.constant_pool[ToInt(name_and_type["name_index"])-1]["bytes"])

			// fmt.Println(class_name, type_name)

			if class_name != "java/io/PrintStream" || type_name != "println" {
				panic("Invokevirtual class and type not implemented yet")
			}

			data := stack.Pop()
			stream := stack.Pop()

			if string(stream["type"]) != "PrintStream" {
				panic("Invoking function java/io/PrintStream/println() is only impelmented on \"PrintStream\" yet")
			}

			if v, ok := data["string"]; ok {
				fmt.Printf("%v\n", string(v))
			} else if v, ok := data["number"]; ok {
				fmt.Printf("%v\n", ToInt(v))
			}

		// fmt.Println(stream, data)
		case 0x10: // bipush
			byt := NextBytes(code, 1)
			stack.Push("number", byt)
		case 0x11: // sipush
			short := NextBytes(code, 2)
			// fmt.Println("sipush", ToInt(short))
			stack.Push("number", short)

		case 0xb1: // return
			fmt.Println("exit code: 0")
			return
		default:
			panic(fmt.Sprintf("Opcode 0x%x not implemented yet", opcode))
		}

		// fmt.Println(stack)
	}
}

func (c *Class) GetMainMethod() MainMethod {

	methods := c.Methods()
	for _, v := range methods {
		if v.name == "main" {
			m := MainMethod{
				TypeMethod: v,
				class:      c,
			}
			return m
		}
	}

	panic("No main method found in class")
}

func (c *Class) parseAttributes(count int, bytes []byte) []TypeAttribute {
	attributes := make([]TypeAttribute, 0)

	for i := 0; i < count; i++ {
		attribute_name_index := ToInt(NextBytes(bytes, 2))
		attribute_length := ToInt(NextBytes(bytes, 4))
		info := NextBytes(bytes, attribute_length)

		attribute := TypeAttribute{}

		c.expectPoolIndexType(attribute_name_index-1, CONSTANT_Utf8)
		attribute.name = string(c.constant_pool[attribute_name_index-1]["bytes"])

		attribute.bytes = info
		attributes = append(attributes, attribute)
	}

	return attributes
}

func (c *Class) Methods() []TypeMethod {
	methods := make([]TypeMethod, 0)
	for _, v := range c.methods {
		method := TypeMethod{}

		name_index := ToInt(v["name_index"])
		c.expectPoolIndexType(name_index-1, CONSTANT_Utf8)
		method.name = string(c.constant_pool[name_index-1]["bytes"])

		access_flags := ToInt(v["access_flags"])
		method.access_flags = ParseFlags(access_flags, METHODS_ACCESS_FLAGS)

		descriptor_index := ToInt(v["descriptor_index"])
		c.expectPoolIndexType(descriptor_index-1, CONSTANT_Utf8)
		method.descriptor = string(c.constant_pool[descriptor_index-1]["bytes"])

		attribute := TypeAttribute{}

		attribute_name_index := ToInt(v["attributes_info-attribute_name_index"])
		c.expectPoolIndexType(attribute_name_index-1, CONSTANT_Utf8)
		attribute.name = string(c.constant_pool[attribute_name_index-1]["bytes"])

		attribute_info := v["attributes_info-info"]
		attribute.bytes = attribute_info

		method.attributes = attribute

		methods = append(methods, method)
	}

	return methods
}
