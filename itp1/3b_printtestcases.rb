#! /usr/bin/ruby

i = 0
$stdin.each_line do |s|
	if s == "0\n"
		i = 0
		next
	end
	i += 1
	printf "Case %d: %s", i, s
end
