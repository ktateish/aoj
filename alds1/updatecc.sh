#!/bin/sh

course=$(basename $(pwd) | tr 'a-z' 'A-Z')
files=$(ls -1 *.cc | sed -e 's/\.cc$//g')
mkdir -p cc
exec > cc/Makefile.am

echo "CXXFLAGS += -Wall -std=c++14"
echo "LIBS += -lm"
echo bin_PROGRAMS = $files

for i in $files
do
	echo "${i}_SOURCES = ${i}.cc"
	echo
	echo "${i}.cc: ../${i}.cc"
	echo "	@rm -f ${i}.cc"
	echo "	@cat ../${i}.cc > ${i}.cc"
	echo "	@chmod 444 ${i}.cc"
	echo
done

cat <<EOM

.%.done: %
	@../../tools/dotest.sh $course $<

EOM

cat <<'EOM'
aojclean:
	rm -f *.received *.input *.expected .*.done *.cc

aojcheck: $(patsubst %,.%.done,$(bin_PROGRAMS))

EOM
