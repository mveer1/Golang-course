package main
import ("strings"; "fmt"; "bufio"; "os")
var p = fmt.Println

func main()  {
  scanner := bufio.NewScanner(os.Stdin)
  //var s string
  p("Enter a string to search ")
  scanner.Scan()
  s := scanner.Text()
  p("the string is:", s)
  if (strings.HasPrefix(s, "i") || strings.HasPrefix(s, "I")) && (strings.HasSuffix(s, "n")|| strings.HasSuffix(s, "N")) && (strings.Contains(s, "a") || strings.Contains(s, "A")) {
    p("Found!")
  } else {
    p("Not found!")
   }

}



package main

import ("fmt"
        "sort"
     sc "strconv")
var p = fmt.Println

func main()  {
  var x string
  slice := make([]int, 0, 3)
  for {
    p("Enter a Integer")
    fmt.Scan(&x)
    if x=="X" {break}
    y,_ := sc.Atoi(x)
    slice = append(slice, y)
   }
  sort.Ints(slice)
  for _, a := range slice{
    fmt.Print(a, " ")
  }
}




package main
import ("fmt"
        "encoding/json"
        "os"
        "bufio")
var p = fmt.Println
func main()  {

  scanner := bufio.NewScanner(os.Stdin)
  p("Enter Your Name")
  scanner.Scan()
  name := scanner.Text()

  p("Enter Your Address ")
  scanner.Scan()
  addr := scanner.Text()

  myMap := map[string]string { "name":name, "address":addr }
  barr, _ := json.Marshal(myMap)
  p(string(barr))
}





package main
//const (
    //maxLength = 20
//)
import ("fmt";"os";"io";"bufio";"reflect";"strings")

var pl = fmt.Println
var pf = fmt.Printf
var p = fmt.Print

type S1 struct{           //struct defination
  fname string        //the size has to be of 20 bytes
  lname string
}
func main()  {
  //ReadLine() or ReadString('\n') are different
  var s S1
  filename := "ab.txt.txt"
  f,_ := os.Open(filename)
  reader := bufio.NewReader(f)
  a := []S1{}
  for {
    line, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {break}
    loc:= strings.Index(line, " ")
    line = strings.TrimSpace(line)
    fname := line[0:loc]
    lname:= line[loc+1:]
    if err != nil {break}
    //pl(fname, lname)
    s.fname = fname
    s.lname = lname
    a = append(a,s)

  }
  // a is the required slice
  pl(reflect.ValueOf(a).Kind())
  for _, v := range a{
    pl(v.fname, v.lname)
  }
}


//same better assigment
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main(){

	var filePath string
	namesSlice := make([]Name, 0)

	fmt.Println("Enter file path:")
	fmt.Scan(&filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("could not open specified file")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names := strings.Split(scanner.Text(), " ")
		namesSlice = append(namesSlice, Name{fname: names[0], lname: names[1]})
	}

	file.Close()

	fmt.Println("names found: ")
	for _,n := range namesSlice {
		fmt.Println(n.fname, n.lname)
	}

}




//json
package main
import ("fmt"
        "encoding/json"
        "os"
        "bufio")
var p = fmt.Println
func main()  {

  scanner := bufio.NewScanner(os.Stdin)
  p("Enter Your Name")
  scanner.Scan()
  name := scanner.Text()

  p("Enter Your Address ")
  scanner.Scan()
  addr := scanner.Text()

  myMap := map[string]string { "name":name, "address":addr }
  barr, _ := json.Marshal(myMap)
  p(string(barr))
}
