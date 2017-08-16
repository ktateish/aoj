#!/bin/sh

course=$(basename $(pwd) | tr 'a-z' 'A-Z')
files=$(ls -1 *.c | sed -e 's/\.c//g')
mkdir -p c
exec > c/Makefile.am

echo "CFLAGS += -Wall -std=c99"
echo "LIBS += -lm"
echo bin_PROGRAMS = $files

for i in $files
do
	echo "${i}_SOURCES = ${i}.c"
	echo
	echo "${i}.c: ../${i}.c"
	echo "	@rm -f ${i}.c"
	echo "	@cat ../${i}.c > ${i}.c"
	echo "	@chmod 444 ${i}.c"
	echo
done

cat <<EOM

.%.done: %
	@../../tools/dotest.sh $course $<

EOM

cat <<'EOM'
aojclean:
	rm -f *.received *.input *.expected .*.done *.c

aojcheck: $(patsubst %,.%.done,$(bin_PROGRAMS))

EOM
