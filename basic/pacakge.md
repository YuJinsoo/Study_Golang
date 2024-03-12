
## 패키지와 변수, 함수

### 패키지
	- 모든 Go 프로그램은 패키지로 구성되어있습니다.
	- 프로그램은 main 패키지에서 실행을 시작합니다.
	- 패키지의 이름은 import 경로의 마지막 요소와 같습니다. 
		- 예를 들어 "math/rand" 패키지는 package rand 문으로 시작하는 파일들로 구성되어있습니다.

	- Go 언어에서는 한 디렉토리 내의 모든 Go 파일이 동일한 패키지 이름을 사용해야 합니다. 
	

### import
	- 외부 패키지를 가져오는 구문?
	- 사용하지 않는 패키지들은 저장시 자동으로 사라진다..
	
	- 패키지를 강제로 유지하기: 
		- 특정 이유로 패키지를 코드에서 직접 사용하지 않더라도 import 구문에 유지하고 싶은 경우, 
		- Go에서는 빈 import를 사용할 수 있습니다. 패키지 경로 앞에 _를 붙여 이를 구현할 수 있습니다. 
			ex) _ "math"
	
- Export 되는 이름
	- 첫 글자가 대문자인 것은 Export되기 때문에 사용할 수 있음
		ex) ftm.Printf(), fmt.Println(), math.Pi 등등


### 함수
	- 변수, 함수 이름 뒤에 타입
	- 복수개의 결과는 := 연산자를 사용해서 가능

- 변수, 함수 선언방식이 C 스타일과 다른 이유
	- c스타일은 함수포인터류의 선언이 복잡한 것들은 알아보기 힘들다.
	- 이름이 먼저가 아니라 타입이 먼저라 더 헷갈림
	https://go.dev/blog/declaration-syntax


### 함수에서 이름이 주어진 반환값
	- 반환값에 이름을 붙일 수 있음. 'naked return'
	- 이름이 정해진 반환값은 함수 내부에서 선언한 변수처럼 사용
	- 리턴을 굳이 해주지 않아도 순서대로 리턴
	- 가독성을 떨어뜨리므로 짧은 함수에서만 사용을 권장

```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

### 가변 인자 함수 (추가)

	- 전달하는 매개변수의 개수를 호출할 때마다 다르게 전달할 수 있는 함수
	- 가변인자함수 특징
		1. n개의 동일한 형의 매개변수를 전달.
		2. 전달된 변수들의 형태는 slice형
		3. "func FuncName(매개변수이름 ...매개변수형) 반환형"으로 선언 (매개변수형앞에 ...을 붙이면 됨)
		4. 매개변수로 슬라이스를 전달 할 수 있음. "함수이름(슬라이스이름...)"

```go
package main

import (
	"fmt"
	"math"
)

func main(){
	s1, s2, s3, s4, s5 := "picachu", "raichu", "pairi", "kame", "hennaHana"
	ss := []string{"a1", "a2", "a3", "a4"}
	fmt.Println(addCaseOne(1, s1, s2, s3, s5)) // picachu, raichu, pairi, hennaHan
	fmt.Println(addCaseOne(2, s1, s2, s4))     // picachu, raichu, kame
	fmt.Println(addCaseOne(3, ss...))          // a1, a2, a3, a4
}

// 가변인자 함수
// 가변인자 함수는 함수의 마지막에 위치해야 한다.
func addCaseOne(aaa int, names ...string) string {
	var result string
	
	// 가변인자로 받은 변수의 타입은 슬라이스
	fmt.Printf("%d - len(names) : %d Type: %T \n", aaa, len(names), names)

	for i := 0; i < len(names); i++ {
		result += names[i]
		if i+1 != len(names) {
			result += ", "
		}
	}

	return result
}
```
	
참고: https://velog.io/@csh6222/Golang-%ED%95%A8%EC%88%98%EC%A0%95%EB%A6%AC

### 익명함수
	- 함수를 선언하면 전역으로 초기화 되면서 메모리를 소모된다
	- 사용할 때에도 메모리 주소를 검색해서 사용함.
	- 반면 익명함수를 호출될 떄마다 메모리에 바로 올라가 실행후 사라짐
	- 1회성으로 사용하는 부분에 활용하면 좋음
	
```go
package main

import "fmt"

func main() {
	// 함수의 이름만 없고 그 외에 형태는 동일
	// 함수의 블록 마지막 브레이스(}) 뒤에 괄호(())를 사용해 함수를 바로 호출
	// 괄호 안에 매개변수를 넣을 수 있음.
	func(){
		fmt.Println("Hello EveryBody")
	}()
	
	func(name string, age int) {
		fmt.Println(fmt.Sprintf("Hello My Name is %s, %d years old" , name, age))
	}("show", 30)

	// 선언 함수는 반환 값을 변수에 초기화함으로써 변수에 바로 할당이 가능
	introducingMessage := func(name string, age int) string{
		return fmt.Sprintf("Hello My Name is %s, %d years old" , name, age)
	}("Hopangmen",100)
	fmt.Println(introducingMessage)

	name, age := "Golang" , 30
	introducingMessage2 := func(n string, age int) string {
		return fmt.Sprintf("Hello My Name is %s, %d years old" , name, age)
	}(name, age)
	fmt.Println(introducingMessage2)

}
```