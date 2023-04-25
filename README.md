To run:

``go run is_nopointer.go``

pasta is_thrash: possui tentativas frustadas de implementar o IS com buckets e com coisas de ponteiros. Tem erro de out of bounds em vetor.

arquivo is_nopointers.go:
NÃO usa buckets e também não tem muito paralelismo. MAS executa até o final.
Infelizmente, a verificação não funciona. Ele não consegue ordenar a segunda metade do array,
