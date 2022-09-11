# Golang

## download

Go 버전 : 1.19.1

* linux

바이너리 파일 다운로드

```shell
curl -o ./go1.19.1.linux-amd64.tar.gz https://storage.googleapis.com/golang/go1.19.1.linux-amd64.tar.gz
```

압축 해제

```shell
sudo tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz
```

PATH 설정 `$HOME/.bash_profile`

```shell
PATH=$PATH:/usr/local/go/bin
```

* mac

m1 : <https://go.dev/dl/go1.19.1.darwin-arm64.pkg>

intel : <https://go.dev/dl/go1.19.1.darwin-amd64.pkg>

---
