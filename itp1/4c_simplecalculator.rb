#! /usr/bin/ruby

$stdin.each_line do |line|
	a, op, b = line.split(/\s+/)
	break if op == "?"
	a = a.to_i
	b = b.to_i

	x = case op
	    when "+"
		    a + b
	    when "-"
		    a - b
	    when "*"
		    a * b
	    when "/"
		    a / b
	    end
	puts x
end

