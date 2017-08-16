#!/bin/sh

echo .gitignore > .gitignore
echo Makefile.am >> .gitignore

echo "SUBDIRS = c" > Makefile.am
sh ./updatec.sh
echo c >> .gitignore

cat >> Makefile.am <<'EOM'

aojcheck aojclean:
	@for i in $(SUBDIRS) ;\
	 do \
		  ( cd $$i ; make $@ ; ) ;\
	 done
EOM
