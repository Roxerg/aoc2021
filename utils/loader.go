package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LoadFile(day string, separator string) (string, []string) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	content, err := ioutil.ReadFile(day + "/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	strContent := string(content)

	return strContent, strings.Split(strContent, separator)
}
