package cgo
import "C"

// #include <stdlib.h>
// #include <stdio.h>
// int* global;
// int* allocate()
// {
//	    global = malloc(3 * sizeof(int));
//      int i;
//      for (i = 0; i < 3; i++) {
//         global[i] = i;
//      }
//      return global;
// }
//
// int test()
// {
//      int i = 0;
//      for (i = 0; i < 3; i++) {
//         printf("%d\n", i);
//      }
//      return 0;
// }
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	f := C.allocate()
	C.test()
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
	i := (*[3]int)(unsafe.Pointer(f))
	i2 := i[:]
	fmt.Println(i2)
	fmt.Println(reflect.TypeOf(i2))
}