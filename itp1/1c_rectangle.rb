#! /usr/bin/ruby

a, b = $stdin.gets.split(/\s+/).map do |x|
	x.to_i
end

printf "%d %d\n", a*b, 2*(a+b)
