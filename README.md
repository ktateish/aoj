My source code for AOJ
======================

Build
-----
```
% ./autogen.sh
% ./configure
% make
```


Note
----
vim-template.c is C template file for Vim.  Save this file as ~/.vim/tmpl/new.c
and put the following line to you .vimrc
```
autocmd BufNewFile *.c 0r ~/.vim/tmpl/new.c
```
