package main

import (
	"fmt"
	"reflect"
	"time"
)

type Result struct {
	Field1 string
	Field2 int64
	Field3 time.Time
}

func main() {

	t := time.Now()

	Result{}

	res1 := []Result{
		Result{"foo", int64(1), t},
		Result{"bar", int64(2), t},
		Result{"baz", int64(3), t},
	}
	res2 := []Result{
		Result{"foo", int64(1), t},
		Result{"bar", int64(2), t},
		Result{"baz", int64(3), t},
	}

	fmt.Printf("Equal? %t\n", reflect.DeepEqual(res1, res2))
}
