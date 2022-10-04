CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
drop table if exists company;
CREATE TABLE company(
    ID uuid DEFAULT uuid_generate_v4 (),
    Name VARCHAR(15) unique NOT NULL,
    Description VARCHAR(3000),
    EmployeesNum NUMERIC NOT NULL,
    Registered BOOLEAN NOT NULL,
    Type Numeric, 
    PRIMARY KEY(ID)
);

insert into company(Name,Description,EmployeesNum,Registered,Type) VALUES('Prva Kompanija','prva napravljena kompanija',10,true,1);
insert into company(Name,EmployeesNum,Registered,Type) VALUES('Druga Kompanija',100,true,1);