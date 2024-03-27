## Concurrency (동시성)

### Goroutines
	- Go 런타임에서 관리되는 경량 쓰레드
	- 런타임에 의해 관리되므로 개발자는 OS 레벨의 스레드와 다르게 직접 반납하거나 종료시키는 과정이 없다.
	- goroutines은 가볍고 stack 공간 할당보다 약간의 비용이 더 듭니다.
	- goroutines은 여러개의 os의 스레드에 의해서 다중처리가 되므로
		- 만약에 하나 goroutines이 I/O를 대기와 같은 이유로 Block이 되더라도
		- 다른 goroutines들은 멈추지 않고 수행
 
	- 왜 사용할까? 다른거랑 무슨 차이가 있는거지? (python의 코루틴과 비교해보자 - 완전다름)
		- context switching 비용이 낮다(일반적으로 16개 레지스터 값 교체(save/restore) 필요. go는 3개(PC, SP, DX) )
		- 약 2KB의 스택 메모리 영역만 필요함.
		
	- https://velog.io/@khsb2012/go-goroutine

``` golang
package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// world를 출력하는 goroutine과 hello를 출력하는 goroutine이 동시에 실행된다.
func main() {
	// 예제1 동시에 돌아가는 출력문
	go say("world")
	say("hello")

	println("=======================================")

	// 예제2 sync.WaitGroup >> TODO - 뭔지 찾아보기
	wg := sync.WaitGroup{}
	wg.Add(1)

	counter := 0

	// 임의함수를 goroutine으로 실행
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000; i++ {
			counter++
		}
	}()

	wg.Wait() // goroutine이 끝날때까지 대기
	fmt.Println("counter:", counter)

}

```


### Channel
	- 채널 연산자인 `<-` 을 통해 값을 주고 받을 수 있는 하나의 분리된 통로
		- 데이터는 화살표 방향대로 흐름
	- 기본적으로 전송과 수신은 다른 한 쪽이 준비될 때까지 block 상태
		- 명시적인 lock이나 조건 변수 없이 goroutine이 synchronous하게 작업될 수 있도록 합니다.
	- map과 slice처럼 사용하기 전에 생성되어야만합니다.
	
```golang
ch := make(chan int) // 속성은 chan, 전달할 데이터 타입
```
		
```golang
channel <- value    // 채널 channel에 value를 전송한다.
value := <-channel  // channel로 부터 값을 받고,
					// 값을 value에 대입한다.
```

	- channel 예제

TODO - channel 상태를 어떻게 구분하는지 case를 알고싶다.

```golang
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```
TODO - 어떤 채널이 먼저 끝나는지 알 수 있나?
	1. 로그로 찍기
	2. sync.waitgroup사용
	3. 결과 채널에 추가 정보 포함.


	
- Buffered Channel
	- 채널의 버퍼 크기를 지정해서 선언할 수 있습니다.
		- make함수의 두 번째 인자에 버퍼 크기 지정
	- 버퍼의 크기를 넘거나, 비어있을 때 값을 참조하면 런타임 에러가 발생합니다.
	
``` golang
func main() {
	// 버퍼 크기만큼만 가지고 있을 수 있음. 
	// 버퍼가 가득 차면 block됨
	ch := make(chan int, 2) // buffer 크기 = 2
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
	ch <- 4
	// ch <- 5 // deadlock. 버퍼에서 값을 꺼내서 빈자리가 날때까지 기다리는듯
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

```


- Range와 Close
	- 전송자는 더 이상 보낼 데이터가 없다는 것을 암시하기 위해 channel을 close할 수 있습니다
	- 수신자는 수신에 대한 표현에 두 번째 매개변수를 할당하여 채널이 닫혔는지 테스트 할 수 있음
		- 채널이 닫혀있다면 ok는 false를 가짐
```
v, ok := <- ch
```
	
	- 절대로 수신자가 아닌 전송자만이 channel을 닫아야 함
		- 닫힌 channel에 전송하면 panic 발생
		
	- 채널을 닫는 동작은 수신자가 더 이상 들어오는 값이 없다는 것을 알아야하는 경우에만 필요합니다.
		- 파일 다루는 것처럼 무조건 해야 하는 것은 아님.

TODO - ????????????????? //channel 을 range하는게 이해안됨
```
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // 위 for문을 돌면 버퍼가 가득 차므로 close를 해줌
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// channel 을 range하는게 이해안됨
	for i := range c {
		fmt.Println(i)
	}
}
```
	- close 예제
	
```golang
func main() {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// 방법1
	// 채널이 닫힌 것을 감지할 때까지 계속 수신

	for {
		if i, success := <-ch; success {
			println(i)
		} else {
			break
		}
	}

	// 방법2
	// 위 표현과 동일한 채널 range 문
	for i := range ch {
		println(i)
	}
}

```


- Select
	- 복수 채널들을 기다리면서 준비된 (데이타를 보내온) 채널을 실행하는 기능을 제공
	- goroutine이 다중 커뮤니케이션 연산에서 대기할 수 있게 합니다.
	- 만약 다수의 case가 준비되는 경우에는 select가 무작위로 하나를 선택합니다.
	- select 는 자기 case들 중 하나가 실행될 때까지 block됩니다.

```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

EXIT:
	for {
		select {
		case <-done1: // done1 채널에 값이 들어온 case
			println("run1 완료")

		case <-done2: // done2 채널에 값이 들어온 case
			println("run2 완료")
			break EXIT
		}
	}
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

```

```golang
// playground 예제
package main

import (
	"fmt"
	"time"
)

func fibonacci_sel(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func select_main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // channel이 준비 될때까지 blocking
		}
		quit <- 0
	}()
	fibonacci_sel(c, quit)
}

```


- Default Selection
	- select 에서의 default case 는 다른 case들이 모두 준비되지 않았을 때 실행됩니다.
	
```
select {
case i := <-c:
	// use i
default:
	// c로부터 값을 받아오는 것이 block된 경우
}
```

	- 예제

```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	// 둘다 채널임 
	// <-chan time.Time
	tick := time.Tick(100 * time.Millisecond)	// 100미리마다 tick 발생
	boom := time.After(500 * time.Millisecond)	// 500미리 이후 tick발생
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

```


- 연습문제
```
// playground에서 돌리기
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// 전달한 tree의 Left, Right를 순회하면서 값을 채널에 전달
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)

	ch2 := make(chan int)
	go Walk(t2, ch2)

	// range for문의 경우 채널이 clsoe되기 전까지 계속 값을 추출한다.
	// close가 없다면 빈 channel에서 값을 빼오려고 하기 때문에 deadlock이 발생한다.

	// 채널에서 값을 읽어서 비교
	// for i := range ch1 { // TODO range이해하고 다시
	for i := <-ch1; i < 10; i++ {
		d := <-ch2
		if i != d {
			fmt.Println(i)
			fmt.Println(d)
			return false
		}
	}

	return true
}

func main() {
	// 매번 돌릴때마다 다르게 나옴 >> 다양한 이진 트리가 있기 때문...
	// var t1 = tree.New(1)
	// var t2 = tree.New(1)
	// fmt.Println(t1)
	// fmt.Println(t2)

	if Same(tree.New(1), tree.New(1)) {
		fmt.Println("Same")
	} else {
		fmt.Println("Different")
	}
}


```

- sync.Mutex
	- mutual exclusion 이라고 불리고, 자료 구조에서 그것의 관습적인 이름은 mutex 라고 함
	- (기본 communication 이외에)충돌을 피하기 위해 오직 하나의 goroutine만이 특정 순간에 특정 변수에 접근할 수 있도록 하고 싶을 때 사용
	
	- Go의 표준 라이브러리는 sync.Mutex와 그것의 두 가지 method를 통해 mutual exclusion을 제공합니다
		- Lock, Unlock

	- Lock과 Unlock으로 감싼 구문.
	- unlock에 defer를 사용하는 경우도 있음
	
```golang
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```
	
- 예제 다음에?





TODO - GO로 oop 흉내내기?
	https://velog.io/@kineo2k/Go-%EC%96%B8%EC%96%B4-OOP-%ED%9D%89%EB%82%B4-%EB%82%B4%EA%B8%B0