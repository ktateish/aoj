#!/bin/sh

files=$(ls -1 *.c | sed -e 's/\.c//g')

exec > Makefile.am

echo bin_PROGRAMS = $files

echo .gitignore > .gitignore
echo Makefile.am >> .gitignore
for i in $files
do
	echo ${i}_SOURCES = ${i}.c
	#echo ${i}_LDADD = -lm
	echo $i >> .gitignore
done
