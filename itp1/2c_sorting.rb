#! /usr/bin/ruby

a, b, c = $stdin.gets.split(/\s+/).map(&:to_i).sort

printf "%d %d %d\n", a, b, c
