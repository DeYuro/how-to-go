package main

import "fmt"

func main() {
	differentPrintf()
}

type Foo struct {
	Bar int
	Baz string
}

func differentPrintf() {
	var st string
	var in int
	var fl float64
	var bl bool
	var ar [42]int
	var sr Foo
	sl := make([]int, 5, 10)
	mp := make(map[string]int)
	ch := make(chan int, 5)
	println("zero values")
	fmt.Printf("st = %v, in = %v, fl = %v, bl = %v, ar = %v, sl = %v, mp = %v, ch = %v, sr = %v \n", st, in, fl, bl, ar, sl, mp, ch, sr)
	st = "foo"
	in = 42
	fl = 42.2
	bl = true
	sl[2] = 5
	mp["foo"] = 2
	ch <- 42
	sr.Bar = 42
	sr.Baz = "Question"
	println("%v after init val")
	fmt.Printf("st = %v, in = %v, fl = %v, bl = %v, ar = %v, sl = %v, mp = %v, ch = %v, sr = %v \n", st, in, fl, bl, ar, sl, mp, ch, sr)
	println("+v after init val")
	fmt.Printf("st = %+v, in = %+v, fl = %+v, bl = %+v, ar = %+v, sl = %+v, mp = %+v, ch = %+v, sr = %+v \n", st, in, fl, bl, ar, sl, mp, ch, sr)
	println("#v after init val")
	fmt.Printf("st = %#v, in = %#v, fl = %#v, bl = %#v, ar = %#v, sl = %#v, mp = %#v, ch = %#v, sr = %#v \n", st, in, fl, bl, ar, sl, mp, ch, sr)
	println("%T")
	fmt.Printf("st = %T, in = %T, fl = %T, bl = %T, ar = %T, sl = %T, mp = %T, ch = %T, sr = %T \n", st, in, fl, bl, ar, sl, mp, ch, sr)
	println("Default format")
	fmt.Printf("st = %s, in = %d, fl = %g, bl = %t, ar = %v, sl = %v, mp = %v, ch = %p, sr = %+v \n", st, in, fl, bl, ar, sl, mp, ch, sr)
}
