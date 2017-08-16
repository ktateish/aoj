#! /usr/bin/ruby

min = 1.0 / 0
max = -min
sum = 0
n = $stdin.gets.to_i
$stdin.gets.split(/\s+/).map(&:to_i).each do |x|
	min = x if x < min
	max = x if max < x
	sum += x
end
printf "%d %d %d\n", min, max, sum
