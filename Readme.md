# go-xxd

A simple `xxd` implementation written in Go.

Developed using the TDD approach by comparing the output of the system `xxd` with the output of `./bin/xxd`.

## Implemented features

 * `-c cols` Format `cols` octets per line. Default 16.
 * `-g bytes` Separate the output of every `bytes` bytes (two hex characters) by a whitespace. Default 2.
 * `-l len` Stop after printing `len` octets. Defaults to -1 (Inf).  


## Example output

`./bin/xxd test/data/very_small.txt`
```
00000000: 5468 6973 2069 7320 616e 2065 7861 6d70  This is an examp
00000010: 6c65 2066 696c 652e 0a41 6c6c 2074 6865  le file..All the
00000020: 2063 6f6e 7465 6e74 7320 6172 6520 7265   contents are re
00000030: 7072 6573 656e 7465 6420 626f 7468 2061  presented both a
00000040: 7320 6865 7820 6461 7461 2061 6e64 2061  s hex data and a
00000050: 7363 6969 2e0a 416c 6c20 6368 6172 6163  scii..All charac
00000060: 7465 7273 2074 6861 7420 776f 756c 6420  ters that would 
00000070: 6d65 7373 2077 6974 6820 7468 6520 6f75  mess with the ou
00000080: 7470 7574 2066 6f72 6d61 7420 6c69 6b65  tput format like
00000090: 2074 6865 206e 6577 6c69 6e65 2063 6861   the newline cha
000000a0: 7261 6374 6572 2061 7265 2072 6570 6c61  racter are repla
000000b0: 6365 6420 7769 7468 2061 2070 6572 696f  ced with a perio
000000c0: 642e 0a42 696e 6172 7920 6368 6172 7320  d..Binary chars 
000000d0: 6172 6520 7368 6f77 6e20 6173 2070 6572  are shown as per
000000e0: 696f 6473 2074 6f6f 2e                   iods too.
```
