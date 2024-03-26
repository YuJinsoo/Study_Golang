## Method와 Interface

### Methods
- Method
	- Go는 클래스가 없음
	- Go만의 방식으로 객체지향 프로그래밍을 지원
	- 메소드는 리시버 인수가 있는 함수입니다.
		- 리시버 인자: 어떤 구조체를 위한 것인지 표기하는 인자
		- 메소드 내부에서 receiver 인자를를 활용할 수 있습니다.
	
	- 메서드는 인스턴스에 dot(.)으로 호출 가능합니다.
	- 같은 타입의 메소드를 정의할 수 있습니다.
	- 리시버가 포인터/값 으로 다르고, 메서드 매개변수가 다르더라도 같은 이름의 메서드를 선언할 수 없음
	```
	package main

	import (
		"fmt"
		"math"
	)

	type Rect struct {
		width, height int
	}

	// func와 메서드 이름 사이에 receiver를 추가
	// area()에는 r라는 이름의 Rect 유형의 리시버가 있음
	func (r Rect) area() int {
		return r.width * r.height
	}
	
	// 리시버 인자가 포인터/값 으로 다르고, 인자가 다르더라도 같은 이름의 메서드를 선언할 수 없음.
	// method Rect.area already declared
	// func (r *Rect) area(i int) int {
	// 	return r.width * r.height
	// }

	func main() {
		r := Rect{10, 20}
		fmt.Println("r area: ", r.area()) // r area:  200
	}

	```
	
	- 구조체가 아닌 형식에 대해서도 메소드를 선언할 수 있음
	- 메소드와 동일한 패키지에 유형이 정의된 수신자가 있는 메소드만 선언할 수 있습니다. ???? 
		- 유형이 다른 패키지 (int 와 같은 빌트인 유형 포함)에 정의된 리시버로 메소드를 선언 할 수 없습니다.
	
	```golang
	package main

	import (
		"fmt"
		"math"
	)

	type MyFloat float64 // 패키지 내에 유형을 정의

	// 같은 패키지 내에서 method 선언가능
	func (f MyFloat) Abs() float64 {
		if f < 0 {
			return float64(-f)
		}
		return float64(f)
	}

	func main() {
		fmt.Println(math.Sqrt2) // 1.4142135623730951
		f := MyFloat(-math.Sqrt2)
		fmt.Println(f.Abs()) // 1.4142135623730951
	}

	```
	
- 포인터 리시버
	- 메서드의 리시버가 포인터 타입인 경우를 말함
		- 리시버 유형이 일부 유형 T 에 대한 리터럴 구문 *T 를 가짐을 의미
		- T 자체는 *int 와 같은 포인터 타입이 될 수 없음
		
	- 값 리시버를 사용하면 메소드 영역에서 값의 복사본에서 작동합니다. 
		- 다른 함수 인수와 동일합니다
	- main 함수에 선언된 Rect 인스턴스의 값을 변경하기 위해서는 포인터 리시버 메서드가 있어야합니다.
		- 값을 바꾸려면 리시버 인자를 *타입으로 지정해야 함.
		- 직접 값을 바꾸는 경우 말고 메서드나 함수를 통해 바꾸려면 그래야 함

```golang
package main

import (
	"fmt"
	"math"
)


type Rect struct {
	width, height int
}

// value receiver
func (r Rect) plusWidth() {
	r.width++
	fmt.Println("plusWidth value receiver: ", r.width)
}

// pointer receiver
func (r *Rect) plusHeight() {
	r.height++
	fmt.Println("plusWidth pointer receiver: ", r.height)
}


func main() {
	r := Rect{width: 10, height: 20}
	fmt.Println("r area: ", r.area()) // r area:  200

	r.plusWidth()                                                     // value receiver method
	r.plusHeight()                                                    // pointer receiver method
	fmt.Printf("r values width: %d, height: %d\n", r.width, r.height) // r values width: 10, height: 21
}
```

- 포인터와 함수
	- 메서드로 만들지 않고 일반 함수에 포인터 인자로 구조체 등을 전달하여 상태를 변경할 수 있습니다
	
```golang
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

```

- 메소드와 포인터 indirection
	- 포인터 리시버가 있는 메서드를 호출할 때 & 연산자를 쓰지 않아도 &v와 같이 변환해서 호출함
	- 값 리시버가 있는 메서드를 호출할 때 pointer로 호출해도 자동으로 *p로 변환해서 호출함
	
	- 일반 함수에 포인터인자라면 반드시 포인터를 전달해줘야 하지만, 메서드의 리시버는 어떻게 호출해도 알맞은 형태로 변경해줍니다.

```golang
package main

import (
	"fmt"
	"math"
)

type Rect struct {
	width, height int
}

func (r Rect) printRect_value() {
	fmt.Println("printRect_value: ", r.width, r.height)
}
func (r *Rect) printRect_pointer() {
	fmt.Println("printRect_pointer: ", r.width, r.height)
}

func main(){
	r2 := Rect{width: 10, height: 20}
	pr := &r2
	r2.printRect_value()   // 값으로 value receiver 호출
	
	// (&r2).printRect_pointer() 로 변환되어 호출됨
	r2.printRect_pointer() // 값으로 pointer receiver 호출
	
	// (*pr).printRect_value() 로 변환되어 호출됨
	pr.printRect_value()   // 포인터로 value receiver 호출
	
	pr.printRect_pointer() // 포인터로 pointer receiver 호출
	// 위 함수 모두 정상 출력됨
}
```

- 값 또는 포인터 리시버 선택
	- 값 리시버 사용이유
		1. 리시버의 복사본을 사용하기 때문에 원본이 변하지 않아야 하는 경우에 사용
			- 복사하기 때문에 메모리가 추가적으로 소모된다.
		
	- 포인터 리시버 사용이유
		1. 메서드가 리시버가 가리키는 값을 수정할 수 있기 때문
		2. 각각의 메서드 call에서의 value 복사 문제를 피하기 위함
			- 리시버로 지정한 타입이 큰 경우 효과적
	

	- 일반적으로 특정 유형의 모든 방법에는 값이나 포인터 리시버가 있어야 하지만, 둘 다 혼합되어서는 안됩니다.
		>>> why??
	
	TODO - 객체 리스트를 순회시킬 떄 메서드는 어떻게 동작할지
		https://jamong-icetea.tistory.com/402



### Interface

TODO - 인터페이스 전반
https://deku.posstree.com/ko/golang/interface/

TODO - 콘크리트 값이란?

- Interface란
	- 메소드의 시그니처 집합
	- type이 구현해야 하는 메서드 원형들을 정의
	- 어떤 인터페이스를 구현한건지 어떻게 알아보는지 궁금 ?????

```golang
// 인터페이스 정의
type Abser interface {
	Abs() float64
}
```
	
- 인터페이스 사용
	- 일반적으로 함수가 파라미터로 인터페이스를 받는 경우
		- 어떤 타입이든 해당 인터페이스를 구현하면 인자로 받을 수 있음

```golang
package main

import "fmt"

// 인터페이스 정의
type Shape interface {
	area() float64
	perimeter() float64
}

// 구조체 정의
type Rectangle struct {
	width, height float64
}

type Cricle struct {
	radius float64
}

// Rectangle 타입에 Shape 인터페이스를 구현
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 Shape 인터페이스를 구현
func (c Cricle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Cricle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// 인터페이스를 구현한 구조체를 인자로 받는 함수
func showArea(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Println(s.area())
	}
}

func main() {
	r := Rectangle{width: 10, height: 4}
	c := Cricle{radius: 5}

	showArea(r, c)
	// 40
	// 78.53981633974483
}

```
	- 인터페이스 사용 시 타입에서 구현한 메서드가 포인터 시리버인지 값 리시버인지 구분합니다.
		- 타입에서 인터페이스 함수를 값 리시버로 정의했으면 인터페이스 변수에 값으로 할당
		- 타입에서 인터페이스 함수를 포인터 리시버로 정의했으면 인터페이스 변수에 포인터로 할당

```golang
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // MyFloat 로 Abser 수현
	fmt.Println(a.Abs())
	
	a = &v // *Vertex 로 Abser 수현
	fmt.Println(a.Abs())
	
	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v 에러발생 : (method Abs has pointer receiver)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```

- 인터페이스 암시적 구현
	- 인터페이스 구현 시 타 언어처럼 `implements` 같은 예약어로 명시적으로 할 필요가 없음
	
- TODO - 어떤 인터페이스를 구현했는지 확인하는 방법이 있나?
- TODO - https://hoonyland.medium.com/%EB%B2%88%EC%97%AD-interfaces-in-go-d5ebece9a7ea


- 인터페이스 값
	- 콘크리트 타입의 튜플이라고 생각할 수 있습니다.
		- (value, type)
		- 아무것도 할당하지 않은 인터페이스는 (<nil>, <nil>) 
	
	- 인터페이스를 구현한 메서드가 포인터 리시버이면 interface 객체로 받을 때도 &로 할당
	- 인터페이스를 구현한 메서드가 값 리시버이면 interface 객체로 받을 때도 값으로 할당

```golang
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I
	describe(i) // (<nil>, <nil>) << 인터페이스 기본값
	// i.M() invalid memory address or nil pointer dereference
	
	i = &T{"Hello"} // 포인터 리시버로 구현해서 주소로 할당
	describe(i) 	// (&{Hello}, *main.T)
	i.M() 			// Hello

	i = F(math.Pi) // 값 리시버로 구현해서 주소로 할당
	describe(i) 	// (3.141592653589793, main.F)
	i.M()			// 3.141592653589793
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

- Nil interface
	- 인터페이스 자체 내부 콘크리트 값이 0인 경우, nil 리시버로 호출됩니다
	- nil 콘크리트 값을 갖는 인터페이스 값 자체가 nil이 아니라는 점에 유의해야 함 (nil, type?)
	- nil 인터페이스로 로 메서드 호출하면 런타임 에러임에 주의
	
	- 일부 언어에서는 이것이 null 포인터 예외를 발생함
	- Go 에서는 nil 리시버로 호출되는 것으로 불리는 매우 좋은 방법을 사용하는 것이 일반적입니다 ???


https://www.hahwul.com/2021/08/09/do-you-know-about-golang-nil/

```golang
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	temp_str string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("this is <nil>")
		return
	}
	fmt.Println(t.temp_str)
}

func main() {
	var i I
	describe(i)	// (<nil>, <nil>)
	
	// 선언만 했을 때에는 nil (둘다 nil일 때는 nil)
	if (i == nil){
		fmt.Println("nil") // nil
	} else {
		fmt.Println("not nil")
	}
	
	var t *T // T를 가리키는 포인터변수.
	i = t
	// 값이 없어서 nil이지만, type은 지정되어있음
	describe(i) // (<nil>, *main.T)
	i.M()		// this is <nil>
	
	// 값이 없어서 nil이지만, type은 지정되어있음
	// 그래서 nil이 아님!
	if (i == nil){
		fmt.Println("nil")
	} else {
		fmt.Println("not nil") // not nil
	}

	i = &T{"hello"}
	describe(i)	// (&{hello}, *main.T)
	i.M()		// hello
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

- 인터페이스 타입 (빈 인터페이스)
	- 아무 메서드를 지칭하지 않는 인터페이스를 empty interface 라고 합니다.
		- 여러 표준 패키지들을 보면 흔히 확인할 수 있음 ( interface{} ) 
	
	- 빈 인터페이스는 어떠한 타입이든 담을 수 있음
		- C#, Java의 object?
	- 알 수 없는 값을 처리하는 데 사용합니다
	
```golang
package main

import "fmt"

func main() {
	var x interface{}
	x = 1
	printInterface(x) // (1, int)
	x = "hello"
	printInterface(x) // (hello, string)
}

func printInterface(v interface{}) {
	fmt.Printf("(%v, %T)\n", v, v)
}
```

- 타입선언 ( Type Assertion )
	- 인터페이스 값의 기초적인 콘크리트 값에 대한 접근을 제공합니다	
	- interface type의 value의 Type을 확인하는 것이다.
	
```golang
t := i.(T)
```
	- 인터페이스 값 i가 콘크리트 타입 T를 가지고 있으면 기본 값인 T값을 변수 t에 할당
		- 만약 i가 T값을 가지지 못하면 이 선언은 panic
	
	- 선언 성공 여부를 판단하기 위해 다음과 같이 선언
		- i 가 T를 갖는다면 t는 underlying value, ok는 true
		- 아니라면 ok는 false가 되고 t는 T라는 유형의 zero값이 됩니다.
		- 선언에 실패해도 panic이 발생하지 않음
```golang
t, ok := i.(T)
```

	- 예제
		- 타입이 맞는지 확인하는 용도로 사용 >> type switch?
		- nil 인터페이스는 변환이안됨.
```golang
package main

import "fmt"

func main(){
	var inter2 interface{}
	tmp2 := inter2
	fmt.Println(tmp2) // <nil>
	// s_tmp2 := inter2.(string) // panic: interface conversion: interface {} is nil, not string
	// fmt.Println(s_tmp2) // <nil>

	var inter interface{} = "hello"
	tmp := inter
	fmt.Println(tmp) // hello

	s := inter.(string)
	fmt.Println(s)

	s, ok := inter.(string)
	fmt.Println(s, ok) // hello true

	f, ok := inter.(float64)
	fmt.Println(f, ok) // 0 false

	// f = inter.(float64) // panic
	// fmt.Println(f)
}
```
	
TODO - ????????콘크리트 타입이라는건 뭐지
TODO - ???????? panic은?
	
- 타입 스위치
	- 여러 타입의 선언을 직렬로 허용하는 구조
	- 일반 스위치문과 같지만, 값이 아닌 타입을 명시하고 값들은 지정된 인터페이스 값에 의해 유지되는 값의 타입과 비교됩니다.
	
	- 타입 스위치의 선언은 타입 선언 i.(T) 와 같은 구문을 가집니다.
	- 인터페이스가 보유하지 않은 타입인 경우(디폴트 케이스) 변수 v는 인터페이스 종류와 상관없이 값이 i와 같음??
	
```golang
package main

import "fmt"

func do_switch(i interface{}) {
	// type switch 에서만 사용 가능한 특별한 구문. i.(type)
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}


func main() {
	do_switch(21)      // Twice 21 is 42
	do_switch("hello") // "hello" is 5 bytes long
	do_switch(true)    // I don't know about type bool!
}

```

- Stringers
	- 내장 인터페이스
	- 인터페이스 중 하나로 fmt 패키지에 의해 정의된 Stringer 입니다.
	
	- 자신을 문자열로 설명할 수 있는 타입입니다. 
	- fmt 패키지 및 기타 여러 패키지는 값을 출력하기 위해 이 인터페이스를 사용합니다.
	
```
// Stringer의 구현은 다음과 같다. String()이라는 메서드를 가지고 있음
type Stringer interface {
	String() string
}
```
	- 사용예
	- String()을 한 것과 그냥 출력한 것이 같다. >> 프린트류 함수들이 String()을 통해 호출한다고 추정
	
```golang
package main

import "fmt"

func main(){
	coffeepot := CoffePot{name: "Coffee", kind: "Espresso"}
	fmt.Println(coffeepot.String()) // Coffee: Espressocoffee pot
	fmt.Println(coffeepot)          // Coffee: Espressocoffee pot

}

type CoffePot struct {
	name string
	kind string
}

func (c CoffePot) String() string {
	return c.name + ": " + c.kind + "coffee pot"
}

```

- Stringer 예제
	- [4]byte 인  타입의 값을 "1.2.3.4" 와 같이 출력되도록 개발하라
	
```golang
package main

import "fmt"

type IPAddr [4]byte

// 	작성함수
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	
	// map의 range 활용
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

	
- Errors
	- Go 프로그램은 'error' 값으로 오류 상태를 표현합니다.
	- error 타입은 fmt.Stringer 와 유사한 내장 인터페이스
	- error 인터페이스로 custom error 타입을 만들 수 있습니다.
	- 예제를 보니 error의 값으로 오류 로직을 처리함.(살짝 귀찮아 질 수도?)

TODO - fmt.Errorf("asd") 는 뭐지?
TODO - Golang의 에러처리 방법 
	- 에러를 값으로 판별해서 처리하는데... 상위?에서 처리하도록 버블링하고싶을 땐 어떻게?????????
TODO - Sprintf등 프린트 함수종류마다 뭐가 다른거니?????????????????????

```golang
// error 인터페이스. Error() 메서드.
type error interface {
	Error() string
}
```
	
	- 오류(err)가 nil 과 같은지 테스트하여 오류를 처리해야합니다.
		- nil error 는 성공을 나타냅니다; nil이 아닌 error 는 실패를 나타냅니다.
		
```golang
i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't convert number: %v\n", err)
	return
}
fmt.Println("Converted integer:", i)
```

	- 실제로는 구조체 타입에 커스텀 에러를 만들 수 있다.
	- 에러 구조체를 생성하고, Error()함수를 구현 > 에러 출력할 메시지.
	- 에러를 발생할 수 있는 경우 리턴을 error 인터페이스로 함.(그럼 커스텀 인터페이스를 받을 수 있음
	
```golang
package main

import (
	"fmt"
	"time"
)

// 커스텀 에러 구조체
type MyError struct {
	When time.Time
	What string
}

// 인터페이스 error를 구현한다.
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// 에러가 발생할 경우에 error(custom Error)를 리턴해줌.
// 에러가 아니면 nil을 리턴해주면 로직처리에서 편함
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

```

```golang
// 다음예제

type overHeatError float64

func (o overHeatError) Error() string {
	return fmt.Sprintf("overheat temperature: %0.2f", float64(o))
}

func checkTemperature(actual float64, criteria float64) error {
	excess := actual - criteria
	if excess > 0 {
		return overHeatError(excess)
	}
	return nil // 에러가 아닌 경우는 nil
}
func main(){
	err := checkTemperature(38.5, 37.5)
	if err != nil {
		fmt.Println(err) // overheat temperature: 1.00
	}
}
```


	
- Errors연습문제

```
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return x, nil
	}
	return 0, ErrNegativeSqrt(x)
}

func main() {
	fmt.Println(Sqrt(2))  // 2 <nil>
	fmt.Println(Sqrt(-2)) // 0 cannot Sqrt negative number: -2
}

```

- Readers
	- io 패키지는 데이터 스트림의 읽기를 나타내는 io.Reader 인터페이스를 지정합니다.
	- io 패키지에 있음
	
```golang
// Reader 인터페이스 구조
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

TODO - golang에서 byte stream을 다루는 방법에는
	https://etloveguitar.tistory.com/100

TODO - reader 개선하는 예제
	https://velog.io/@yeonjoo7/go%EC%96%B8%EC%96%B4%EC%9D%98-io%ED%8C%A8%ED%82%A4%EC%A7%80
	
	
```golang
//예제
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!") // reader생성

	b := make([]byte, 8) // byte slice 생성. reader를 읽어서 저장할 공간
	for {
		// n은 읽어낸 byte수, err는 결과
		// 더이상 읽은 것이 없으면 err은 io.EOF ("EOF")를 뱉음
		n, err := r.Read(b) 
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

// n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
// b[:n] = "Hello, R"
// n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
// b[:n] = "eader!"
// n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
// b[:n] = ""

```

- Reader 예제 -1 
	- ASCII 문자의 무한 스트림을 방출하는 Reader 유형 구현 'A'.
TODO - 아직 이해 못함
```
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

```

- Reader 예제 -2
	- rot13Reader
	-
https://godsman.tistory.com/entry/golang-A-Tour-of-Go-61-%EC%97%B0%EC%8A%B5-%EB%AC%B8%EC%A0%9C

	- ROT13 치환 암호화는 입력 알파벳을 13만큼 이동시킨 것으로, 
	- a -> a+13 = n, b -> b+13 = o로 치환합니다. 
	- 알파벳 z가 넘어갈 경우에는 다시 a 부터 시작하도록 순환합니다. 
	- p -> p+13 = z+3 = c

TODO - 아직 이해 못함

```golang
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13byte(sb byte) byte {
	s := rune(sb)
	if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
		sb += 13    
	}
	if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
		sb -= 13    
	}    
	return sb
}

func (rot13 rot13Reader) Read(data []byte) (len int, err error) {
	len, err = rot13.r.Read(data)
	for i := 0; i < len; i++ {
		data[i] = rot13byte(data[i])    
	}    
	return
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

```


- Images
	- Package image 는 Image 인터페이스를 정의합니다:
	https://pkg.go.dev/image#RGBA
```golang
// 구현
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```
	 - 예제
	 
```
package main

import (
	"fmt"
	"image"
)

func main() {
	// image.Rect image.Rectangle 를 반환
	// image.NewRGBA *image.RGBA 를 반환
	// NewRGBA는 image에 이미 구현되어 있는 함수일 것입니다.
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA()) // 픽셀의 R, G, B, A값이 uint32 로 출력
	fmt.Println(m.At(0, 1).RGBA())
	fmt.Println(m.At(0, 2).RGBA())
}

// RGBA 구현체
type RGBA struct {
	// Pix holds the image's pixels, in R, G, B, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect Rectangle
}

```

- Images 연습문제
	- 일단 풀긴햇는데... 찾아보기
	
```
package main

import "image"
import "image/color"
import "golang.org/x/tour/pic"

type Image struct {
	width, height int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}


func main() {
	m := Image{500, 500}
	pic.ShowImage(m)
}

```
