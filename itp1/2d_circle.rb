#! /usr/bin/ruby

w, h, x, y, r = $stdin.gets.split(/\s+/).map(&:to_i)

if r <= x && x <= w-r && r <= y && y <= h-r
	puts "Yes"
else
	puts "No"
end
