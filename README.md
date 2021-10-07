## Projeto-Go-Armazenamento-Dados
 Projeto criado com intuito de resolver a problemática oferecendo: armazenamento, processamento e disponibilização de dados, sempre considerando que tudo deve estar conforme as boas práticas de segurança em TI. 
 No projeto existem três grandes bases de dados externas que organizam as seguintes informações:
### Base A
 * CPF
 * Nome
 * Endereço
 * Lista de dívidas
### Base B
 * Idade
 * Lista de bens (Imóveis, etc)
 * Endereço
 * Fonte de renda
### Base C
* Última consulta do CPF em um Bureau de crédito (Serasa e outros).
* Movimentação financeira nesse CPF.
* Dados relacionados a última compra com cartao de crédito vinculado ao CPF.
### Ferramentas
Utilizando a linguagem de progamação da google GoLang, postgreSQL com SSL

### UML
Na imagem abaixo é apresentada uma possivel representaçao de dados pos processamento para aproveitamentos das bases A e B

![UML Diagram](https://user-images.githubusercontent.com/36166925/136266150-391d61c6-1c3c-4172-9629-9a76f3be512a.jpg)

### Execução do projeto

Para executar o projeto basta rodar os seguintes comandos:
```
cd controle-credito
go run main.go
```
