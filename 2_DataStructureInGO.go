

______________________________________________________________________________________________________
1. Arrays
  fixed-length series of elements of a chosen type
  elements initialized to zero value //unlike other langs

  i. Declaration
      var x [5]int
      x[2] = 4
      //array literal is an array with predefined values like x ->
      var x [5]int = [5]{1,2,3,4,5}
      "..." for size in array literal(the second [5] above) infers size from number of initializers
      x:= [...]int{1,2,3,4,5}

  ii. Iterating
                through arrays //using for loop with the range keyword
  x:=[3]int{1,2,3}
  for i, v range x{       //i is the index, and v is the value in the array x, because ->
    //range returns two values, index and element at index
    fmt.Printf("index %d has value %d", i, v)
  }
______________________________________________________________________________________________________
2. Slices (more fexlible than arrays)
  is a window on an underlying array
  variable size, up to the whole array (only can change forward in the array, pointer remains static)
  "Pointer" indicates the start of the slice    *
  "Length" is the number of elements in slice   len()
  "capacity" is the maximum number of elements  cap()

  i. Declaration
    arr:= [...]string{"a", "b", "c", "d", "e", "f", "g"}  //not a slice, its a array
    s1 := arr[1:3]   //1 and 2 and not 3
    s2 := arr[2:5]   //not 5 only 2,3,4
    s2_5 := arr[]   //takes complete arr and hence is a slice literal (later)

  ii. Properties
    fmt.Printf(len(s1), cap(s2))
    changing a slice, changes the underlying array and any other slices overlapping
    slice literal is where the slice points to the start of the array and length = capacity

    s3 := []int{1,2,3,4,5}   //empty brackets means complete array is sliced and hence we know its slice literal
    note - if we put ... instead it would become array literal thats the difference
    three ways to make a slice, 1. using a array 2. directly initialize as slice literal 3. using make() (later), generally used.
  iii. Variable slices
    we can create a slice (and array) using make()
    sli = make(type , 10) 10 is length or capacity, type could be []int
    sli = make([]int, 10, 15)  10 is length and 15 is capacity
  iv. Append
    we can increase the size of a slice using append(), adds to the end, inserts into underlying arrays too
      it can increase the size of the array too, if neccesary
    sli = make([]int, 0)
    sli = append(sli, 100)
______________________________________________________________________________________________________
3 Hash Tables
  Contains key, value pairs  //key has to be unique
  Hash function is used to compute the slot for a key //behind the scenes though
______________________________________________________________________________________________________
4. Maps, golang implementation of hashtables
    var idMap map[keytype]valuetype  //string and int for example
    idMap = make(map[keytype]valuetype)
    i. Define a map literal:
        idMap := map[string]int {
          "joe":12 }
    ii. Accesing maps
        fmt.Println(idMap["joe"])  //prints zero if key not present
        idMap["jane"]=45   //adding a keyvalue pair
        delete(idMap, "joe")  //deletes
    iii. Two-value assigment tests for existence of the key
        id,p := idMap["joe"]  //"joe", id pair is added or updated to iMap, p is 1 if "joe" is present already
        //otherwise p is 0, when joe is not present in the iMap already, //p is presence of key
        len() returns number of values
    iv. Iterating through map, using for loop and range keyword
        for key, value := range idMap{
          fmt.Println(key, val)
        }
_____________________________________________________________________________________________________
5. Struct
  Aggregate data type.
    type struct Person {
      name string         //name address phone, all are called fields
      addr string
      phone string
    }

    var p1 Person
    p1.name = "joe"
    x = p1.addr

    //another way to make struct is new()
    p2 := new(Person) //initializes fields to zero
    // can initialize using a struct literal
    p3 := Person(name:"joe", addr: "a st.", phone: "123")
______________________________________________________________________________________________________




questions
type P struct {
    x string
    y int
}
func main() {
  b := P{"x", -1}
  a := [...]P{P{"a", 10},
        P{"b", 2},
        P{"c", 3}}
  for _, z := range a {
    if z.y > b.y {
      b = z
    }
  }
  fmt.Println(b.x)
}
//this function is basically to find the largest among the values of a array/slice of a specific struct


func main() {
  s := make([]int, 0, 3)
  s = append(s, 100)
  fmt.Println(len(s), cap(s))
}
output: 1 3
