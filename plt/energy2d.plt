#!/usr/bin/gnuplot
set loadpath "~/.gnuplot/"
load "./viridis.pal"
################
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
###########
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output './img/k1k3.pdf'
set view map
set size ratio 1
set xrange[1:50]
set yrange[1:50]
set xlabel "k1"
set ylabel "k3"
set object 1 rect from graph 0, graph 0 to graph 1, graph 1 back
set object 1 rect fc rgb "black" fillstyle solid 1.0

splot "./dat/k1k3.dat" using 1:3:4 with points pointtype 5 pointsize 0.3 palette linewidth 1

