package sampletest

import (
	"io"
)

/* -> Before
func ReadFileContents(f *os.File, numBytes int) ([]byte, error){
	defer f.Close()
	data := make([]byte, numBytes)
	_, err := f.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
*/

func ReadContents(rc io.ReadCloser, numBytes int) ([]byte, error){
	defer rc.Close()
	data := make([]byte, numBytes)
	_, err := rc.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
Source : https://www.youtube.com/watch?v=LEnXBueFBzk

- "Accept interfaces, return structs"
- We used io.ReadCloser, but don't be afraid to define your own  interface
*/