package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// Data ...
type Data struct {
	Content interface{}
}

// ConfigObj ...
type ConfigObj interface {
	LoadConfiguration(string) ([]byte, error)
	GetString(string) string
	GetInt(string) int
}

// LoadConfiguration ...
func (s *Data) LoadConfiguration(file string) ([]byte, error) {

	// configObj := make(map[string]interface{})

	cfgFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	s.Content = cfgFile

	return cfgFile, nil
}

// GetString ...
func (s *Data) GetString(str string) (ans string) {

	// var dataI interface{}
	dataMap := make(map[string]interface{})
	sliceByte := s.Content.([]byte)

	err := json.Unmarshal(sliceByte, &dataMap)
	if err != nil {
		fmt.Println("something went wrong!")
	}

	if strings.Contains(str, ".") {

		sliceStr := strings.SplitAfter(str, ".") // this creates a slice of two strings

		// To remove the character "." which will be attached to the first string in the slice
		sliceStr[0] = strings.Replace(sliceStr[0], ".", "", -1)

		switch v := dataMap[sliceStr[0]].(type) {

		case interface{}:
			ansI := dataMap[sliceStr[0]].(map[string]interface{})

			ansValue, ok := ansI[sliceStr[1]].(string)
			ans = ansValue

			if !ok {
				fmt.Printf("type: %v is not a string\n", v)
			}
		default:

			ans = ""
			fmt.Printf("type: %v is unknown\n", v)
		}

	} else {

		// using type switches for several assertions
		switch v := dataMap[str].(type) {
		case string:
			ans = dataMap[str].(string)
			err = nil
		case interface{}:

		default:
			ans = ""
			fmt.Printf("type: %v is unknown\n", v)
		}

	}

	return ans
}

// GetInt ...
func (s *Data) GetInt(str string) (ans int) {

	dataMap := make(map[string]interface{})
	sliceByte := s.Content.([]byte)

	err := json.Unmarshal(sliceByte, &dataMap)
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	if strings.Contains(str, ".") {

		sliceStr := strings.SplitAfter(str, ".") // this creates a slice of two strings

		// To remove the character "." which will be attached to the first string in the slice
		sliceStr[0] = strings.Replace(sliceStr[0], ".", "", -1)

		switch v := dataMap[sliceStr[0]].(type) {

		case interface{}:
			ansI := dataMap[sliceStr[0]].(map[string]interface{})

			ansValue, ok := ansI[sliceStr[1]].(float64)
			ans = int(ansValue)

			if !ok {
				fmt.Printf("type: %v is not a float64\n", v)
			}
		default:

			ans = 0
			fmt.Printf("type: %v is unknown\n", v)
		}
	} else {

		// using type switches for several assertions
		switch v := dataMap[str].(type) {
		case float64:
			answer := dataMap[str].(float64)
			ans = int(answer)
			err = nil
		default:
			fmt.Printf("type: %v is unknown\n", v)
		}

	}

	return ans
}
