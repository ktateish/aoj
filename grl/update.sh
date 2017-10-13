#!/bin/sh

echo "SUBDIRS =" > Makefile.am
echo .gitignore > .gitignore
echo Makefile.am >> .gitignore

echo "SUBDIRS += go" >> Makefile.am
sh ./updatego.sh
echo go >> .gitignore
