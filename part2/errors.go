package main

import "fmt"
import "errors"

type f func(int) (int, error)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Cannot work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

//implement interface method
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Cannot work with 42"}
	}
	return arg + 3, nil
}

func loop(fun f) {
	for _, i := range []int{7, 42} {
		if _, e := fun(i); e != nil {
			fmt.Println("failed:", e)
		}
	}
}

func main() {
	fmt.Println("f1")
	loop(f1)
	fmt.Println("f2")
	loop(f2)
}
