# Docker Compose para Stess test

Este projeto utiliza o Docker para rodar um script que faz um teste de estresse em uma URL específica, sendo passado via parametro:

```URL: Qual a URL será testada```

```REQUESTS: Quantidade de requests que devem ser feitas```

```CONCURRENCY: Quantas threads devem serem abertas```

## Como Rodar o Projeto

Siga os passos abaixo para rodar o ambiente com Docker:

### 2. Construir e Subir os Contêineres

Execute o comando abaixo para construir o container:

```docker build -t stress-test -f Dockerfile.dev```

Este comando irá:
- Construir a imagem da aplicação a partir do Dockerfile.

Execute o comando abaixo para iniciar o app:

```docker run -p 8080:8080 -e URL=http://google.com.br -e REQUESTS=100 -e CONCURRENCY=10```

Este comando iniciará o app passando os 3 parametros como variáveis de ambiente para dentro do container.

Caso nao seja passado uma URL para o app, o servico nao inicializará.
Os outros dois parametros são opcionais e caso nao sejam setados, terá o valor default de "1".

Ao final do script, será exibido um status report informando a quantidade de requests realizadas, 
e um agrupamento com seus respectivos status code.