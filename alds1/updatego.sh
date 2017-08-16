#!/bin/sh

course=$(basename $(pwd) | tr 'a-z' 'A-Z')
files=$(ls -1 *.go | sed -e 's/\.go//g')
mkdir -p go
exec > go/Makefile.am

echo "CFLAGS += -Wall -std=c99"
echo "LIBS += -lm"
echo bin_PROGRAMS = $files

for i in $files
do
	echo "${i}_SOURCES = ${i}.go"
	echo
	echo "${i}: ../${i}.go"
	echo "	go build -o ${i} ../${i}.go"
	echo
done

cat <<EOM

SUFFIXES = .go

.%.done: %
	@../../tools/dotest.sh $course $<

EOM

cat <<'EOM'
aojclean:
	rm -f *.received *.input *.expected .*.done *.c

aojcheck: $(patsubst %,.%.done,$(bin_PROGRAMS))


EOM
