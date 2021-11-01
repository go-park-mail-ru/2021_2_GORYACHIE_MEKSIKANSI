CREATE TABLE IF NOT EXISTS general_user_info
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


CREATE TABLE IF NOT EXISTS restaurant (
    id serial PRIMARY KEY,
    owner int,
    FOREIGN KEY (owner) REFERENCES general_user_info (id) ON DELETE CASCADE,
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

CREATE TABLE IF NOT EXISTS cookie (
    id serial PRIMARY KEY,
    client_id int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    session_id varchar(92) NOT NULL,
    date_life timestamp NOT NULL,
    csrf_token varchar(92) NOT NULL
);

CREATE TABLE IF NOT EXISTS host (
    id serial PRIMARY KEY,
    client_id int UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS client (
    id serial PRIMARY KEY,
    client_id int UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    date_birthday timestamp DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS courier (
    id serial PRIMARY KEY,
    client_id int UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS worker (
    id serial PRIMARY KEY,
    client_id int UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS manager (
    id serial PRIMARY KEY,
    client_id int UNIQUE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS card (
    id serial PRIMARY KEY,
    client_id int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    number BIGINT NOT NULL,
    month varchar(2) NOT NULL,
    year varchar(2) NOT NULL,
    alias text
);

CREATE TABLE IF NOT EXISTS feedback (
    id serial PRIMARY KEY,
    author int,
    restaurant int,
    FOREIGN KEY (author) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    text text,
    created timestamp DEFAULT NOW() NOT NULL,
    emoji int NOT NULL,
    grade int NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS favorite_restaurant (
    id serial PRIMARY KEY,
    author int,
    restaurant int,
    FOREIGN KEY (author) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS event (
    id serial PRIMARY KEY,
    restaurant int,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    start_date timestamp DEFAULT NOW() NOT NULL,
    end_date timestamp NOT NULL,
    name text NOT NULL,
    description text NOT NULL
);

CREATE TABLE IF NOT EXISTS restaurant_category (
    id serial PRIMARY KEY,
    category text NOT NULL,
    restaurant int,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dishes (
    id serial PRIMARY KEY,
    name text NOT NULL,
    cost int,
    count int,
    restaurant int,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    description text NOT NULL,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    kilocalorie int NOT NULL,
    carbohydrates double precision NOT NULL,
    weight int NOT NULL,
    category_dishes text NOT NULL,
    category_restaurant text NOT NULL,
    avatar text DEFAULT '/uploads/',
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS address_user (
    id serial PRIMARY KEY,
    client_id int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    city text NOT NULL,
    street text NOT NULL,
    house text NOT NULL,
    flat int DEFAULT 0,
    porch int DEFAULT 0,
    floor int DEFAULT 0,
    intercom text DEFAULT '',
    comment text DEFAULT '',
    alias text DEFAULT '',
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS structure_dishes (
    id serial PRIMARY KEY,
    name text DEFAULT '' NOT NULL,
    cost int,
    food int,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    carbohydrates double precision NOT NULL,
    kilocalorie int,
    count_element int NOT NULL,
    changed boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS radios (
   id serial PRIMARY KEY,
   name text DEFAULT '' NOT NULL,
   food int,
   FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS structure_radios (
    id serial PRIMARY KEY,
    name text DEFAULT '' NOT NULL,
    radios int,
    FOREIGN KEY (radios) REFERENCES radios (id) ON DELETE CASCADE,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    carbohydrates double precision NOT NULL,
    kilocalorie int NOT NULL
);

CREATE TABLE IF NOT EXISTS promocode (
    id serial PRIMARY KEY,
    restaurant int,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    name text NOT NULL,
    order_sale int DEFAULT 0 NOT NULL,
    delivery_sale int DEFAULT 0 NOT NULL,
    end_date timestamp NOT NULL,
    start_date timestamp DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS promocode_on_food (
    id serial PRIMARY KEY,
    sale int,
    promocode int,
    food int,
    FOREIGN KEY (promocode) REFERENCES promocode (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_user (
    id serial PRIMARY KEY,
    client_id int,
    courier_id int,
    address_id int,
    restaurant_id int,
    promocode_id int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (courier_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (address_id) REFERENCES address_user (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant_id) REFERENCES restaurant (id) ON DELETE CASCADE,
    FOREIGN KEY (promocode_id) REFERENCES promocode (id) ON DELETE CASCADE,
    comment text DEFAULT '' NOT NULL,
    status text DEFAULT '' NOT NULL,
    method_pay int NOT NULL
);

CREATE TABLE IF NOT EXISTS order_list (
    id serial PRIMARY KEY,
    order_id int,
    food int,
    FOREIGN KEY (order_id) REFERENCES order_user (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cart (
    id serial PRIMARY KEY,
    number_item int NOT NULL,
    client_id int,
    food int,
    count_food int DEFAULT 1 NOT NULL,
    restaurant_id int,
    FOREIGN KEY (restaurant_id) REFERENCES restaurant (id) ON DELETE CASCADE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cart_structure_food (
    id serial PRIMARY KEY,
    checkbox int,
    client_id int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (checkbox) REFERENCES structure_dishes (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS cart_radios_food (
    id serial PRIMARY KEY,
    radios_id int,
    radios int,
    client_id int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (radios_id) REFERENCES radios (id) ON DELETE CASCADE,
    FOREIGN KEY (radios) REFERENCES structure_radios (id) ON DELETE CASCADE
);
