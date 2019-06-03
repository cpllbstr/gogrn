#!/usr/bin/gnuplot
set loadpath "~/.gnuplot/"
load "./set1.pal"
################
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
################
out = system("echo $out")
fil = system("echo $file")
set terminal pdfcairo  enhanced color font 'Helvetica,10'
set output out
set xlabel "T" offset 0,1.25  ,graph 1
set multiplot layout 2,1 margins 0.05,0.95,.1,.9 spacing 1,0.075
#set xr[0:10]
set key right top
set ylabel "X" offset 2.5,0,0
plot fil using 1:2 title 'X1' w l ls 3 lw 2, fil using 1:3 title 'X2' w l ls 4 lw 2, fil using 1:4 title 'X3' w l ls 5 lw 2,  fil using 1:5 title 'avX' w l ls 1 lw 1,
set ylabel "V" offset 5,0,0
plot fil using 1:6 title 'V1' w l ls 3 lw 2, fil using 1:7 title 'V2' w l ls 4 lw 1, fil using 1:8 title 'V3' w l ls 5 lw 2, fil using 1:9 title 'avV' w l ls 1 lw 1,
unset multiplot
pause -1