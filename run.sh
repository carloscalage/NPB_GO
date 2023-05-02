# #CLASSES DE PROBLEMA:
# #S = 24
# #W = 25
# #A = 28
# #B = 30
# #C = 32
# #D = 36
# #E = 40

# #rodando o programa 20 vezes
# N=20

# sum1=0

# C=28
# printf "Running serial EP\n"
# for i in $(seq 1 $N); do
#     start=$(date +%s%N) #nanosegundos
#     go run ep.go $C
#     end=$(date +%s%N)    
#     sum1=$((end-start))
# done


# avg=$(echo "$sum1 / $N" | bc -l)
# ml=$(echo "$sum1 / 1000000" | bc -l)
# avgsec=$(echo "$ml / $N / 60" | bc -l)

# printf "time elapsed in nanoseconds: %0.0f, time elapsed in miliseconds: %0.1f\n" $avg $ml
# printf "average time in seconds: %0.1f\n" $avgsec

read -p "Qual kernel deseja executar? (EP, IS): " KERNEL

read -p "Em qual linguagem deseja executar? (GO, C): " LANG

file_name=`echo "print '$KERNEL'.lower()" | python`
echo $file_name

echo "Running $KERNEL - $CLASS in $LANG"

N=30

echo >> results/log.csv

for CLASS in S W A B C D E; do
    echo $CLASS
    if [ $LANG == "GO" ];
        then
            for i in $(seq 1 $N); do
            echo $i
            start=$(date +%s%N)
            go run ${KERNEL}/${KERNEL}_serial.go $CLASS >> output.txt
            end=$(date +%s%N) 
            time=$((end-start))
            echo "${KERNEL},${CLASS},serial,${LANG},${time}" >> results/log.csv
            done
    fi

    if [ $LANG == "C" ];
        then
            cd GMAP/NPB-SER
            make clean
            make ${KERNEL}
            for i in $(seq 1 $N); do
            echo $i
            start=$(date +%s%N)
            ./bin/${KERNEL}.${CLASS} >> output.txt
            end=$(date +%s%N) 
            time=$((end-start))
            echo "${KERNEL},${CLASS},serial,${LANG},${time}" >> results/log.csv
            done
    fi

done