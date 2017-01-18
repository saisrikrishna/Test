package main

import (
        "bufio"
        "fmt"
	"os"
        "os/exec"
	"strings"
)

func main() {

        out, err := exec.Command("ls").Output()
        if err != nil {
                fmt.Println(err)
        }
        fmt.Printf("Files in the Directory are : \n%s\n", out)
        fmt.Print("Enter first file name with extension : ")
        var fname1 string
        fmt.Scanln(&fname1)
        fmt.Print("Enter second file name with extension : ")
        var fname2 string
        fmt.Scanln(&fname2)

	f1, er1 := os.Open(fname1)
	if er1!=nil{
		fmt.Println("Can't open File : ",er1)
	}
	sc1 := bufio.NewScanner(f1)

	f2, er2 := os.Open(fname2)
	if er2!=nil{
		fmt.Println("Can't open File : ",er2)
	}
	sc2 := bufio.NewScanner(f2)
	i := 0 
	c:=0
	for sc1.Scan() && sc2.Scan(){
		l1 := sc1.Text()
		l2 := sc2.Text()
		i++
		if strings.Compare(l1,l2)!=0 {
			fmt.Println(l1)
			fmt.Println(l2)
			fmt.Println("---------------------------------------------------")
			c++
		}
	}
	fmt.Println(c)
	fmt.Println(i)
}
