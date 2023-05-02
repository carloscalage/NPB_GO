#!/bin/bash

read -p "Qual kernel deseja executar? (EP, IS): " KERNEL

read -p "Em qual linguagem deseja executar? (GO, C): " LANG

#file_name='echo "print '${KERNEL}'.lower()" | python'
#echo $file_name

echo "Running $KERNEL - $CLASS in $LANG"

N=30

echo >> results/log.csv

for CLASS in S W A B C D E; do
    echo $CLASS
    if [ $LANG = GO ]; then
            for i in $(seq 1 $N); do
            echo $i
            start=$(date +%s%N)
            go run ${KERNEL}/${KERNEL}_serial.go $CLASS >> output.txt
            end=$(date +%s%N) 
            time=$((end-start))
            echo "${KERNEL},${CLASS},serial,${LANG},${time}" >> results/log.csv
            done
    fi

    if [ $LANG = C ]; then
            cd GMAP/NPB-SER
            make clean
            make ${KERNEL}
            for i in $(seq 1 $N); do
            echo $i
            start=$(date +%s%N)
            ./bin/is.${CLASS} >> output.txt
            end=$(date +%s%N) 
            time=$((end-start))
            echo "${KERNEL},${CLASS},serial,${LANG},${time}" >> results/log.csv
            done
    fi

done