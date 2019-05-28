#!/usr/bin/gnuplot
set loadpath "~/.gnuplot/"
load "./viridis.pal"
set key off
set title "3dstfs"
set xyplane 0
set xlabel "k1"
set ylabel "k2"
set zlabel "k3"
splot './dat/3stf.dat' using 1:2:3:($4<0.01 ? $4:1/0) with dots ls 6
#with points pointtype 1 pointsize 2 palette linewidth 1 (($4>0.1)&&($4<0.21) ? $4:1/0)
pause -1