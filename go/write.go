package main

import (
  "encoding/gob"
  "os"
)

func main() {
  foo1 := makeFoo()
  bar1 := makeBar()
  baz1 := makeBaz()

  w := gob.NewEncoder(os.Stdout)
  okOrPanic(w.Encode(foo1))
  okOrPanic(w.Encode(bar1))
  okOrPanic(w.Encode(baz1))
}

// type.go

type Foo struct {
  A int
  B string
  C []float32
  D []string
}

type Bar struct {
  E map[string]int
  F Foo
  G []Foo
}

type Baz struct {
  H map[int]Foo
  I map[string]Bar
}

func makeFoo() Foo {
  return Foo{
    A: 1,
    B: "foo1",
    C: []float32{1.1, 2.2, 3.3333333333333333},
    D: []string{"abc", "def", "ghi"},
  }
}

func makeBar() Bar {
  return Bar{
    E: map[string]int{
      "abc": 1,
      "def": 2,
      "ghi": 3,
    },
    F: makeFoo(),
    G: []Foo{makeFoo(), makeFoo(), makeFoo()},
  }
}

func makeBaz() Baz {
  return Baz{
    H: map[int]Foo{
      1: makeFoo(),
      2: makeFoo(),
      3: makeFoo(),
    },
    I: map[string]Bar{
      "abc": makeBar(),
      "def": makeBar(),
    },
  }
}

func okOrPanic(err error) {
  if err != nil {
    panic(err)
  }
}
