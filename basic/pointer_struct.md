## pointer, struct

### pointer
- 포인터는 값의 메모리 주소를 가지고 있습니다.
- `&` 연산자로 주소를 반환할 수 있고, 포인터 변수의 값입니다.
    - 포인터 주소에 값을 초기화할때 사용
- `*` 연산자는 포인터가 가리키는 주소의 값을 나타냅니다. (역참조, 간접참조)
    - 포인터로 값을 확인하거나 수정할 때 사용
- zero value : nil
- C언어와 다르게 Go는 포인터 산술을 지원하지 않습니다.
    - pointer 산술이란? 포인터에 +- 와 포인터끼리 빼기연산 등
		
```golang
package main

import "fmt"

func main() {
	i, j := 42, 2701
	//s := "hi"
	
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	
	// p = &s // 다른 타입을 가리키면 에러
	fmt.Printf("value: %d, Type: %T", p, p) // value: 824633794640, Type: *int
}

```

### struct
TODO - 구조체 상세설명
	https://goldenrabbit.co.kr/2021/10/07/%E1%84%80%E1%85%AE%E1%84%8C%E1%85%A9%E1%84%8E%E1%85%A6%E1%84%8B%E1%85%A6-%E1%84%89%E1%85%A2%E1%86%BC%E1%84%89%E1%85%A5%E1%86%BC%E1%84%8C%E1%85%A1%E1%84%85%E1%85%B3%E1%86%AF-%E1%84%83%E1%85%AE%E1%86%AF/

- Struct (구조체)
	- 구조체는 필드의 집합체입니다. (필드의 컨테이너)
	- 코드의 결합도, 의존성을 낮게 만들고 응집도를 높게 만드는 역할을 합니다.
	- 클래스처럼 메서드를 가지지 않습니다. (메서드는 인터페이스로 선언후 따로 구현)
	- Go는 전통적인 OOP 개념인 클래스, 상속, 객체 개념이 없음
	- struct는 기본적으로 mutable 개체로서 필드값이 변화할 경우 (별도로 새 개체를 만들지 않고) 해당 개체 메모리에서 직접 변경
	- 다른 함수의 파라미터로 넘긴다면, Pass by Value에 따라 복사해서 전달됨.
	- 다른함수에서 객체를 수정해야 하는 경우 Pass by reference로 전달해야 함.(포인터로 전달)

- Struct Fields	
	- dot으로 필드 접근 가능
	- 구조체 pointer도 *연산 없이 값에 접근이 가능하다. (*p).X == p.X
	
```golang
// 선언 방식
type Vertex struct {
	X int
	Y int
}

// 이렇게 선언하면 외부에서 참조 불가능
type vertex sturct {
	X, Y int
}
```

- Struct 구조체 인스턴스 생성하는 방법
	1. `구조체이름{}` 으로 생성
		- 구조체 선언시 필드 순서대로 값 입력
		- 필드 이름에 매핑해서 입력
	2. `new()` 메서드를 사용하는 방법
		- `new()`는 객체의 포인터를 리턴합니다.
	3. 하나씩 모두 설정해주는 방법
		
```golang
type person Struct{
	name string
	age int
}
var p1 person
p1 = person{"bob", 20}				// 값을 순차적으로 넣어 생성
p2 = person{name: "Ssss", age: 44}  // 필드명을 지정해서 생성

p := new(person)
p.name = "jake"
(*p).age = 14
```

- 구조체 생성자 함수
	- 생성자 함수 없음. 따로 정의해야 함
	- 구조체를 생성해서 필드를 초기화하는 메서드를 작성하고 생성한 객체의 포인터를 리턴하는 함수를 작성
	- 구조체 안에 map같은게 있다면, 반드시 초기화를 해줘야 값을 넣어줄 수 있기 때문에 생성자 함수를 만들어서 초기화를 해준다.
	
```golang
type dict struct {
	data map[int]string
}

// 맵은 초기화해주지 않으면 바로 사용할 수 없기 때문에
// 생성자 함수에서 초기화해서 구조체의 포인터를 리턴해줌
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func main() {
	dict := newDict()
	dict.data[1] = "A"
	(*dict).data[2] = "b"
	fmt.Println(dict.data[1]) // A
	fmt.Println(dict.data[2]) // b
}
``` 
	
- Pointers to Structs
	- `*` 연산(역참조 연산) 안해줘도 읽어오고, 수정할 수 있음

```golang
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	fmt.Println((*p).X)
	fmt.Println(p.X)
}
```

- Struct Literals
	- 필드 값을 나열하여 새로 할당된 구조체 값을 나타냅니다.
	- 필드가 선언된 순서대로 값을 입력하면 새로운 구조체 객체가 생긴다는 말인 듯
	- 구조체 값을 생성하는 방법.
```golang
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
	fmt.Println(&v1, p)
	fmt.Println(&v1 == p) // false 다른 객체
}

```

https://deku.posstree.com/ko/golang/struct/

- 중첩된 구조체 Nested Struct
	- 구조체가 다른 구조체 타입의 변수를 가지고 있는 것

- Embedded Field
	- 구조체가 다른 구조체 타입을 변수선언 없이 가지고 있는 것
	- 임베디드 필드(다른 구조체 이름)를 변수처럼 사용할 수 있음

```golang
package main

import "fmt"

type ClassInfo struct {
	Name string
	Code int
}

// 중첩 구조체 (nested struct)
type CourseInfo struct {
	Class ClassInfo // Embedded Field
	Open  bool
}

func main() {
	var s = CourseInfo{
		Class: ClassInfo{Name: "math", Code: 101},
		Open:  true,
	}

	fmt.Printf("Type: %T, value: %v\n", s, s) // Type: main.CourseInfo, value: {{math 101} true}

	var s_embed = CourseInfoEmbed{
		ClassInfo{Name: "math", Code: 101},
		true,
	}

	fmt.Println(s_embed.ClassInfo)                        // {math 101}
	fmt.Println(s_embed.ClassInfo.Code)                   // 101
	fmt.Println(s_embed.ClassInfo.Name)                   // math
	fmt.Println(s_embed.Open)                             // true
	fmt.Printf("Type: %T, value: %v\n", s_embed, s_embed) // Type: main.CourseInfoEmbed, value: {{math 101} true}
}

// embeded field
type CourseInfoEmbed struct {
	ClassInfo // 다른 구조체를 상속하듯 가지고있음.
	Open bool
}

```