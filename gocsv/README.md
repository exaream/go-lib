# CSV reader for Shift_JIS made of Go

## Install
```shell
$ git clone https://github.com/exaream/go-snippet.git
$ cd go-snippet/gocsv
$ curl https://www.post.japanpost.jp/zipcode/dl/kogaki/zip/ken_all.zip -o testdata/ken_all.zip
$ unzip testdata/ken_all.zip -d testdata
```

```shell
$ cd go-snippet/gocsv/cmd/gocsv
$ go run main.go
```

## Directory Structure

```
go-snippet
├── README.md
└── gocsv
    ├── README.md
    ├── cmd
    │   └── gocsv
    │       └── main.go
    ├── csv.go
    ├── go.mod
    ├── go.sum
    └── testdata
```

## References
- https://zenn.dev/mattn/articles/fd545a14b0ffdf
- https://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html

## TODO
- [ ] github.com/gocarina/gocsv
- [ ] Tests
- [ ] Workspace

