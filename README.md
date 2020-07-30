# sandbox

Go install

* 참조  
https://golang.org/doc/install  
&nbsp;
* 인스톨러 다운로드
https://golang.org/dl/

   - 위 사이트에서 본인의 시스템에 맞는 인스톨러를 다운로드하여 설치 실행

   - 윈도우의 경우, 설치 기본경로 `C:¥Go¥`로 지정

   - 인스톨 완료후 환경변수 설정 되었는지 확인

   - 아래 커맨드 입력하여 Go설치여부 확인

``` bash
> go version
```
&nbsp;
* GOPATH지정 

   - 시스템변수 설정에서 지정
     * 윈도우 : GOPATH <- `C:¥workspace¥go`
     * Mac : `export GOPATH=~/workspace/go`


* 폴더구성
``` cmd
C:¥worksapce
└─go
    ├─bin
    └─src
```
* git clone
```
git clone https://github.com/ijnk2553/sandbox.git
```
* module download
&nbsp;
```
cd C:¥worksapce¥go¥src¥sandbox

go mod download 
```

* module download
&nbsp;
```
cd C:¥worksapce¥go¥src¥sandbox

go mod download 
```

* hello world 실행
&nbsp;
```
cd hello

go run ./hello.go 

> hello world
```

* launch.json 설정

```
{
    // IntelliSense를 사용하여 가능한 특성에 대해 알아보세요.
    // 기존 특성에 대한 설명을 보려면 가리킵니다.
    // 자세한 내용을 보려면 https://go.microsoft.com/fwlink/?linkid=830387을(를) 방문하세요.
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {},
            "args": []
        }
    ]
}
```
