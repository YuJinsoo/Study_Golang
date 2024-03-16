

## 흐름제어

### 반복문: for
	- 단 하나의 for문
	- while도 for로 대체함
	- for 초기화; 조건; 사후구문 {} 
		- 초기화와 사후구문을 생각해서 while 기능 구현 가능
		- 모두 없으면 forever-for로 무한루프 생성
		
```golang
func main(){
	sum := 0

	// 기본적인 for 구문. ()가 없고 구분을 구분하는 {} 만 있음.
	// 초기화 ; 조건; 사후구문
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 45

	// 초기화 구문과 사후 구문은 필수가 아닙니다. 조건만 작성 가능
	// ;; 는 생략됨
	// ;; 가 생략되기 때문에 C의 while같은 구문은 for로 쓰입니다.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum) // 1440

	// 무한루프 만들기
	// forever-for 문
	// 반복 조건을 생략하면 무한루프 for 생성
	for {
	}
}
```

### 조건문: if
	- ( ) 괄호로 둘러쌓일 필요는 없지만, { } 괄호는 필수입니다
	
	- 짧은 구문 if
		- 조건문 전에 수행할 짧은 구문으로 시작할 수 있음. 조건문과 `;` 로 구분
		- 짧은 구문에서 선언된 변수들은 오직 if 문의 끝까지만 스코프가 존재합니다.
			- if나 else 구분에서만 사용 가능
		
```golang
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
```

- if else
	- if 문 안에서 선언한 변수들은 어떤 else에서도 사용할 수 있습니다.
```golang
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// v 사용 가능
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 여기서는 v를 사용할 수 없음
	return lim
}
```

- else if
	- else if 는 다음과 같이 사용하 수 있다.
	
```
var a = 17

	if a < 5 {
		fmt.Println("a < 5")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("a >= 10")
	}
```
	
### Switch 문
	- 문자열 switch 가능
	- switch case는 상수일 필요가 없으며 그 값들은 정수일 필요도 없습니다.
	- 블록 종료시 break문이 필요 없음.
```golang
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

- Switch  평가 순서
	- 위에서부터 순차적으로 진행합니다.
	- 일치하는 case가 있으면 그 이후로는 평가하지 않습니다.
	
- 조건 없는 Switch
	- switch true와 동일합니다.
	- 아주 긴 if-else를 구현하는데 아주 깔끔한 방식입니다.

```golang
t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}
```

- break and fallthrough
	- 타 언어에서 break가 없으면 다른 블록도 실행되지만, Go는 break없어도 블록이 종료됨
	- 하지만 다음 케이스를 실행하고 싶다면 `fallthrough`구문을 사용하면 됩니다.
	
```golang
package main

import "fmt"

func main() {
	a := 2

	switch a {
		case 1:
			fmt.Println("a == 1")
		case 2:
			fmt.Println("a == 2")
			fallthrough
		case 3:
			fmt.Println("a == 3")
		default:
			fmt.Println("a is not 1, 2 or 3")
	}
	// a == 2
	// a == 3
}
```

- 여러개의 값에서 같은 블록을 동작시킬 수 있습니다.

```golang
package main

import "fmt"

func main() {
	switchResult(1)
	switchResult(2)
	switchResult(3)
	switchResult(4)
	switchResult(5)
	// a == 1 or a == 2
	// a == 1 or a == 2
	// a == 3 or a == 4
	// a == 3 or a == 4
	// a is not 1, 2, 3 or 4
}

func switchResult(a int) {
	switch a {
			case 1, 2:
				fmt.Println("a == 1 or a == 2")
			case 3, 4:
				fmt.Println("a == 3 or a == 4")
			default:
				fmt.Println("a is not 1, 2, 3 or 4")
		}
}
```


### Defer
	- 자신을 포함하는 함수가 반환되기 전에 실행하도록 예약한다. (연기된 함수라고 볼 수 있음)
	- 연기된 호출의 인자는 즉시 평가되지만 그 함수 호출은 자신을 둘러싼 함수가 종료할 때까지 수행되지 않습니다.
	- 종료 직전에 defer로 호출된 함수가 실행됩니다.
	
	- stack과 같이 동작함.
		- 함수 내에서 defer로 등록한 것들은 함수가 끝날때 맨 마지막에 추가된 동작부터 수행합니다.
		- 후입선출 LIFO
		- https://go.dev/blog/defer-panic-and-recover

https://velog.io/@whdnjsdyd111/GO-2-2.-%EA%B3%A0%EB%9E%AD-%EA%B8%B0%EB%B3%B8%EB%AC%B8%EB%B2%95-%ED%8C%A8%ED%82%A4%EC%A7%80

- defer 스택 쌓기
	- defer는 스택처럼 동작합니다.
	- 한 함수에서 여러 defer를 선언하면, 함수 종료 직전 가장 마직막에 선언한 defer 부터 실행됩니다.
	- 후입선출!
	
	- 어디에 사용된는 걸까?
		- 파일 관련 동작
			- 함수가 종료되기 전 반드시 close를 해야 하는데, 조건에 따라 return 시점이 다르면 매번 close해주어야 하는 번거로움
			- https://bugoverdose.github.io/development/go-defer-basics/
			
```golang
// 예제추가필요
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

```

TODO - panic과 recover!
