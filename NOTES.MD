 
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

go run main.go \
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
%d -> decimal\
%f -> float\
%.3f -> float (3 decimal points)\
%s -> string\
%q -> double-quoted string\
%v -> value (any)\
%#v -> a Go-syntax representation of the value\
%T -> value Type\
%t -> bool (true or false)\
%p -> pointer (address in base 16, with leading 0x)\
%c -> char (rune) represented by the corresponding Unicode code point
```

# Convert String to Numbers and vice versa
```golang
String(99)  // c  (ascii)\
fmt.Sprintf("%f", 44.2) // "44.2000"

// strconv package from string to Numbers 
strconv.ParseFloat("3.123", 64) // 3.123 (float) 

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
if (n == m) {  // true, same num of elements, order and values

}
```

# Slice
**Reference type**
Reference Type
Slice header:
  - beginning addr (first element) of the backing array
  - len of the slice
  - capacity, len calc. from starting addr of the point 
    to end of the backing array
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

E.g. Appending
s1 := []string{"a","b","c","d","e"}
s2 := []string{}
s2 = append(s2, s1[0:2]...)         // s2=["a","b"]

E.g. Cutting
s1 := []int{1,2,3,4,5}
s2 := s1[0,3]  // s2=[0,1,2]
l := len(s2)   // l=3 (the slice has 3 elements)
c := cap(s2)   // c=5 (backing array has 5 elements)
//*Note: capacity measured from starting element*

E.g. Capacity from starting element
s1 := []int{1, 2, 3, 4, 5}  // addr: 0x00 len(s1)=5 cap(s1)=5
n1 := s1[1:3]     // addr: 0x08 (1*8 bytes) n1=[2,3] len(s1)=2 cap(s1)=4
n2 := s1[1:4]     // addr: 0x08 (1*8 bytes) n1=[2,3,4] len(s1)=3 cap(s1)=4
n3 := s1[2:5]     // addr: 0x10 (2*8 bytes) n1=[3,4,5] len(s1)=3 cap(s1)=3
```

**Copy slice:**
```golang
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
	result := make([]int, 0, maxValue)
	result = append(result, 88)
	result2 := result
	for i := 0; i < maxValue; i++ {
		if i%2 == 0 {
			fmt.Printf("appending '%d': %s   %s\n", i, getSliceHeader(&result), getSliceHeader(&result2))
      // Appending to result, change the len.
      // Note that len in result2 is NOT changing, even though 
      // result and result2 are pointing to the same
      // backing array adress.
			result = append(result, i)
			fmt.Printf("appended '%d': %s   %s\n", i,	getSliceHeader(&result), getSliceHeader(&result2))
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

/*
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
*/

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
  - Map Header --> data structure
  - When a map is assigned to another map (m1 := m2), the header of m2 is copied to m1.
  - Both map headers point to the same data structure
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
m2 := make(map[string]int)	           // create another map
for k,v := range m1 {                  // clone by copying each key/value
  m2[k] = v
}
```

# Structs