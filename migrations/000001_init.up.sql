CREATE TABLE IF NOT EXISTS adverts (
    id serial PRIMARY KEY,
    title varchar(200) NOT NULL,
    description varchar(1000),
    photos varchar(50) [],
    price int,
    creation_date bigint NOT NULL
);
