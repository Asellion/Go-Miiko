drop database `Miiko`;
create database if not exists `Miiko` default character set utf8;
use `Miiko`;

-- Drop View
-- drop view `love`;

-- Drop Table
drop table if exists `pins`;
drop table if exists `welcome`;

-- Create Table

-- Servers
create table if not exists `welcome` (
	`server` varchar(32) primary key,
	`channel` varchar(32) not null
) engine=InnoDB default charset=utf8;

-- Pins
create table if not exists `pins` (
	`server` varchar(32) not null,
	`message` varchar(32) primary key,
	`member` varchar(32) not null
) engine=InnoDB default charset=utf8;

-- Views
-- drop view `love`;
-- drop view `pins-count`;

-- Pins Count
create view `pins-count` as
select `server`, `member`, count(`message`) as `count`
from `pins`
group by `server`, `member`
order by `server`, `count` desc
;

-- Love
create view `love` as
select `server`, `member`, `count`
from `pins-count` as `p1`
where `p1`.`count` = (
	select max(`count`) as `max`
	from `pins-count` as `p2`
	where `p1`.`server` = `p2`.`server`
)
group by `server`
;

-- Test Values

-- Servers
INSERT INTO `welcome`(`server`, `channel`) VALUES("1", "1");
INSERT INTO `welcome`(`server`, `channel`) VALUES("2", "2");
INSERT INTO `welcome`(`server`, `channel`) VALUES("3", "3");
INSERT INTO `welcome`(`server`, `channel`) VALUES("4", "4");
INSERT INTO `welcome`(`server`, `channel`) VALUES("5", "5");

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
