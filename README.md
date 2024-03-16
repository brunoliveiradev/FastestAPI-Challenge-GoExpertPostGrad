# GoExpertPostGrad-Challenge
Repository for GoExpertPostGrad-challenge: A curated collection of technical challenges from the Postgraduate program, focused on advanced Go programming concepts. Here, you'll find exercises and solutions.

# Fastest API Response Challenge

[Leia isso em português](#desafio-da-resposta-mais-rápida-de-api)

## Challenge Description

In this challenge, you will use what we have learned about Multithreading and APIs to fetch the fastest result between two distinct APIs. The requests will be made simultaneously to the following APIs:

- `https://brasilapi.com.br/api/cep/v1/{cep}`
- `http://viacep.com.br/ws/{cep}/json/`

## Requirements

- Accept the response from the API that delivers the fastest result and discard the slower one.
- The result should be displayed on the command line with the address data, as well as which API sent it.
- Limit the response time to 1 second. Otherwise, a timeout error should be displayed.

## How to Run

1. Clone this repository.
2. Navigate to the project directory.
3. Run the command `go run main.go` inside the '`src/cmd`' folder to start the local server. This command compiles and executes your Go code that initiates the HTTP server, typically configured to listen on port `8080`.
4. **Make a Request**:
    - To request address information using a CEP, use a browser, `curl`, or any HTTP client, pointing to `http://localhost:8080/api/cep/{CEP_HERE}`, replacing `{CEP_HERE}` with the desired CEP.

   **Example using `curl`**:
   ```sh
   curl http://localhost:8080/api/cep/01001000
    ```
5. The result will be displayed in the terminal, showing the address data and which API sent it.

Please ensure your server is running if you're testing this locally, or replace `./api/cep/` with the appropriate command or URL prefix if you're making requests to a deployed API.

## Contributing

We welcome contributions! Please open an issue or submit a pull request for any improvements or additional features you would like to add.

---

# Desafio da Resposta Mais Rápida de API

## Descrição do Desafio

Neste desafio, você usará o que aprendemos sobre Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas. As requisições serão feitas simultaneamente para as seguintes APIs:

- `https://brasilapi.com.br/api/cep/v1/{cep}`
- `http://viacep.com.br/ws/{cep}/json/`

## Requisitos

- Acatar a resposta da API que entregar o resultado mais rápido e descartar a mais lenta.
- O resultado deve ser exibido no command line com os dados do endereço, bem como qual API o enviou.
- Limitar o tempo de resposta em 1 segundo. Caso contrário, deve ser exibido o erro de timeout.

## Como Executar

1. Clone este repositório.
2. Navegue até o diretório do projeto.
3. Execute o comando `go run main.go` dentro do diretorio '`src/cmd`'. Este comando compila e executa o seu código Go que inicia o servidor HTTP, tipicamente configurado para escutar na porta `8080`.
4. Para solicitar informações de endereço usando um CEP, utilize um navegador, curl ou qualquer cliente HTTP, apontando para http://localhost:8080/api/cep/{CEP_AQUI}, substituindo {CEP_AQUI} pelo CEP desejado.

   **Exemplo usando `curl`**:
   ```sh
   curl http://localhost:8080/api/cep/01001000
    ```
5. O resultado será exibido no terminal, mostrando os dados do endereço e qual API o enviou.

Certifique-se de que seu servidor esteja rodando se você estiver testando isso localmente, ou substitua `./api/cep/` pelo comando apropriado ou prefixo de URL se você estiver fazendo solicitações a uma API implantada.

## Contribuindo

Contribuições são bem-vindas! Por favor, abra uma issue ou envie um pull request para quaisquer melhorias ou funcionalidades adicionais que você gostaria de adicionar.
