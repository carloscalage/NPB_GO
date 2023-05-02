#!/bin/bash

read -p "Qual kernel deseja executar? (EP, IS): " KERNEL

read -p "Em qual linguagem deseja executar? (GO, C): " LANG

#file_name='echo "print '${KERNEL}'.lower()" | python'
#echo $file_name

echo "Running $KERNEL - $CLASS in $LANG"

N=30

echo >> results/log.csv

if [ $LANG = C ]; then
    cd GMAP/NPB-OMP
    make clean
fi

if [ $LANG = GO ]; then
    cd IS
fi

for CLASS in S W A B C D E; do
    echo $CLASS
    if [ $LANG = GO ]; then
        for i in $(seq 1 $N); do
        echo $i
        go build
        start=$(date +%s%N)
        ./${KERNEL} ${CLASS}
        end=$(date +%s%N) 
        time=$((end-start))
        echo "${KERNEL},${CLASS},serial,${LANG},${time}" >> ../results/log.csv
        done
    fi

    if [ $LANG = C ]; then
        make ${KERNEL} CLASS=${CLASS}
        for i in $(seq 1 $N); do
        echo $i
        start=$(date +%s%N)
        ./bin/is.${CLASS}
        end=$(date +%s%N) 
        time=$((end-start))
        echo "${KERNEL},${CLASS},serial,${LANG},${time}" >> ../../results/log.csv
        done
    fi

done