package main

import (
	"bytes"
	"container/list"
	"errors"
	"fmt"
	"os" // "io/ioutil"
	"path/filepath"
)

func main() {
	// fmt.Println(strings.Contains("test", "es"))

	var buf bytes.Buffer
	buf.Write([]byte("test"))
	// fmt.Println(buf.Bytes())

	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(stat.Size())
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))

	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	err = errors.New("My error")
	fmt.Println(err)

	var l list.List
	l.PushBack(1)
	l.PushBack(2)
	for it := l.Back(); it != nil; it = it.Prev() {
		fmt.Println(it.Value)
	}
}
