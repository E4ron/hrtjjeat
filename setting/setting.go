package setting

import (
	"encoding/json"
	"fmt"
	"os"
)

type Setting struct {
	Address string
	Port    string
	DbHost  string
	DbPort  string
	DbUser  string
	DbPass  string
	DbName  string
}

var options Setting

func Load(filename string) *Setting {
	file, e := os.Open(filename)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	defer file.Close()
	stat, e := file.Stat()
	if e != nil {
		fmt.Println(e)
		return nil
	}
	bytes := make([]byte, stat.Size())
	n, e := file.Read(bytes)
	if e != nil || n != int(stat.Size()) {
		fmt.Println(e)
		return nil
	}
	e = json.Unmarshal(bytes, &options)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	return &options
}

func Save(filename string, s *Setting) {
	bytes, e := json.Marshal(s)
	if e != nil {
		fmt.Println(e)
		return
	}
	e = os.WriteFile(filename, bytes, 0777)
	if e != nil {
		fmt.Println(e)
	}
}
