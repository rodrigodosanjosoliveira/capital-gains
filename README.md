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

---

## Como executar

### 1. Localmente com Go

#### Dependências:
- Go 1.21 ou superior
- `go.mod` e `go.sum` com dependências resolvidas (`go mod tidy`)

#### Rodando:

**Com JSON via argumento:**

```bash
go run ./cmd/app '[{"operation":"buy", "unit-cost":10, "quantity":10000},{"operation":"sell", "unit-cost":25, "quantity":1000}]'
```

**Ou com JSON via entrada padrão (stdin):**

```bash
echo '[{"operation":"buy", "unit-cost":10, "quantity":10000},{"operation":"sell", "unit-cost":25, "quantity":1000}]' \
| go run ./cmd/app
```

---

### 2. Usando Docker (imagem pública no Docker Hub)

#### Executar com argumento:

```bash
docker run --rm devbychoice/ganho-de-capital \
'[{"operation":"buy", "unit-cost":10, "quantity":10000},{"operation":"sell", "unit-cost":25, "quantity":1000}]'
```

#### Executar com entrada via stdin:

```bash
echo '[{"operation":"buy", "unit-cost":10, "quantity":10000},{"operation":"sell", "unit-cost":25, "quantity":1000}]' \
| docker run -i --rm devbychoice/ganho-de-capital
```

---

## Exemplo de entrada

```json
[
  { "operation": "buy",  "unit-cost": 10.00, "quantity": 10000 },
  { "operation": "sell", "unit-cost": 25.00, "quantity": 1000 }
]
```

### Saída esperada

```json
[
  { "tax": 0.0 },
  { "tax": 1000.0 }
]
```

