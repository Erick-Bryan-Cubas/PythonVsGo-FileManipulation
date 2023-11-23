# PythonVsGo-FileManipulation

## Comparação entre Python e Go para Operações em Arquivos Excel

Este repositório contém scripts e resultados de uma comparação entre as linguagens Python e Go em relação a operações em arquivos Excel. O objetivo é avaliar o desempenho e a eficiência de ambas as linguagens em tarefas comuns de processamento de planilhas.

## Resultados

Os resultados da comparação foram obtidos em uma máquina com a seguinte configuração:

- Modelo do Sistema: Acer Nitro AN515-44
- Processador: AMD64 Family 23 Model 96 ~2900 Mhz
- Memória Física Total: 7.549 MB

Os resultados incluem:

- Tempo decorrido em segundos para cada operação em ambas as linguagens.
- Uso da CPU antes e depois de cada operação em ambas as linguagens.
- Uso de memória antes e depois de cada operação em ambas as linguagens.

Abaixo estão os resultados de desempenho para cada operação:

### Concatenação de Arquivos Excel

- Go:
  - Tempo decorrido: 26.96 segundos
  - Uso de CPU Antes: 0.67%, Depois: 3.86%
  - Uso de Memória Antes: 83.00%, Depois: 86.00%

- Python:
  - Tempo decorrido: 746.41 segundos
  - Uso de CPU Antes: 2.3%, Depois: 3.4%
  - Uso de Memória Antes: 82.7%, Depois: 81.4%

### Soma Cumulativa de Números em Arquivos Excel

- Go:
  - Tempo decorrido: 24.92 segundos
  - Uso de CPU Antes: 0%, Depois: 4.34%
  - Uso de Memória Antes: 6241 MB, Depois: 6311 MB

- Python:
  - Tempo decorrido: 250.90 segundos
  - Uso de CPU Antes: 0.0%, Depois: 5.5%
  - Uso de Memória Antes: 87.8%, Depois: 80.2%

### Leitura de Arquivos Excel

- Go:
  - Tempo decorrido: 13.63 segundos
  - Uso de CPU Antes: 0%, Depois: 7.30%
  - Uso de Memória Antes: 6260 MB, Depois: 6376 MB

- Python:
  - Tempo decorrido: 7.61 segundos
  - Uso de CPU Antes: 3.8%, Depois: 5.3%
  - Uso de Memória Antes: 80.3%, Depois: 81.3%

### Contagem de Linhas Únicas em Arquivos Excel

- Go:
  - Tempo decorrido: 24.93 segundos
  - Uso de CPU Antes: 0%, Depois: 4.55%
  - Uso de Memória Antes: 86%, Depois: 83%

- Python:
  - Tempo decorrido: 259.97 segundos
  - Uso de CPU Antes: 3.8%, Depois: 5.6%
  - Uso de Memória Antes: 81.3%, Depois: 81.6%


## Gráfico de Colunas

## Tempo Decorrido

A seguir, apresentamos um gráfico de colunas que compara o tempo decorrido em segundos para cada operação em Python e Go:

![Gráfico de Tempo Decorrido - Python vs Go](https://github.com/Erick-Bryan-Cubas/PythonVsGo-FileManipulation/blob/main/images/tempo_decorrido_python_vs_go.png?raw=true)

### Uso de CPU

Gráfico comparativo do uso de CPU antes e depois das operações para Python e Go:

![Gráfico de Uso de CPU - Python vs Go](https://github.com/Erick-Bryan-Cubas/PythonVsGo-FileManipulation/blob/main/images/cpu_uso_python_vs_go.png?raw=true)

### Uso de Memória

Gráfico comparativo do uso de memória para Python e Go:

![Gráfico de Uso de Memória - Python vs Go](https://github.com/Erick-Bryan-Cubas/PythonVsGo-FileManipulation/blob/main/images/memoria_uso_python_vs_go.png?raw=true)

## Como Executar

Para executar os scripts e realizar a comparação por conta própria, siga as instruções no arquivo de cada script.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir problemas, enviar solicitações de pull ou melhorar a documentação.

## Licença

Este projeto está licenciado sob a Licença MIT - consulte o arquivo [LICENSE](LICENSE) para obter detalhes.

---

**Nota:** Os resultados são baseados nas condições e configurações específicas em que os testes foram executados. Os resultados podem variar dependendo do ambiente e dos recursos do sistema.
