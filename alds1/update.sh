#!/bin/sh

echo .gitignore > .gitignore
echo Makefile.am >> .gitignore

echo "SUBDIRS = c" > Makefile.am
sh ./updatec.sh
echo c >> .gitignore

echo "SUBDIRS = cc" > Makefile.am
sh ./updatecc.sh
echo cc >> .gitignore

echo "SUBDIRS += go" >> Makefile.am
sh ./updatego.sh
echo go >> .gitignore

cat >> Makefile.am <<'EOM'

aojcheck aojclean:
	@for i in $(SUBDIRS) ;\
	 do \
		  ( cd $$i ; make $@ ; ) ;\
	 done
EOM
