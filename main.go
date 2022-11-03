package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("缺少下載的網址")
	}

	url := os.Args[1]
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fileName := strconv.Itoa(int(time.Now().Unix())) + ".png"
	err = ioutil.WriteFile(fileName, body, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("save file", fileName)
}
