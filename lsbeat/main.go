package main

import (
	"os"
	"fmt"
	"io/ioutil"
)
func main() {
    //apply run path "." without argument.
    if len(os.Args) == 1 {
        listDir(".")
    } else {
        listDir(os.Args[1])
    }
}

func listDir(dirFile string) {
    files, _ := ioutil.ReadDir(dirFile)
    for _, f := range files {
        t := f.ModTime()
        fmt.Println(f.Name(), dirFile+"/"+f.Name(), f.IsDir(), t, f.Size())

        if f.IsDir() {
            listDir(dirFile + "/" + f.Name())
        }
    }
}

