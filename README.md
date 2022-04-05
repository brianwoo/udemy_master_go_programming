### *Notes from Course: Master Go (Golang) Programming:The Complete Go Bootcamp 2022* 
### *- by Andrei Dumitrescu*
<br/><br/>


# Go directory structure:
```
go/src
  /pkg
  /bin
```
  
Create project directory under src.

# Compile and build
```bash
# init go module: 
go mod init

go run main.go
go build -o main.go
```


# GoFmt
```bash
gofmt -w -l <directory>
gofmt -w <file>
# or just do: go fmt

# -w write changes to file
# -l show which file got formatted
```


# FMT Package
Println - added newline at the end\
Printf  - no added newline

```
%d ->   decimal
%f ->   float
%.3f -> float (3 decimal points)
%s ->   string
%q ->   double-quoted string
%v ->   value (any)
%#v ->  a Go-syntax representation of the value
%T ->   value Type
%t ->   bool (true or false)
%p ->   pointer (address in base 16, with leading 0x)
%c ->   char (rune) represented by the corresponding Unicode code point
```

# Convert String to Numbers and vice versa
```golang
String(99)                // c  (ascii)
fmt.Sprintf("%f", 44.2)   // "44.2000"

// strconv package from string to Numbers 
strconv.ParseFloat("3.123", 64)   // 3.123 (float64) 

strconv.Atoi("-50") // -50
strconf.Itoa(20)    // "20"
```

# Loops
```golang
numbers := []int{1,2,3,4}
for i:=0; i <= len(numbers); i++ {}  // for
for n < 5 { n++ }                    // while loop
for i,v := range numbers {}          // foreach
for {}                               // infinite loop
```

*NOTE:* if numbers is nil, range will just skip it


# Arrays
**Ellipsis operator:** 
```golang
[...]int[1,2,3,4,5,6]  // Any number of elements
```

**Multi dimension array:** 
```golang
array := [2][3]int{
  [1,2,3],
  [4,5,6],
}
```

**Array are copied when assigned:**
```golang
m := 3[int]{1,2,3}
n := m   // n is a copy of m, not a reference
```

**Comparing arrays:**
```golang
if (n == m) {  // true if same num of elements, order and values

}
```

# Slice
**Reference type**\
Slice header:
  - pointing to beginning addr (first element) of the backing array
  - len of the slice
  - capacity, len calc. from starting element 
    to end element of the backing array
  - when one slice is assigned to another slice s1 := s2, s2's slice header is copied to s1. 

```golang
var cities []string  // when cities not initialized, it's nil
cities[0] = "London" // error, cannot assign to a nil slice
```

**Initialize a slice (2 ways):**
```golang
numbers1 := []int{1,2,3,4}
numbers2 := make([]int, 4)
```

**Comparing slices:**
```golang
n, m := []int{1,2,3}, []int{1,2,3}
if (n == m) {  // ERROR! cannot compare slices. 
               // slice can only be compared to nil.
}              // Need to use a loop to compare slices
```

**Appending an element to a slice:** \
*NOTE: Appending creates a new slice* \
*If no more space in the backing array, a new backing array will be created* 
```golang
numbers1 = append(numbers1, 4)     // append returns a new slice
numbers1 = append(numbers1, 5,6,7) // append can take more than 1 value
n := []int{8,9}
numbers1 = append(numbers1, n...)  // ellipsis like spread in JS
                                   // ... turns slice into individual values

E.g. Appending
s1 := []string{"a","b","c","d","e"}
s2 := []string{}
s2 = append(s2, s1[0:2]...)         // s2=["a","b"]

E.g. Cutting
s1 := []int{1,2,3,4,5}
s2 := s1[0,3]  // s2=[1,2,3]
l := len(s2)   // l=3 (the slice has 3 elements)
c := cap(s2)   // c=5 (backing array has 5 elements)
// *Note: capacity measured from starting element*

E.g. Capacity from starting element (looking at slice header)
s1 := []int{1, 2, 3, 4, 5}  // addr: 0x00 len(s1)=5 cap(s1)=5
n1 := s1[1:3]     // addr: 0x08 (1*8 bytes) n1=[2,3] len(n1)=2 cap(n1)=4
n2 := s1[1:4]     // addr: 0x08 (1*8 bytes) n1=[2,3,4] len(n2)=3 cap(n2)=4
n3 := s1[2:5]     // addr: 0x10 (2*8 bytes) n1=[3,4,5] len(n3)=3 cap(n3)=3
```

**Copy slice:**
```golang
src := []int{1,2,3}
dst := src                    // this only copies the slice header

// To really copy a slice
src := []int{1,2,3}
dst := make([]int, len(src))  // if len is 0, nothing is copied
nn := copy(dst, src)          // nn is the number of elements copied
```

**Get a slice from an array**
```golang
a := [5]int{1,2,3,4,5}
b := a[1:3]   // b is [2,3], ending index is exclusive
c := a[2:]    // c is [3,4,5], from index to the end
d := a[:3]    // d is [1,2,3], from 0 index to index 2
e := a[:]     // e is [1,2,3,4,5]
```

**Backing array**
```golang
// Changing an element in a slice affects other slices derived from it
s1 := []int{1,2,3,4,5}
s3,s4 := s1[0:2], s1[1:3]   // s3=[1,2], s4=[2,3]
s3[1] = 60                  // s3=[1,60], s4=[60,3], s1=[1,60,3,4,5]

// Changing an array affects slices derived from it
a1 := [4]int{1,2,3,4}
s1,s2 := a1[0:2], a1[1:3]   // s1=[1,2], s2=[2,3]
a1[1] = 20                  // s1=[1,20], s2=[20,3]
```

**Bonus: element appending demo**
```golang
func main() {
	maxValue := 20
	result := make([]int, 0, maxValue) // cap is 20
	result = append(result, 88)
	result2 := result
	for i := 0; i < maxValue; i++ {
		if i%2 == 0 {
			fmt.Printf("appending '%d': %s   %s\n", i, 
				getSliceHeader(&result), getSliceHeader(&result2))
			// Appending to result changes the len.
			// Note that len in result2 is NOT changing 
			// (coz result2 has a diff header), even though 
			// result and result2 are pointing to the same
			// backing array adress.
			result = append(result, i)
			fmt.Printf("appended '%d': %s   %s\n", i, 
				getSliceHeader(&result), getSliceHeader(&result2))
		}
	}

	// Even though result and result2 are 
	// pointing to the same backing array adress,
	// result2 Len remains the same when appending to result.
	// This explains why result and result2 are different
	// when they are printed out.
	fmt.Println(result)
	fmt.Println(result2)
}

// https://stackoverflow.com/a/54196005/463785
func getSliceHeader(slice *[]int) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}


Output:
appending '0': &{Data:824634777600 Len:1 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '0': &{Data:824634777600 Len:2 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '2': &{Data:824634777600 Len:2 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '2': &{Data:824634777600 Len:3 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '4': &{Data:824634777600 Len:3 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '4': &{Data:824634777600 Len:4 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '6': &{Data:824634777600 Len:4 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '6': &{Data:824634777600 Len:5 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '8': &{Data:824634777600 Len:5 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '8': &{Data:824634777600 Len:6 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '10': &{Data:824634777600 Len:6 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '10': &{Data:824634777600 Len:7 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '12': &{Data:824634777600 Len:7 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '12': &{Data:824634777600 Len:8 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '14': &{Data:824634777600 Len:8 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '14': &{Data:824634777600 Len:9 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '16': &{Data:824634777600 Len:9 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '16': &{Data:824634777600 Len:10 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appending '18': &{Data:824634777600 Len:10 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
appended '18': &{Data:824634777600 Len:11 Cap:20}   &{Data:824634777600 Len:1 Cap:20}
[88 0 2 4 6 8 10 12 14 16 18]
[88]

```

# String
```golang
// Regular string, need to escape quotes: 
s1 := "abcde \"esc string\""
// Raw String 
s2 := `abcde "even in quote"`
// Multi-line String
s3 := `
line 1
line 2
line 3
`

// Concat Strings
s4 := "I love " + "Go Programming"

// Note: Strings are IMMUTABLE
// Error:
s4[0] = 'U'
```

**String - Ascii & Rune characters**
```golang
// 'A' is a char aka a rune (int32)
c1 := 'A'  
fmt.Printf("c1 Size: %T, Value: %d", c1, c1) // Size: int32, Value: 65

// A string is a slice of uint8 or int32
s1a := "Go is cool!"
fmt.Println(s1a[0])         // Ascii code 71
fmt.Println(string(s1a[0])) // string() makes a ascii code a string: "G"
fmt.Println(s1a[0:2])       // s1a[0:2] is a slice i.e. a string: "Go"

// A char in a string can be uint8 (1 byte ASCII) 
// or int32 (unicode) type.
// For unicode, it can be a variable 1-4 bytes size.
s2 := "AB"                                             // Ascii chars
for i := 0; i < len(s2); i++ {
  fmt.Printf("s2 Type: %T, Value: %c\n", s2[i], s2[i]) // A & B are uint8
}

s1 := "世界"                                           // Unicode chars
fmt.Printf("Len of s1: %d\n", len(s1))                // Len of s1: 6

// Because unicode chars can be variable size,
// we cannot print each char byte-by-byte
// Fail:
for i := 0; i < len(s1); i++ {
  fmt.Printf("%c", s1[i])
}

// Success (** PREFERED option **):
for _, r := range s1 {
  fmt.Printf("%c", r)              // 世界 (3 bytes each)
}

// Success (option 2):
for i := 0; i < len(s1); {
  r, size := utf8.DecodeRuneInString(s1[i:])
  fmt.Printf("%c", r)              // 世界 (3 bytes each)
  i += size
}
```

**Slicing String - Ascii & Rune characters**
```golang
// Slicing a string works only when strings contains ASCII chars.
// That's because slicing returns bytes (not runes)
// E.g. ASCII string:
s1 := "abcdefghijkl"
fmt.Println(s1[2:5]) // --> cde


// Unicode string (contains runes which are >1 byte)
s2 := "我喜歡編程視頻遊戲"
fmt.Println(s2[0:2]) // -> � - the unicode representation of 
                     //         bytes from index 0 to 1

// To slice a unicode string
// 1. convert the string into a rune slice
rs2 := []rune(s2)         // string --> []int32

// 2. slice it from the rune slice
s3 := string(rs2[0:3])    // s3 = "我喜歡"
```

**String functions**
```golang
result := strings.Contains("I love food", "love")  // result=true
result := strings.Contains("I love food", "lovex") // result=false

// Any of the characters (x or y) matches
result := strings.ContainsAny("I love food", "vy") // result=true

// ContainsRune (contains a single char)
result := strings.ContainsRune("I love food", 'e') // result=true

// Count how many of "o"
result := strings.Count("I love food", "o")      // result=3
result := strings.Count("I love food", "oo")     // result=1
result := strings.Count("Five", "")              // result=5 (4+1 chars)

// Lower and Upper
result := strings.ToLower("I loVe fOOd")         // result="i love food"
result := strings.ToUpper("I loVe fOOd")         // result="I LOVE FOOD"

// Comparison
// Can use == for case sensitve, for case insensitive comparison:
result := strings.EqualFold("GO", "go")          // result=true

//Repeating 
result := strings.Repeat("Go", 5)                // GoGoGoGoGo

// Replace and ReplaceAll
// 2 means first 2 occurences. -1 means all.
result := strings.Replace("1.2.3.4", ".", "-", 2)  // 1-2-3.4
result := strings.Replace("1.2.3.4", ".", "-", -1) // 1-2-3-4
result := strings.ReplaceAll("1.2.3.4", ".", "-")  // 1-2-3-4

// Split: into a slice of strings
result := strings.Split("1.2.3", ".") // []string{"1","2","3"}
result := strings.Split("12 34", "")  // []string{"1", "2", " ", "3", "4"}

slice := []string{"1", "2", "3"}
result := strings.Join(slice, ":")      // "1:2:3"

// Fields: split by space & newline into a slice
result := strings.Fields("1 2 \n 3 4")  // []string{"1", "2", "3", "4"}

// Trim: remove leading and trailing "space-like" char of the string
result := strings.TrimSpace("\t 1 2 3 \n")     // "1 2 3"

// Trim: remove specific chars in a string
result := strings.Trim("\t1.2.3! \n", " \t\n") // "1.2.3!"
```

# Map
- Reference Type
  - Map Header (addr of data structure) --> data structure
  - When a map is assigned to another map (m1 := m2), the header of m2 is copied to m1.
  - Both map headers m1 & m2 then point to the same data structure
- Cannot use Float as a key
- A Map cannot be compared to another Map. It can only be compared to nil.

**Declaring and Initializing a Map**
```golang
var m1 map[string]string     // m1=nil
var m1 map[[]int]string      // Error: slice not comparable, can't be key
var m1 map[[3]int]string     // Array is comparable, ok to be key

// initializing a map
m1 := map[string]string{}      // initializing a Map (Option 1)
m1 := make(map[string]string)  // initializing a Map (Option 2)
m1 := map[string]string{       // initializing a Map with values
  "key1": "value1",
  "key2": "value3",
}
```

**Setting, Deleting and Getting values from a Map**
```golang
var m1 map[string]string     // m1=nil
s1 := m1["anykey"]           // s1="", even though map is nil (Surprise!)
m1["anykey"] = "value"       // Error: cannot assign value to nil map

// To distinguish Map key not exist or value is actually empty
v, ok := m1["anykey"]        // if key exists, ok returns true

// looping over a Map
for k,v := range m1 {        // k=key, v=value
}

// Deleting a key in a Map
delete(m1, "anykey")         // delete "anykey" in the map
```

**Comparing Maps**
```golang
m1 := map[string]string{"a":"x"}
m2 := map[string]string{"a":"y"}
if m1 == m2 {                     // Error! Cannot compare 2 maps
                                  // Can only compare a map to nil
}

// Workaround, using Sprintf
s1 := fmt.Sprintf("%s", m1)       // "map[a:x]"
s2 := fmt.Sprintf("%s", m2)       // "map[a:y]"
if (s1 == s2) {                   // Can compare 2 strings

}
```

**Cloning Map**
```golang
m1 := map[string]int{"A": 40, "B":25}
m2 := m1                               // This only copies the map header

m1 := map[string]int{"A": 40, "B":25}
m2 := make(map[string]int)	           // create another map
for k,v := range m1 {                  // clone by copying each key/value
  m2[k] = v
}
```

# Structs
**Creating Struct**
```golang
// Defining a struct
type s1 struct {
  f1 string
  f2 int
  f3, f4 string
}

// Initializing a struct
myS1 := s1{"xx", 22, "yy", "zz"}  // just by order (not a good way)
myS1 := s1{f1: "xx", f2: 22, f3: "yy", f4: "zz"}  // by field name, (good)
myS1 := s1{f1: "xx"}     // init only f1, the rest are 0 and empty string


// Anonymous struct, not from a type
myS2 := struct {
  f1     string
  f2     int
  f3, f4 string
}{
  f1: "xx",
  f2: 22,
  f3: "yy",
  f4: "zz",
}
```

**Retrieving / Updating Struct**
```golang
type s1 struct {
  f1 string
  f2 int
  f3, f4 string
}

// Retrieving values
myS1 := s1{f1: "xx"} 
field1 := myS1.f1                // field1="xx"
field1 := myS1.nonExistingField  // Error

// Updating values
myS1.f2 = 12
```

**Comparing Struct**
```golang
myS1 := s1{f1: "xx", f2: 22, f3: "yy", f4: "zz"}
myS2 := s1{f1: "xx", f2: 22, f3: "yy", f4: "zz"}
result := myS1 == myS2     // true
```

**Copying a Struct**
```golang
myS1 := s1{f1: "xx", f2: 22, f3: "yy", f4: "zz"}
myS2 := myS1        // this copies the entire myS1 struct over to myS2
```

**Embedded Struct**
```golang
type embedded struct {
  ef1 string
  ef2 int
}

type s1 struct {
  f1         string
  f2         int
  embeddedF3 embedded
}

myS1 := s1{
  f1: "xx",
  f2: 22,
  embeddedF3: embedded{
    ef1: "yy",
    ef2: 33,
  },
}

result := myS1.embeddedF3.ef1      // retrieving
myS1.embeddedF3.ef1 = "zz"         // setting

```

**Bonus: Anonymous Struct Fields**
```golang
type s1 struct {
  f1 string
  int
}

as1 := s1{"xx", 22}        // OK
as2 := s1{f1:"xx", 22}     // Error: mixture of field:value and value elements in struct
result := as1.f1           // "xx"
as1.f1 = "yy"
```

# Functions
- No Function overloading
- Use camelCase
- No Default arguments in GoLang
- Arguments passed by value

**Function, arguments and return type**
```golang
func f1() {
}

func f1(a1 int, a2 int) {
}

func f2(a1, a2 int) {           // shorthand arguments
}

func f2(a1, a2 int) float64 {   // returning a float64
}

func f2(a1, a2 int) (int, int) {   // returning multiple values
}

func f2(a1, a2 int) (s int) {   // s is a named return value
  s = a1 + a2
  return                        // this is called returning a naked value
}

```

**Variadic Functions - 0 or more args**
```golang
func f1(a1 ...int) {
	fmt.Printf("%#v\n", a1)  // a1 is represented as an int slice
}

f1(1, 2, 3, 4)   // a1 arg is []int{1, 2, 3, 4}
f1()             // a1 arg is []int(nil)

args := []int{1, 2, 3, 4}
f1(args...)      // a1 arg is []int{1, 2, 3, 4}

// NOTE: Only the last arg can be variadic
func f1(a1 string, a2 ...int) {
}
```

**Variadic Functions - modifying passed in arg's value**
```golang
func f1(a1 ...int) {
  a1[0] = 88               // modifying arg 0's value
}

args := []int{1, 2, 3, 4}
f1(args...)                // args is a slice which is passed in as a ref
result := args             // []int{88, 2, 3, 4}
```

**Defer statement and LIFO**
```golang
func f1() {
	fmt.Println("f1")
}

func f2() {
	fmt.Println("f2")
}

func f3() {
	fmt.Println("f3")
}

func f4() {
	fmt.Println("f4")
}

func main() {
	defer f1()     // defer works like a stack, f1()
	f2()
	f3()
	defer f4()     // f4() is now at the top of the stack
}

Output:
f2
f3
f4            // f4() is executed first before f1() 
f1            // because it's at the top of the stack
```

**Anonymous Functions**
```golang
func(msg string) {      // anonymous func has no name so we have to
  fmt.Println(msg)      // self executing it.
}("Anonymous function!")         

// Create a function that return a function
func increment(x int) func() int {
	return func() int {    // Returning a function that returns an int
		x++
		return x
	}
}

// To execute:
incrementFunc := increment(5)
result := incrementFunc()         // result=6
result := incrementFunc()         // result=7
result := incrementFunc()         // result=8
```

# Pointers
- A pointer is a variable that stores an address of another variable
- GoLang does not support pointer arithmetic.

**Pointer illustration**
```ocaml
PointerA (at 0xE020)         |---->    VariableA (at 0x10250)
value: 0x10250         -------         value: "Batman"
```

**Pointer and address**
```golang
v := "Batman"
result := &v                           // &v = get the addr of v
fmt.Printf("%T %v\n", result, result)  // *string 0x10250
fmt.Printf("%p\n", result)      // %p = pointer, addr of the var: 0x10250
fmt.Printf("%p\n", &result)     // Addr of the pointer 0xE020

// Create a pointer without assigning anything to it
var ptr1 *float64                   // * in front of type means pointer type
fmt.Printf("%v %p\n", ptr1, &ptr1)  // nil, 0xABCD

// Create a pointer with a new keyword
ptr1 := new(int)
x := 100
ptr1 = &x
fmt.Printf("%T %v\n", ptr1, ptr1)   // *int 0xABCD
fmt.Printf("Addr of x: %p\n", &x)   // Addr of x: 0xABCD

*ptr1 = 33      // Change the value via pointer, same as x = 33
                // * in front of a variable means dereferencing
```

**A Pointer to a Pointer**
```ocaml
pp1 (at 0x1111)     |----> p1 (at 0xBBBB)     |-----> v (at 0x2222)
value: 0xBBBB   -----      value: 0x2222  -----       value: "Batman"
```
```golang
v := "Batman"
p1 := &v
pp1 := &p1
print("val of p1: %v, addr of p1: %v\n", p1, &p1)     // 0x2222 0xBBBB
print("val of pp1: %v, addr of pp1: %v\n", pp1, &pp1) // 0xBBBB 0x1111

// Dereferencing the pointers
fmt.Printf("*p1 is %v\n", *p1)         // "Batman"
fmt.Printf("*pp1 is %v\n", *pp1)       // 0x2222

// Double dereferencing
fmt.Printf("**pp1 is %v\n", **pp1)     // "Batman"
```

**Comparing pointers**
```golang
v := "Batman"
p1 := &v
x := "Batman"
p2 := &x
p3 := &x

result := p1 == p2  // false, p1 and p2 are pointing to a diff var
result := p3 == p2  // true, p2 and p3 are pointing to the same var
```

**Passing pointers to function**
```golang
// char int, float, bool, string, struct are value types
// Need to pass by pointer to modify in a function
func change(p *int) {
	*p = 888
}

func main() {
	x := 10
	p := &x
	change(p)
	fmt.Printf("After change: %v\n", *p)    // 888
}
```

**Passing pointer->struct to function (Shortcut on struct)**
```golang
type product struct {
	name string
	amt  int
}

func changeProduct(p *product) {
	(*p).name = "blah"
	p.amt = 888        // Shortcut. p.amt is the same as (*p).amt
}

func main() {
	prod := product{
		name: "Game",
		amt:  25,
	}

	changeProduct(&prod)
	fmt.Println(prod)
}
```


**Passing Slice and Map to function**
```golang
// char, int, float, bool, string, struct are value types
// slice and map are reference type
func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 11
	m["b"] = 12
	m["c"] = 13
}

func main() {
	// Slice
	s := []int{1, 2, 3}
	changeSlice(s)
	fmt.Println(s)              // []int{2,3,4}

	// Map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	changeMap(m)
	fmt.Println(m)              // map[a:11 b:12 c:13]
}

```

# Methods and Interfaces
- Receiver type cannot be a pointer type or interface type
  - No *bool, *int, *float, *string types, etc...
  - Map and Slice are a reference type but that is OK.
  
**Receiver Method**
```golang
type names []string

// Receiver method
func (n names) print() {
	for i, name := range n {
		fmt.Printf("%d %s\n", i, name)
	}
}

func main() {
	friends := names{"Tony", "Steve"}
	friends.print()      // receiver method (option 1)
	names.print(friends) // or receiver function (option 2)
}
```

**Receiver Method with a Point Receiver**
```golang
type car struct {
	brand string
	price int
}

// Recommendation:
// Use pointer receiver if we want to change the data or
// if the data is large
func (c *car) changeCar(newBrand string, newPrice int) {
	// With pointer receiver (car is a struct), GoLang is smart 
	// enough to automatically switch from c.brand to (*c).brand
	c.brand = newBrand
	c.price = newPrice
}

func main() {
	myCar := car{brand: "Toyota", price: 32000}

	// With pointer receiver, GoLang is smart enough
	// to automatically switch from myCar.changeCar() to 
    	// (&myCar).changeCar()
	myCar.changeCar("Tesla", 60000)
	fmt.Println(myCar)
}
```

**Interfaces**
- Unlike Java, you don't have to explicitly have a class implement an interface. If a type implements all the interface methods, this type has  implemented the interface.
```golang
type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func print(s shape) {
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func (c circle) area() float64         { return pi * c.radius * c.radius }
func (c circle) perimeter() float64    { return 2 * pi * c.radius }
func (r rectangle) area() float64      { return r.width * r.height }
func (r rectangle) perimeter() float64 { return 2*r.width + 2*r.height }

func main() {
	c := circle{radius: 5.0}
	r := rectangle{width: 4, height: 3}
	print(c)
	print(r)
}
```

**Type Assertions - brings out the concrete type underlying the interface**
```golang
type shape interface {
	area() float64
	perimeter() float64
}

type circle struct{ radius float64 }
type rectangle struct{ width, height float64 }

func (c circle) area() float64         { return pi * c.radius * c.radius }
func (c circle) perimeter() float64    { return 2 * pi * c.radius }
func (r rectangle) area() float64      { return r.width * r.height }
func (r rectangle) perimeter() float64 { return 2*r.width + 2*r.height }

// volume() exists for circle but not in interface
func (c circle) volume() float64       { return 1.3 * pi * math.Pow(c.radius, 3) }

func main() {
	var c shape = circle{radius: 5.0}
	var r shape = rectangle{width: 4, height: 3}

	r.area()

    // c.volume()       // Even though there is a method volume(), shape
                        // can't access the method because it's not listed
                        // in the interface.
	c.(circle).volume() // To access the method, we need Type Assertion

  // More sophisticated way: check before calling volume()
  circle, ok := c.(circle)
	if ok == true {
		circle.volume()
	}
}
```

**Type Switches - like a regular switch but for types**
```golang
type shape interface {
	area() float64
	perimeter() float64
}

func main() {
  var c shape = circle{radius: 5.0}

  switch value := c.(type) {   // type is a keyword
  case circle:
    fmt.Printf("%#v has circle type\n", value)
  case rectangle:
    fmt.Printf("%#v has rectangle type\n", value)
  }
}
```

**Embedded Interfaces - to get around an interface extends other interfaces**
```golang
type shape interface {
	area()
}

type color interface {
	color()
}

// shape and color are embedded in allFeatures
type allFeatures interface {
	shape
	color
	otherMethod()
}

type circle struct{ radius float64 }

// Implement all required methods for allFeatures
func (c circle) area()  {}
func (c circle) color() {}
func (c circle) otherMethod() {}

func print(a allFeatures) {}

func main() {
	var c = circle{radius: 5.0}
	print(c)     // because we have implemented all allFeatures methods 
	             // for c, c can be passed to print() which expects
	             // an allFeatures object.
}
```

**Empty Interface - an interface that can store any value**
- if an operation takes a concrete type as an arg, it can't to take an empty interface (even though an empty interface can be any type)
- To use an operation on a empty interface value, we need to use type assertion to convert the empty interface to a concrete type
- Use empty interface only when it's necessary
```golang
var empty interface{}

empty = 5
empty = "Go"
empty = []int{4, 5, 6}

result := len(empty)         // Error: len takes a concrete type
result := len(empty.([]int)) // After type assertion, it works

fmt.Println(empty)           // OK, Println takes interface{}
```

# Concurrency - GoRoutine
- GoRoutines don't use native threads in the kernel. They are managed by the Go Runtime
- Cheaper than threads

**Check GoRoutine info**
```golang
var p = fmt.Println
p("main execution started")
p("No. of CPUs:", runtime.NumCPU())                // 8
p("No. of Gorountines:", runtime.NumGoroutine())   // 1, main() count as 1

p("OS:", runtime.GOOS)      // linux
p("OS:", runtime.GOARCH)    // amd64

// Max num of CPUs that can run simultaneously
p("GOMAXPROCS:", runtime.GOMAXPROCS(0)) 
```

**WaitGroups - GoRoutine synchronization**
- GoRoutines run concurrently with main() and main() doesn't wait for GoRoutines to finish. 
- WaitGroups can be used so that main() waits for GoRoutines to finish. Otherwise, main() finishes and exits the program before one or more GoRoutines to be able to finish execution.
- wg.Add()  - Add a GoRoutine in the group
- wg.Wait() - Wait for waitgroup to finish
- wg.Done() - Let waitgroup know this routine is done
```golang
func f1(wg *sync.WaitGroup) { // wg is passed as a pointer
	fmt.Println("f1(goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		// sleep for a second to simulate an expensive task.
		time.Sleep(time.Second)

	}
	fmt.Println("f1 execution finished")
	wg.Done() // Tell GoLang that f1() is done
}

func f2() {
	fmt.Println("f2 execution started")
	time.Sleep(time.Second)

	for i := 5; i < 8; i++ {
		fmt.Println("f2(), i=", i)

	}
	fmt.Println("f2 execution finished")

}

var p = fmt.Println

func main() {
	var wg sync.WaitGroup
	// Need Add() so Golang knows how many go routines
	// to wait for Done()
	wg.Add(1)
	go f1(&wg)

	p("No. of Gorountines:", runtime.NumGoroutine())

	f2()
	wg.Wait() // wait for waitgroup to finish
	p("main execution stopped")

}
```

**Race Condition**
- 2 or more goRoutines which read/write the same resource could create a race condition if they are not synchronize correctly.
- Command line flag to detect race condition: **-race** (go run -race main.go)

**Mutex - Concurrency Synchronization**
- var m sync.Mutex
- m.Lock()
- m.Unlock()
```golang
func main() {

	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var shared int = 0 // Declaring a shared value

	var m sync.Mutex   // 1. Declaring a mutex.

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()   // 2. Lock the access to shared
			shared++
			m.Unlock() // 3. Unlock shared after it's incremented
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()  // Using defer is another option
			shared--
			wg.Done()
		}()
	}

	wg.Wait()

	// Printing the final value of shared
	// The final final of shared will be always 0
	fmt.Println(shared)
}
```

**Channel**
- Channel is like a pointer
- Channel is used to both receive and send information
- Receving channel blocks when it's waiting

***Defining and Initializing a channel***
```golang
var ch chan int         // defining a new channel ch (nil)
ch = make(chan int)     // initialize ch, ch is a pointer type

ch := make(chan int)     // chan that can do both sending & rece 
ch := make(<- chan int)  // receiving only chan (can't send to ch)
ch := make(chan <- int)  // sending only chan (can't rece from ch)
ch := make(chan int 3)   // buffered channel with 3 avail slots

c <- 10     // SEND to channel
num := <- c // RECEIVE from channel

close(c)    // close channel c
// NOTE: send to a closed channel will cause a panic
// NOTE: receive from a closed channel will receive a 0 value
```

***Simple example***
```golang
func wait(n int, ch chan int) {
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Printf("Waited %d seconds\n", n)
	ch <- n
}

func main() {
	ch := make(chan int)
	defer close(ch)

	for i := 5; i > 0; i-- {
		go wait(i, ch)
		n := <-ch // waits and blocks
		fmt.Printf("Main notified. %d seconds wait finished\n", n)
	}
}

OUTPUT:
Waited 5 seconds
Main notified. 5 seconds wait finished
Waited 4 seconds
Main notified. 4 seconds wait finished
Waited 3 seconds
Main notified. 3 seconds wait finished
Waited 2 seconds
Main notified. 2 seconds wait finished
Waited 1 seconds
Main notified. 1 seconds wait finished
```

**Buffered vs Unbuffered Channels**
- Unbuffered Channel 
  - No Capacity (1 in, 1 out, synchronized)
  - Send to channel: can't send more than 1 item
  - Rece from channel: blocks when chan is empty or waiting to be filled
- Buffered Channel
  - Has Capacity defined
  - Send to channel: can send until reaches capacity
  - Rece from channel: blocks when chan is empty or waiting to be filled

**Select Statement - let you wait on multiple channel operations**
```golang
func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {     // blocks until it receive from the chan
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

OUTPUT:
received one
received two
```

## Files
**Methods**
```golang
// Create a new file
var newFile *os.File
newFile, err := os.Create("a.txt")

// Trucate a file
err = os.Truncate("a.txt", 0)

// Open a file
file, err := os.Open("a.txt")
file, err := os.OpenFile("a.txt", os.O_APPEND, 0644)  // option 2
file, err := os.OpenFile("a.txt", os.O_CREATE | os.O_APPEND, 0644) 

// Write to file
bytes := []byte("Go Lang")
bytesWrote, err := file.Write(bytes)

// Write to file using ioutil (WriteFile auto close the file)
// WriteFile also auto truncates and creates the file if not exists
err := ioutil.WriteFile("b.txt", bytes, 0644)

// Write to file using Buffered Writer (bufio pkg)
bWriter := bufio.NewWriter(file)
bytesWrote, err := bWriter.Write(bytes)
bytesWrote, err := bWriter.WriteString("Go Lang")
bWriter.Available()  // buffer size avail in bytes (default 4096)
bWriter.Buffered()   // num of bytes not flushed to disk yet
bWriter.Flush()      // flush data to disk (or else not written to disk)
bWriter.Reset(bWriter) // reset the buffer before flushing to disk

// Read content from file (specify by size to read)
bytes := make([]byte, 2)  // alloc 2 bytes
numBytes, err := io.ReadFull(file, bytes)  // only read 2 bytes

// Read all content from file
data, err := io.ReadAll(file)  // data is string type, not []byte

// Read file (auto close file)
bytes, err = ioutil.ReadFile("b.txt") // bytes is []byte type

// Close file
defer file.Close()

// Get file info
var fileInfo os.FileInfo
fileInfo, err = os.Stat("a.txt")
fileInfo.Name()  // file name from the fileInfo
fileInfo.Size()  // file size from the fileInfo
fileInfo.IsDir() // is the file a directory?
fileInfo.Mode()  // file permissions
os.IsNotExist(err) // Check if file exists

// Rename a file
err = os.Rename("oldPath.txt", "newPath.txt")

// Remove a file
err = os.Remove("a.txt")
```

**Scanner - Read file line by line**
```golang
// Read file line-by-line
scanner := bufio.NewScanner(file)
ok := scanner.Scan()
scanner.Text() // first line read from file

// To read entire file line-by-line
for scanner.Scan() {
	line := scanner.Text()   // line is string 
	bytes := scanner.Bytes() // or: bytes is []byte
	print(line)
}
// check scanner error
if err := scanner.Err(); err != nil {
	log.Fatal
}

// To read entire file word-by-word
scanner.Split(bufio.ScanWords)
for scanner.Scan() {
	...
}

// To read entire file letter-by-letter
scanner.Split(bufio.ScanRunes)
for scanner.Scan() {
	...
}
```

**Scanner - Read from user input**
```golang
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()

text := scanner.Text()    // text is string
bytes := scanner.Bytes()  // bytes is []byte

// Continuously read from input, until "exit" is entered
for scanner.Scan() {
	text = scanner.Text()
	if text == "exit" {
		break
	}
}
```


## Init() function
- init() runs before main()
- cannot call init() explicitly
- can have multiple init() functions
	- execution of init() according to the order as written in the file
<br/><br/>

## Command Line Arguments (os.Args)
- os.Args is type []strings
- Args[0] is the executable
- Args[1..N] are the arguments
<br/><br/>

## Go Module and Packages
Default Go workspace is under ~/go \
To start with a project, create a directory under src.

```bash
# To create a module
mkdir master_go
cd master_go
go mod init

# To create a module outside of the GOPATH
mkdir master_go
cd master_go
go mod init master_go  # need to give it a name

# To create a module for github
mkdir master_go
cd master_go
go mod init github.com/bwoo/master_go
```

Or to create a project with a sub-module:
```bash
# To create a sub-module under a project
mkdir -p project/mymodule1
cd project/mymodule1
go mod init

cd ../..

mkdir -p project/mymodule2
cd project/mymodule2
go mod init

# Directory structure:
└── project
    ├── mymodule1
    │   ├── go.mod
    │   ├── main.go
    │   └── mypackage
    │       └── foo.go
    └── mymodule2
        ├── go.mod
        └── mypackage2
            └── bar.go      
```

**Use a package within the SAME module**
- just import module1 and package. \
e.g. projects/mymodule1/main.go:

```golang
package main

import "project/mymodule1/mypackage"

func main() {
	mypackage.PrintPkg1()
```

**Use a package within from a DIFFERENT local module**
- Step 1: manually add in mymodule1/go.mod:
```golang
module project/mymodule1

go 1.18

require "project/mymodule2" v0.0.0
replace "project/mymodule2" => "../mymodule2"
```
- Step 2: import module2 and package:
```golang
package main

import (
	"project/mymodule1/mypackage"
	"project/mymodule2/mypackage2"
)

func main() {
	mypackage.PrintPkg1()
	mypackage2.PrintPkg2()
}
```

**Use a package from GITHUB or other VCS**
- E.g. https://github.com/ddadumitrescu/hellomod
- hellomod is the name of the package (also the module name)
- we import the package in the code
- run: go get github.com/ddadumitrescu/hellomod
- hellomod is saved to $GOPATH/pkg/mod/github.com/...
```golang
package main

import (
	"github.com/ddadumitrescu/hellomod"
)

func main() {
	hellomod.SayHello()
	hellomod.Salut()
}
```





