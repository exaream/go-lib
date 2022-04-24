# Go Sample
Go's sample using workspace mode.

## Install
```shell
$ git clone https://github.com/exaream/go-snippet.git
$ cd go-snippet/sample
$ curl https://www.post.japanpost.jp/zipcode/dl/kogaki/zip/47okinaw.zip -o testdata/47okinaw.zip
$ unzip testdata/47okinaw.zip -d testdata
```

## Run
```shell
$ cd go-snippet/sample/cmd/sheet
$ go run main.go
```

```
sample
├── README.md
├── cmd
│   └── sample
│       └── main.go
├── go.mod
└── testdata
```

## References
- https://zenn.dev/kimuson13/articles/go-workspace-mode-impressions

## TODO
- [x] Workspace
- [ ] Comment
