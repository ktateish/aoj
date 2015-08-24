#! /usr/bin/ruby

a, b, c = $stdin.gets.split(/\s+/).map &:to_i

if a < b && b < c
	puts "Yes"
else
	puts "No"
end
