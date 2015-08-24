#!/bin/sh

files=$(ls -1 *.c | sed -e 's/\.c//g')
rbfiles=$(ls -1 *.rb  2>/dev/null)

exec > Makefile.am

echo "CFLAGS += -Wall"
echo "LIBS += -lm"
echo bin_PROGRAMS = $files
if [ -n "$rbfiles" ]; then
	echo dist_bin_SCRIPTS = $rbfiles
	chmod +x $rbfiles
fi

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

%.rb.done: %.casesfetchdone %.rb
	@../tools/dotest.sh -ruby $<

aojclean:
	rm -f *.output *.input *.myout *.tset \
		*.tmp *.casesfetchdone *.done

aojcheck: $(patsubst %,%.done,$(bin_PROGRAMS)) $(patsubst %,%.done,$(dist_bin_SCRIPTS))

EOM
