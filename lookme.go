package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"golang.org/x/crypto/ssh"
)

type Mikortik struct {
	Username string `json:"user"`
	Password string `json:"pass"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

var ver string = "0.01 beta"

func about() {
	fmt.Printf("This Create By Erwin Kurniawan\n\tEmail  : rootshaxor@gmail.com\n\tGithub : github.com/rootshaxor\n\tBlog   : https://rootshaxor.github.io\n")
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <command>", os.Args[0])
	} else if os.Args[1] == "about" {
		about()
		return
	} else if os.Args[1] == "version" {
		fmt.Println(ver)
	}

	GooS := runtime.GOOS
	Goarch := runtime.GOARCH

	client, session, err := connectToHost()
	if err != nil {
		log.Fatalln(err)
	}

	out, err := session.CombinedOutput(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s %s Running on %s %s \n-------------------------------------------\n", os.Args[0], ver, GooS, Goarch)
	fmt.Println(string(out))
	client.Close()

}

func connectToHost() (*ssh.Client, *ssh.Session, error) {

	// Json Area
	JsonFile, err := os.Open("config.json")
	if err != nil {
		log.Println(err)
	}

	defer JsonFile.Close()

	ValueJson, _ := ioutil.ReadAll(JsonFile)

	var RB Mikortik

	json.Unmarshal(ValueJson, &RB)

	// Json End
	user := RB.Username
	pass := RB.Password
	host := RB.Host + ":" + RB.Port

	// fmt.Printf("Password : ")
	// bytePass, err := terminal.ReadPassword(syscall.Stdin)

	// pass := string(bytePass)

	sshConf := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConf.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConf)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
