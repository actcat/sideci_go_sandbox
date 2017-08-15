package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type BrokenStruct2 struct {
	hoge string `aabbcc`
}

// dirty MarshalJSON
func (b *BrokenStruct2) MarshalJSON() int {
	return 0
}
func (s *BrokenStruct2) Hoge() {
	return
}
