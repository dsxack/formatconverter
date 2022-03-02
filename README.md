# Format converter

### Install

```sh
go install github.com/dsxack/formatconverter/cmd/formatconverter@latest
```

### Usage

```sh
# Basic usage
formatconverter convert ./source.json ./destination.yaml
formatconverter convert ./source.yaml ./destination.json

# Specify destination format
formatconverter convert ./source.yaml ./destination.superjson -d json

# Convert all files in source directory into destination directory to json format
formatconverter convert ./source_dir ./destination_dir -d json

# Serve convert by http server
formatconverter serve
formatconverter serve --addr ":8000"
curl "localhost:8080?dstFormat=yaml" --data @test.json
```