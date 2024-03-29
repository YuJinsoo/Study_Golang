
# Go lang 학습
- google이 만든 언어
- 빠른 성능, 안정성, 편의성, 쉬운 프로그래밍을 목표로 제작된 언어

playground: 
widi docs : https://wikidocs.net/163777

### 장단점

- 장점
    1. 속도
        - Go는 C를 기반으로 만든 컴파일 언어이기 때문에, 인터프리터 언어(ex. Python)에 비해 월등히 빠른 퍼포먼스를 보인다.
        - 네이티브 바이너리를 결과물로 제공하기 때문에, 배포 입장에서도 매우 빠르게 처리할 수 있다.

    2. Garbage Collection
        - Go 역시 auto gc를 지원하기 때문에, 개발자 입장에선 매우 편리하다.
        - C/C++의 포인터로 인한 Memory Leak 발생이 없다.

    3. 진입장벽이 낮다
        - 위의 GC와도 연관되는 장점이다.
        - Go는 혁신적으로 불필요한 문법들을 가차없이 가지치기하였다. (ex) while 문 삭제
        - 특히, 멀티스레드는 go 키워드를 사용해서 1줄로 구현해 낼 수 있다!

    4. 컨벤션의 통일
        - 코딩 컨벤션을 일체화시켰다.
        - Go는 `{` 는 반드시 함수 옆에 위치해야 한다" 처럼 컨벤션을 fix시켜 버렸고,
        - 이에 스트레스 받을 개발자들을 위하여 컨벤션에 맞춰 코드를 수정해주는 'go fmt' 기능을 제공하고 있다.

- 단점


### 장단점 - gpt

- 장점

1. 간결하고 명확한 문법: 
	- Go는 간결하고 이해하기 쉬운 문법을 가지고 있어, 코드를 빠르게 작성하고 유지보수하기 쉽습니다.
2. 뛰어난 동시성 처리: 
	-Go는 고루틴(Goroutines)을 통해 경량 스레드를 사용하여 동시에 여러 작업을 처리할 수 있게 합니다. 이는 동시성 프로그래밍을 간단하고 효율적으로 만들어 줍니다.
3. 빠른 컴파일 속도: 
	- Go는 매우 빠른 컴파일 속도를 자랑합니다. 이는 개발자의 생산성을 높여 줍니다.
4. 강력한 표준 라이브러리: 
	- Go는 네트워킹, 문자열 처리, 그리고 I/O와 같은 다양한 작업을 위한 풍부한 표준 라이브러리를 제공합니다.
5. 크로스 플랫폼 개발 지원: 
	- Go는 다양한 운영 체제와 아키텍처에서 실행될 수 있도록 크로스 컴파일을 지원합니다.
	
- Golang의 단점
1. 제네릭의 부재(최근에 추가됨): 
	- 이전에 Go는 제네릭을 지원하지 않았으나, Go 1.18 버전에서 제네릭을 도입하여 이 단점을 어느 정도 해결했습니다.
2. 인터페이스와 타입 시스템: 
	- Go의 인터페이스와 타입 시스템은 때때로 더 복잡하거나 유연하지 못한 설계를 초래할 수 있습니다.
3. 컬렉션 연산의 한계: 
	- Go는 map, slice, array 외에 다양한 컬렉션 연산을 기본적으로 제공하지 않아, 때로는 라이브러리를 추가로 찾거나 직접 구현해야 할 수 있습니다.


### 주로 사용되는 분야
1. 클라우드 서비스 및 인프라스트럭처: 
	- Google Cloud, Docker, Kubernetes와 같은 클라우드 기반 인프라와 서비스의 개발에 널리 사용됩니다.
2. 마이크로서비스 아키텍처: 
	- Go의 간결함과 병행성 지원 덕분에, 분산 시스템과 마이크로서비스 기반의 아키텍처 구축에 이상적입니다.
3. 데이터베이스 및 스토리지 시스템: 
	- InfluxDB, CockroachDB와 같은 고성능 데이터베이스 시스템 개발에도 사용됩니다.

### 블록체인에서의 Golang의

- Golang은 블록체인 개발에 널리 사용되는 언어 중 하나입니다. 
- Ethereum의 Go 구현체인 Geth(Go Ethereum)이 Go로 작성되었습니다. 
- Golang의 성능, 병행성 지원, 그리고 네트워킹 기능은 분산 네트워크와 블록체인 애플리케이션의 개발에 매우 적합합니다. 
- 이 외에도 많은 블록체인 프로젝트와 플랫폼이 Go를 사용하여 개발되고 있어, 
- 이 분야에서 Golang의 기여도는 상당히 높다고 할 수 있습니다. 
- Golang은 블록체인 기술의 발전과 함께 계속해서 중요한 역할을 할 것으로 예상됩니다.

> Golang의 블록체인 개발 적합성

	1. 성능과 효율성: 
		- Golang은 컴파일 언어로서 높은 실행 속도와 효율성을 제공합니다. 
		- 블록체인 네트워크에서는 대량의 데이터 처리와 빠른 트랜잭션 처리가 중요한데, Go는 이러한 요구 사항을 충족시킵니다.

	2. 병행 처리 능력: 
		- Go의 고루틴과 채널을 통한 병행 처리 기능은 동시에 다수의 트랜잭션을 처리하고, 블록체인 네트워크의 노드 간 통신을 효율적으로 관리하는 데 이상적입니다. 이는 특히 대규모 분산 시스템에서 중요합니다.

	3. 네트워킹 지원: 
		- Golang은 강력한 네트워킹 라이브러리를 제공합니다. 
		- 이는 P2P 네트워크 구축, 데이터 동기화, 노드 간 통신 등 블록체인 네트워크의 핵심 요소를 구현하는 데 필수적입니다.

> 주요 블록체인 프로젝트 및 플랫폼

	1. Ethereum의 Go 구현체 (Geth): 
		- Geth는 Ethereum 블록체인을 구현한 Go 언어 버전입니다. 
		- 이는 개발자들이 Ethereum 네트워크에 참여하고, 스마트 컨트랙트를 개발 및 배포할 수 있게 하는 핵심 도구 중 하나입니다. 
		- Geth의 고성능 및 병행 처리 기능은 블록체인 애플리케이션 개발에 매우 유용합니다.

	2. Hyperledger Fabric: 
		- Hyperledger Fabric은 기업용 블록체인 플랫폼으로, 고도의 보안성, 확장성, 그리고 모듈식 아키텍처를 제공합니다. 
		- Fabric의 일부 구성 요소는 Go로 개발되었으며, 이는 기업 환경에서 블록체인 솔루션을 개발하는 데 Go 언어가 얼마나 유용한지 보여줍니다.

	3. Cosmos SDK: 
		- Cosmos는 인터블록체인 커뮤니케이션을 위한 프레임워크를 제공하는 프로젝트로, 다양한 블록체인이 서로 소통하고 상호 작용할 수 있게 합니다. 
		- Cosmos SDK는 주로 Go로 작성되어 있으며, 개발자가 쉽게 자신만의 블록체인 애플리케이션을 구축할 수 있도록 지원합니다.


### 설치 및 사용법

- go 설치 후

- 버전확인
	go version

- 환경변수 확인
	go env 
	- GOPATH: JAVA의 WORKSPACE와 같다고 생각하면 된다.
	- GOROOT: 어디에서나 Go Command를 호출할 수 있다.

- 빌드명령어
	go build test.go
	- 빌드 결과로 exe 파일 생성됨
	

- 파일실행 명령어
	go run test.go
