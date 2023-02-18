CREATE TABLE item(
    id serial primary key,
    name varchar(20),
    value integer not null
);

DROP TABLE item;