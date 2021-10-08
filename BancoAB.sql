CREATE DATABASE BASE_A;

USE BASE_A

CREATE TABLE Pessoa( 
	Cpf varchar(14) NOT NULL PRIMARY KEY, 
	Nome varchar(150),
	DataNascimento timestamp,
	DataCriacao timestamp
);

CREATE TABLE Endereco( 
	Id int NOT NULL PRIMARY KEY, 
	Logadouro varchar(150),
	Numero varchar(3),
	Bairro varchar(150),
	Cidade varchar(150),
	Estado varchar(2),
	Complemento varchar(150),
	PontoReferencia varchar(150),
	DataCriacao timestamp,
	Cpf varchar(14),
	FOREIGN KEY (Cpf) REFERENCES Pessoa(Cpf)
);



CREATE TABLE BemMaterial( 
	Id int NOT NULL PRIMARY KEY, 
	Descricao varchar(250),
	Valor Numeric(12,2),
	DataCriacao timestamp,
	Cpf varchar(14),
	FOREIGN KEY (Cpf) REFERENCES Pessoa(Cpf)
);

CREATE TABLE Divida( 
	Id int NOT NULL PRIMARY KEY, 
	Valor Numeric(12,2)
	DataCriacao timestamp,
	Cpf varchar(14),
	FOREIGN KEY (Cpf) REFERENCES Pessoa(Cpf)
);

CREATE TABLE Renda( 
	Id int NOT NULL PRIMARY KEY, 
	Descricao varchar(250)
	Valor Numeric(12,2)
	DataCriacao timestamp,
	Cpf varchar(14),
	FOREIGN KEY (Cpf) REFERENCES Pessoa(Cpf)	
);


