# PSNeoway
Serviço desenvolvido para o processo seletivo da Neoway.

Autor: Felipe Pereira de Souza Duarte




Como executar:

Rode o seguinte comando para preparar o ambiente (deploy container e configurar banco):

make setup-env

Execute o programa com o seguinte comando:

make run DB_URI="postgres://service:service@localhost:5432/postgres?sslmode=disable" FILE="arquivoTeste/base_teste.txt"

A variavel "DB_URI" deve conter as informações necessárias para a conexão com o banco. Altere caso necessário para encaixar com a sua configuração.

A variavel "FILE" deve apontar para o arquivo a ser importado no banco.

*Para simular um maior volume de dados, está configurado para importar o arquivo especificado 10 vezes.



Utitizei uma metodolia de bulk inserts através de transactions exemplificada no seguinte site:

link referência para bulk inserts no banco: https://medium.com/@amoghagarwal/insert-optimisations-in-golang-26884b183b35
