# GoExpertPostGrad-Challenge

Repository for GoExpertPostGrad-challenge: A curated collection of technical challenges from the Postgraduate program,
focused on advanced Go programming concepts. Here, you'll find exercises and solutions.

# Fastest API Response Challenge

[Leia isso em português](#desafio-da-resposta-mais-rápida-de-api)

## Challenge Description

In this challenge, you will use what we have learned about Multithreading and APIs to fetch the fastest result between
two distinct APIs. The requests will be made simultaneously to the following APIs:

- `https://brasilapi.com.br/api/cep/v1/{cep}`
- `http://viacep.com.br/ws/{cep}/json/`

## Requirements

- Accept the response from the API that delivers the fastest result and discard the slower one.
- The result should be displayed on the command line with the address data, as well as which API sent it.
- Limit the response time to 1 second. Otherwise, a timeout error should be displayed.

## How to Run

1. Clone this repository.
2. Navigate to the project directory.
3. Run `go mod tidy` to ensure all dependencies are properly installed.
4. Start the local server by running `go run ./cmd/api/main.go`. This will start the server on port `8080`.
5. **Making a Request**:
    - To request address information using a CEP, use a browser, `curl`, or any HTTP client, pointing
      to `http://localhost:8080/api/cep/{CEP_HERE}`, replacing `{CEP_HERE}` with the desired CEP.

   **Example using `curl`**:
   ```sh
   curl http://localhost:8080/api/cep/01001000
    ```
6. The result will be displayed in the terminal, showing the address data and which API sent it.
7. The field `source` will indicate which API sent the response.

## Accessing Swagger Documentation

To access the Swagger-generated API documentation, navigate to http://localhost:8080/swagger/ after starting the server.
This will provide an interactive UI where you can view and test the API endpoint.

## Contributing

We welcome contributions! Please open an issue or submit a pull request for any improvements or additional features you
would like to add.

---

# Desafio da Resposta Mais Rápida de API

## Descrição do Desafio

Neste desafio, você usará o que aprendemos sobre Multithreading e APIs para buscar o resultado mais rápido entre duas
APIs distintas. As requisições serão feitas simultaneamente para as seguintes APIs:

- `https://brasilapi.com.br/api/cep/v1/{cep}`
- `http://viacep.com.br/ws/{cep}/json/`

## Requisitos

- Acatar a resposta da API que entregar o resultado mais rápido e descartar a mais lenta.
- O resultado deve ser exibido no command line com os dados do endereço, bem como qual API o enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, deve ser exibido o erro de timeout.

## Como Executar

1. Clone este repositório.
2. Navegue até o diretório do projeto.
3. Execute `go mod tidy` para garantir que todas as dependências estejam devidamente instaladas.
4. Execute o comando `go run main.go` dentro do diretorio '`src/cmd/api`'. Este comando compila e executa o seu código
   Go que inicia o servidor HTTP, tipicamente configurado para escutar na porta `8080`.
5. Para solicitar informações de endereço usando um CEP, utilize um navegador, curl ou qualquer cliente HTTP, apontando
   para http://localhost:8080/api/cep/{CEP_AQUI}, substituindo {CEP_AQUI} pelo CEP desejado.

   **Exemplo usando `curl`**:
   ```sh
   curl http://localhost:8080/api/cep/01001000
    ```
6. O resultado será exibido no terminal, mostrando os dados do endereço e qual API o enviou.

## Acessando a Documentação Swagger

Para acessar a documentação da API gerada pelo Swagger, navegue até http://localhost:8080/swagger/index.html após
iniciar o servidor. Isso proporcionará uma UI interativa onde você pode visualizar e testar os endpoints da API.

## Contribuindo

Contribuições são bem-vindas! Por favor, abra uma issue ou envie um pull request para quaisquer melhorias ou
funcionalidades adicionais que você gostaria de adicionar.
