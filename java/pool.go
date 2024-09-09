package java

import "fmt"

type CpInfo map[string][]byte
type Interface map[string][]byte
type FieldInfo map[string][]byte
type MethodInfo map[string][]byte
type AttributeInfo map[string][]byte

func (cp_info CpInfo) String() string {
	s := ""
	switch ToInt(cp_info["tag"]) {
	case CONSTANT_Class:
		s += fmt.Sprintf("CONSTANT_Class = {\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("\t name_index = %v\n", ToInt(cp_info["name_index"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_Fieldref:
		s += fmt.Sprintf("CONSTANT_Fieldref\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("\t class_index = %v\n", ToInt(cp_info["class_index"]))
		s += fmt.Sprintf("\t name_and_type_index = %v\n", ToInt(cp_info["name_and_type_index"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_Methodref:
		s += fmt.Sprintf("CONSTANT_Methodref\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("\t class_index = %v\n", ToInt(cp_info["class_index"]))
		s += fmt.Sprintf("\t name_and_type_index = %v\n", ToInt(cp_info["name_and_type_index"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_InterfaceMethodref:
		s += fmt.Sprintf("CONSTANT_InterfaceMethodref\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_String:
		s += fmt.Sprintf("CONSTANT_String\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("\t string_index = %v\n", ToInt(cp_info["string_index"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_Integer:
		s += fmt.Sprintf("CONSTANT_Integer\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_Float:
		s += fmt.Sprintf("CONSTANT_Float\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_Long:
		s += fmt.Sprintf("CONSTANT_Long\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_Double:
		s += fmt.Sprintf("CONSTANT_Double\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_NameAndType:
		s += fmt.Sprintf("CONSTANT_NameAndType\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("\t name_index = %v\n", ToInt(cp_info["name_index"]))
		s += fmt.Sprintf("\t descriptor_index = %v\n", ToInt(cp_info["descriptor_index"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_Utf8:
		s += fmt.Sprintf("CONSTANT_Utf8 = {\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("\t length = %v\n", ToInt(cp_info["length"]))
		s += fmt.Sprintf("\t bytes = \"%v\"\n", string(cp_info["bytes"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_MethodHandle:
		s += fmt.Sprintf("CONSTANT_MethodHandle\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_MethodType:
		s += fmt.Sprintf("CONSTANT_MethodType\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
		s += fmt.Sprintf("}\n")
	case CONSTANT_InvokeDynamic:
		s += fmt.Sprintf("CONSTANT_InvokeDynamic\n")
		s += fmt.Sprintf("}\n")
		s += fmt.Sprintf("\t tag = %v\n", ToInt(cp_info["tag"]))
	}
	return s
}

func (mi MethodInfo) String() string {
	s := ""
	for k, v := range mi {
		if k == "attributes_info-info" {
			s += fmt.Sprintf("\n  %v: %v", k, v)
			continue
		}
		s += fmt.Sprintf("\n  %v: %v", k, ToInt(v))
	}
	return s + "\n"
}
