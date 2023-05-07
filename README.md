### Implementações NPB em Go

Grupo: Ariam Bartsch, Carlos Calage, João Belmonte.

Repositório para a cadeira de TECVII do professor Gerson Cavalheiro, onde a ideia era explorar diferentes linguagens de programação que fornecem ferramentas para implementação de paralelismo e concorrência em comparação com a linguagem de programação C. Nosso grupo decidiu utilizar a linguagem de programação Go.

Para a comparação, utilizamos os problemas propostos pela NASA chamados NPB. Aqui fizemos a implementação do kernel EP e IS, e mantivemos uma cópia do repositório do GMAP com a implementação dos mesmos problemas em CPP para fins de comparação.

#### Benchmark

Para rodar o benchmark nas duas linguagens de forma intercalada, basta rodar o script shell `run.sh` com o comando `sh run.sh`.

Na sequência, o script perguntará qual o kernel deverá roda e quantos núcleos deve usar. Após isso, ele irá compilar o programa em cada uma das linguagens e após isso, irá rodar 30 vezes cada um de forma intercalada e irá colocar os resultados no csv presente na pasta results.

Para nossas comparações, há o arquivo log.csv que possui o tempo que cada uma das implementações levou para concluir.

Também há um notebook python contendo o código que gera os plots baseado nesses resultados.

### Estrutura de Pastas

Para as nossas implementações em Go, temos as pastas EP e IS na raíz e para as implementações de CPP dentro da pasta [GMAP](https://github.com/GMAP/NPB-CPP) e estamos utilizando o NPB-OMP para comparação.