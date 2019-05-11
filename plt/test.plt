#!/usr/bin/gnuplot
set loadpath "~/.gnuplot/"
load "./viridis.pal"
set xrange [0:43]
splot './dat/three.dat' using 1:2:3:(($4>0.6)&&($4<0.61) ? $4:1/0) with dots
#with points pointtype 1 pointsize 2 palette linewidth 1
pause -1