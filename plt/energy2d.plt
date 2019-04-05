#!/usr/bin/gnuplot
set loadpath "~/.gnuplot/"
load "./viridis.pal"
################
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
###########
set terminal postscript eps enhanced color font 'Helvetica,10'
set output './img/en2dm2m3.eps'
#set output "./img/heatmapm1m2.svg"
set view map
set size ratio 1
set xrange[1:50]
set yrange[1:50]
set xlabel "m2"
set ylabel "m3"
set object 1 rect from graph 0, graph 0 to graph 1, graph 1 back
set object 1 rect fc rgb "black" fillstyle solid 1.0

splot "./dat/enm1m2.dat" using 1:2:3 with points pointtype 5 pointsize 0.3 palette linewidth 1