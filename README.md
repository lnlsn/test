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
