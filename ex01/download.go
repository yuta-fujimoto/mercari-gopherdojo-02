package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/sync/errgroup"
)

func download(info *contentInfo, client *http.Client) ([]*os.File, error) {
	saveFiles := make([]*os.File, info.RoutineCnt)
	for i := int64(0); i < info.RoutineCnt; i++ {
		var err error
		saveFiles[i], err = ioutil.TempFile("", "download")
		if err != nil {
			return saveFiles, err
		}
		defer saveFiles[i].Close()
	}

	eg, ctx := errgroup.WithContext(context.TODO())
	for i := int64(0); i < info.RoutineCnt; i++ {
		min := info.BytesPerRoutine * i
		max := min + info.BytesPerRoutine
		if i == info.RoutineCnt-1 {
			max = min + info.LastBytes
		}
		i := i
		output := saveFiles[i]
		defer output.Close()

		eg.Go(func() error {
			req, err := http.NewRequest("GET", info.Url, nil)
			if err != nil {
				return fmt.Errorf("create new request: %w", err)
			}

			req = req.WithContext(ctx)
			range_header := "bytes=" + strconv.FormatInt(min, 10) + "-" + strconv.FormatInt(max-1, 10)
			req.Header.Add("Range", range_header)

			req.Close = true
			resp, err := client.Do(req)
			if err != nil {
				return fmt.Errorf("send http request: %w", err)
			}

			defer resp.Body.Close()

			if _, err = io.Copy(output, resp.Body); err != nil {
				return fmt.Errorf("copy to temp file: %w %s", err, output.Name())
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return saveFiles, err
	}
	return saveFiles, nil
}
