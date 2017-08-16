#!/bin/sh

course=$(basename $(pwd) | tr 'a-z' 'A-Z')
files=$(ls -1 *.rb | sed -e 's/\.rb$//g')
mkdir -p ruby
exec > ruby/Makefile.am

echo bin_PROGRAMS = $files

for i in $files
do
	echo "${i}_SOURCES = ${i}.rb"
	echo
	echo "${i}: ../${i}.rb"
	echo "	@rm -f ${i}"
	echo "	@cat ../${i}.rb > ${i}"
	echo "	@chmod 555 ${i}"
	echo
done

cat <<EOM

SUFFIXES = .rb

.%.done: %
	@../../tools/dotest.sh $course $<

EOM

cat <<'EOM'
aojclean:
	rm -f *.received *.input *.expected .*.done *.c

aojcheck: $(patsubst %,.%.done,$(bin_PROGRAMS))


EOM
