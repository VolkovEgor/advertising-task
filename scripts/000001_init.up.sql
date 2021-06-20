CREATE TABLE IF NOT EXISTS adverts (
    id serial PRIMARY KEY,
    title varchar(200) NOT NULL,
    description varchar(1000),
    photos varchar(50) [],
    price int,
    creation_date bigint NOT NULL
);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 1', 'Description 1', '{"link1", "link2", "link3"}', 10000, 1);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 2', 'Description 2', '{"link1", "link2"}', 30000, 2);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 3', 'Description 3', '{"link1", "link2", "link3", "link4"}', 60000, 3);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 4', 'Description 4', '{"link1", "link2", "link3"}', 50000, 4);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 5', 'Description 5', '{"link1"}', 20000, 5);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 6', 'Description 5', '{"link1"}', 25000, 6);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 7', 'Description 5', '{"link1"}', 35000, 7);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 8', 'Description 5', '{"link1"}', 20000, 8);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 9', 'Description 5', '{"link1"}', 21000, 9);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 10', 'Description 5', '{"link1"}', 15000, 10);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 10', 'Description 5', '{"link1"}', 16000, 11);

INSERT INTO adverts (title, description, photos, price, creation_date)
VALUES ('Advert 10', 'Description 5', '{"link1"}', 55000, 12);