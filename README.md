# Format converter

### Install

```sh
go install github.com/dsxack/formatconverter/cmd/formatconverter@latest
```

### Usage

```sh
# Basic usage
format-converter convert ./source.json ./destination.yaml
format-converter convert ./source.yaml ./destination.json

# Specify destination format
format-converter convert ./source.yaml ./destination.superjson -d json

# Convert all files in source directory into destination directory to json format
format-converter convert ./source_dir ./destination_dir -d json

# Serve convert by http server
format-converter serve
format-converter serve --addr ":8000"
curl "localhost:8080?dstFormat=yaml" --data @test.json
```