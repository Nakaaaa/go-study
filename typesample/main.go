package main

import "fmt"

func main() {
	i := interface{}("Expert")
	s := i.(string)
	fmt.Println(s)

	n, ok := i.([]byte)
	fmt.Println(n, ok)

	typeSwitch(i)

	do := func() error {
		return &ErrConflictEntity{}
	}

	switch do().(type) {
	case nil:
		// do nothing
	case *ErrNoSuchEntity:
		fmt.Println("error no such entity")
	case *ErrConflictEntity:
		fmt.Println("error conflict entity")
	default:
		fmt.Println("unknown error")
	}
}

func typeSwitch(i interface{}) {
	switch i.(type) {
	case int, int8, int16, int32, int64:
		fmt.Println("This is integer,", i)
	case string:
		fmt.Println("This is string,", i)
	default:
		fmt.Printf("This is unknown type, %T\n", i)
	}
}

// エラー
type ErrNoSuchEntity struct{ error }
type ErrConflictEntity struct{ error }
