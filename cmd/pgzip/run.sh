#!/bin/sh

input=/usr/share/dict/words

cat "${input}" | cat                                  | wc -c
cat "${input}" | ENV_ENCODE_LEVEL=Store       ./pgzip | wc -c
cat "${input}" | ENV_ENCODE_LEVEL=Fast        ./pgzip | wc -c
cat "${input}" | ENV_ENCODE_LEVEL=Best        ./pgzip | wc -c
cat "${input}" | ENV_ENCODE_LEVEL=Default     ./pgzip | wc -c
cat "${input}" | ENV_ENCODE_LEVEL=            ./pgzip | wc -c
cat "${input}" | ENV_ENCODE_LEVEL=HuffmanOnly ./pgzip | wc -c
cat "${input}" | ENV_ENCODE_LEVEL=Constant    ./pgzip | wc -c
