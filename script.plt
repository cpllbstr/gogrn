set style line 1 lt 1 lw 1 pt 1 linecolor rgb "red"
set style line 2 lt 1 lw 1 pt 1 linecolor rgb "green"
set style line 3 lt 1 lw 1 pt 1 linecolor rgb "blue"

set grid xtics lc rgb '#555555' lw 1 lt 0
set grid ytics lc rgb '#555555' lw 1 lt 0
set xrange[0:10]
plot "./output.txt" using 1:2 title 'X1' w l ls 1, "./output.txt" using 1:3 title 'X2' w l ls 2, "./output.txt" using 1:4 title 'X3' w l ls 3, 
pause -1