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



## A sql file example (track sql file)

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

## Another sql file example (flat sql file)

filename: V01_create_table.sql

```SQL
CREATE TABLE table_name (
    column1 datatype,
    column2 datatype,
    column3 datatype,
   ....
);
```
