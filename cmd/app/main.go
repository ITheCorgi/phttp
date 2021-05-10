package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

const DefaultParallel = 10

var Parallel int

type URLReq struct {
	URL  string
	Hash string
}

func init() {
	flag.IntVar(&Parallel, "parallel", DefaultParallel, "-parallel <count>")
}

func GetFlags() []string {
	flag.Parse()
	if flag.NArg() < 2 {
		fmt.Println("Arguments must been provided")
		os.Exit(-1)
	}
	return flag.Args()
}

func ParallelURLReq(urls []string, limit int) error {
	sem := make(chan struct{}, limit)
	res := make(chan *URLReq)

	defer func() {
		close(sem)
		close(res)
	}()

	for i, url := range urls {
		go func(i int, url string) {
			sem <- struct{}{}
			resp, _ := http.Get(url)
			defer resp.Body.Close()
			res <- &URLReq{url, GetMD5Hash(url)}
			<-sem
		}(i, url)
	}
	var cnt int
	for {
		msg1 := <-res
		cnt++
		fmt.Printf("%s %s\n", msg1.URL, msg1.Hash)
		if cnt == len(urls) {
			break
		}
	}
	return nil
}

func GetMD5Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func main() {
	var urls = GetFlags()
	runtime.GOMAXPROCS(Parallel)

	err := ParallelURLReq(urls, Parallel)
	if err != nil {
		fmt.Println(err)
	}
}
