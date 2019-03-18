# gc-test

### Para rodar o projeto basta:

1. ```git clone https://github.com/lucassscaravelli/gc-test.git```
2. ```cd gc-test```
3. ```docker-compose up```
4. Acessar a url ```http://localhost:8080/static```

### Na página WEB você pode tentar:

* Criar um novo torneio (pequeno formulário na parte superior do site)
* Obter a tabela de um torneio
* Processar a fase de grupos clicando no botão que contém a instrução (simula todas as partidas da fase)
* Visualizar a tabela de grupos e suas partidas (clicando no botão no card do grupo)
* Avançar a fase do playoff clicando no botão que contém a instrução até a partida final, e visualizar os resultados

### Observações:

* Os times são gerados aleatoriamente (na criação do primeiro torneio)

### Arquitetura

* Banco de dados: Postgres
* Backend: Golang
* Frontend: ReactJS

* A página WEB é hosteada pelo backend no endpoint ```/static/```
