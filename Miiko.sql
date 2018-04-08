drop database `miiko`;
create database if not exists `miiko` default character set utf8;
use `miiko`;

-- Drop View
-- drop view `love`;

-- Drop Table
drop table if exists `pins`;
drop table if exists `servers`;

-- Create Table

-- Servers
create table if not exists `servers` (
	`server` varchar(32) primary key,
	`welcome` varchar(32) not null
) engine=InnoDB default charset=utf8;

-- Pins
create table if not exists `pins` (
	`server` varchar(32) not null,
	`message` varchar(32) primary key,
	`member` varchar(32) not null
) engine=InnoDB default charset=utf8;

-- Views

-- Love
create view `love` as
select `server`, `member`
from (
	select `server`, `member`, count(`message`) as `count`
	from `pins`
	group by `server`, `member`
) as `pins-count`
group by `server`
;

-- Test Values

-- Servers
INSERT INTO `servers`(`server`) VALUES("1");
INSERT INTO `servers`(`server`) VALUES("2");
INSERT INTO `servers`(`server`) VALUES("3");
INSERT INTO `servers`(`server`) VALUES("4");
INSERT INTO `servers`(`server`) VALUES("5");

-- Pins
INSERT into `pins` VALUES ("1", "1", "1");
INSERT into `pins` VALUES ("2", "2", "2");
INSERT into `pins` VALUES ("2", "3", "3");
INSERT into `pins` VALUES ("3", "4", "1");
INSERT into `pins` VALUES ("3", "5", "2");
INSERT into `pins` VALUES ("3", "6", "3");
INSERT into `pins` VALUES ("4", "7", "1");
INSERT into `pins` VALUES ("4", "8", "2");
INSERT into `pins` VALUES ("4", "9", "3");
INSERT into `pins` VALUES ("4", "10", "1");
INSERT into `pins` VALUES ("5", "11", "2");
INSERT into `pins` VALUES ("5", "12", "3");
INSERT into `pins` VALUES ("5", "13", "1");
INSERT into `pins` VALUES ("5", "14", "2");
INSERT into `pins` VALUES ("5", "15", "3");
