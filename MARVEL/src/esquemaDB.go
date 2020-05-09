create database if not exists agenda;
use agenda;
create table if not exists agenda(
    id bigint unsigned not null auto_increment,
    nombre varchar(255) not null,
    direccion varchar(255) not null,
    correo_electronico varchar(255) not null,
    primary key(id)
);