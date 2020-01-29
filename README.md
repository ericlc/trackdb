# trackdb

trackdb is an application to execute and track the execution of SQL in databases. It is designed to enhance the application deployment that depends of databases changes.

## Objectives

trackdb follow some principles guidelines

- never modifies a sql script
- hide sensitive output data of the execution
- do not expose database password as plain text
- execute mutiple sql files per request
- compatible with major rdbms such as: oracle, sql server, ibm db2, mysql and postgres

## Trackdb syntax

Trackdb has two methods of executing a sql file:

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

### Example

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

Trackdb uses SQL comments to capture the properties of execution.

Track example:
- --track:01 (V)

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


