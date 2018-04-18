// Created by cgo - DO NOT EDIT

//line c1.go:1
package main

// typedef int (*intFunc) ();
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import _ "unsafe"
import "fmt"

func main() {
	f := _Ctype_intFunc((_Cgo_ptr(_Cfpvar_fp_fortytwo)))
	fmt.Println(int((_Cfunc_bridge_int_func)(f)))
	// Output: 42
}