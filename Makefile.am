PROB_DIRS= itp1 alds1
TOOL_DIRS= tools
SUBDIRS= $(TOOL_DIRS) $(PROB_DIRS)

aojcheck aojclean:
	@for i in $(PROB_DIRS) ;\
	 do \
		  ( cd $$i ; make $@ ; ) ;\
	 done
