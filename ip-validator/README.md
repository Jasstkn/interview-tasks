# task

Write validator for IPv4 address.

## Implementations

1. Via `for` cycle + few validations with `if/else`
2. Regular expressions
3. Python

## Tests

Run

```bash
$ go test validator.go validator_test.go

ok      command-line-arguments  0.579s
```

or for Python

```bash
python3 -m unittest test_validator.py -v
```

## Benchmarks

Execute:

```bash
$ go test validator.go validator_test.go -bench . -benchmem

goos: darwin
goarch: amd64
cpu: VirtualApple @ 2.50GHz
BenchmarkIPValidator-8          1000000000               0.0000022 ns/op               0 B/op          0 allocs/op
BenchmarkIPValidatorRegexp-8    1000000000               0.0001655 ns/op               0 B/op          0 allocs/op
PASS
ok      command-line-arguments  0.189s
```
