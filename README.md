## Projeto Controle de Crédito
 Projeto criado com intuito de resolver a próblemática oferecendo: armazenamento, processamento e disponibilização de dados, sempre considerando que tudo deve estar conforme as boas práticas de segurança em TI. 
 No projeto existem três grandes bases de dados externas que organizam as seguintes informações:
### Base A
 • CPF
 • Nome
 • Endereço
 • Lista de dívidas
### Base B
 • Idade
 • Lista de bens (Imóveis, etc)
 • Endereço
 • Fonte de renda
### Base C
• Última consulta do CPF em um Bureau de crédito (Serasa e outros).
• Movimentação financeira nesse CPF.
• Dados relacionados a última compra com cartao de crédito vinculado ao CPF.
### Ferramentas
Utilizando a linguagem de progamação da google GoLang, postgreSQL com SSL

### UML
Na imagem abaixo é apresentada uma possivel representaçao de dados pos processamento para aproveitamentos das bases A e B

![UML Diagram](https://user-images.githubusercontent.com/36166925/136266150-391d61c6-1c3c-4172-9629-9a76f3be512a.jpg)

### Considerações finais
Gostaria de ter colocado a Base B e a Base C no Elastic para otimização mas infelizmente me vi impossibilitada pela escolha infeliz do beego que o ORM não da suporte ao Elastic. Tentei também ativar o sslmode para cripitografar as informações do postegres mas não consegui, acredito que se tivesse tentado por mais tempo teria sido possível, voltarei a esse projeto futuramente e o terminarei.
