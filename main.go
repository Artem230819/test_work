package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)


func Split(r rune) bool {
	return r == ' ' || r == ',' || r == '.' || r == '!' || r == '?' || r == '"' || r == ';' || r == ':' || r == '[' || r == ']' || r == '(' || r == ')' || r == '\n' || r == '\r' || r == '\t'
}


func main()  {

	var url string
	fmt.Print("Введите адрес: ")
	fmt.Scan(&url)
			client := &http.Client{}

			resp, err := client.Get(url)
			if err != nil {
				panic(err)
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(resp.Body)
			responseData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			responseString := string(responseData)
			responseStringSplit := strings.FieldsFunc(responseString, Split)
			file, err := os.Create("test.txt")
			if err != nil {
				panic(err)
			}
			for _, v := range responseStringSplit{
				sum := 0
				for i, j := range responseStringSplit{
					if j == v {
						sum += 1
						responseStringSplit = append(responseStringSplit[:i], responseStringSplit[i + 1:]...)
					}
				}
				if sum != 0 {
					_, err := fmt.Fprintf(file, "%s - %d \n", v, sum)
					if err != nil {
						panic(err)
						return
					}
				}
			}
			defer func(file *os.File) {
				err := file.Close()
				if err != nil {
					panic(err)
				}
			}(file)

}