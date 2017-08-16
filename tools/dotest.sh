#!/bin/sh

col_ng='\033[1;31m'
col_ok='\033[1;32m'
col_off='\033[0m'

c="$1"
e="$2"
p1=$(basename "$e" | sed -Ee 's/^([0-9]+).*/\1/')
p2=$(basename "$e" | sed -Ee 's/^([0-9]+)([^_]+)_.*/\2/' | tr 'a-z' 'A-Z')
p="${p1}_$p2"
i=0
n=$(aoj testcase -l ${c}_$p)
rc=0
printf "Checking %s for %s_%s\n" $e $c $p
while [ $i -lt $n ]
do
	printf "  case %d: " $i
	aoj check -c $i ${c}_$p $e
	rc=$((rc + $?))
	if [ $rc -ne 0 ]; then
		aoj testcase -i $i ${c}_$p > $e.$i.input
		aoj testcase -o $i ${c}_$p > $e.$i.expected
		./$e < $e.$i.input > $e.$i.received
		printf "${col_ng}NG${col_off}\n"
		git diff --quiet --no-index $e.$i.expected $e.$i.received
	else
		printf "${col_ok}OK${col_off}\n"
		rm -f $e.$i.input $e.$i.received $e.$i.expected
	fi
	i=$((i+1))
done

if [ $rc -eq 0 ]; then
	touch .$e.done
fi
exit $rc
