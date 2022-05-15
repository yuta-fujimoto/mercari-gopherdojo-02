package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "net/http/httputil"
	"os"
	// "strconv"
	"sync"
)

func main() {
	routines := flag.Int64("p", 5, "number of routines")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		return
	}

	var wg sync.WaitGroup
	res, err := http.Head(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
	bytesPerRoutine := res.ContentLength / (*routines)
	LastBytes := bytesPerRoutine + res.ContentLength%(*routines)
	body := (make([][]byte, *routines))
	fmt.Println("head request done", res.ContentLength)

	var i int64
	for i = 0; i < *routines; i++ {
		wg.Add(1)

		min := bytesPerRoutine * i
		max := min + bytesPerRoutine
		if i == *routines-1 {
			max = min + LastBytes
		}

		go func(min int64, max int64, i int64) {
			client := &http.Client{}
			req, err := http.NewRequest("GET", args[0], nil)
			if err != nil {
				log.Fatal(err.Error())
			}
			// range_header := "bytes=" + strconv.FormatInt(min, 10) + "-" + strconv.FormatInt(max-1, 10)
			// req.Header.Add("Range", range_header)
			// fmt.Println(range_header)
			resp, err := client.Do(req)
			
			if err != nil {
				log.Fatal(err.Error())
			}
			defer resp.Body.Close()
			reader, err := ioutil.ReadAll(res.Body)

			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			body[i] = reader
			fmt.Printf("%#v\n", string(reader))
			wg.Done()
		}(min, max, i)
	}
	wg.Wait()
	fmt.Println("dowload complete")
	for i = 0; i < *routines; i++ {
		err = ioutil.WriteFile("result", body[i], os.ModeDevice)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	}
}
