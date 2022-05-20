package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	// "os/exec"
	"path/filepath"
)


func run(fn string, routineCnt int64) error {
	client := &http.Client{}

	info, err := getContentInfo(client, fn, routineCnt)
	if err != nil {
		return err
	}
	
	saveFiles, err := download(info, client)
	defer func() {
		for i := 0; i < len(saveFiles); i++ {
			// cat, _ := exec.Command("cat", createSubfilename(i)).Output()
			// fmt.Printf("%#v\n", string(cat))

			os.Remove(saveFiles[i].Name())
		}
	}()
	if  err != nil {
		return err
	}
	
	dstfile, err := os.OpenFile(filepath.Base(info.Url), os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer dstfile.Close()

	for i := int64(0); i < routineCnt; i++ {
		subfile, err := os.Open(saveFiles[i].Name())
		if err != nil {
			return fmt.Errorf("open temp files: %w", err)
		}
		defer subfile.Close()
		if _, err = io.Copy(dstfile, subfile); err != nil {
			return fmt.Errorf("write to dst file: %w", err)
		}
	}
	return  nil
}

func main() {
	routineCnt := flag.Int64("p", 2, "number of routines")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		os.Exit(1)
	}

	if err := run(args[0], *routineCnt); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
