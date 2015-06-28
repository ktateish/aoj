#!/bin/sh

get_pid () {
	d=$(pwd)
	pset=$(basename $d | tr a-z A-Z)
	pname=$(echo $1 | sed -Ee 's/_.*//; s/(.)$/_\1/;' |tr a-z A-Z)
	echo ${pset}_$pname
}

baseurl='http://analytic.u-aizu.ac.jp:8080/aoj/testcase_header.jsp?id='
pid=$(get_pid $1)
ofile="$(basename $1 .c).tset"

echo "Fetching test cases info. for ${pid}"
curl -s "${baseurl}${pid}" > "$ofile"
