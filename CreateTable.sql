DROP TABLE IF EXISTS restaurant, general_user_info, host, client, cookie, courier CASCADE;

CREATE TABLE cart(
                     id serial,
                     FOREIGN KEY (client_id) INTEGER REFERENCES general_user_info (id) ON DELETE CASCADE,
                     FOREIGN KEY (restaurant) INTEGER REFERENCES restaurant (id) ON DELETE CASCADE
);

CREATE TABLE general_user_info
(
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    password varchar(64) NOT NULL,
    salt varchar(5) NOT NULL,
    phone varchar(15) UNIQUE NOT NULL,
    email text UNIQUE,
    avatar text DEFAULT '/uploads/',
    date_registration timestamp DEFAULT NOW() NOT NULL,
    deleted boolean DEFAULT false
);


CREATE TABLE restaurant (
    id serial PRIMARY KEY,
    owner INTEGER,
    FOREIGN KEY (owner) REFERENCES general_user_info (id)ON DELETE CASCADE,
    name text NOT NULL,
    description text NOT NULL,
    created timestamp DEFAULT NOW() NOT NULL,
    deleted boolean DEFAULT false,
    avatar text DEFAULT '/uploads/',
    min_price int DEFAULT 0,
    price_delivery int NOT NULL,
    min_delivery_time int,
    max_delivery_time int,
    city text NOT NULL,
    street text NOT NULL,
    house text NOT NULL,
    floor int,
    rating double precision,
    location text
);

CREATE TABLE cookie (
    id serial PRIMARY KEY,
    client_id INTEGER,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    session_id text NOT NULL,
    date_life timestamp NOT NULL,
    csrf_token varchar(64) NOT NULL
);

CREATE TABLE host (
    id serial PRIMARY KEY,
    client_id INTEGER UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE
);

CREATE TABLE client (
    id serial PRIMARY KEY,
    client_id INTEGER UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    date_birthday timestamp NOT NULL
);

CREATE TABLE courier (
    id serial PRIMARY KEY,
    client_id  INTEGER UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE
);

CREATE TABLE worker (
    id serial PRIMARY KEY,
    client_id INTEGER UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE
);

CREATE TABLE manager (
    id serial PRIMARY KEY,
    client_id INTEGER UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE
);

CREATE TABLE card (
    id serial PRIMARY KEY,
    client_id INTEGER,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    number BIGINT NOT NULL,
    month varchar(2) NOT NULL,
    year varchar(2) NOT NULL,
    alias: text
);

CREATE TABLE feedback (
    id serial PRIMARY KEY,
    author INTEGER,
    restaurant INTEGER,
    FOREIGN KEY (author) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    text text,
    created timestamp DEFAULT NOW() NOT NULL,
    emoji int NOT NULL,
    grade int NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE favorite_restaurant (
    id serial PRIMARY KEY,
    author INTEGER,
    restaurant INTEGER,
    FOREIGN KEY (author) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE
);


CREATE TABLE event (
    id serial PRIMARY KEY,
    restaurant INTEGER,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    start_date timestamp DEFAULT NOW() NOT NULL,
    end_date NOT NULL,
    name text NOT NULL,
    description text NOT NULL
);

CREATE TABLE restaurant_category (
    id serial PRIMARY KEY,
    category int NOT NULL,
    restaurant INTEGER,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
);

CREATE TABLE dishes (
    id serial PRIMARY KEY,
    name text NOT NULL,
    restaurant INTEGER,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    description text NOT NULL,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    carbohydrates double precision NOT NULL,
    category text NOT NULL,
    avatar text DEFAULT '/uploads/',
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE address_user (
    id serial PRIMARY KEY,
    client_id INTEGER,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    city text NOT NULL,
    street text NOT NULL,
    house int NOT NULL,
    flat int,
    porch int,
    floor int,
    intercom int,
    comment text DEFAULT '',
    alias text DEFAULT '',
    location text,
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE struture_dishes (
    id serial PRIMARY KEY,
    element text DEFAULT '' NOT NULL,
    food INTEGER,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    carbohydrates double precision NOT NULL,
    count INTEGER NOT NULL,
    changed boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE order_list (
    id serial PRIMARY KEY,
    order int INTEGER,
    food INTEGER,
    FOREIGN KEY (order) REFERENCES order (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE
);

CREATE TABLE promocode (
    id serial PRIMARY KEY,
    restaurant int INTEGER,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    name text NOT NULL,
    order_sale DEFAULT 0 NOT NULL,
    delivery_sale INTEGER DEFAULT 0 NOT NULL,
    end_date timestamp NOT NULL,
    start_date timestamp DEFAULT NOW()
);

CREATE TABLE promocode_on_food (
    id serial PRIMARY KEY,
    sale int INTEGER,
    promocode INTEGER,
    food INTEGER,
    FOREIGN KEY (promocode) REFERENCES promocode (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes(id) ON DELETE CASCADE
);

CREATE TABLE order (
    id serial PRIMARY KEY,
    client_id INTEGER,
    courier_id INTEGER,
    address_id INTEGER,
    restaurant_id INTEGER,
    promocode_id INTEGER,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (courier_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (address_id) REFERENCES address_user (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant_id) REFERENCES restaurant (id) ON DELETE CASCADE,
    FOREIGN KEY (promocode_id) REFERENCES promocode (id) ON DELETE CASCADE,
    comment text DEFAULT '' NOT NULL,
    status text DEFAULT '' NOT NULL,
    method_pay INTEGER NOT NULL,

);

INSERT INTO general_user_info (name, email, phone, password, salt) VALUES ('root', 'root@root', 88888888888, 'ca2e080a74ed1590cd141171c20e164d40d058fb45817c7b59f83159d059a6c0', 'salt')
INSERT INTO client (client_id, date_birthday) VALUES (1, NOW())
