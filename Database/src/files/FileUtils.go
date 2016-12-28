package files

import (
	"os"
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
	res "./../resources"
	"errors"
)

func CreateFile(filename string, fields [] string) {
	file, _ := os.Create(filename + res.FILE_EXTENSION)
	writer := bufio.NewWriter(file)
	defer file.Close()
	fmt.Fprintf(writer, strings.Join(fields, " ") + "\n")
	writer.Flush()
}

func SaveToFile(filename string, entity map[string]string ) {
	f, _ := os.OpenFile(filename + res.FILE_EXTENSION,  os.O_RDWR|os.O_APPEND, 0660)
	defer f.Close()
	save, _ := json.Marshal(entity)
	f.WriteString(string(save) + "\n")
}

func ReadCollectionType(filename string) (string, error) {
	if file, err := os.Open(filename + res.FILE_EXTENSION); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		if scanner.Scan() {
			return scanner.Text(), nil
		} else {
			return res.FILE_NOT_FOUND, errors.New(res.FILE_NOT_FOUND)
		}
	} else {
		return res.FILE_NOT_FOUND, errors.New(res.FILE_NOT_FOUND)
	}
}
