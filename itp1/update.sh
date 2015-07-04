#!/bin/sh

files=$(ls -1 *.c | sed -e 's/\.c//g')

exec > Makefile.am

echo "CFLAGS += -Wall"
echo "LIBS += -lm"
echo bin_PROGRAMS = $files

echo .gitignore > .gitignore
echo Makefile.am >> .gitignore
echo '*.myout' >> .gitignore
echo '*.output' >> .gitignore
echo '*.input' >> .gitignore
echo '*.tset' >> .gitignore
echo '*.casesfetchdone' >> .gitignore
echo '*.done' >> .gitignore
echo '*.tmp' >> .gitignore
for i in $files
do
	echo ${i}_SOURCES = ${i}.c
	#echo ${i}_LDADD = -lm
	echo $i >> .gitignore
done

cat <<'EOM'

.PRECIOUS: %.tset %.casesfetchdone %.output %.input

%.tset: %.c
	@../tools/testset.sh $<

%.casesfetchdone: %.tset
	@../tools/cases.sh $<

%.done: %.casesfetchdone %
	@../tools/dotest.sh $<

aojclean:
	rm -f *.output *.input *.myout *.tset \
		*.tmp *.casesfetchdone *.done

aojcheck: $(patsubst %,%.done,$(bin_PROGRAMS))

EOM
