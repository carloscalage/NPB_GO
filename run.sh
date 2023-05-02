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

read -p "Qual kernel deseja executar? (EP, IS)" KERNEL

read -p "Qual classe de problema deseja executar? (S, W, A, B, C, D, E)" CLASS

echo "Running $KERNEL in GO"

N = 30

for i in $(seq 1 $N); do
  echo $i
  start=$(date +%s%N/1000000)
  go run ${KERNEL}/${KERNEL}_serial.go $CLASS
  end=$(date +%s%N/1000000) 
  time=$((end-start))
  echo >> results/log.csv
  echo "${KERNEL},${CLASS},serial,GO,${time}" >> results/log.csv
done