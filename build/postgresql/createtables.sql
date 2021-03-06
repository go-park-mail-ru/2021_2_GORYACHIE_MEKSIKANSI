DROP INDEX IF EXISTS
    restaurant_fts, restaurant_category_fts;

CREATE TABLE IF NOT EXISTS general_user_info
(
    id SERIAL PRIMARY KEY,
    name text NOT NULL,
    password varchar(64) NOT NULL,
    salt varchar(5) NOT NULL,
    phone varchar(15) UNIQUE NOT NULL,
    email text UNIQUE,
    avatar text DEFAULT '/default/defaultUser.jpg',
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
    avatar text DEFAULT '/default/defaultRestaurant.jpg',
    min_price int DEFAULT 0,
    price_delivery int NOT NULL,
    min_delivery_time int,
    max_delivery_time int,
    city text NOT NULL,
    street text NOT NULL,
    house text NOT NULL,
    floor int,
    rating double precision,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    fts TSVECTOR
);

CREATE INDEX restaurant_fts ON restaurant (fts);

CREATE TABLE IF NOT EXISTS restaurant_category (
   id serial PRIMARY KEY,
   category text NOT NULL,
   restaurant int,
   place int,
   FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
   fts TSVECTOR
);

CREATE INDEX restaurant_category_fts ON restaurant_category (fts);

CREATE TABLE IF NOT EXISTS cookie (
    id serial PRIMARY KEY,
    client_id int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    session_id varchar(92) NOT NULL,
    date_life timestamp NOT NULL,
    csrf_token varchar(92) NOT NULL,
    websocket varchar(40)
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

CREATE TABLE IF NOT EXISTS review (
    id serial PRIMARY KEY,
    author int,
    restaurant int,
    FOREIGN KEY (author) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    text text,
    date_create timestamp DEFAULT NOW() NOT NULL,
    rate int NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS favorite_restaurant (
    id serial PRIMARY KEY,
    client int,
    restaurant int,
    position int,
    FOREIGN KEY (client) REFERENCES general_user_info (id) ON DELETE CASCADE,
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

CREATE TABLE IF NOT EXISTS dishes (
    id serial PRIMARY KEY,
    name text NOT NULL,
    cost int,
    count int,
    restaurant int,
    place int,
    place_category int,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    description text NOT NULL,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    kilocalorie int NOT NULL,
    carbohydrates double precision NOT NULL,
    weight int NOT NULL,
    category_dishes text NOT NULL,
    category_restaurant text NOT NULL,
    avatar text DEFAULT '/default/defaultDishes.jpg',
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS address_user (
    id serial PRIMARY KEY,
    client_id int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    city text NOT NULL,
    street text NOT NULL,
    house text NOT NULL,
    flat text DEFAULT '',
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
    place int,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    carbohydrates double precision NOT NULL,
    kilocalorie int NOT NULL,
    count_element int NOT NULL,
    changed boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS radios (
   id serial PRIMARY KEY,
   name text DEFAULT '' NOT NULL,
   place int,
   food int,
   FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS structure_radios (
    id serial PRIMARY KEY,
    name text DEFAULT '' NOT NULL,
    radios int,
    place int,
    FOREIGN KEY (radios) REFERENCES radios (id) ON DELETE CASCADE,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    carbohydrates double precision NOT NULL,
    kilocalorie int NOT NULL
);

CREATE TABLE IF NOT EXISTS promocode (
    id serial PRIMARY KEY,
    code text NOT NULL,
    type int DEFAULT 0 NOT NULL,
    restaurant int,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    name text NOT NULL,
    description text NOT NULL,
    start_date timestamp DEFAULT NOW(),
    end_date timestamp NOT NULL,
    avatar text NOT NULL,
    free_delivery bool,
    cost_for_free_dish int,
    free_dish_id int,
    FOREIGN KEY (free_dish_id) REFERENCES dishes (id) ON DELETE CASCADE,
    cost_for_sale int,
    sale_percent int,
    sale_amount int,
    time_for_sale timestamp,
    sale_in_time_percent int,
    sale_in_time_amount int
);

CREATE TABLE IF NOT EXISTS order_user (
    id serial PRIMARY KEY,
    client_id int,
    courier_id int,
    address_id int,
    restaurant_id int,
    promo_code text,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (courier_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (address_id) REFERENCES address_user (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant_id) REFERENCES restaurant (id) ON DELETE CASCADE,
    comment text DEFAULT '' NOT NULL,
    status int DEFAULT 1 NOT NULL,
    method_pay text NOT NULL,
    date_order timestamp DEFAULT NOW(),
    dCost int,
    sumCost int,
    check_run boolean DEFAULT true
);

CREATE TABLE IF NOT EXISTS order_list (
    id serial PRIMARY KEY,
    order_id int,
    food int,
    place int,
    count_dishes int,
    item_number int,
    FOREIGN KEY (order_id) REFERENCES order_user (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_structure_list (
    id serial PRIMARY KEY,
    order_id int,
    place int,
    food int,
    structure_food int,
    list_id int,
    FOREIGN KEY (list_id) REFERENCES order_list (id) ON DELETE CASCADE,
    FOREIGN KEY (order_id) REFERENCES order_user (id) ON DELETE CASCADE,
    FOREIGN KEY (structure_food) REFERENCES structure_dishes (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_radios_list (
    id serial PRIMARY KEY,
    order_id int,
    place int,
    radios_id int,
    radios int,
    food int,
    list_id int,
    FOREIGN KEY (list_id) REFERENCES order_list (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE,
    FOREIGN KEY (order_id) REFERENCES order_user (id) ON DELETE CASCADE,
    FOREIGN KEY (radios_id) REFERENCES radios (id) ON DELETE CASCADE,
    FOREIGN KEY (radios) REFERENCES structure_radios (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS cart_user (
    id SERIAL PRIMARY KEY,
    client_id int UNIQUE,
    promo_code text,
    restaurant int,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant) REFERENCES restaurant (id) ON DELETE CASCADE,
    UNIQUE(client_id, promo_code)
);

CREATE TABLE IF NOT EXISTS cart_food (
    id serial PRIMARY KEY,
    place int,
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
    place int,
    food int,
    cart_id int,
    FOREIGN KEY (cart_id) REFERENCES cart_food (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (checkbox) REFERENCES structure_dishes (id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS cart_radios_food (
    id serial PRIMARY KEY,
    radios_id int,
    radios int,
    place int,
    client_id int,
    food int,
    cart_id int,
    FOREIGN KEY (cart_id) REFERENCES cart_food (id) ON DELETE CASCADE,
    FOREIGN KEY (food) REFERENCES dishes (id) ON DELETE CASCADE,
    FOREIGN KEY (client_id) REFERENCES general_user_info (id) ON DELETE CASCADE,
    FOREIGN KEY (radios_id) REFERENCES radios (id) ON DELETE CASCADE,
    FOREIGN KEY (radios) REFERENCES structure_radios (id) ON DELETE CASCADE
);
