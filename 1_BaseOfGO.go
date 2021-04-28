Index - Basics, Datatypes, Pointers, Scope, Garbage collection, Control statements
iota, contants, go tools,
        Packages (Unicode, Strings, Strconv, fmt, os, bufio)
ascii, unicode, runes




GO Programming rough basics
Go, C C++ java  are compiled languages
python java are interpreted languages

1. Garbage collection - (like interpreted languages) us automatic memory management
2. Weakly Object oriented - no class, only struct s (group of data) with associated methods
    for structs, no inheritance, no constructors, no generics
3. Concurrency (because speed is needed)
    1. moores law states 18 months, transistors x2
    2. parrallelism is increasing number of cores, Go uses it
   Concurrency is management of multiple tasks at the same time
4. Go routine tasks represent Concurrency
    channels are used to communicate between tasks
    select enables task synchronisation

// Go tools
1. go build //creates a executable
2. go doc //print documentation for the package
3. go fmt //formats source code files
4. go get //download packages and installs them
5. go list //lists all installed packages
6. go run //runs tests using files ending in "__test.go"
_______________________________________________________________________________________________________________
Data types
  //declaration of variables
  var x int  //has the value 0
  var x,y int
  var x int = 100
  var x = 100
   // after var x int
      x=100
   x:=10 //short declaration

types of int- int8 int16 int32 int64  uint8 uint16 uint32 uint64 (u for unsigned)
note- //all int32 int16 int provide same accuracy
1. type conversion
var x float64 = 123.45
var y float32 = 1.23e2   //just to show it can be declared this way
x=y //would fail
x = float64(y)  // T() operator where t is the name of the datatype

float32 ~6 digit precision
float64 ~15 digit precision

var z complex128 = complex(1,2)

//constants
expression whose value is known at compile time

const x = 1.3 //going from righthandside, compiler decides that its type is floating point
const (
  y=4
  z="hi"
)

//iota
its like enum (enumeration) in c
generates a set of related but distinct constants, actual value not important,
  it just should be different
 e.g. days of the week
   type Grades int  //making a type called grades, aliasing
   const (
     A Grades = iota // just at the top
     B
     C
     D
     F
    )
    each constant is assigned to a unique integer, "generally" starts with A=1, B=2, ...
______________________________________________________________________________________________________
Pointers in go
operators
1. & returns the address of a var/function  // you put this infront of a var/function
2. * returns the data at an Object (derefrencing) //you put this infront of a pointer

var x int = 1
var y int
var ip *int  //ip is pointer to int
ip = &x
y = *ip
      func main() {
          var x int
          var y *int
          z := 3
          y = &z
          x = &y
        }   //does not compile

// note alternate way to create var
new(type) create a variable and returns a pointer to the variable
ptr := new(int)
*ptr = 3

______________________________________________________________________________________________________
Scope
1. universal blocks //all go source code
2. packages blocks
3. file blocks
4. if for swtich blocks

go follows lexical scopping (first look in the same block then move above to a higher)
______________________________________________________________________________________________________
Deallocating Space
stack memory - dedicated to function calls (so local variables are stored here)
  when the function ends, its deallocated
heap memory - we have to explicitly deallocate it ( in .c using free(x); )

//garbage collection
func foo() *int  {
  x:=1
  return &x
}
func main()  {
  var y *int
  y=foo()
  fmt.printf("%d", *y)
}   //here x wont get deallocated because its refrence is present in y ig
______________________________________________________________________________________________________
Control Statements
if x>5 {
    fmt.Printf("y")
  } else {
    ...
}

for //may have an initialization and update operation
    for <init>; <cond>; <update> {   //init once before. cond everytime before, update et end
        <statements>
     }
             1. for i:=0; i<10; 1++{}
             2. i = 0
                for i<10{
                  ...
                  i++
                }
              3. for {...} //infinite
switch
      switch <expression> {
      case <condition>:
        }
        switch x {
        case 1:
          fmt.Printf("case 1")
        case 2:
          fmt.Printf("case 2")
        default:
          fmt.printf("nocase")
        }
        // no need of break

    tagless/caseless switch
        switch  {
        case x>1:
          ...
        case x<-1:
          ...
        default:
          ...
        } //once anyone comes tru, it breaks the switch block, if none is tru only then default
break and continue
      for i<10{
          i++
          if i==5 {
            break
            //continue
          }
          fmt.printf("continue will skip one iteration")
        }
______________________________________________________________________________________________________
______________________________________________________________________________________________________
Packages
one package must be named main and should have a function main, only this packages
gets converted into a exe
______________________________________________________________________________________________________
1. fmt package  (overview of doc contained)
Printing
  The Print() function simply prints a list of variables to STDOUT with default formatting.
  The Printf() function allows you to specify the formatting using a format template.
  The Println() function works like Print() except it inserts spaces between the variables and appends a new line at the end.

  fmt.printf("hi"+x)
  fmt.printf("hi %s", x)

Scanning
  scan reads user input, takes a pointer as an argument, typed data is written to pointer
  It returns two things
    1. number of scanned items
    2. nil in case of no error, error code in case of a error

    fmt.printf("Enter number of apples")
    num, err := fmt.scan(&appleNum)   //note the & too
    //num has value 1, error has nil if correctly entered, applenum will have what we enter


// ascii and unicode
ascii is a 8 bit per character representation
unicode is a 32 bit per character representation
UTF-8 is variable length
"unicode codepoints" or we say "rune" in Go are characters, like 6 runes in x:="hel lo"

______________________________________________________________________________________________________
2. unicode package (complete doc contained)
  strings are basically a array (or raid) of runes
  runes are divided in many categories
all this give 1 or 0 as outputs (returns)
    //func In(r rune, ranges ...*RangeTable) bool   //rangetable is 16bit, 32bit etc
      In reports whether the rune is a member of one of the ranges.
    //func Is(rangeTab *RangeTable, r rune) bool   //rangetable is 16bit, 32bit etc
      Is reports whether the rune is in the specified table of ranges.
    //func IsControl(r rune) bool
      IsControl reports whether the rune is a control character.
      The C (Other) Unicode category includes more code points such as surrogates; use Is(C, r) to test for them.
    func IsDigit(r rune) bool
      IsDigit reports whether the rune is a decimal digit.
    func IsGraphic(r rune) bool
      IsGraphic reports whether the rune is defined as a Graphic by Unicode.
      Such characters include letters, marks, numbers, punctuation, symbols, and spaces, from categories L, M, N, P, S, Zs.
    func IsLetter(r rune) bool
      IsLetter reports whether the rune is a letter (category L).
    func IsLower(r rune) bool
      IsLower reports whether the rune is a lower case letter.
    //func IsMark(r rune) bool
      IsMark reports whether the rune is a mark character (category M).
    //func IsNumber(r rune) bool
      IsNumber reports whether the rune is a number (category N)..
    //func IsPrint(r rune) bool
      IsPrint reports whether the rune is defined as printable by Go.
      Such characters include letters, marks, numbers, punctuation, symbols,
      and the ASCII space character, from categories L, M, N, P, S and the ASCII space character.
      This categorization is the same as IsGraphic except that the only spacing character is ASCII space, U+0020.
    func IsPunct(r rune) bool
      IsPunct reports whether the rune is a Unicode punctuation character (category P).
    func IsSpace(r rune) bool
      IsSpace reports whether the rune is a space character as defined by Unicode's White Space property; in the Latin-1 space this is
      '\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
      Other definitions of spacing characters are set by category Z and property Pattern_White_Space.
    //func IsSymbol(r rune) bool
      IsSymbol reports whether the rune is a symbolic character.
    //func IsTitle(r rune) bool
      IsTitle reports whether the rune is a title case letter.
    func IsUpper(r rune) bool
      IsUpper reports whether the rune is an upper case letter.
  conversions
    ToUpper(r rune)
    ToLower(r rune)
______________________________________________________________________________________________________
3. strings package  (complete doc contained)
string manipulation (modified strings are 'returned', as strings are immutable)

strings.Contains("test", "es")
s.ContainsAny("test", "sm")    *details down
s.Compare("a", "b")    //0 if a==b, -1 if a<b, 1 if a>b
s.Count("test", "t")
s.EqualFold("Go", "go") bool //true, basically for case-insensitivity
s.Fields("abc   nc ") []string  //returns ["abc", "nc"]  //empty slice if only space
s.FieldsFunc(s string, f func(rune) bool) []string   *details down //as per f
s.HasPrefix("test", "te")
s.HasSuffix("test", "st")
s.Index("test", "est"))    //returns 1 here, -1 is didnt find
s.IndexAny(s, chars string) int   *details down  |  s.IndexByte(s string, c byte) int  just ''  instead of ""  | s.IndexFunc(s string, f func(rune) bool) int  *details down  | s.LastIndex(s, substr string) int   //returns last index instead of first index of the substr found   //above has variantions to byte, any, func
s.Join([]string{"a", "b"}, "-")
s.Repeat("na", 2))  //It panics if count is negative or if the result of (len(s) * count) overflows.
s.Replace("fooo", "o", "0", -1)   replaces everywhere | s.ReplaceAll(str, oldstr, newstr string) //same as -1
s.Replace("fooo", "ni", "in", 2)  //replaces at first two places //s.Replace("fooo", "o", "0", 1)  at first place
s.Split(s, sep string) []string  |  s.SplitAfter(s, sep string) []string   *details down  |  s.SplitAfterN(s, sep string, n int) []string   *details down  |  s.SplitN(s, sep string, n int) []string
s.ToLower("TEST")
s.ToUpper("test")
s.TrimSpace(s) - whitespace from front and back removed
s.Trim("¡¡¡Hello, Gophers!!!", "!¡"))   //Hello, Gophers
s.TrimFunc(s string, f func(rune) bool) string  *details down   |   Trimleft,right,leftfunc,rightfunc,space,suffix,prefix

-------------------------details down---------------------------
P(s.Contains("seafood", ""))
P(s.Contains("", ""))
//both are true HERE

p(s.ContainsAny("failure", "ui"))  //true
p(s.ContainsAny("foo", ""))         //false
p(s.ContainsAny("", ""))          //false
------------------------------------------------------------------------------------------------
s.FieldsFunc
func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}


s.IndexAny("chicken", "aeiouy") //2
s.IndexAny("crwth", "aeiouy")  //1
------------------------------------------------------------------------------------------------
s.IndexFunc
func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	s.IndexFunc("Hello, 世界", f) //7
	s.IndexFunc("Hello, world", f) } //-1

------------------------------------------------------------------------------------------------
func Map
s.Map(mapping func(rune) rune, s string) string
  //Map returns a copy of the string s with all its characters modified according to the mapping function.
   //If mapping returns a negative value, the character is dropped from the string with no replacement.
  func main() {
  	rot13 := func(r rune) rune {
  		switch {
  		case r >= 'A' && r <= 'Z':
  			return 'A' + (r-'A'+13)%26
  		case r >= 'a' && r <= 'z':
  			return 'a' + (r-'a'+13)%26
  		}
  		return r
  	}
  	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))
  }

  'Gjnf oevyyvt naq gur fyvgul tbcure...

------------------------------------------------------------------------------------------------
s.SplitAfter(s, sep string) []string
SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.
  If s does not contain sep and sep is not empty, SplitAfter returns a slice of length 1 whose only element is s.
  If sep is empty, SplitAfter splits after each UTF-8 sequence. If both s and sep are empty, SplitAfter returns an empty slice.
  It is equivalent to SplitAfterN with a count of -1.
strings.SplitAfter("a,b,c", ","))  //["a," "b," "c"]

s.SplitAfterN(s, sep string, n int) []string
SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.
The count determines the number of substrings to return:
n > 0: at most n substrings; the last substring will be the unsplit remainder.
n == 0: the result is nil (zero substrings)
n < 0: all substrings
s.SplitAfterN("a,b,c", ",", 2))   //["a," "b,c"]

s.SplitN(s, sep string, n int) []string
SplitN slices s into substrings separated by sep and returns a slice of the substrings between those separators.
The count determines the number of substrings to return:
...(same as above) example lite

------------------------------------------------------------------------------------------------
TrimFunc returns a slice of the string s with all leading and trailing Unicode code points c satisfying f(c) removed.
func main() {
	fmt.Print(strings.TrimFunc("¡¡¡Hello123, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}
output - Hello123, Gophers
______________________________________________________________________________________________________
4. Strconv Package  (overview of doc contained)
(conversions to and from string representation of basic data types)
    Atoi(s) - ascii to integer
    Itoa(s) - converts int(base 10) to string
    FormatFloat (f, fmt, prec, bitSize) - converts floating point number to a string
    ParseFloat (s, bitSize) - Converts a string to a floating point number
func main() {
  i, _ := strconv.Atoi("10")
  y := i * 2
  fmt.Println(y)
  }
//20 ??
------------------------------------------------------------------------------------------------
Numeric Conversions
i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)
------------------------------------------------------------------------------------------------
ParseBool, ParseFloat, ParseInt, and ParseUint convert strings to values:

b, err := strconv.ParseBool("true")
f, err := strconv.ParseFloat("3.1415", 64)
i, err := strconv.ParseInt("-42", 10, 64)
u, err := strconv.ParseUint("42", 10, 64)
The parse functions return the widest type (float64, int64, and uint64),
but if the size argument specifies a narrower width the result can be converted to that narrower type without data loss:
s := "2147483647" // biggest int32
i64, err := strconv.ParseInt(s, 10, 32)
...
i := int32(i64)

------------------------------------------------------------------------------------------------
FormatBool, FormatFloat, FormatInt, and FormatUint convert values to strings:

s := strconv.FormatBool(true)
s := strconv.FormatFloat(3.1415, 'E', -1, 64)
s := strconv.FormatInt(-42, 16)
s := strconv.FormatUint(42, 16)
AppendBool, AppendFloat, AppendInt, and AppendUint are similar but append the formatted value to a destination slice.
______________________________________________________________________________________________________
5. Package os  (overview contained)
Package os provides a platform-independent interface to operating system functionality.
The design is Unix-like, although the error handling is Go-like; failing calls return values of type error rather than error numbers.
Often, more information is available within the error.
For example, if a call that takes a file name fails, such as Open or Stat, the error will
  include the failing file name when printed and will be of type *PathError, which may be unpacked for more information.
The os interface is intended to be uniform across all operating systems.
Features not generally available appear in the system-specific package syscall.

Here is a simple example, opening a file and reading some of it.

file, err := os.Open("file.go") // For read access.
if err != nil {
	log.Fatal(err)
}
If the open fails, the error string will be self-explanatory, like

open file.go: no such file or directory
The file's data can then be read into a slice of bytes. Read and Write take their byte counts from the length of the argument slice.

data := make([]byte, 100)
count, err := file.Read(data)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("read %d bytes: %q\n", count, data[:count])
Note: The maximum number of concurrent operations on a File may be limited by the OS or the system.
The number should be high, but exceeding it may degrade performance or cause other issues.
______________________________________________________________________________________________________
6. Buffio (check some vdo or something)
to scan a complete line without breaking, these three lines
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()
s := scanner.Text()
______________________________________________________________________________________________________
7. Sort 
