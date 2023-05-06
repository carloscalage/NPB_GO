#!/bin/bash

read -p "Qual kernel deseja executar? (EP, IS): " KERNEL
read -p "Quantos cores gostaria de usar? " NUM_CORES

N=30

echo >> results/log.csv

# Compiling files

# Building GO
cd $KERNEL
go build

# Building C
cd ../GMAP/NPB-OMP
make clean

for CLASS in S W A B; do
    make ${KERNEL} CLASS=${CLASS}
done

cd ../../

# Running All The Classes 
for CLASS in S W A B; do
    echo $CLASS

    for i in $(seq 1 $N); do
        echo $i

	echo C
        start=$(date +%s%N)
        ./GMAP/NPB-OMP/bin/is.$CLASS
        end=$(date +%s%N)
        time=$((end-start))
        echo "${KERNEL},${CLASS},C,${time}" >> results/log.csv

        echo GO
        start=$(date +%s%N)
        ./${KERNEL}/${KERNEL} ${CLASS} ${NUM_CORES}
        end=$(date +%s%N)
        time=$((end-start))
        echo "${KERNEL},${CLASS},GO,${time}" >> results/log.csv
    done
done
