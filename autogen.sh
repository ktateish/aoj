#!/bin/sh

find * -type f -name update.sh | while read i
do
	dir=$(dirname $i)
	(cd $dir ; ./update.sh)
done

autoreconf -if
