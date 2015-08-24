#! /usr/bin/ruby

$stdin.each_line do |s|
	x, y = s.split(/\s+/).map(&:to_i).sort
	break if x == 0 && y == 0
	printf "%d %d\n", x, y
end
