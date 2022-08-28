package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// File ReadWrite

func ReadByOSReadFile() {
	dat, err := os.ReadFile("./main.go")
	// ioutil.ReadFile等同于os.ReadFile
	//dat, err := ioutil.ReadFile("./main.go")
	check(err)
	fmt.Print(string(dat))
}

func ReadBySeek() {
	f, err := os.Open("./main.go")
	defer f.Close()
	check(err)
	// 将文件指针移动到 从文件头开始，偏移8位
	// 返回偏移的位置
	o2, err := f.Seek(8, 0)
	check(err)

	// 从当前文件指针位置开始读4个字节
	b2 := make([]byte, 4)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))
}

func ReadByBufio() {
	f, err := os.Open("./main.go")
	check(err)
	defer func() {
		err := f.Close()
		check(err)
	}()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			break
		}
		check(err)
		fmt.Print(line)
	}
}

func WriteByOpenFile() {
	// 如果文件不存在，则创建；写入时先请空文件再开始写
	f, err := os.OpenFile("tmp.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	check(err)
	defer f.Close()
	// 写入bytes
	f.Write([]byte("f.Write \n"))
	// 写入string
	f.WriteString("f.WriteString")

	// 通过bufio.NewWriter写入（基于缓存）
	writer := bufio.NewWriter(f)
	for i := 0; i < 3; i++ {
		writer.WriteString("bufio.NewWriter\n") // 先写入缓存
	}
	//writer.Flush() // 落盘
}

func WriteByIOUtil() {
	err := os.WriteFile("tmp.txt", []byte("Write by os.WriteFile\n"), 0644)

	//err := ioutil.WriteFile("tmp.txt", []byte("Write by ioutil.WriteFile\n"), 0644)
	check(err)
}

// JsonFile ReadWrite

func ReadJsonFileToMap(filename string) {
	jsonFile, err := os.Open(filename)
	check(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println(result["users"])
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}
type Users struct {
	Users []User `json:"users"`
}

func ReadJsonFileToStructAndDumpJson(filename string) {
	jsonFile, err := os.Open(filename)
	check(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
	f, _ := json.MarshalIndent(&users, "", "\t")
	ioutil.WriteFile("users_dump.json", f, 0644)
}

type Conf struct {
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

func WriteYaml() {
	conf := Conf{
		"test", "test",
	}
	data, _ := yaml.Marshal(&conf)
	fmt.Printf("%v\n", string(data))
	ioutil.WriteFile("config.yaml", data, 0644)
}

func main() {

	//ReadByOSReadFile()
	//ReadBySeek()
	//ReadByBufio()
	//ReadByIOUtil()
	//WriteByOpenFile()
	//WriteByIOUtil()
	//ReadJsonFileToStructAndDumpJson("users.json")
	WriteYaml()
}
