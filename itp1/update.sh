#!/bin/sh

echo .gitignore > .gitignore
echo Makefile.am >> .gitignore

echo "SUBDIRS = c" > Makefile.am
sh ./updatec.sh
echo c >> .gitignore

echo "SUBDIRS += ruby" >> Makefile.am
sh ./updateruby.sh
echo ruby >> .gitignore

cat >> Makefile.am <<'EOM'

aojcheck aojclean:
	@for i in $(SUBDIRS) ;\
	 do \
		  ( cd $$i ; make $@ ; ) ;\
	 done
EOM
