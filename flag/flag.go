package main

import (
	"flag"
	"fmt"
)
// ./flag -h
//Usage of ./flag:
//-otherbool
//flag.Bool example
//-otherstring string
//flag.String example (default "other default value")
//-somebool
//boolVar example
//-somestring string
//stringVar example (default "default string value")

func main()  {
	var b1 bool
	flag.BoolVar(&b1, "somebool", false, "boolVar example")
	b2 := flag.Bool("otherbool", false, "flag.Bool example")

	fmt.Printf("BoolVar: %T, flag.Bool: %T: types \n", b1, b2)
	fmt.Printf("BoolVar: %t, flag.Bool: %t: before parse\n", b1, *b2)

	var s1 string
	flag.StringVar(&s1, "somestring", "default string value", "stringVar example")
	s2 := flag.String("otherstring", "other default value", "flag.String example")
	fmt.Printf("StringVar: %T, flag.String: %T: types \n", s1, s2)
	fmt.Printf("StringVar: %s, flag.String: %s: before parse\n", s1, *s2)

	flag.Parse()
	fmt.Printf("BoolVar: %t, flag.Bool: %t: after parse\n", b1, *b2)
	fmt.Printf("StringVar: %s, flag.String: %s: after parse\n", s1, *s2)
}