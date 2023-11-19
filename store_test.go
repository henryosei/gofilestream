package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpiture"
	pathkey := CASPathTransformFun(key)
	expectedOriginal := "d869e9397855c9d4802bd8887d8c5c749f85d989"
	expectedPathName := "d869e/93978/55c9d/4802b/d8887/d8c5c/749f8/5d989"
	if pathkey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathkey.PathName, expectedPathName)
	}
	if pathkey.Orignal != expectedOriginal {
		t.Errorf("have %s want %s", pathkey.Orignal, expectedOriginal)
	}
	fmt.Println(pathkey.PathName)
}
func TestScore(t *testing.T) {
	opts := StoreOpts{PathTransformFunc: CASPathTransformFun}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("some jpeg byte"))
	if err := s.writeStream("myspecialpicture", data); err != nil {

	}
}
