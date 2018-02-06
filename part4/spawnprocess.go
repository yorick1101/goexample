package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	ipCmd := exec.Command("ipconfig")
	//hold bytes
	ipOut, err := ipCmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(">ipconfig", string(ipOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep \n goodbye grep"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("grep hello:", string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(lsOut))

}
