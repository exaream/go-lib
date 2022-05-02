# CSV reader for Shift_JIS made of Go

## Install
```shell
$ git clone https://github.com/exaream/go-lib.git
$ cd go-lib/csvutil
$ curl https://www.post.japanpost.jp/zipcode/dl/kogaki/zip/47okinaw.zip -o testdata/47okinaw.zip
$ unzip testdata/47okinaw.zip -d testdata
```

## Run
```shell
$ cd go-lib/csvutil/cmd/csv
$ go run main.go
```

```shell
$ cd go-lib/csvutil/cmd/zip
$ go run main.go
```

## Directory Structure

```
csvutil
├── cmd
│   ├── csv
│   │   └── main.go
│   ├── csv2
│   │   ├── country.csv
│   │   └── main.go
│   └── zip
│       └── main.go
├── testdata
├── README.md
├── csv.go
├── go.mod
├── go.sum
├── gocsv.go
└── zip.go
```

## References
- https://zenn.dev/mattn/articles/fd545a14b0ffdf
- https://www.post.japanpost.jp/zipcode/dl/kogaki-zip.html
