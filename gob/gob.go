package main

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Example struct {
	Name     string
	Relation Relation
	Scalar   uint32
}

type Relation struct {
	Foo int
	Bar string
	Baz []int
}

func main() {
	fmt.Println("Similar struct example")
	similar()
	fmt.Println("Diff names struct example")
	diffName()
	fmt.Println("Diff relation struct example")
	diffRelation()
	fmt.Println("Diff scalar struct example")
	diffScalar()
}

func similar()  {
	ex1 := Example{
		Name:     "Common",
		Relation: Relation{
			Foo: 1,
			Bar: "one",
			Baz: []int{1,2,3,4,5},
		},
		Scalar:   10,
	}

	ex2 := Example{
		Name:     "Common",
		Relation: Relation{
			Foo: 1,
			Bar: "one",
			Baz: []int{1,2,3,4,5},
		},
		Scalar:   10,
	}

	var a bytes.Buffer
	if err := gob.NewEncoder(&a).Encode(ex1); err != nil {
		log.WithError(err).Fatal("Unable to encode", ex1)
	}

	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(ex2); err != nil {
		log.WithError(err).Fatal("Unable to encode", ex2)
	}

	fmt.Printf("%x\n",md5.Sum(a.Bytes()))
	fmt.Printf("%x\n",md5.Sum(b.Bytes()))
}

func diffName()  {
	ex1 := Example{
		Name:     "Common",
		Relation: Relation{
			Foo: 1,
			Bar: "one",
			Baz: []int{1,2,3,4,5},
		},
		Scalar:   10,
	}

	ex2 := Example{
		Name:     "Not common",
		Relation: Relation{
			Foo: 1,
			Bar: "one",
			Baz: []int{1,2,3,4,5},
		},
		Scalar:   10,
	}

	var a bytes.Buffer
	if err := gob.NewEncoder(&a).Encode(ex1); err != nil {
		log.WithError(err).Fatal("Unable to encode", ex1)
	}

	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(ex2); err != nil {
		log.WithError(err).Fatal("Unable to encode", ex2)
	}

	fmt.Printf("%x\n",md5.Sum(a.Bytes()))
	fmt.Printf("%x\n",md5.Sum(b.Bytes()))
}

func diffRelation()  {
	ex1 := Example{
		Name:     "Common",
		Relation: Relation{
			Foo: 1,
			Bar: "one",
			Baz: []int{1,2,3,4,5},
		},
		Scalar:   10,
	}

	ex2 := Example{
		Name:     "Common",
		Relation: Relation{
			Foo: 1,
			Bar: "one",
			Baz: []int{1,2,3,4},
		},
		Scalar:   10,
	}

	var a bytes.Buffer
	if err := gob.NewEncoder(&a).Encode(ex1); err != nil {
		log.WithError(err).Fatal("Unable to encode", ex1)
	}

	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(ex2); err != nil {
		log.WithError(err).Fatal("Unable to encode", ex2)
	}

	fmt.Printf("%x\n",md5.Sum(a.Bytes()))
	fmt.Printf("%x\n",md5.Sum(b.Bytes()))
}

func diffScalar()  {
	ex1 := Example{
		Name:     "Common",
		Relation: Relation{
			Foo: 1,
			Bar: "one",
			Baz: []int{1,2,3,4,5},
		},
		Scalar:   10,
	}

	ex2 := Example{
		Name:     "Common",
		Relation: Relation{
			Foo: 1,
			Bar: "one",
			Baz: []int{1,2,3,4,5},
		},
		Scalar:   11,
	}

	var a bytes.Buffer
	if err := gob.NewEncoder(&a).Encode(ex1); err != nil {
		log.WithError(err).Fatal("Unable to encode", ex1)
	}

	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(ex2); err != nil {
		log.WithError(err).Fatal("Unable to encode", ex2)
	}

	fmt.Printf("%x\n",md5.Sum(a.Bytes()))
	fmt.Printf("%x\n",md5.Sum(b.Bytes()))
}