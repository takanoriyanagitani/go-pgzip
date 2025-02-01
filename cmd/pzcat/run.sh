#!/bin/sh

sample=/usr/share/dict/words
input=./sample.d/sample.gz

geninput(){
	echo generating input...
	cat "${sample}" | ../pgzip/pgzip > "${input}"
}

test -f "${input}" || geninput

cat "${input}" |    zcat | wc -c
cat "${input}" | ./pzcat | wc -c
