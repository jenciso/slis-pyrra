package slogen

import (
	"crypto/md5"
	"fmt"
	"math"
	"os"

	"gopkg.in/yaml.v2"
)

func GenerateHashID(s string) string {
	data := md5.Sum([]byte(s))
	hashid := fmt.Sprintf("%x", data)
	return hashid[:8]
}

func Yaml2String(data any) string {
	yamlFile, err := yaml.Marshal(&data)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(string(yamlFile))
}

func PrintYamlResult(filename string, data string) {
	if filename != "" {
		WriteYamlResult(filename, data)
	} else {
		fmt.Printf("%s", data)
	}
}

func WriteYamlResult(filename string, data string) {
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func TruncateString(str string, length int) string {
	if length <= 0 {
		return ""
	}
	truncated := ""
	count := 0
	for _, char := range str {
		truncated += string(char)
		count++
		if count >= length {
			break
		}
	}
	return truncated
}
