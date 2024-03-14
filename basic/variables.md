## 기본 자료형
	Go 에는 6가지 정도로 분류 가능한 자료형이 있습니다.

	- bool(참/거짓)
		- bool
	- 정수형 타입
		int, int8, int16, int32, int64
		uint, uint8, uint16, uint32, uint64
		uintptr
	- Float 및 복소수 타입
		float32, float64
	- 복소수 
		complex64, complex128
		복소수를 저장해요. 
		1 + 2i 와 같이 저장
		각각의 실수부, 허수부를 호출할 때에는 real() 함수와 imag()함수를 이용해요
	- 문자열 타입
		- string
	- 기타 타입
		byte: uint8과 동일, 바이트 코드에 사용
		rune: int32과 동일, 유니코드 코드포인트(문자에 부여한 고유한 숫자값)에 사용한다

	- u가 접두사로 붙은건 unsigned를 의미하고, 접미사로 붙은 숫자는 비트(bit)를 나타내요
	숫자가 붙지 않은 자료형(int, uint, uintptr)은 32-bit 시스템에서는 32bit, 64-bit 시스템에서는 64bit 길이에요
	특별히 정수의 크기나 부호를 지정할 이유가 없다면 int를 쓰면 됩니다.


## 변수
	- var 예약어로 선언합니다.
	- `:=` 로 선언시 var를 생략할 수 있습니다. (함수 밖에서는 := 사용할 수 없음)
	- 여러개를 한 번에 선언할 수 있습니다. 타입은 뒤에
	- 이미 선언된 변수는 C 나 C++ 과 같이 변수에 정적으로 자료형이 고정됩니다
	
```golang
var a int = 13
b := 123
var c, python, java bool // zero values 할당됨
```
	
	- 한 번에 변수 초기화가 가능합니다.
	- 이 경우 변수는 초기값의 타입을 가지게 됩니다.
	
```golang
var c, python, java = true, false, "no!"
```

### zero values
	- 명시적인 초기값 없이 선언된 변수는 그것의 zero value가 할당됩니다.
	- 숫자 : 0
	- boolean: false
	- string: ""
	- pointer : nil

- inference type
	- 명시적인 type을 정의하지 않고 변수를 선언할 때, 그 변수 type은 오른 편에 있는 값으로부터 유추하는 것을 의미

```golang
v := "hello world!" 				// v가 string 타입으로 유추됨
fmt.Printf("v is of type %T\n", v) 	// v is of type string
```

### 형변환
	- 타입을 변환하는 기능
	
```golang
const a = 1
b := float32(a);
println(b);
```

### 상수
	- const 예약어로 선언합니다.
	- string, boolean, 숫자가 가능합니다. :=으로 선언할 수 없습니다.

```golang
const a int = 1
const b string = "hello"
```

- 숫자형 상수
```golang
package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int64) int64 { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// cannot use Big (untyped int constant 1267650600228229401496703205376) as int value in argument to fmt.Println (overflows)
	// 위 에러가 발생하지만 값은 매우 정확하게 알고 있음.
	// fmt.Println(Big)
	
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

```
