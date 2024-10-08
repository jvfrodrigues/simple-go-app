# Aplicação de Roteamento de Viagem

## Como executar a aplicação

1. Certifique-se de ter o Go instalado em sua máquina (versão 1.22 ou superior).
2. Clone o repositório:
   ```
   git clone https://github.com/jvfrodrigues/simple-go-app
   cd simple-go-app
   ```
3. Execute a aplicação:
   - Para o modo CLI:
     ```
     go run main.go cli input-routes.csv
     ```
   - Para o modo API:
     ```
     go run main.go api input-routes.csv
     ```

## Interfaces da aplicação

### CLI

A interface CLI permite que você consulte a melhor rota entre dois pontos.

Exemplo de uso:

```
please enter the route: GRU-CDG
best route: GRU - BRC - SCL - ORL - CDG > $40
```

### API REST

A API REST oferece dois endpoints:

1. Consulta de melhor rota:

   - Método: GET
   - URL: `/route?from=ORIGEM&to=DESTINO`
   - Exemplo: `GET /route?from=GRU&to=CDG`

2. Registro de nova rota:
   - Método: POST
   - URL: `/route/add`
   - Corpo da requisição (JSON):
     ```json
     {
       "from": "GRU",
       "to": "CDG",
       "cost": 75
     }
     ```

## Estrutura dos arquivos/pacotes

```
.
├── cmd
│   ├── api.go
│   └── cli.go
├── internal
│   ├── model
│   │   └── route.go
│   │   └── graph.go
│   ├── infra
│   │   └── csv.go
│   └── service
│       ├── graph.go
│       └── api.go
├── main.go
└── README.md
```

- `cmd`: Contém os pontos de entrada para diferentes modos de execução (CLI e API).
- `internal`: Contém pacotes internos da aplicação.
- `model`: Define as estruturas de dados usadas na aplicação.
- `infra`: Lida com operações de infraestrutura, como leitura/escrita de CSV.
- `service`: Contém a lógica de negócio principal da aplicação.

## Decisões de design adotadas

1. **Separação de responsabilidades**: A aplicação foi dividida em pacotes distintos para melhorar a organização e manutenibilidade do código.

2. **Injeção de dependências**: Os serviços são criados com suas dependências injetadas, facilitando testes e flexibilidade.

3. **Uso de estruturas de dados eficientes**: Utilizamos um grafo para representar as rotas, permitindo um algoritmo de busca eficiente.
