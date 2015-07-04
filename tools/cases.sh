#!/bin/sh

get_pid () {
	d=$(pwd)
	pset=$(basename $d | tr a-z A-Z)
	pname=$(echo $1 | sed -Ee 's/_.*//; s/(.)$/_\1/;' |tr a-z A-Z)
	echo ${pset}_$pname
}

which jq > /dev/null
if [ $? -ne 0 ]; then
	echo "$0: jq is needed" >&2
	exit 1
fi

if [ -z "$1" ]; then
	echo "Usage: $0 <problem name>.tset" >&2
	exit 1
fi

if [ ! -f "$1" ]; then
	echo "$0: not found: $1" >&2
	exit 1
fi

basefile=$(basename $1 .tset)
baseurl='http://analytic.u-aizu.ac.jp:8080/aoj/testcase.jsp'
pid=$(get_pid $1)
setfile=$1

availablep=$(jq .available "$setfile")

if [ -z "$availablep" -o $availablep -ne 1 ]; then
	echo "$0: test set is not available for $1" >&2
	exit 1
fi

echo -n "Fetching cases for $basefile "
i=0
istat=$(jq ".input[$i]" "$setfile")
ostat=$(jq ".output[$i]" "$setfile")
while [ "$istat" != "null" -a "$ostat" != "null" ]
do
	i=$((i + 1))
	caseurl="${baseurl}?id=${pid}&case=${i}"
	prefix="${basefile}.case${i}"
	echo -n .
	if [ ! -f "${prefix}.input" ]; then
		curl -fs "${caseurl}&type=in" > "${prefix}.input.tmp" \
			&& mv -f "${prefix}.input.tmp" "${prefix}.input"
	fi
	if [ ! -f "${prefix}.output" ]; then
		curl -fs "${caseurl}&type=out" > "${prefix}.output.tmp" \
			&& mv -f "${prefix}.output.tmp" "${prefix}.output"
	fi
	if [ ! -f "${basefile}.case${i}.input" -o ! -f "${basefile}.case${i}.output" ]; then
		echo "Cannot get files for case $i" >&2
		exit 1
	fi

	istat=$(jq ".input[$i]" "$setfile")
	ostat=$(jq ".output[$i]" "$setfile")
done
echo " done"
echo $((i)) > ${basefile}.casesfetchdone
