package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	f, err := os.OpenFile("/var/tmp/a/counter", os.O_RDWR, 0)
	i := 0
	if err != nil {
		err = os.MkdirAll("/var/tmp/a", os.FileMode(0755))
		if err != nil {
			panic(err)
		}
		f, err = os.Create("/var/tmp/a/counter")
		if err != nil {
			panic(err)
		}
		f.Write([]byte("0"))
		f.Close()
	} else {
		a, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		f.Truncate(0)
		f.Seek(0, io.SeekStart)
		i, err = strconv.Atoi(string(a))
		if err != nil {
			panic(err)
		}
		i++
		_, err = f.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			panic(err)
		}
		f.Close()
	}
	err = os.Mkdir("/var/tmp/a/"+strconv.Itoa(i), os.FileMode(0755))
	if err != nil {
		panic(err)
	}
	fmt.Println("/var/tmp/a/" + strconv.Itoa(i))
}
