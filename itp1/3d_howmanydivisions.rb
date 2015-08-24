#! /usr/bin/ruby

a, b, c = $stdin.gets.split(/\s+/).map &:to_i

cnt = 0
(a..b).each do |i|
	cnt += 1 if c % i == 0
end
puts cnt
