# Generated parser for LLVM IR assembly

This project generates lexers and parsers for LLVM IR assembly from an [EBNF grammar](https://github.com/llir/grammar/blob/master/ll.tm) using [Textmapper](https://github.com/inspirer/textmapper).

## Installation

```bash
# Clone repo and submodules.
git clone --recursive https://github.com/llir/ll

# Install textmapper.
cd ll/tools/textmapper/tm-tool
./gen.sh

# Generate LLVM IR parser.
cd ../../..
make
```

## License

The `llir/ll` project is dual-licensed to the [public domain](UNLICENSE) and under a [zero-clause BSD license](LICENSE). You may choose either license to govern your use of `llir/ll`.
