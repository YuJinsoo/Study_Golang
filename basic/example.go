package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// context 라이브러리란?

// 그냥 계속 실행할 계획이면..
func tick_forever() {
	
}

// 시작 - 취소예제
func start_tick(ctx context.Context, duration time.Duration) {
	// time.Tick은 정지할 수 없음
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case t := <- ticker.C:
			fmt.Printf("주기적으로 실행할 작업 수행: %v \n", t)
		case <- ctx.Done():
			fmt.Println("종료")
			return
		}
	}
}

func main() {
	fmt.Println("start")
	ticker := time.Tick(5*time.Second)

	go func() {
		for t := range ticker {
			fmt.Println("Tick at %v\n", t)
		}
	}()
	
	// 시작 - 취소예제
	// ctx, cancel := context.WithCancel(context.Background())
	// go start_tick(ctx, 5*time.Second)

	// go func() {
	// 	// 특정시간 뒤에 취소됨
	// 	time.Sleep(20*time.Second)
	// 	cancel()
	// }()
}

// 지연 미들웨어 적용.
func delay_middleware(next http.Handle) http.Handle {
	return http.HandleFunc( func(w http.ResponseWriter, r * http.Response) {
		time.Sleep(1*time.Second)
		next.ServeHTTP(w, r)
	})
}

func my_handler(w http.ResponseWriter, r *http.Request){}

func was_main(){
	r := mux.NewRouter()

	// 미들웨어를 사용하여 모든 요청에 대해 지연을 추가합니다.
	r.Use(DelayMiddleware)

	// 특정 경로에 대한 요청을 핸들링합니다.
	r.HandleFunc("/", my_handler)

	http.ListenAndServe(":8080", r)
}

// rate limiting
// 사용 방법은 use

// "golang.org/x/time/rate"
var limiter = rate.NewLimiter(1,3) // 초당 1개 요청, 버스트 사이즈 3

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// rate.NewLimiter(r, b)는 레이트 리미터를 생성
// 여기서 r은 초당 요청 수를 의미하고, b는 버스트 사이즈를 의미
// 버스트 사이즈는 짧은 시간 동안 허용되는 최대 요청 수
// RateLimitMiddleware 미들웨어는 각 요청이 들어올 때마다 레이트 리미터의 Allow 메소드를 호출하여 요청이 허용되는지 확인
// 요청이 허용되지 않으면, 429 Too Many Requests 응답을 반환

