package pickStr

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

const (
	fileName = "pickStr/data"
)

type wordInfo struct {
	wordList []string
	listLen int
	isUsed []bool
}

var wi wordInfo 

func Init() error {
	fp, err := os.Open(fileName)

	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		wi.wordList = append(wi.wordList, scanner.Text())
	}
	
	wi.listLen = len(wi.wordList)
	// initialized with false
	wi.isUsed = make([]bool, wi.listLen)
	rand.Seed(time.Now().UnixNano())

	return nil
}

func Pick() string {
	var i int
	for {
		i = rand.Intn(wi.listLen)
		if !wi.isUsed[i] {
			wi.isUsed[i] = true
			return wi.wordList[i]
		}
	}
}
