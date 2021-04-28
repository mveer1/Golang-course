RFCs

p1 := Person(name:"joe", addr:"a st.")
has a equivalent json format as
{"name":"joe, "addr":"a st."}
 
json properties:
all unicode
types include, array, structs and primitive too

JSON marshalling, is generating JSON representation from an object
p1 := Person(name:"joe", addr:"a st.")
//Marshall function in json package
barr, err := json.Marshal(p1)
//unmarshalling
var p2 person
err:= json.Unmarshal(barr, &p2)   //note the ampersand
HERE, barr is byte array, i.e. the original json object
pointer to go object is passed to Unmarshal()
err is nil is everything works







file access, ioutil
linear access and not random access
basic operations
import "ioutil"
  open  //get handle for access
  dat, e := ioutil.ReadFile("test.txt")  //read bytes into []byte (dat here), it will read and close the file with it. all built in the same function
  err := WriteFile("outfile.txt", dat, 0777)  //write []byte into outfile.txt, 0777 is permission (unix style permision byte, anyone can do anything), cant append to a file, it always creates a new file
  close
  seek //moves read/write head

  io/ioutil package has this above basic functions

file access, os
os.Open() return a file descriptor
os.Read()
os.Write()
os.Close()
e.g.
    f, err := os.Open("dt.txt")
    barr := make([]byte, 10)  //that many characters we wants, could be less
    nb, err := f.Read(barr) //nb is number of bytes, check that once
    f.Close()

    f,err := os.Create("out.txt")
    barr := []byte{1,2,3}
    nb, err := f.Write(barr)  //write a []byte (byte array) i.e. any unicode sequence
    nb, err := f.WriteString("hi")
