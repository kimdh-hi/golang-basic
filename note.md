## go basic

vscode 실행: `ctrl+F5`
터미널 싫행: `go run file.go`

### go 실행 및 빌드
**vscode 실행**: `ctrl+F5`

**CLI  실행**
```bash
go run file.go

go run -x file.go # 필요 패키지, 의존성 등 출력

go build # 실행파일(exe) 생성

go build -o myapp.exe # 실행파일의 이름을 지정하여 빌드
```

**다른 플랫폼(운영체제)를 타겟으로 빌드**
```bash
# windows application
GOOS=windows GOARCH=amd64 go build -o windowsapp.exe 

# linux application
GOOS=linux GOARCH=amd64 go build -o linuxapp

# macos application
GOOS=darwin GOARCH=amd64 go build -o macapp
```

### 소스코드 포맷팅 gofmt

```bash
go fmt # 현재 디렉토리의 go 파일에 저장시 자동 포멧팅 적용

gofmt -w file.go

gofmt -w -l 디렉토리명
````


***
참고

go.mod file not found 에러
```bash
go env -w GO111MODULE=auto
```