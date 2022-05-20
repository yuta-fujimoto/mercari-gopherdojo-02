package main

import (
	"errors"
	"net/http"
)

type contentInfo struct {
	BytesPerRoutine int64
	LastBytes int64
	RoutineCnt int64 
	Url string 
}

func getContentInfo(client *http.Client, url string, routineCnt int64) (*contentInfo, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Header.Get("Accept-Ranges") != "bytes" {
		return nil, errors.New("does not support range request")
	}
	if resp.ContentLength <= 0 {
		return nil, errors.New("invalid content length")
	}
	return &contentInfo{
		BytesPerRoutine: resp.ContentLength / routineCnt,
		LastBytes: resp.ContentLength / routineCnt + resp.ContentLength % routineCnt,
		RoutineCnt: routineCnt,
		Url: url,
	} , nil
}
