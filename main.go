package main

import "github.com/nichttoxisch/go-jvm/java"


func main() {
	var filename string = "sources/Main.class"

	bytes := java.ReadFile(filename)

	class := java.Class{}
	class.ParseFromBytes(bytes)

	// fmt.Println(class)

	// fmt.Println("Access flags:")
	// for _, v := range class.Flags() {
	// 	fmt.Printf("  %v\n", v)
	// }

	// fmt.Println("Classes:")
	// for _, v := range class.Classes() {
	// 	fmt.Printf("  %v\n", v)
	// }

	// fmt.Println("Strings:")
	// for _, v := range class.Strings() {
	// 	fmt.Printf("  %v\n", v)
	// }

	main := class.GetMainMethod()
	main.Execute()
}
