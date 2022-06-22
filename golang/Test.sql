use test_api;
create table Articles(
	Id int ,
	Title varchar(30),
    Descript varchar(255),
	Content Blob,
    primary key(Id)
);

insert into Articles Values(1,"the house","Some blode wrote this","haha");
insert into Articles Values(2,"the house_2","Some blode wrote this","haha");
insert into Articles Values(3,"the house_3","Some blode wrote this","haha");
insert into Articles Values(4,"the house_4","Some blode wrote this","haha");
insert into Articles Values(5,"the house_5","Some blode wrote this","haha");
insert into Articles Values(6,"the house_6","Some blode wrote this","haha");
insert into Articles Values(7,"the house_7","Some blode wrote this","haha");
insert into Articles Values(8,"the house_8","Some blode wrote this","haha");
insert into Articles Values(9,"the house_9","Some blode wrote this","haha");
insert into Articles Values(10,"the house_10","Some blode wrote this","haha");
insert into Articles Values(11,"the house_11","Some blode wrote this","haha");
insert into Articles Values(12,"the house_12","Some blode wrote this","haha");
