#! /usr/bin/ruby

S = $stdin.gets.to_i

m, s = S.divmod 60
h, m = m.divmod 60
printf "%d:%d:%d\n", h, m, s
