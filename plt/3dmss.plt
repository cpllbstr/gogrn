#!/usr/bin/gnuplot
set loadpath "~/.gnuplot/"
load "./viridis.pal"
#set terminal pdfcairo  enhanced color font 'Helvetica,10'
#set output './img/3dmss.pdf'
set key off
set title "3dmss"
set grid
set xlabel "m1"
set ylabel "m2"
set zlabel "m3"
set xyplane 0 
splot './dat/3mssnew.dat' using  1:2:3:4 with dots palette
# ($4<0.01 ? $4:1/0) with points pointtype 1 pointsize 2 palette linewidth 1 (($4>0.1)&&($4<0.21) ? $4:1/0) || (($4<0.3) ? $4:1/0)
pause -1