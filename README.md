# Go json2yaml

I've been working with quite a bit of YAML lately. Bunch of configs were in JSON format and am finding myself needing to convert them. This is a super quick hack.

## Install

```
go get -u github.com/kalyan02/go-json2yaml
```

## Usage
```
go-json2yaml -input input.json > output.yaml
```

## Build

1. Checkout source
2. Run `dep ensure`
3. Run `go build github.com/kalyan02/go-json2yaml`

## License

BSD
