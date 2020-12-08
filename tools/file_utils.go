package tools

import (
	"bufio"

	"os"
)

/**
文件读取
*/
func FileLinesScanner(fileName string) ([]string, error) {
	var results []string
	file, err := os.Open(fileName)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		results = append(results, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		//log.Fatal(err)
		return nil, err
	}
	return results, nil
}
