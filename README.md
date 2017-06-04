# gouter
gouter is simple golang routing library

## Requirement
Go 1.8

## Installation
```shell
go get github.com/kpango/gouter
```

## Example
```go
	route := gouter.New().
		AddRoute("index", "/", http.MethodGet, handler.Index).
		AddRoute("protogen", "/proto", http.MethodPost, handler.ProtoGen).
		GetRouter()

	http.ListenAndServe("port", route)
```

## Contribution
1. Fork it ( https://github.com/kpango/gouter/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## Author
[kpango](https://github.com/kpango)

## LICENSE
gouter released under MIT license, refer [LICENSE](https://github.com/kpango/gouter/blob/master/LICENSE) file.
