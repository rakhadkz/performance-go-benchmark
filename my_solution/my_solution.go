package my_solution

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
)

func MySolution(out io.Writer) {

	// Reading the whole file
	f, err := ioutil.ReadFile("my_text.txt")

	if err != nil {
		panic(err)
	}

	// Creating Buffered channel
	output := make(chan []byte, 2000)

	// + Multithreading
	go readBytes(f, output)
	getBytes(output, out)
}

func readBytes(input []byte, output chan []byte) {

	// Creating slice with assigned CAPACITY
	var slice = make([]byte, 1000)
	for _, v := range input {
		
		if isChar(v) {
			if v >= 65 && v <= 90 {
				v += 32
			}
			slice = append(slice, v)
		} else if len(slice) != 0 {
			output <- slice
			slice = []byte{}
		}
	}
	close(output)
}

type word struct {
	data    []byte
	counter int
}

func getBytes(input chan []byte, out io.Writer) {
	var words []*word
	for data := range input {
		if words == nil {
			words = append(words, &word{
				data:    data,
				counter: 1,
			})
		} else {
			isFound := false
			for i, v := range words {
				if bytes.Equal(v.data, data) {
					words[i].counter++
					isFound = true
					break
				}
			}
			if !isFound {
				words = append(words, &word{
					data:    data,
					counter: 1,
				})
			}
		}
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].counter > words[j].counter
	})
	for i := 0; i < 20; i++ {
		_, _ = fmt.Fprintf(out, "%v %v\n", words[i].counter, string(words[i].data))
	}
}

func isChar(byteVal byte) bool {
	if (byteVal >= 65 && byteVal <= 90) || (byteVal >= 97 && byteVal <= 122) {
		return true
	}
	return false
}
