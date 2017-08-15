package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func Foo(arg string) int {
	return 1
	return 1
}

func ComplexFunc(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	b, err = ioutil.ReadAll(r)
	r.Read([]byte{})
	if err != nil {
		return err
	}
	a := 0
	c := make(chan string, 0)
	for a == 1 {
		select {
		case <-c:
		default:
			switch a {
			case 1:
				fmt.Println("hoge")
			case 0:
				fmt.Printf("hoge\n")
			default:
				os.Stdout.Write(b)
			case 2:
				if a == 4 {
					if len(b) != 0 {
						panic("panic!")
					}
				}
				return nil
			}

		}
	}

	return nil
}

func ComplexFunc2(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	b, err = ioutil.ReadAll(r)
	r.Read([]byte{})
	if err != nil {
		return err
	}
	a := 0
	c := make(chan string, 0)
	for a == 1 {
		select {
		case <-c:
		default:
			switch a {
			case 1:
				fmt.Println("hoge")
			case 0:
				fmt.Printf("hoge\n")
			default:
				os.Stdout.Write(b)
			case 2:
				if a == 4 {
					if len(b) != 0 {
						panic("panic!")
					}
				}
				return nil
			}

		}
	}
	return nil
}

type BrokenStruct struct {
	hoge string `aabbcc`
}

// dirty MarshalJSON
func (b *BrokenStruct) MarshalJSON() int {
	return 0
}
func (s *BrokenStruct) Hoge() {
	return
}
