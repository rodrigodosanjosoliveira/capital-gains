# Capital Gains Calculator

## Descrição
O projeto calcula os impostos sobre ganhos de capital com base em operações de compra e venda de ativos. A solução foi desenvolvida em Go, com foco em modularidade, simplicidade e extensibilidade. Ele suporta operações de compra e venda, respeitando limites de isenção, perdas acumuladas e outras regras fiscais.

## Decisões de Projeto
### 1. **Separação de responsabilidades:**
- A solução foi dividida em pacotes dentro da pasta internal/ para organizar responsabilidades:
  - `calculator/`: Lógica de cálculo de impostos.
  - `models/`: Definição das structs que representam operações e impostos.
  - `io/`: Manipulação de entrada e saída, como leitura do JSON de stdin e escrita no stdout.
### 2. **Transparência referencial:**
- A função `CalculateCapitalGains` foi projetada para ser pura, retornando sempre os mesmos resultados para os mesmos inputs, sem efeitos colaterais.
### 3. **Testes abrangentes:**
- Testes de unidade verificam a funcionalidade de cada componente.
- Testes de integração validam a execução de ponta a ponta, cobrindo desde a entrada até a saída.
### 4. **Transparência referencial:**
- **Pipeline de dados:** A entrada é lida, transformada em structs Go, processada pelo motor de cálculo e serializada para saída.

## Justificativa para o Uso de Bibliotecas
O projeto utiliza apenas as bibliotecas padrão do Go para cumprir os requisitos:
1. `encoding/json`: Para serialização e desserialização de JSON.
   - Escolhida por ser nativa, eficiente e amplamente usada em projetos Go.
2. `os` e `bufio`: Para manipulação de entrada e saída padrão.
   - Adequadas para o fluxo baseado em stdin e stdout.
3. `testing`: Para testes unitários e de integração.
   - Parte da biblioteca padrão, evitando a necessidade de dependências externas.

## Instruções para Compilar e Executar o Projeto

### Pré-requisitos
- Go 1.20 ou superior

### Compile o Projeto
1. **No diretório raiz do projeto, execute:**
```bash
go build -o capital-gains cmd/app/main.go
```
2. **Execute o programa:**
```bash
echo '[{"operation":"buy", "unit-cost":10.00, "quantity":100},{"operation":"sell", "unit-cost":15.00, "quantity":50},{"operation":"sell", "unit-cost":15.00, "quantity":50}]' | ./capital-gains
```
ou utilizando um arquivo de entrada:
```bash
./capital-gains < input.txt
```

### Execute os Testes
```bash
go test ./... -v
```