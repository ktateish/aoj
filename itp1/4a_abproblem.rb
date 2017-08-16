#! /usr/bin/ruby

a, b = $stdin.gets.split(/\s+/).map(&:to_i)

printf "%d %d %5f\n", a/b, a%b, 1.0 * a / b
