#!/usr/bin/gnuplot
set loadpath "~/.gnuplot/"
load "./set1.pal"
################
set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
################
set term svg enhanced background rgb 'white'
set output "./img/energyM[0].svg"
fil = "./dat/energy.dat"

plot fil using 1:3 title 'M1' w l ls 3 lw 2, 
reset