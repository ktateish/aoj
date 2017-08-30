#!/bin/sh

echo .gitignore > .gitignore
echo Makefile.am >> .gitignore

echo "SUBDIRS = c" > Makefile.am
sh ./updatec.sh
echo c >> .gitignore

echo "SUBDIRS += ruby" >> Makefile.am
sh ./updateruby.sh
echo ruby >> .gitignore

echo "SUBDIRS += go" >> Makefile.am
sh ./updatego.sh
echo go >> .gitignore
