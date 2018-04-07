drop database `miiko`;
create database if not exists `miiko` default character set utf8;
use `eldarya`;

drop table if exists `welcome`;
create table if not exists `welcome` (
	`server` varchar(32) primary key,
	`channel` varchar(32) not null
) engine=InnoDB default charset=utf8;