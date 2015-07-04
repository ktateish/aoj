#!/bin/sh

if [ -z "$1" ]; then
	echo "Usage: $0 <problem name>.casesfetchdone" >&2
	exit 1
fi

if [ ! -f "$1" ]; then
	echo "$0: not found: $1" >&2
	exit 1
fi

prog=$(basename $1 .casesfetchdone)
n=$(cat $1)

col_ng='\033[1;31m'
col_ok='\033[1;32m'
col_de='\033[0m'
has_error=false
errors=""
echo -n "Checking cases for ${prog} "

i=1
while [ $i -le $n ]
do
	in="${prog}.case${i}.input"
	out="${prog}.case${i}.output"
	myout="${prog}.case${i}.myout"
	diffcmd="diff -q"
	if [ -f "${prog}.fuzzy" ]; then
    diffcmd="$(dirname $0)/difftool"
	fi

	./"$prog" < "$in" > "$myout"
	if [ $? -ne 0 -o ! -f "${myout}" ]; then
		echo -n "R"
		has_error=true
		errors="$errors $in"
		i=$((i+1))
		continue
	fi
	${diffcmd} "$out" "$myout"
	if [ $? -ne 0 ]; then
		echo -n "W"
		has_error=true
		errors="$errors $in"
		i=$((i+1))
		continue
	fi
	echo -n "."
	i=$((i+1))
done
if $has_error; then
	echo -e " ${col_ng}NG${col_de}"
	echo "Following cases got fault."
	for e in "$errors"
	do
		echo $e
	done
else
	echo -e " ${col_ok}OK${col_de}"
	echo ok > ${prog}.done
fi
