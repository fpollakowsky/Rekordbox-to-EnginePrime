package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	xmlSource, err := ioutil.ReadFile("collection.xml")

	if err != nil {
		log.Fatalln(err)
	}

	_1 := strings.Replace(string(xmlSource), "Rating=\"51\"", "Rating=\"5\"", -1)
	_2 := strings.Replace(_1, "Rating=\"102\"", "Rating=\"15\"", -1)
	_3 := strings.Replace(_2, "Rating=\"153\"", "Rating=\"25\"", -1)
	_4 := strings.Replace(_3, "Rating=\"204\"", "Rating=\"35\"", -1)
	_5 := strings.Replace(_4, "Rating=\"255\"", "Rating=\"45\"", -1)

	f, err := os.Create("collection_exported.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(_5)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println("Finished \n Total written lines: ",l)
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(15 * time.Second)
}