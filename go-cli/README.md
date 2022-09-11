# Go Sample CLI

Golang `cobra` 모듈을 이용한 CLI 만들기

---

## Preparation

작업 디렉토리 준비

```shell
mkdir <sample-cli>
cd <sample-cli>
```

go mod 준비

* 모통 모듈명은 `github.com/<yourname>/<application>`
  
```shell
go mod init <sample-cli>
```

## Install

`cobra` 라이브러리 인스톨

```shell
go get -u github.com/spf13/cobra@latest
```

라이브러리 import

```shell
import "github.com/spf13/cobra"
```

`cobra-cli` 다운로드

* `cobra` 애플리케이션 부트스트랩
  
```shell
go install github.com/spf13/cobra-cli@latest
```

`cobra-cli` alias 설정

```shell
alias cobra='$GOPATH/bin/cobra-cli'
```

---

## Run

cobra 프로젝트 생성

```shell
cobra-cli init
```

## Download

```shell
/usr/local/bin/<build-application>
```

## Build

바이너리 빌드

```shell
# building the program for intel macs
GOOS=darwin GOARCH=amd64 go build -o <application-name> main.go 
# building the program for M1 macs
GOOS=darwin GOARCH=arm64 go build -o <application-name> main.go 
# building the program for 64 bits amd/intel linux
GOOS=linux GOARCH=amd64 go build -o <application-name> main.go
```

바이너리 압축

```shell
tar -xzvf <application-name>.<GOOS-GOARCH>.tar.gz <application-name.exec>
```

압축 풀기

```shell
sudo tar -C /usr/local/bin -xzvf <application-name>.<GOOS-GOARCH>.tar.gz
```
