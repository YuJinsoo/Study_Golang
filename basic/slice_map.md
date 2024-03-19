## slice, map

### Slices 슬라이스
	- 슬라이스는 배열의 요소들을 동적인 크기로, 유연하게 볼 수 있습니다.
	- 슬라이스 생성 방법
		- var a []int
		- a := []int{1,2,4,5}
		- var a = make([]int, 4, 5) // 길이 4 용량 5 인 슬라이스 (용량은 생략 가능)
	- 배열을 슬라이싱 연산 한 것은 슬라이스!
		- 슬라이싱 연산 array[start_index : end_index]
	
```golang
primes := [6]int{2, 3, 5, 7, 11, 13} // 이건 배열

var s []int = primes[1:4] // 배열 선언에서 크기를 지정하지 않으면 슬라이스가 됨
fmt.Println(s)

primes_2 := [6]int{2, 3, 5, 7, 11, 13}
var s2 []int = primes_2[1:4] 	// 슬라이스 연산해서 슬라이스 반환
fmt.Println(s2) 			//[3 5 7]
s2[1] = 99
fmt.Println(s2)           	//[3 99 7]
fmt.Println(primes_2)     	//[2 3 99 7 11 13] // 원본 배열도 변경됨
fmt.Printf("Type: %T", s) 	// Type: []int
```

- Slices are like references to arrays (배열을 참조하는 슬라이스)
	- 슬라이스는 어떤 데이터도 저장할 수 없습니다. 
	- 단지 기본 배열의 한 영역을 참조하는 형태
	- 슬라이스의 요소를 변경하면 기본 배열의 해당 요소가 수정됩니다.
``` golang
names := [4]string{
	"John",
	"Paul",
	"George",
	"Ringo",
}
fmt.Println(names)

a := names[0:2]
b := names[1:3]
fmt.Println(a, b)

b[0] = "XXX"
fmt.Println(a, b)
fmt.Println(names)
// Type: []int[John Paul George Ringo]
// [John Paul] [Paul George]
// [John XXX] [XXX George]
// [John XXX George Ringo]
```
	
- Slice literals
	- 슬라이스를 리터럴을 주어 생성하는 방법
	
	- 슬라이스 생성에는 세가지 방식이 있음
		- 배열 혹은 슬라이스에서 슬라이싱 연산으로 생성
		- 리터럴을 주어 생성
		- make함수로 생성
```golang
package main
import "fmt"
 
func main() {
    s := []int{0, 1, 2, 3, 4, 5} // 리터럴 값으로 slice 생성.
    s = s[2:5]  
    fmt.Println(s) //2,3,4 출력
}
```

- Slice 요소 추가
	- 배열과 다르게 자유롭게 요소를 추가할 수 있다.
	- 추가 append() 함수

```golang
// 추가예제
temp_s := []int{0, 1}

temp_s2 := append(temp_s, 2)
fmt.Println("append from: ", temp_s) // [0 1] append해도 원본은 변하지 않음
fmt.Println("append to: ", temp_s2)  // [0 1 2]
// 같은 배열에 추가하고싶으면 append()의 리턴을 같은 슬라이스에 할당하면 됨

temp_s = append(temp_s, 3, 4, 5)
fmt.Println("여러개 어펜드: ", temp_s) // [0 1 3 4 5]

```

- slice 용량이 늘어나는 원리
	- TODO - Underlying Array 란??????????????????????????
	- 슬라이스 s 의 길이와 용량은 len(s) 와 cap(s) 식으로 얻을 수 있습니다.
	- 용량이 다 찼을 때 append() 따위의 함수로 추가된다면 용량이 늘어납니다.
``` golang
// 슬라이스.... 용량증가
	slice_a := make([]int, 2, 3)      // 길이가 2, 용량이 3인 슬라이스 생성. 길이만큼 0으로 초기화됨
	fmt.Println("slice_a: ", slice_a) // [0, 0]

	for i := 1; i <= 15; i++ {
		slice_a = append(slice_a, i)
		fmt.Println("slice len: ", len(slice_a), " / slice cap: ", cap(slice_a))
	}
	// cap이 가득 찰 때마다 2배의 용량으로 늘리는 것을 확인할 수 있다.
	// Underlying array를 생성하고 기존 배열 값들을 새 배열에 복사한 후 다시 슬라이스를 할당한다고 함.
	// Underlying array란??
```

- Appending to a slice (슬라이스에 요소 추가하기)
	- append 함수
``` golang
// 슬라이스 요소 추가하기
var s_add []int
printSlice(s_add) // len=0 cap=0 []

s_add = append(s_add, 0)
printSlice(s_add) //len=1 cap=1 [0]

s_add = append(s_add, 1)
printSlice(s_add) // len=2 cap=2 [0 1]

s_add = append(s_add, 2, 3, 4)
printSlice(s_add) // len=5 cap=6 [0 1 2 3 4]
// cap이 부족하면 더 늘어나는 것을 확인할 수 있다.
```

- Slice 요소 삭제
	- 내장 삭제 함수가 존재하지 않음
	
	- remove()라는 함수를 만들어서 사용
		- 특정 인덱스를 제외하는 서브 슬라이스 2개를 합친 슬라이스를 반환하는 함수.
		- 순서가 유지됨
		
	- 다른방법
		- 원소의 맨 뒤의 값을 삭제할 인덱스 위치에 넣은 뒤 길이를 1개 줄임
		- 위의 방식보다 빠르지만 순서가 보장되지 않음
	
	- copy 함수를 사용해서 새로운 슬라이스를 반환하는 방식으로 삭제할 수 있다.
		- `copy(slice[index:], slice[index+1:])`
	
	- 아래 예제에서 함수에서 slice를 복사해서 사용하는 이유
		- 복사하지 않으면 원본 슬라이스가 가리키는 배열을 조작하기 때문
		- 슬라이스를 조작할 때, 슬라이스의 내부 배열에 대한 수정이 일어날 수 있으며, 이는 원본 슬라이스에도 영향을 줌
		- 이 경우 append()가 결과 슬라이스를 구성하기 위해 원본 슬라이스의 내부 배열을 재사용할 수 있음
```golang
// 순서보장 삭제
func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

// 빠르지만 순서 미보장
func remove_un(slice []int, index int) []int {
	var copy_slice = make([]int, len(slice), cap(slice))
	copy(copy_slice, slice)
	copy_slice[index] = copy_slice[len(copy_slice)-1]
	return copy_slice[:len(copy_slice)-1]
}

// slice 요소삭제
	var slice_del = []int{1, 2, 3, 4, 5}
	fmt.Println("slice_del: ", slice_del) // [1 2 3 4 5]

	slice_del_after := remove(slice_del, 2)
	fmt.Println("slice_del_after: ", slice_del_after) // [1 2 4 5]

	fmt.Println("slice_del: ", slice_del) // [1 2 3 4 5] 
	slice_del_after2 := remove_un(slice_del, 2)
	fmt.Println("slice_del_after2: ", slice_del_after2) // [[1 2 5 4]
```

- Slice 병합 (두 개의 슬라이스 합치기)
	- 원소 추가와 같이 append()를 사용한다.

```golang
// slice 병합
// 새로운 슬라이스를 만들어서 병합하는 방법
sliceA := []int{1, 2, 3}
sliceB := []int{4, 5, 6}

sliceC := append(sliceA, sliceB...)
fmt.Println("sliceC: ", sliceC) // [1 2 3 4 5 6]
fmt.Println("sliceA: ", sliceA) // [1 2 3]
fmt.Println("sliceB: ", sliceB) // [4 5 6]
```

- Slice 복사
	- 복사의 경우 copy() 함수를 사용합니다.
	- 요소 추가의 가능성이 있으므로 cap을 2배로 주어 생성한 뒤 복사(필수X)
	- copy(복사본slice, 원본slice)
	- Slice는 참조연산이므로 원본에 영향이 가지 않게 하려면 복사해서 사용해야 함
	
```golang
// slice 복사

slice_ori := []int{1, 2, 3, 4, 5}
slice_copy := make([]int, len(slice_ori), cap(slice_ori)*2)
copy(slice_copy, slice_ori)
fmt.Println("slice_ori: ", slice_ori)   // [1 2 3 4 5]
fmt.Println("slice_copy: ", slice_copy) // [1 2 3 4 5]

```

- Slice default
	- 슬라이싱 연산 동작
	- 상한 또는 하한을 생략하면, 슬라이싱할 때 기본 값을 사용할 수 있습니다. 
	- 하한의 경우 기본 값은 0이고, 상한의 경우 슬라이스의 길이입니다.
	- 슬라이싱 연산에 음수나 용량을 초과하는 인덱스를 입력하면 에러를 뱉는다.
	
	```golang
	// 슬라이스 연산 기본 값
	s_1 := []int{2, 3, 5, 7, 11, 13}

	s_1 = s_1[1:4]
	fmt.Println(s_1) // [3 5 7]

	s_1 = s_1[:2]    // 0부터~ 2까지
	fmt.Println(s_1) // [3 5]

	s_1 = s_1[1:]    // 1부터~ 끝까지
	fmt.Println(s_1) // [5]
	```
	
	
- Nil slices (nil 슬라이스)
	- 슬라이스의 zero value는 nil 입니다.
	- nil 슬라이스의 길이와 용량은 0이며, 기본 배열을 가지고 있지 않습니다.
	- 내부적으로 사용하는 배열이 없는 경우에만 nil로 판단하는 것으로 추정
		- 즉 길이가 0 이더라도 용량이 0 이 아니면 nil이 아님
		- 내부적으로 사용하는(가리키는) 배열이 없는 경우에만 nil로 판단하는 것으로 추정

``` golang
package main

import "fmt"

func main() {
	// 선언만 한 경우 cap이 0이기 때문에 nil
	// nil slice : nil 값임.
	var s []int
	fmt.Println(s, len(s), cap(s))	// [] 0 0
	if s == nil {
		fmt.Println("nil!")
	}
	
	// 길이가 0 이더라도 용량이 있으면 배열을 가리키고 있으므로 nil이 아님
	var s2 = make([]int, 0, 10)
	fmt.Println(s2, len(s2), cap(s2))	// [] 0 10
	if s2 == nil {
		fmt.Println("nil!")
	} else {
		fmt.Println("not nil!")
	}
}
```
	
- Creating a slice with make (make 함수로 슬라이스 만들기)
	- make(타입, 길이, 용량)
		- 용량을 생략하면 길이 = 용량으로 생성됨
	
```
// make함수로 슬라이스 생성하기
	// int의 zero value인 0으로 초기화됨
	a := make([]int, 5)
	printSlice2("a", a) // a len=5 cap=5 [0 0 0 0 0]

	// 길이를0, capa를 ㅈ정해주면 빈 슬라이스가 생성됨
	b := make([]int, 0, 5)
	printSlice2("b", b) // b len=0 cap=5 []

	// 슬라이싱하면 zerovalue 할당됨
	c := b[:2]
	printSlice2("c", c) // c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice2("d", d) // d len=3 cap=3 [0 0 0]

	// e := c[2:6] // slice bounds out of range [:6] with capacity 5
	// printSlice2("e", e)
```

- Slices of slices (슬라이스의 슬라이스)
	- 슬라이스는 다른 슬라이스를 포함하여 모든 타입을 담을 수 있습니다.

```
// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

// 슬라이싱 연산 시 범위를 넘어가면 에러가 발생함
// 음수 인덱스 안됨
// panic: runtime error: slice bounds out of range [6:5]
test_slice := []int{1, 2, 3, 4, 5}
fmt.Println("test_slice: ", test_slice[:len(test_slice)])
// fmt.Println("test_slice: ", test_slice[len(test_slice)+1:]) // panic
// fmt.Println("test_slice: ", test_slice[:-1]) // 음수 안됨
```
	
### 슬라이스 내부 동작
	- 내부적으로 사용하는 배열 부분 영역에 대한 메타 정보를 가지고 있음
	- 배열에 대한 포인터, 배열 길이, 배열 용량 에 대한 필드를 가지고 있음
		- 이러한 특징 때문에 
	- 공식 블로그참고: https://go.dev/blog/slices-intro


### Range
	- range 는 요소들을 순회하는 것을 돕는 역할
	- 배열, 슬라이스, 맵, 문자열, 채널 이 올 수 있고, 종류마다 반환하는 값이 다름
		- 배열, 슬라이스, 문자열:	인덱스, 값
		- 맵:					키,     값
		- 채널:					채널값
		- 이때 반환하는 값은 인덱스(혹은 키) 값의 복사본
		- 채널의 경우 close()가 있어야 끝났다는 것을 인지하여 에러가 안 남
	- python의 enumerate 같음
	- range 왼쪽의 = 또는 := 기준으로 왼쪽 영역은 이터레이션 변수(iteration variables)라고 함
	- range expression(range 우측)이 채널인 경우는 최대 1개의 이터레이션 변수가 가능하고, 나머지는 최대 2개까지 가능하다.

```golang
package main

import (
	"fmt"
)
func main() {
	// range  사용 구문
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	
	// range를 사용하면 for문에서 array나 slice의 index와 value 를 함께 순회할 수 있습니다.
	for i, v := range "hello world" {
		fmt.Printf("%d = %v\n", i, v)
		// 0 = 104 으로 나옴. 아스키코드로 출력되는 것 같다.
	}

	// 유니코드인 경우 인덱스가 1씩이 아니라 크기만큼 증가..
	for i, r := range "→👍👎🌮🗂HelloWorld!안녕세상아!😊🚀🔥📝." {
		fmt.Print(i, " ")      // i 자리수가 여러개 뛰어넘음 왜? >>
		fmt.Println(string(r)) // ascii값을 string으로 변환
	}
	// 0 →
	// 3 👍
	// 7 👎
	// 11 🌮
	// 15 🗂
	// 19 H
	// 20 e

	// 둘 중 하나만 원하는 경우 index만 받거나, _ 로 처리할 수 있습니다.
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// map의 경우 key, value를 순회합니다.
	var mm = map[int]string{1: "a", 2: "b", 3: "c"}
	for k, v := range mm {
		fmt.Println("key: ", k, " value: ", v)
	}
}
```
	
	

### Maps 맵
	- 키를 값에 매핑합니다.
	- zero value는 nil
	- nil 맵은 키도 없고, 키를 추가할 수도 없습니다.
		- make로 만들어서 초기화 시켜주던가 방법이 필요함.
	
```golang
package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main(){
	// var m = map[key속성]value속성

	m := make(map[string]Vertex) 	// make로 만든 것은 nil이 아님.
	fmt.Println(m)                	// map[]

	if m == nil {
		fmt.Println("nil!")
	} else {
		fmt.Println("not nil!") // not nil!
	}

	var m2 map[string]int // 선언만 한 경우에는 nil이다.
	fmt.Println(m2)       // map[]
	if m2 == nil {
		fmt.Println("nil!") // nil!
	} else {
		fmt.Println("not nil!")
	}

	// m2["Tset"] = 1  // panic: assignment to entry in nil map
	// fmt.Println(m2) //  참조하는 순간 panic이 발생한다.
}
```

- Maps Literals
	- 구조체 리터럴과 같지만 key가 필요합니다.

```golang
type Vertex_2 struct {
	Lat, Long float64
}

main(){
// map literal
	var m_lit = map[string]Vertex_2{
		"Bell Labs": Vertex_2{
			40.68433, -74.39967,
		},
		"Google": Vertex_2{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m_lit) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}

```
	- 최상위 타입이 타입의 이름인 경우 요소에서 삭제할 수 있습니다.
```golang
	// Vertex라는 구조체 이름을 생략.
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
```
	
- Mutating Maps
	- 맵에 요소 추가하기
		m[key] = value
	- 검색하기
		elem = m[key]
	- 제거하기
		delete(m, key)
	- 존재하는지 확인
		elem, ok = m[key]
		- 만약 key 가 m 안에 있다면, ok 는 true 입니다. 아니라면, ok 는 false 입니다.
		- 만약 key 가 맵 안에 없다면, elem 은 map의 요소 타입의 zero value입니다. (흠...)

```
package main

import "fmt"

func main(){
	m_test1 := make(map[string]string)
	m_test1["one"] = "test"
	m_test1["two"] = "test2"
	m_test1["three"] = "test3"
	m_test1["four"] = "test4"
	
	fmt.Println(m_test1)		// map[four:test4 one:test three:test3 two:test2]
	fmt.Println(m_test1["one"]) // test2
	delete(m_test1, "one")
	delete(m_test1, "one") 		// 여러번 삭제해도 에러는 안남.
	fmt.Println(m_test1)   		// map[four:test4 three:test3 two:test2]

	v, ok := m_test1["one"]
	fmt.Println("The value:", v, "Present?", ok) // The value:  Present? false
}

```

TODO - map에서 값으로 key 찾는 법은?


- Map 연습문제
	- strings라는 패키지...
		Split
		Fields
		
```
package main

import (
	"fmt"
	"strings"
	//"golang.org/x/tour/wc" // playground용
)

func main(){
	ttt := strings.Split("foo bar baz", " ")  // 구분자로 딱딱 나눠서
	tttt := strings.Fields("   foo bar  baz") // 공백 다 없애서 구분해서 줌. 심지어 2칸도 하나로 나눠줌

	fmt.Println(ttt)  // [foo bar baz]
	fmt.Println(tttt) // [foo bar baz]

	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
var ret_map = make(map[string]int)

for _, v := range strings.Fields(s) {
	ret_map[v] += 1
}
return ret_map
}

```