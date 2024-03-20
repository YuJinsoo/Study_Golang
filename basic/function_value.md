
### Function value

- Function values (함수값)
	- 함수들도 값입니다. 그것들은 다른 값들과 마찬가지로 인자로 전달될 수 있습니다.
	- 함수 값은 함수의 인수나 반환 값으로 사용될 수 있습니다.
	- 변수에 함수를 저장하는 것도 가능 (익명함수 활용 방법)
	
```golang
package main

import (
	"fmt"
	"math"
)

// 인자로 함수를 전달하는 함수
func compute(func_parmeter func(float64, float64) float64) float64 {
	return func_parmeter(3, 4)
}

func main() {
	// 익명함수를 변수에 할당.
	// hypot는 float64 인자 2개를 받고 float64를 반환하는 익명함수
	// 내부 동작은 math.Sqrt(x*x + y*y) 를 실행하여 결과반환
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // 13

	fmt.Println(compute(hypot)) // 5  // hypot 함수 심볼이 compute의 인자에 전달될 수 있음.
	fmt.Println(compute(math.Pow)) // 81
}

```
	
- Function closures(함수 클로저)
	- Go의 함수들은 클로저일 수도 있음
	- 함수의 외부로부터 오는 변수를 참조하는 함수 값
	- 왜 사용할까????
		1. 함수의 캡슐화
		2. 불필요한 전역변수를 없애고, 변수를 공유할 수 있게 된다.
			https://hwan-shell.tistory.com/339
	
	- 예제해석
		- 변수에 adder 함수를 할당하고 for문에서 각 adder함수를 실행함
		- 이때 pos와 neg에 할당한 함수는 adder내부에 리턴(더하는 함수) 뿐이지만, adder(밖 범위)의 sum 변수를 기억하고 있음
			- pos와 neg에서 각각 기억
		- 함수는 변수에 "bound(바운드)" 됩니다. 각 클로저는 sum 변수에 바운드 된 것
	
```golang
package main

import (
	"fmt"
	"math"
)

// 클로저함수예제
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
	// 0 0
	// 1 -2
	// 3 -6
	// 6 -12
	// 10 -20
	// 15 -30
	// 21 -42
	// 28 -56
	// 36 -72
	// 45 -90
}

```

- 클로저 연습문제
	- 리턴하는 함수 밖의 변수를 활용한다고 생각해서 해결
```
package main

import "fmt"

func fibonacci() int {
	// main이 동작하는 피보나치 수열 클로저 함수
	// 이번 항과 다음 항을 기억하는 정보를 클로저 변수로 만들자
	now, next = 0, 1
	// fibonacci()의 리턴 형태를 맞춘 함수. 사실 둘다 int 없어도 되긴 함
	return func() int {
		ret := now
		now, next = next, now + next
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```
