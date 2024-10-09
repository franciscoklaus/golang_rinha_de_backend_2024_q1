CREATE DATABASE IF NOT EXISTS api;

USE api;

DROP TABLE IF EXISTS clientes;
DROP TABLE IF EXISTS transacoes;


CREATE TABLE clientes (
    id int auto_increment primary key,
    limite int NOT NULL,
    saldo_inicial int NOT NULL,
    saldo_atual int NOT NULL
);

INSERT INTO clientes (limite, saldo_inicial, saldo_atual) VALUES (100000, 0,0);
INSERT INTO clientes (limite, saldo_inicial, saldo_atual) VALUES (80000, 0,0);
INSERT INTO clientes (limite, saldo_inicial, saldo_atual) VALUES (1000000, 0,0);
INSERT INTO clientes (limite, saldo_inicial, saldo_atual) VALUES (10000000, 0,0);
INSERT INTO clientes (limite, saldo_inicial, saldo_atual) VALUES (500000, 0,0);


CREATE TABLE transacoes (
    valor INTEGER NOT NULL,
    tipo char(1) NOT NULL,
    descricao varchar(10) NOT NULL,
    realizada_em timestamp default current_timestamp(),
    cliente_id INTEGER NOT NULL REFERENCES clientes(id)
)
