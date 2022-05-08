package gosteam

import (
	"fmt"
	"io"
)

/*
	ONLY FOR TESTING, REMOVE LATER
*/

func Stringify(body io.ReadCloser) (string, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}
