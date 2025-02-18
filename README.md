# go-aplicacao-cli
## Instruções de Uso

Para usar a aplicação de linha de comando, siga os passos abaixo:

1. Clone o repositório:
  ```sh
  git clone https://github.com/seu-usuario/go-aplicacao-cli.git
  ```

2. Navegue até o diretório do projeto:
  ```sh
  cd go-aplicacao-cli
  ```

3. Compile a aplicação:
  ```sh
  go build
  ```

4. Execute um dos comandos da aplicação:
  ```sh
  go run main.go servidores --host amazon.com.br
  ```
  ```sh
  go run main.go ip --host amazon.com.br
  ```

5. Para ver mais opções disponíveis, use o comando de ajuda:
  ```sh
  go run main.go --help
  ```