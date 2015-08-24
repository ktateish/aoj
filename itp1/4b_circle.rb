#! /usr/bin/ruby

r = $stdin.gets.to_f

printf "%.6f %.6f\n", Math::PI*r*r, 2*Math::PI*r
