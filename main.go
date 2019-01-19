package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Which encrypt algorithm? (sha256 OR md5) \n> ")
	algorithm, _ := reader.ReadString('\n')
	algorithm = strings.Replace(algorithm, "\n", "", -1)

	fmt.Print("Which password? \n> ")
	pwd, _ := reader.ReadString('\n')
	pwd = strings.Replace(pwd, "\n", "", -1)

	switch algorithm {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(pwd)))
		break
	case "md5":
		fmt.Printf("%x\n", md5.Sum([]byte(pwd)))
		break
	}

}
