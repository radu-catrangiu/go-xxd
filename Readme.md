# go-xxd

A simple `xxd` implementation written in Go.

Developed using the TDD approach by comparing the output of the system `xxd` with the output of `./bin/xxd`.

## Implemented features

 * `-c cols` Format `cols` octets per line. Default 16.
 * `-g bytes` Separate the output of every `bytes` bytes (two hex characters) by a whitespace. Default 2.
