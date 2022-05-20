package download

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	// "os/exec"
	"path/filepath"
)


func Run(fn string, routineCnt int64) error {
	client := &http.Client{}

	info, err := getContentInfo(client, fn, routineCnt)
	if err != nil {
		return err
	}
	
	saveFiles, err := download(info, client)
	defer func() {
		for _, sf := range saveFiles {
			os.Remove(sf.Name())
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
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		os.Exit(1)
	}

	if err := Run(args[0], 2); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
