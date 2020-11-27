# PSNeoway
Serviço desenvolvido para o processo seletivo da Neoway.

Autor: Felipe Pereira de Souza Duarte




Como executar:

Rode o seguinte comando para preparar o ambiente (deploy container e configurar banco):

```sh make setup-env```

Execute o programa com o seguinte comando:

```sh make run DB_URI="postgres://service:service@localhost:5432/postgres?sslmode=disable" FILE="arquivoTeste/base_teste.txt"```

A variavel "DB_URI" deve conter as informações necessárias para a conexão com o banco. Altere caso necessário para encaixar com a sua configuração.

A variavel "FILE" deve apontar para o arquivo a ser importado no banco.

*Para simular um maior volume de dados, use a variavel NUM_EXECUTIONS e informe o numero de vezes o arquivo deve ser inserido no banco. Ex.:

```sh make run DB_URI="postgres://service:service@localhost:5432/postgres?sslmode=disable" FILE=arquivoTeste/base_teste.txt NUM_EXECUTIONS=10```

*Irá inserir o arquivo 10 vezes.


Para verificar a performance executei o programa configurado para inserir o arquivo 10 vezes. (NUM_EXECUTIONS=10)

Nesta condição o programa encerra em ~10 segundos, inserindo 1 arquivo (~50k registros) por segundo.

[INFO] Inserido 10 arquivo(s) em 9.90 segundos.

Utitizei uma metodoligia de bulk inserts através de transactions exemplificada no seguinte site (é o segundo método descrito pelo autor):

link referência para bulk inserts no banco: https://medium.com/@amoghagarwal/insert-optimisations-in-golang-26884b183b35
