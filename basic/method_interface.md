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
