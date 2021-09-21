CREATE TABLE IF NOT EXISTS owners (
    id SERIAL PRIMARY KEY,
    name text,
    gender text
);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    owner_id int REFERENCES owners,
    name text,
    tier text
);

INSERT INTO owners(name, gender) VALUES('Jeremy', 'M');
INSERT INTO owners(name, gender) VALUES('Brandon', 'M');
INSERT INTO owners(name, gender) VALUES('Alex', 'F');
INSERT INTO products(owner_id, name, tier) VALUES(1, 'WidgetFactory', 't0');
INSERT INTO products(owner_id, name, tier) VALUES(1, 'WidgetShipper', 't1');
INSERT INTO products(owner_id, name, tier) VALUES(2, 'FooBlaster', 't1');
INSERT INTO products(owner_id, name, tier) VALUES(3, 'BarGrabber', 't2');
