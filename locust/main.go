package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	//fmt.Println("Starting Locust ", os.Args[7], " ", os.Args[5], " ", os.Args[9])
	cmdLine := "locust -f traffic.py --headless --users " + os.Args[7] + " -r " + os.Args[5] + " -t " + os.Args[9] + " -H http://34.70.59.65.nip.io/input"
	//fmt.Println(cmdLine)
	cmd := exec.Command("cmd.exe", "/c", "start "+cmdLine)
	err := cmd.Run()
	fmt.Printf("%s, error:%v \n", cmdLine, err)
}
