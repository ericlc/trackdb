# trackdb

trackdb is an application to execute and track the execution of SQL in databases. It is designed to enhance the application deployment that depends of databases changes.

## Objectives

trackdb follow some principles guidelines

- never modifies a sql script
- hide sensitive output data of the execution
- do not expose database password as plain text
- execute mutiple sql files per request
- compatible with major rdbms such as: oracle, sql server, ibm db2, mysql and postgres

## trackdb syntax

trackdb has two methods of executing a sql file:

1. flat SQL file
2. track SQL file

and has two types of execution:

1. version execution (V) - checks if that SQL has already been executed
2. repeatable execution (R) - does not check if that SQL has already been executed

### Flat SQL File

trackdb uses the filename to capture the properties of execution.

filename example: V01__create_table.sql

- V - type of execution: version (checks if that SQL has already been executed)
- 01 - number of execution
- create_table - description of the file
- .sql - file extension

##### Example

filename: V01_create_table.sql

```SQL
CREATE TABLE table_name (
    column1 datatype,
    column2 datatype,
    column3 datatype,
   ....
);
```

### Track SQL File

##### Example

```SQL
--trackdb

--track:01 (v)
CREATE TABLE table_name (
    column1 datatype,
    column2 datatype,
    column3 datatype,
   ....
);

--track:02 (v)
CREATE TABLE table_name2 (
    column1 datatype,
    column2 datatype,
    column3 datatype,
   ....
);
```


