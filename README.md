# trackdb

Trackdb is an application to execute and track the execution of SQL in databases. It is designed to enhance the application deployment that depends of databases changes.

## Objectives

Trackdb follow some principles guidelines

- never modifies a SQL script
- hide sensitive output data of the execution
- do not expose database password as plain text
- execute mutiple SQL files per request
- compatible with major rdbms such as: oracle, sql server, ibm db2, mysql and postgres

## Trackdb syntax

Trackdb has two methods of executing a SQL file:

1. Flat SQL file
2. Track SQL file

And has two types of execution:

1. Version execution (V) - checks if that SQL has already been executed
2. Repeatable execution (R) - does not check if that SQL has already been executed

### Flat SQL File

Trackdb uses the filename to capture the properties of execution.

Filename example: V01__create_table.sql
- V - type of execution: version (checks if that SQL has already been executed)
- 01 - number of execution
- double underscore
- create_table - description of the file
- .sql - file extension

Filename example: R02__create_proc.sql
- R - type of execution: repeatable (does not check, always execute)
- 02 - number of execution
- double underscore
- create_proc - description of the file
- .sql - file extension

#### Example

Filename: V01_create_table.sql

```SQL
CREATE TABLE table_name (
    column1 datatype,
    column2 datatype,
    column3 datatype,
   ....
);
```

Filename: R02_create_proc.sql

```SQL
CREATE OR REPLACE PROCEDURE remove_emp (employee_id NUMBER) AS
   tot_emps NUMBER;
   BEGIN
      DELETE FROM employees
      WHERE employees.employee_id = remove_emp.employee_id;
   tot_emps := tot_emps - 1;
   END;
/
```

### Track SQL File

Trackdb uses SQL comments to capture the properties of execution. The filename does not have any restrictions.

Track example:
- --track:01 (V)
- --track:02 (R)

Filename: changesv01.sql

#### Example

```SQL
--trackdb

--track:01 (V)
CREATE TABLE table_name (
    column1 datatype,
    column2 datatype,
    column3 datatype,
   ....
);

--track:02 (R)
CREATE OR REPLACE PROCEDURE remove_emp (employee_id NUMBER) AS
   tot_emps NUMBER;
   BEGIN
      DELETE FROM employees
      WHERE employees.employee_id = remove_emp.employee_id;
   tot_emps := tot_emps - 1;
   END;
/
```
