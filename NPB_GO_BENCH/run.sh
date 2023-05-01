#!/bin/bash


#CLASSES DE PROBLEMA:
#S = 24
#W = 25
#A = 28
#B = 30
#C = 32
#D = 36
#E = 40


#rodando o programa 20 vezes
N=30

sum1=0
echo "program name | class size | thread amount | time" >> timerA.txt


CLASS=28
CORES=8

printf "Running parallel, same threads as cores EP\n" 
printf "---------------------------------------------------------------\n" >> timer.txt
for i in $(seq 1 $N); do
    start=$(date +%s%N) #nanosegundos
    ep_parrallel/ep_parrallel $CLASS $CORES
    end=$(date +%s%N)    
    sum1=$(((end-start)/1000000))
    printf "go_par.go $CLASS $CORES $sum1 \n" >> timerA.txt
done


avg=$(echo "$sum1 / $N" | bc -l)
ml=$(echo "$sum1 / 1000000" | bc -l)
avgsec=$(echo "$ml / $N / 60" | bc -l)

printf "time elapsed in nanoseconds: %0.0f, time elapsed in miliseconds: %0.1f\n" $avg $ml
printf "average time in seconds: %0.1f\n" $avgsec

