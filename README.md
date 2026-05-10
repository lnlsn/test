# test

A CLI tool for converting strings between different case formats.

## Usage

```
./test [--upper | --camel | --snake | --capital] <word> [word ...]
```

## Flags

| Flag        | Description                        | Example output     |
|-------------|------------------------------------|--------------------|
| `--upper`   | Convert to UPPER CASE              | `HELLO WORLD`      |
| `--camel`   | Convert to camelCase               | `helloWorld`       |
| `--snake`   | Convert to snake_case              | `hello_world`      |
| `--capital` | Convert to Capital Case            | `Hello World`      |

If no flag is provided, the input is echoed back unchanged.

## Examples

```sh
./test --upper "hello world"
# HELLO WORLD

./test --camel "hello world"
# helloWorld

./test --snake "HelloWorld"
# hello_world

./test --capital "hello_world"
# Hello World
```

The converter understands camelCase, snake_case, hyphen-case, and space-separated input.

## Testing

Run all unit tests:

```sh
go test ./...
```

Run with verbose output to see each test case:

```sh
go test -v ./...
```

Run a specific test function:

```sh
go test -run TestToSnakeCase ./...
```

## Benchmarks

Run all benchmarks (unit tests are skipped):

```sh
go test -bench=. -benchmem ./...
```

Run a specific benchmark:

```sh
go test -bench=BenchmarkToSnakeCase -benchmem ./...
```

Control how long each benchmark runs (default is `1s`):

```sh
go test -bench=. -benchmem -benchtime=5s ./...
```

### Reading benchmark output

```
BenchmarkToSnakeCase/helloWorldFooBar-4   1240622   973.6 ns/op   232 B/op   9 allocs/op
```

| Column        | Meaning                                     |
|---------------|---------------------------------------------|
| `1240622`     | Number of iterations run                    |
| `973.6 ns/op` | Average time per operation                  |
| `232 B/op`    | Average heap bytes allocated per operation  |
| `9 allocs/op` | Average number of allocations per operation |

## Authors

- LnLsn <to@lnlsn.xyz>


## License

Just take a look at LICENSE file.
