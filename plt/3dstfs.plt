#!/usr/bin/gnuplot
set loadpath "~/.gnuplot/"
load "./viridis.pal"
set xyplane at 0
splot './dat/3stf.dat' using 1:2:3:4 with dots palette 
#with points pointtype 1 pointsize 2 palette linewidth 1 (($4>0.1)&&($4<0.21) ? $4:1/0)
pause -1