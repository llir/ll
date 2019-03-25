# Generated parser for LLVM IR assembly

This project generates lexers and parsers for LLVM IR assembly from an [EBNF grammar](https://github.com/llir/grammar/blob/master/ll.tm) using [Textmapper](https://github.com/inspirer/textmapper).

## Installation

```bash
# Clone repo and submodules
git clone --recursive https://github.com/llir/ll

# Install textmapper
cd ll/tools/textmapper/tm-tool
./gen.sh

# Generate LLVM IR parser.
cd ../../..
make
```
