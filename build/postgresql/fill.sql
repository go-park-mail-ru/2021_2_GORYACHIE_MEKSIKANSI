--
--


SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
--

CREATE TABLE public.address_user (
    id integer NOT NULL,
    client_id integer,
    city text NOT NULL,
    street text NOT NULL,
    house text NOT NULL,
    flat text DEFAULT ''::text,
    porch integer DEFAULT 0,
    floor integer DEFAULT 0,
    intercom text DEFAULT ''::text,
    comment text DEFAULT ''::text,
    alias text DEFAULT ''::text,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE public.address_user OWNER TO root;

--
--

CREATE SEQUENCE public.address_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.address_user_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.address_user_id_seq OWNED BY public.address_user.id;


--
--

CREATE TABLE public.card (
    id integer NOT NULL,
    client_id integer,
    number bigint NOT NULL,
    month character varying(2) NOT NULL,
    year character varying(2) NOT NULL,
    alias text
);


ALTER TABLE public.card OWNER TO root;

--
--

CREATE SEQUENCE public.card_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.card_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.card_id_seq OWNED BY public.card.id;


--
--

CREATE TABLE public.cart_food (
    id integer NOT NULL,
    place integer,
    number_item integer NOT NULL,
    client_id integer,
    food integer,
    count_food integer DEFAULT 1 NOT NULL,
    restaurant_id integer
);


ALTER TABLE public.cart_food OWNER TO root;

--
--

CREATE SEQUENCE public.cart_food_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cart_food_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.cart_food_id_seq OWNED BY public.cart_food.id;


--
--

CREATE TABLE public.cart_radios_food (
    id integer NOT NULL,
    radios_id integer,
    radios integer,
    place integer,
    client_id integer,
    food integer,
    cart_id integer
);


ALTER TABLE public.cart_radios_food OWNER TO root;

--
--

CREATE SEQUENCE public.cart_radios_food_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cart_radios_food_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.cart_radios_food_id_seq OWNED BY public.cart_radios_food.id;


--
--

CREATE TABLE public.cart_structure_food (
    id integer NOT NULL,
    checkbox integer,
    client_id integer,
    place integer,
    food integer,
    cart_id integer
);


ALTER TABLE public.cart_structure_food OWNER TO root;

--
--

CREATE SEQUENCE public.cart_structure_food_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cart_structure_food_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.cart_structure_food_id_seq OWNED BY public.cart_structure_food.id;


--
--

CREATE TABLE public.cart_user (
    id integer NOT NULL,
    client_id integer,
    promo_code text,
    restaurant integer
);


ALTER TABLE public.cart_user OWNER TO root;

--
--

CREATE SEQUENCE public.cart_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cart_user_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.cart_user_id_seq OWNED BY public.cart_user.id;


--
--

CREATE TABLE public.client (
    id integer NOT NULL,
    client_id integer,
    date_birthday timestamp without time zone
);


ALTER TABLE public.client OWNER TO root;

--
--

CREATE SEQUENCE public.client_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.client_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.client_id_seq OWNED BY public.client.id;


--
--

CREATE TABLE public.cookie (
    id integer NOT NULL,
    client_id integer,
    session_id character varying(92) NOT NULL,
    date_life timestamp without time zone NOT NULL,
    csrf_token character varying(92) NOT NULL,
    websocket character varying(40)
);


ALTER TABLE public.cookie OWNER TO root;

--
--

CREATE SEQUENCE public.cookie_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cookie_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.cookie_id_seq OWNED BY public.cookie.id;


--
--

CREATE TABLE public.courier (
    id integer NOT NULL,
    client_id integer
);


ALTER TABLE public.courier OWNER TO root;

--
--

CREATE SEQUENCE public.courier_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.courier_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.courier_id_seq OWNED BY public.courier.id;


--
--

CREATE TABLE public.dishes (
    id integer NOT NULL,
    name text NOT NULL,
    cost integer,
    count integer,
    restaurant integer,
    place integer,
    place_category integer,
    description text NOT NULL,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    kilocalorie integer NOT NULL,
    carbohydrates double precision NOT NULL,
    weight integer NOT NULL,
    category_dishes text NOT NULL,
    category_restaurant text NOT NULL,
    avatar text DEFAULT '/default/defaultDishes.jpg'::text,
    deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE public.dishes OWNER TO root;

--
--

CREATE SEQUENCE public.dishes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dishes_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.dishes_id_seq OWNED BY public.dishes.id;


--
--

CREATE TABLE public.event (
    id integer NOT NULL,
    restaurant integer,
    start_date timestamp without time zone DEFAULT now() NOT NULL,
    end_date timestamp without time zone NOT NULL,
    name text NOT NULL,
    description text NOT NULL
);


ALTER TABLE public.event OWNER TO root;

--
--

CREATE SEQUENCE public.event_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.event_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.event_id_seq OWNED BY public.event.id;


--
--

CREATE TABLE public.favorite_restaurant (
    id integer NOT NULL,
    client integer,
    restaurant integer,
    "position" integer
);


ALTER TABLE public.favorite_restaurant OWNER TO root;

--
--

CREATE SEQUENCE public.favorite_restaurant_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.favorite_restaurant_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.favorite_restaurant_id_seq OWNED BY public.favorite_restaurant.id;


--
--

CREATE TABLE public.general_user_info (
    id integer NOT NULL,
    name text NOT NULL,
    password character varying(64) NOT NULL,
    salt character varying(5) NOT NULL,
    phone character varying(15) NOT NULL,
    email text,
    avatar text DEFAULT '/default/defaultUser.jpg'::text,
    date_registration timestamp without time zone DEFAULT now() NOT NULL,
    deleted boolean DEFAULT false
);


ALTER TABLE public.general_user_info OWNER TO root;

--
--

CREATE SEQUENCE public.general_user_info_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.general_user_info_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.general_user_info_id_seq OWNED BY public.general_user_info.id;


--
--

CREATE TABLE public.host (
    id integer NOT NULL,
    client_id integer
);


ALTER TABLE public.host OWNER TO root;

--
--

CREATE SEQUENCE public.host_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.host_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.host_id_seq OWNED BY public.host.id;


--
--

CREATE TABLE public.manager (
    id integer NOT NULL,
    client_id integer
);


ALTER TABLE public.manager OWNER TO root;

--
--

CREATE SEQUENCE public.manager_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.manager_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.manager_id_seq OWNED BY public.manager.id;


--
--

CREATE TABLE public.order_list (
    id integer NOT NULL,
    order_id integer,
    food integer,
    place integer,
    count_dishes integer,
    item_number integer
);


ALTER TABLE public.order_list OWNER TO root;

--
--

CREATE SEQUENCE public.order_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_list_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.order_list_id_seq OWNED BY public.order_list.id;


--
--

CREATE TABLE public.order_radios_list (
    id integer NOT NULL,
    order_id integer,
    place integer,
    radios_id integer,
    radios integer,
    food integer,
    list_id integer
);


ALTER TABLE public.order_radios_list OWNER TO root;

--
--

CREATE SEQUENCE public.order_radios_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_radios_list_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.order_radios_list_id_seq OWNED BY public.order_radios_list.id;


--
--

CREATE TABLE public.order_structure_list (
    id integer NOT NULL,
    order_id integer,
    place integer,
    food integer,
    structure_food integer,
    list_id integer
);


ALTER TABLE public.order_structure_list OWNER TO root;

--
--

CREATE SEQUENCE public.order_structure_list_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_structure_list_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.order_structure_list_id_seq OWNED BY public.order_structure_list.id;


--
--

CREATE TABLE public.order_user (
    id integer NOT NULL,
    client_id integer,
    courier_id integer,
    address_id integer,
    restaurant_id integer,
    promo_code text,
    comment text DEFAULT ''::text NOT NULL,
    status integer DEFAULT 1 NOT NULL,
    method_pay text NOT NULL,
    date_order timestamp without time zone DEFAULT now(),
    dcost integer,
    sumcost integer,
    check_run boolean DEFAULT true
);


ALTER TABLE public.order_user OWNER TO root;

--
--

CREATE SEQUENCE public.order_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.order_user_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.order_user_id_seq OWNED BY public.order_user.id;


--
--

CREATE TABLE public.promocode (
    id integer NOT NULL,
    code text NOT NULL,
    type integer DEFAULT 0 NOT NULL,
    restaurant integer,
    name text NOT NULL,
    description text NOT NULL,
    start_date timestamp without time zone DEFAULT now(),
    end_date timestamp without time zone NOT NULL,
    avatar text NOT NULL,
    free_delivery boolean,
    cost_for_free_dish integer,
    free_dish_id integer,
    cost_for_sale integer,
    sale_percent integer,
    sale_amount integer,
    time_for_sale timestamp without time zone,
    sale_in_time_percent integer,
    sale_in_time_amount integer
);


ALTER TABLE public.promocode OWNER TO root;

--
--

CREATE SEQUENCE public.promocode_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.promocode_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.promocode_id_seq OWNED BY public.promocode.id;


--
--

CREATE TABLE public.radios (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    place integer,
    food integer
);


ALTER TABLE public.radios OWNER TO root;

--
--

CREATE SEQUENCE public.radios_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.radios_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.radios_id_seq OWNED BY public.radios.id;


--
--

CREATE TABLE public.restaurant (
    id integer NOT NULL,
    owner integer,
    name text NOT NULL,
    description text NOT NULL,
    created timestamp without time zone DEFAULT now() NOT NULL,
    deleted boolean DEFAULT false,
    avatar text DEFAULT '/default/defaultRestaurant.jpg'::text,
    min_price integer DEFAULT 0,
    price_delivery integer NOT NULL,
    min_delivery_time integer,
    max_delivery_time integer,
    city text NOT NULL,
    street text NOT NULL,
    house text NOT NULL,
    floor integer,
    rating double precision,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    fts tsvector
);


ALTER TABLE public.restaurant OWNER TO root;

--
--

CREATE TABLE public.restaurant_category (
    id integer NOT NULL,
    category text NOT NULL,
    restaurant integer,
    place integer,
    fts tsvector
);


ALTER TABLE public.restaurant_category OWNER TO root;

--
--

CREATE SEQUENCE public.restaurant_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.restaurant_category_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.restaurant_category_id_seq OWNED BY public.restaurant_category.id;


--
--

CREATE SEQUENCE public.restaurant_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.restaurant_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.restaurant_id_seq OWNED BY public.restaurant.id;


--
--

CREATE TABLE public.review (
    id integer NOT NULL,
    author integer,
    restaurant integer,
    text text,
    date_create timestamp without time zone DEFAULT now() NOT NULL,
    rate integer NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE public.review OWNER TO root;

--
--

CREATE SEQUENCE public.review_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.review_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.review_id_seq OWNED BY public.review.id;


--
--

CREATE TABLE public.structure_dishes (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    cost integer,
    food integer,
    place integer,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    carbohydrates double precision NOT NULL,
    kilocalorie integer NOT NULL,
    count_element integer NOT NULL,
    changed boolean DEFAULT false NOT NULL,
    deleted boolean DEFAULT false NOT NULL
);


ALTER TABLE public.structure_dishes OWNER TO root;

--
--

CREATE SEQUENCE public.structure_dishes_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.structure_dishes_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.structure_dishes_id_seq OWNED BY public.structure_dishes.id;


--
--

CREATE TABLE public.structure_radios (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    radios integer,
    place integer,
    protein double precision NOT NULL,
    falt double precision NOT NULL,
    carbohydrates double precision NOT NULL,
    kilocalorie integer NOT NULL
);


ALTER TABLE public.structure_radios OWNER TO root;

--
--

CREATE SEQUENCE public.structure_radios_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.structure_radios_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.structure_radios_id_seq OWNED BY public.structure_radios.id;


--
--

CREATE TABLE public.worker (
    id integer NOT NULL,
    client_id integer
);


ALTER TABLE public.worker OWNER TO root;

--
--

CREATE SEQUENCE public.worker_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.worker_id_seq OWNER TO root;

--
--

ALTER SEQUENCE public.worker_id_seq OWNED BY public.worker.id;


--
--

ALTER TABLE ONLY public.address_user ALTER COLUMN id SET DEFAULT nextval('public.address_user_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.card ALTER COLUMN id SET DEFAULT nextval('public.card_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.cart_food ALTER COLUMN id SET DEFAULT nextval('public.cart_food_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.cart_radios_food ALTER COLUMN id SET DEFAULT nextval('public.cart_radios_food_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.cart_structure_food ALTER COLUMN id SET DEFAULT nextval('public.cart_structure_food_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.cart_user ALTER COLUMN id SET DEFAULT nextval('public.cart_user_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.client ALTER COLUMN id SET DEFAULT nextval('public.client_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.cookie ALTER COLUMN id SET DEFAULT nextval('public.cookie_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.courier ALTER COLUMN id SET DEFAULT nextval('public.courier_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.dishes ALTER COLUMN id SET DEFAULT nextval('public.dishes_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.event ALTER COLUMN id SET DEFAULT nextval('public.event_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.favorite_restaurant ALTER COLUMN id SET DEFAULT nextval('public.favorite_restaurant_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.general_user_info ALTER COLUMN id SET DEFAULT nextval('public.general_user_info_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.host ALTER COLUMN id SET DEFAULT nextval('public.host_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.manager ALTER COLUMN id SET DEFAULT nextval('public.manager_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.order_list ALTER COLUMN id SET DEFAULT nextval('public.order_list_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.order_radios_list ALTER COLUMN id SET DEFAULT nextval('public.order_radios_list_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.order_structure_list ALTER COLUMN id SET DEFAULT nextval('public.order_structure_list_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.order_user ALTER COLUMN id SET DEFAULT nextval('public.order_user_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.promocode ALTER COLUMN id SET DEFAULT nextval('public.promocode_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.radios ALTER COLUMN id SET DEFAULT nextval('public.radios_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.restaurant ALTER COLUMN id SET DEFAULT nextval('public.restaurant_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.restaurant_category ALTER COLUMN id SET DEFAULT nextval('public.restaurant_category_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.review ALTER COLUMN id SET DEFAULT nextval('public.review_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.structure_dishes ALTER COLUMN id SET DEFAULT nextval('public.structure_dishes_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.structure_radios ALTER COLUMN id SET DEFAULT nextval('public.structure_radios_id_seq'::regclass);


--
--

ALTER TABLE ONLY public.worker ALTER COLUMN id SET DEFAULT nextval('public.worker_id_seq'::regclass);


--
--

INSERT INTO public.address_user VALUES (1, 1, 'Москва', 'Вязов', '2', '28', 2, 5, '28K', 'Есть злая собака', 'Мой дом', 500, 500, false);


--
--



--
--

INSERT INTO public.cart_food VALUES (295, 0, 1, 1, 1, 1, 1);


--
--



--
--



--
--



--
--

INSERT INTO public.client VALUES (1, 1, '2021-12-21 22:12:10.689883');


--
--

INSERT INTO public.cookie VALUES (1, 1, '1', '2022-12-21 22:12:10.689', '', NULL);


--
--



--
--

INSERT INTO public.dishes VALUES (1, 'Тако', 60, 1000, 1, 0, 0, '', 1, 1, 224, 1, 1, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (2, 'Пряник', 60, 1000, 1, 0, 1, '', 1, 1, 126, 1, 1, 'К чаю', 'К чаю', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (3, 'Чёрный бургер', 60, 1000, 1, 1, 0, '', 1, 1, 361, 1, 1, 'горячее', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (4, 'Пицца Ассорти', 60, 1000, 1, 0, 2, '', 1, 1, 1024, 1, 1, 'горячее', 'Пиццы', 'https://www.koolinar.ru/all_image/recipes/156/156543/recipe_7b4bb7f7-1d42-428a-bb0a-3db8df03093a.jpg', false);
INSERT INTO public.dishes VALUES (5, 'Кофе', 60, 1000, 1, 0, 3, '', 1, 1, 90, 1, 1, 'горячее', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (6, 'Картошка Фри', 60, 1000, 1, 2, 0, '', 1, 1, 232, 1, 1, 'горячее', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (7, 'Картошка по деревенски', 60, 1000, 1, 3, 0, '', 1, 1, 172, 1, 1, 'Горячее', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (8, 'МакКомбо', 256, 1000, 1, 0, 4, '', 1, 1, 5036, 1, 1, 'Горячее', 'Комбо', 'https://www.eatthis.com/wp-content/uploads/sites/4/2019/05/mcdonalds-fries-food-lights.jpg', false);
INSERT INTO public.dishes VALUES (9, 'Утреннее комбо', 99, 1000, 1, 1, 4, '', 1, 1, 4708, 1, 1, 'Горячее', 'Комбо', 'https://imageproxy.ru/img/crop/1380x920/https/xn--h1ame.xn--80adxhks/storage/app/uploads/public/5e2/700/f07/5e2700f079c4c587329799.jpg', false);
INSERT INTO public.dishes VALUES (10, 'Аппетитное комбо', 150, 1000, 1, 2, 4, '', 1, 1, 3575, 1, 1, 'Горячее', 'Комбо', 'https://www.iphones.ru/wp-content/plugins/wonderm00ns-simple-facebook-open-graph-tags/fbimg.php?img=https%3A%2F%2Fwww.iphones.ru%2Fwp-content%2Fuploads%2F2018%2F08%2FBurgerN.jpg', false);
INSERT INTO public.dishes VALUES (11, 'Универсальное комбо', 100, 1000, 1, 3, 4, '', 1, 1, 1500, 1, 1, 'Горячее', 'Комбо', 'https://eda.yandex.ru/images/3667559/9724883e03ae48c2b6a1e28c5b9ea111-680x500.jpeg', false);
INSERT INTO public.dishes VALUES (12, 'Рождественский кекс', 350, 1, 2, 4, 0, '', 1, 1, 820, 1, 160, 'Универсальное', 'Путешествие в зиму', 'https://www.gastronom.ru/binfiles/images/20151221/b3c8862d.jpg', false);
INSERT INTO public.dishes VALUES (13, 'Блинчики с икрой', 370, 1000, 2, 0, 0, '', 1, 1, 848, 1, 130, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336390947_m650_60910237c1eb8bd3ef90a2d181c4d679663e37c1358fd4586a0e77c75e217acf.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (14, 'Имбирно-медовый пряник', 180, 1000, 2, 1, 0, '', 1, 1, 928, 1, 80, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391648_m650_ad26bba2487253cb6864ab0b3d1d8f3108088a1de250d4ac93e732148cbd2f5f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (15, 'Какао пряничное', 220, 1000, 2, 2, 0, '', 1, 1, 648, 1, 300, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391796_m650_24b78286ca08ad2c2f1f1ed3bcc700d3435092f87021c4fb31d2f8b74da8af37.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (16, 'Сельдь под шубой', 350, 1000, 2, 3, 0, '', 1, 1, 987, 1, 260, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391001_m650_f1c5402c3ed4959fa51d2b82cc7b3538ee70353e835047281033d8cefa75f18b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (17, 'Капучино', 220, 1, 2, 4, 0, '', 1, 1, 78, 1, 300, 'Универсальное', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (18, 'Ролл Цезарь', 270, 1000, 2, 0, 1, '', 1, 1, 798, 1, 190, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490974_m650_a97f85e487af6dffb2a2af88f7fdc71d2d9ee2354ec24601c56a58c0e9540b71.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (19, 'Ролл Филадельфия', 429, 1000, 2, 1, 1, '', 1, 1, 425, 1, 600, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263431_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (20, 'Ролл с ростбифом', 350, 1000, 2, 2, 1, '', 1, 1, 772, 1, 150, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/15838/334459325_m650_7388a0b3c360c2df62ce1525a0053de791fd4cfb1aa70726fa0ea2724839c8ca.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (21, 'Клаб-сэндвич', 370, 1000, 2, 3, 1, '', 1, 1, 770, 1, 300, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490973_m650_ef5daba5a610ee2eabd16c6b8668295390ff4a214e1ffc78a1cc1bfba492097e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (22, 'Мохито', 330, 1000, 2, 0, 2, '', 1, 1, 56, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490992_m650_e3496f99e11ebf227a739e1b89926dc31cfe30495eeedf9424d7919963dbc982.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (23, 'Лимонад Домашний', 270, 1000, 2, 1, 2, '', 1, 1, 230, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490988_m650_fb2879f40902be792dd258eaee85ee9fa0fcffcf228aae32db7b4a9dfcf24a70.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (24, '7Up', 150, 1000, 2, 2, 2, '', 1, 1, 265, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/324535611_m650_7572828939a42371b8fd488d1b0d998755e8148aa2db61abdabb8c439ffd3b39.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (25, 'Пепси', 150, 1000, 2, 3, 2, '', 1, 1, 255, 1, 250, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263484_m650_9d9f9f789f6abc645d15bae17c5c3084a79df0143183553744f4641c8c7afbbd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (48, 'Тако', 60, 1000, 5, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (26, 'Шаурма домашняя', 199, 1, 3, 2, 0, 'У этого блюда есть абсолютно всё', 1, 1, 602, 1, 330, 'Универсальное', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542260_m650_b2b212d56f1bc91abe68636af87f1e185e719dfdd01409b7888c3c2884e71d04.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (27, 'Шаурма с тунцом', 229, 1000, 3, 0, 0, '', 1, 1, 697, 1, 260, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542767_m650_9321049bda74a7ef4ad30c7d93ccaf8b8ef4fceb3ce09e67b57e9bf5d42bc897.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (28, 'Бургер с куриной котлетой', 229, 1000, 3, 1, 0, '', 1, 1, 407, 1, 200, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/335894906_m650_e5d32a976fbbc466608c4a9625ec1d73e9034b37d6e9611f2bc4a5eaa427e36f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (29, 'Американо', 150, 1000, 3, 0, 1, '', 1, 1, 458, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263353_m650_f4265fc18ab291c9b77d274364519df8c59d3c398facdc47eaf993cc74f32994.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (30, 'Капучино', 130, 1000, 3, 1, 1, '', 1, 1, 927, 1, 200, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (31, 'Латте', 170, 1000, 3, 2, 1, '', 1, 1, 907, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490932_m650_cb85b5a866a3afaa2c1f08b52b425f6e78fb0891ee033676524d6d90b4c6b2fd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (32, 'Фильтр-кофе', 150, 1000, 3, 3, 1, '', 1, 1, 207, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263356_m650_750ecfdcd43338312299910b4a7777c0466586c170bdef71faa0be4dce816437.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (33, 'Блин Ватрушка', 165, 1000, 3, 0, 2, '', 1, 1, 342, 1, 196, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485249_m650_861f7ff88a9073dfcb8fec2d28d8bc72ea710f7caa343417ae6667a1e516630d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (34, 'Блин с сахаром', 87, 1000, 3, 1, 2, '', 1, 1, 335, 1, 124, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485888_m650_ebfd220955c6f29d73ae8df3ed98ed9a406101b3cf030df5b2573959addaad3d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (35, 'Блин с шоколадным кремом', 151, 1000, 3, 2, 2, '', 1, 1, 327, 1, 166, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485270_m650_36ced98a972f785c015b545212a58d23b3cce201603a556b5c998dbd53d4a8ce.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (36, 'Блин с вишневым вареньем', 93, 1000, 3, 3, 2, '', 1, 1, 292, 1, 126, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485286_m650_7334048751f2aa703a75266d1503ac44e8dd2929a429858ef468c59982d5a45e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (37, 'Шаурма домашняя', 199, 1, 4, 2, 0, 'У этого блюда есть абсолютно всё', 1, 1, 602, 1, 330, 'Универсальное', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542260_m650_b2b212d56f1bc91abe68636af87f1e185e719dfdd01409b7888c3c2884e71d04.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (38, 'Шаурма с тунцом', 229, 1000, 4, 0, 0, '', 1, 1, 697, 1, 260, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542767_m650_9321049bda74a7ef4ad30c7d93ccaf8b8ef4fceb3ce09e67b57e9bf5d42bc897.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (39, 'Бургер с куриной котлетой', 229, 1000, 4, 1, 0, '', 1, 1, 407, 1, 200, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/335894906_m650_e5d32a976fbbc466608c4a9625ec1d73e9034b37d6e9611f2bc4a5eaa427e36f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (40, 'Американо', 150, 1000, 4, 0, 1, '', 1, 1, 458, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263353_m650_f4265fc18ab291c9b77d274364519df8c59d3c398facdc47eaf993cc74f32994.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (41, 'Капучино', 130, 1000, 4, 1, 1, '', 1, 1, 927, 1, 200, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (42, 'Латте', 170, 1000, 4, 2, 1, '', 1, 1, 907, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490932_m650_cb85b5a866a3afaa2c1f08b52b425f6e78fb0891ee033676524d6d90b4c6b2fd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (43, 'Фильтр-кофе', 150, 1000, 4, 3, 1, '', 1, 1, 207, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263356_m650_750ecfdcd43338312299910b4a7777c0466586c170bdef71faa0be4dce816437.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (44, 'Блин Ватрушка', 165, 1000, 4, 0, 2, '', 1, 1, 342, 1, 196, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485249_m650_861f7ff88a9073dfcb8fec2d28d8bc72ea710f7caa343417ae6667a1e516630d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (45, 'Блин с сахаром', 87, 1000, 4, 1, 2, '', 1, 1, 335, 1, 124, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485888_m650_ebfd220955c6f29d73ae8df3ed98ed9a406101b3cf030df5b2573959addaad3d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (46, 'Блин с шоколадным кремом', 151, 1000, 4, 2, 2, '', 1, 1, 327, 1, 166, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485270_m650_36ced98a972f785c015b545212a58d23b3cce201603a556b5c998dbd53d4a8ce.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (47, 'Блин с вишневым вареньем', 93, 1000, 4, 3, 2, '', 1, 1, 292, 1, 126, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485286_m650_7334048751f2aa703a75266d1503ac44e8dd2929a429858ef468c59982d5a45e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (49, 'Пряник', 70, 1000, 5, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (50, 'Чёрный бургер', 139, 1000, 5, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (51, 'Кофе', 149, 1000, 5, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (52, 'Coca-cola', 65, 1000, 5, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (53, 'Fanta', 60, 1000, 5, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (54, 'Sprite', 65, 1000, 5, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (55, 'Картошка Фри', 60, 1000, 5, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (56, 'Картошка по деревенски', 60, 1000, 5, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (57, 'Блюдо со стейком', 256, 1000, 5, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (58, 'Сёмга', 800, 1, 5, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (59, 'Цезарь', 5, 1, 5, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (60, 'Бефстроганов', 1000, 1000, 5, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (61, 'Греческий', 230, 150, 6, 4, 0, '', 1, 1, 520, 1, 5, 'Универсальное', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/15838/328171056_m650_cfcbfe9ab6ebf66ca620792736fc485d3c0a91072b9143258e3c902a05717650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (62, 'Салат Крабовичок', 170, 1000, 6, 0, 0, '', 1, 1, 979, 1, 150, 'Пиццы', 'Салаты', 'https://www.povarenok.ru/data/cache/2013oct/16/04/532964_47569-710x550x.jpg', false);
INSERT INTO public.dishes VALUES (63, 'Теплый салат с говядиной', 370, 1000, 6, 1, 0, '', 1, 1, 926, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (64, 'Салат Бригантина', 350, 1000, 6, 2, 0, '', 1, 1, 901, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (65, 'Цезарь с креветками', 400, 1000, 6, 3, 0, '', 1, 1, 640, 1, 250, 'Пиццы', 'Салаты', 'https://2d-recept.com/wp-content/uploads/2019/05/salat-cezar-s-krevetkami-foto.jpg', false);
INSERT INTO public.dishes VALUES (66, 'Салат Винегрет', 360, 1, 6, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 270, 1, 250, 'Универсальное', 'Салаты', 'https://upload.wikimedia.org/wikipedia/commons/a/a7/Vinegret_cleaned.jpg', false);
INSERT INTO public.dishes VALUES (67, 'Борщ', 250, 1000, 6, 0, 1, '', 1, 1, 873, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128261_m650_c4d2c43bb3f55ec4004d87459556a79120b53b6e2528c352a439cd282ae22f0e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (68, 'Сборная солянка мясная', 350, 1000, 6, 1, 1, '', 1, 1, 335, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128262_m650_2c8a69834bd188549d9dad91e69845b4a67c58a2ca50f19d755ca9a191cd1299.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (69, 'Суп-лапша домашняя', 200, 1000, 6, 2, 1, '', 1, 1, 621, 1, 300, 'Пиццы', 'Супы', 'https://www.maggi.ru/data/images/recept/img640x500/recept_4495_czgj.jpg', false);
INSERT INTO public.dishes VALUES (70, 'Картофельный суп', 200, 1000, 6, 3, 1, '', 1, 1, 743, 1, 300, 'Пиццы', 'Супы', 'https://s1.eda.ru/StaticContent/Photos/120131084848/120213174933/p_O.jpg', false);
INSERT INTO public.dishes VALUES (71, 'Синнабон', 120, 1000, 6, 0, 2, '', 1, 1, 777, 1, 150, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334172939_m650_02814baf09423cd3bd412593eb3dfab2ecd2316537045b96989be181002b2ae9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (72, 'Трайфл Сникерс', 250, 1000, 6, 1, 2, '', 1, 1, 535, 1, 180, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333386700_m650_6cb3cf5e3c893441037dc08ab87054615e947637dcbaf3e6849c642aa1aab7c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (73, 'Кейк попс', 50, 1000, 6, 2, 2, '', 1, 1, 901, 1, 50, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334173008_m650_7c7a29300b7b1a11be154554e2120641a75e49ba543c884cc5ebdb060f682918.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (74, 'Шоколадный фондан', 180, 1000, 6, 3, 2, '', 1, 1, 213, 1, 150, 'Универсальное', 'Десерты', 'https://www.gastronom.ru/binfiles/images/20200121/bf504e3d.jpg', false);
INSERT INTO public.dishes VALUES (75, 'Тако', 60, 1000, 7, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (76, 'Пряник', 70, 1000, 7, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (77, 'Чёрный бургер', 139, 1000, 7, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (78, 'Кофе', 149, 1000, 7, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (79, 'Coca-cola', 65, 1000, 7, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (80, 'Fanta', 60, 1000, 7, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (81, 'Sprite', 65, 1000, 7, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (82, 'Картошка Фри', 60, 1000, 7, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (83, 'Картошка по деревенски', 60, 1000, 7, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (84, 'Блюдо со стейком', 256, 1000, 7, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (85, 'Сёмга', 800, 1, 7, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (86, 'Цезарь', 5, 1, 7, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (87, 'Бефстроганов', 1000, 1000, 7, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (88, 'Шашлык из свиной мякоти', 350, 1000, 8, 4, 0, 'Подаётся с луком и зеленью', 1, 1, 649, 1, 200, 'Универсальное', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128163_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (89, 'Шашлык из стейка семги', 540, 1000, 8, 0, 0, '', 1, 1, 1500, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128173_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (90, 'Шашлык из мякоти баранины', 570, 1000, 8, 1, 0, '', 1, 1, 924, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128168_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (91, 'Шашлык из куриного филе', 330, 1000, 8, 2, 0, '', 1, 1, 1234, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128165_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (92, 'Люля-кебаб из курицы', 230, 1000, 8, 3, 0, '', 1, 1, 996, 1, 180, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128170_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (93, 'Овощи на мангале', 230, 1000, 8, 0, 1, '', 1, 1, 750, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128174_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (94, 'Картофель с салом', 160, 1000, 8, 1, 1, '', 1, 1, 526, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128176_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (95, 'Грибы на мангале', 230, 1000, 8, 2, 1, '', 1, 1, 233, 1, 150, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128175_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (96, 'Лимонад Натахтари дюшес', 110, 1000, 8, 0, 2, '', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128181_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (97, 'Coca-cola', 110, 1000, 8, 1, 2, '', 1, 1, 110, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128183_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (98, 'Лимонад Натахтари тархун', 110, 1000, 8, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128182_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (99, 'Лимонад Натахтари барбарис', 110, 1000, 8, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://fermer-63.ru/upload/iblock/10e/10ed883d25ae6ffc445c96c519394779.jpg', false);
INSERT INTO public.dishes VALUES (100, 'Ролл с лососем', 219, 1000, 9, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (101, 'Ролл с огурцом', 149, 1000, 9, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (532, 'Sprite', 65, 1000, 41, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (102, 'Крабик Hot запеченный мини ролл', 189, 1000, 9, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (103, 'Лосось унаги', 319, 1000, 9, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (104, 'Ролл Аяши запеченный', 259, 1000, 9, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (105, 'Ролл Горячий лосось запеченный', 369, 1, 9, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (106, 'Вегетарианский Вок', 289, 1000, 9, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (107, 'Классический Вок', 349, 1000, 9, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (108, 'Сытный Вок', 369, 1000, 9, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (109, 'По-Китайский Вок', 349, 1000, 9, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (110, 'BonAqua', 79, 1000, 9, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (111, 'Coca-cola', 99, 1000, 9, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (112, 'Морс из клюквы', 99, 1000, 9, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (113, 'Морс из черной смородины', 99, 1000, 9, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (114, 'Шашлык из свиной мякоти', 350, 1000, 10, 4, 0, 'Подаётся с луком и зеленью', 1, 1, 649, 1, 200, 'Универсальное', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128163_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (115, 'Шашлык из стейка семги', 540, 1000, 10, 0, 0, '', 1, 1, 1500, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128173_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (116, 'Шашлык из мякоти баранины', 570, 1000, 10, 1, 0, '', 1, 1, 924, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128168_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (117, 'Шашлык из куриного филе', 330, 1000, 10, 2, 0, '', 1, 1, 1234, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128165_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (118, 'Люля-кебаб из курицы', 230, 1000, 10, 3, 0, '', 1, 1, 996, 1, 180, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128170_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (119, 'Овощи на мангале', 230, 1000, 10, 0, 1, '', 1, 1, 750, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128174_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (120, 'Картофель с салом', 160, 1000, 10, 1, 1, '', 1, 1, 526, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128176_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (121, 'Грибы на мангале', 230, 1000, 10, 2, 1, '', 1, 1, 233, 1, 150, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128175_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (122, 'Лимонад Натахтари дюшес', 110, 1000, 10, 0, 2, '', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128181_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (123, 'Coca-cola', 110, 1000, 10, 1, 2, '', 1, 1, 110, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128183_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (124, 'Лимонад Натахтари тархун', 110, 1000, 10, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128182_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (125, 'Лимонад Натахтари барбарис', 110, 1000, 10, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://fermer-63.ru/upload/iblock/10e/10ed883d25ae6ffc445c96c519394779.jpg', false);
INSERT INTO public.dishes VALUES (126, 'Ролл с лососем', 219, 1000, 11, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (127, 'Ролл с огурцом', 149, 1000, 11, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (128, 'Крабик Hot запеченный мини ролл', 189, 1000, 11, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (129, 'Лосось унаги', 319, 1000, 11, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (130, 'Ролл Аяши запеченный', 259, 1000, 11, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (131, 'Ролл Горячий лосось запеченный', 369, 1, 11, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (132, 'Вегетарианский Вок', 289, 1000, 11, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (133, 'Классический Вок', 349, 1000, 11, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (134, 'Сытный Вок', 369, 1000, 11, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (135, 'По-Китайский Вок', 349, 1000, 11, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (136, 'BonAqua', 79, 1000, 11, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (137, 'Coca-cola', 99, 1000, 11, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (138, 'Морс из клюквы', 99, 1000, 11, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (139, 'Морс из черной смородины', 99, 1000, 11, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (140, 'Ролл с лососем', 219, 1000, 12, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (141, 'Ролл с огурцом', 149, 1000, 12, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (142, 'Крабик Hot запеченный мини ролл', 189, 1000, 12, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (143, 'Лосось унаги', 319, 1000, 12, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (144, 'Ролл Аяши запеченный', 259, 1000, 12, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (145, 'Ролл Горячий лосось запеченный', 369, 1, 12, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (172, 'Пицца Карначина', 590, 1000, 14, 3, 0, '', 1, 1, 895, 1, 540, 'Пиццы', 'Пиццы', 'https://cdnn21.img.ria.ru/images/98976/61/989766135_0:100:2000:1233_1920x0_80_0_0_4a3e7af4d4ab43307a68343c059cc57d.jpg', false);
INSERT INTO public.dishes VALUES (146, 'Вегетарианский Вок', 289, 1000, 12, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (147, 'Классический Вок', 349, 1000, 12, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (148, 'Сытный Вок', 369, 1000, 12, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (149, 'По-Китайский Вок', 349, 1000, 12, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (150, 'BonAqua', 79, 1000, 12, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (151, 'Coca-cola', 99, 1000, 12, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (152, 'Морс из клюквы', 99, 1000, 12, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (153, 'Морс из черной смородины', 99, 1000, 12, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (154, 'Пицца Пепперони', 470, 1, 13, 4, 0, '', 1, 1, 952, 1, 420, 'Универсальное', 'Пицца', 'https://s1.eda.ru/StaticContent/Photos/120131085053/171027192707/p_O.jpg', false);
INSERT INTO public.dishes VALUES (155, 'Пицца Ассорти', 429, 1000, 13, 0, 0, '', 1, 1, 660, 1, 600, 'Пиццы', 'Пиццы', 'https://sp-ao.shortpixel.ai/client/q_lossless,ret_img/https://cookery.site/wp-content/uploads/2019/09/photo_92647025-1-678x381.jpg', false);
INSERT INTO public.dishes VALUES (156, 'Пицца Морская', 800, 1000, 13, 1, 0, '', 1, 1, 584, 1, 600, 'Пиццы', 'Пиццы', 'https://pizzarini.info/wp-content/uploads/2018/03/pitstsa-morskaya-1-1.jpg', false);
INSERT INTO public.dishes VALUES (157, 'Пицца Карбонара', 540, 1000, 13, 2, 0, '', 1, 1, 581, 1, 420, 'Пиццы', 'Пиццы', 'https://static.1000.menu/img/content/13174/picca-karbonara_1432215044_0_max.jpg', false);
INSERT INTO public.dishes VALUES (158, 'Пицца Карначина', 590, 1000, 13, 3, 0, '', 1, 1, 895, 1, 540, 'Пиццы', 'Пиццы', 'https://cdnn21.img.ria.ru/images/98976/61/989766135_0:100:2000:1233_1920x0_80_0_0_4a3e7af4d4ab43307a68343c059cc57d.jpg', false);
INSERT INTO public.dishes VALUES (159, 'Пицца 4 сезона', 590, 1, 13, 4, 0, '', 1, 1, 952, 1, 580, 'Универсальное', 'Пицца', 'https://www.citypizza.ru/upload/img/shop/pizza/BORT/resize/4-%D1%81%D0%B5%D0%B7%D0%BE%D0%BD%D0%B0-%D0%B1%D0%BE%D1%80%D1%82-listThumb.jpg', false);
INSERT INTO public.dishes VALUES (160, 'Ролл с угрем', 180, 1000, 13, 0, 1, '', 1, 1, 469, 1, 120, 'Пиццы', 'Роллы', 'https://lh3.googleusercontent.com/proxy/uqHy2GczztN6r4P_eH2lySHzJWY57YsPW-vhEuU9GF46-i2_mf8qBESx4VDgPk4Vck5v7HTfIzpS13PG3puLRnbGfCMw-z4qYPFkKjiNMLiabJhV-7R9ncW31eb2X9gCBeywUfBvapgC9v8KGkNYQakwrbPtaiMZb8__LAvccU9bbP7NoV2zFMDuvM8ZDr9FXqP5wlQHHqDOngrSSISAbFkg8I2uIeBsYggTPQayCdO6D5NhwNf9I8w6-ZLnpLRvTScL5Hg4ifyw_7ARlpDmw2yfx0Ypjw', false);
INSERT INTO public.dishes VALUES (161, 'Дыхание дракона', 360, 1000, 13, 1, 1, '', 1, 1, 356, 1, 120, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011355_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (162, 'Горячий краб', 290, 1000, 13, 2, 1, '', 1, 1, 378, 1, 320, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011356_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (163, 'Запеченный угорь', 330, 1000, 13, 3, 1, '', 1, 1, 654, 1, 200, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011359_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (164, 'Сок в ассортименте', 100, 1000, 13, 0, 2, '', 1, 1, 102, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664250_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (165, 'Coca-cola', 65, 1000, 13, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664245_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (166, 'Fanta', 60, 1000, 13, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082269_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (167, 'Sprite', 65, 1000, 13, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082268_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (168, 'Пицца Пепперони', 470, 1, 14, 4, 0, '', 1, 1, 952, 1, 420, 'Универсальное', 'Пицца', 'https://s1.eda.ru/StaticContent/Photos/120131085053/171027192707/p_O.jpg', false);
INSERT INTO public.dishes VALUES (169, 'Пицца Ассорти', 429, 1000, 14, 0, 0, '', 1, 1, 660, 1, 600, 'Пиццы', 'Пиццы', 'https://sp-ao.shortpixel.ai/client/q_lossless,ret_img/https://cookery.site/wp-content/uploads/2019/09/photo_92647025-1-678x381.jpg', false);
INSERT INTO public.dishes VALUES (170, 'Пицца Морская', 800, 1000, 14, 1, 0, '', 1, 1, 584, 1, 600, 'Пиццы', 'Пиццы', 'https://pizzarini.info/wp-content/uploads/2018/03/pitstsa-morskaya-1-1.jpg', false);
INSERT INTO public.dishes VALUES (171, 'Пицца Карбонара', 540, 1000, 14, 2, 0, '', 1, 1, 581, 1, 420, 'Пиццы', 'Пиццы', 'https://static.1000.menu/img/content/13174/picca-karbonara_1432215044_0_max.jpg', false);
INSERT INTO public.dishes VALUES (173, 'Пицца 4 сезона', 590, 1, 14, 4, 0, '', 1, 1, 952, 1, 580, 'Универсальное', 'Пицца', 'https://www.citypizza.ru/upload/img/shop/pizza/BORT/resize/4-%D1%81%D0%B5%D0%B7%D0%BE%D0%BD%D0%B0-%D0%B1%D0%BE%D1%80%D1%82-listThumb.jpg', false);
INSERT INTO public.dishes VALUES (174, 'Ролл с угрем', 180, 1000, 14, 0, 1, '', 1, 1, 469, 1, 120, 'Пиццы', 'Роллы', 'https://lh3.googleusercontent.com/proxy/uqHy2GczztN6r4P_eH2lySHzJWY57YsPW-vhEuU9GF46-i2_mf8qBESx4VDgPk4Vck5v7HTfIzpS13PG3puLRnbGfCMw-z4qYPFkKjiNMLiabJhV-7R9ncW31eb2X9gCBeywUfBvapgC9v8KGkNYQakwrbPtaiMZb8__LAvccU9bbP7NoV2zFMDuvM8ZDr9FXqP5wlQHHqDOngrSSISAbFkg8I2uIeBsYggTPQayCdO6D5NhwNf9I8w6-ZLnpLRvTScL5Hg4ifyw_7ARlpDmw2yfx0Ypjw', false);
INSERT INTO public.dishes VALUES (175, 'Дыхание дракона', 360, 1000, 14, 1, 1, '', 1, 1, 356, 1, 120, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011355_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (176, 'Горячий краб', 290, 1000, 14, 2, 1, '', 1, 1, 378, 1, 320, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011356_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (177, 'Запеченный угорь', 330, 1000, 14, 3, 1, '', 1, 1, 654, 1, 200, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011359_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (178, 'Сок в ассортименте', 100, 1000, 14, 0, 2, '', 1, 1, 102, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664250_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (179, 'Coca-cola', 65, 1000, 14, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664245_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (180, 'Fanta', 60, 1000, 14, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082269_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (181, 'Sprite', 65, 1000, 14, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082268_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (182, 'Шашлык из свиной мякоти', 350, 1000, 15, 4, 0, 'Подаётся с луком и зеленью', 1, 1, 649, 1, 200, 'Универсальное', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128163_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (183, 'Шашлык из стейка семги', 540, 1000, 15, 0, 0, '', 1, 1, 1500, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128173_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (184, 'Шашлык из мякоти баранины', 570, 1000, 15, 1, 0, '', 1, 1, 924, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128168_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (185, 'Шашлык из куриного филе', 330, 1000, 15, 2, 0, '', 1, 1, 1234, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128165_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (186, 'Люля-кебаб из курицы', 230, 1000, 15, 3, 0, '', 1, 1, 996, 1, 180, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128170_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (187, 'Овощи на мангале', 230, 1000, 15, 0, 1, '', 1, 1, 750, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128174_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (188, 'Картофель с салом', 160, 1000, 15, 1, 1, '', 1, 1, 526, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128176_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (189, 'Грибы на мангале', 230, 1000, 15, 2, 1, '', 1, 1, 233, 1, 150, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128175_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (190, 'Лимонад Натахтари дюшес', 110, 1000, 15, 0, 2, '', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128181_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (191, 'Coca-cola', 110, 1000, 15, 1, 2, '', 1, 1, 110, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128183_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (192, 'Лимонад Натахтари тархун', 110, 1000, 15, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128182_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (193, 'Лимонад Натахтари барбарис', 110, 1000, 15, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://fermer-63.ru/upload/iblock/10e/10ed883d25ae6ffc445c96c519394779.jpg', false);
INSERT INTO public.dishes VALUES (194, 'Ролл с лососем', 219, 1000, 16, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (195, 'Ролл с огурцом', 149, 1000, 16, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (196, 'Крабик Hot запеченный мини ролл', 189, 1000, 16, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (197, 'Лосось унаги', 319, 1000, 16, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (198, 'Ролл Аяши запеченный', 259, 1000, 16, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (199, 'Ролл Горячий лосось запеченный', 369, 1, 16, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (200, 'Вегетарианский Вок', 289, 1000, 16, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (201, 'Классический Вок', 349, 1000, 16, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (202, 'Сытный Вок', 369, 1000, 16, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (203, 'По-Китайский Вок', 349, 1000, 16, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (204, 'BonAqua', 79, 1000, 16, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (205, 'Coca-cola', 99, 1000, 16, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (206, 'Морс из клюквы', 99, 1000, 16, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (207, 'Морс из черной смородины', 99, 1000, 16, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (208, 'Рождественский кекс', 350, 1, 17, 4, 0, '', 1, 1, 820, 1, 160, 'Универсальное', 'Путешествие в зиму', 'https://www.gastronom.ru/binfiles/images/20151221/b3c8862d.jpg', false);
INSERT INTO public.dishes VALUES (209, 'Блинчики с икрой', 370, 1000, 17, 0, 0, '', 1, 1, 848, 1, 130, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336390947_m650_60910237c1eb8bd3ef90a2d181c4d679663e37c1358fd4586a0e77c75e217acf.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (210, 'Имбирно-медовый пряник', 180, 1000, 17, 1, 0, '', 1, 1, 928, 1, 80, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391648_m650_ad26bba2487253cb6864ab0b3d1d8f3108088a1de250d4ac93e732148cbd2f5f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (211, 'Какао пряничное', 220, 1000, 17, 2, 0, '', 1, 1, 648, 1, 300, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391796_m650_24b78286ca08ad2c2f1f1ed3bcc700d3435092f87021c4fb31d2f8b74da8af37.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (212, 'Сельдь под шубой', 350, 1000, 17, 3, 0, '', 1, 1, 987, 1, 260, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391001_m650_f1c5402c3ed4959fa51d2b82cc7b3538ee70353e835047281033d8cefa75f18b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (213, 'Капучино', 220, 1, 17, 4, 0, '', 1, 1, 78, 1, 300, 'Универсальное', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (214, 'Ролл Цезарь', 270, 1000, 17, 0, 1, '', 1, 1, 798, 1, 190, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490974_m650_a97f85e487af6dffb2a2af88f7fdc71d2d9ee2354ec24601c56a58c0e9540b71.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (215, 'Ролл Филадельфия', 429, 1000, 17, 1, 1, '', 1, 1, 425, 1, 600, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263431_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (216, 'Ролл с ростбифом', 350, 1000, 17, 2, 1, '', 1, 1, 772, 1, 150, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/15838/334459325_m650_7388a0b3c360c2df62ce1525a0053de791fd4cfb1aa70726fa0ea2724839c8ca.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (217, 'Клаб-сэндвич', 370, 1000, 17, 3, 1, '', 1, 1, 770, 1, 300, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490973_m650_ef5daba5a610ee2eabd16c6b8668295390ff4a214e1ffc78a1cc1bfba492097e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (218, 'Мохито', 330, 1000, 17, 0, 2, '', 1, 1, 56, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490992_m650_e3496f99e11ebf227a739e1b89926dc31cfe30495eeedf9424d7919963dbc982.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (219, 'Лимонад Домашний', 270, 1000, 17, 1, 2, '', 1, 1, 230, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490988_m650_fb2879f40902be792dd258eaee85ee9fa0fcffcf228aae32db7b4a9dfcf24a70.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (220, '7Up', 150, 1000, 17, 2, 2, '', 1, 1, 265, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/324535611_m650_7572828939a42371b8fd488d1b0d998755e8148aa2db61abdabb8c439ffd3b39.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (221, 'Пепси', 150, 1000, 17, 3, 2, '', 1, 1, 255, 1, 250, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263484_m650_9d9f9f789f6abc645d15bae17c5c3084a79df0143183553744f4641c8c7afbbd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (222, 'Пицца Пепперони', 470, 1, 18, 4, 0, '', 1, 1, 952, 1, 420, 'Универсальное', 'Пицца', 'https://s1.eda.ru/StaticContent/Photos/120131085053/171027192707/p_O.jpg', false);
INSERT INTO public.dishes VALUES (223, 'Пицца Ассорти', 429, 1000, 18, 0, 0, '', 1, 1, 660, 1, 600, 'Пиццы', 'Пиццы', 'https://sp-ao.shortpixel.ai/client/q_lossless,ret_img/https://cookery.site/wp-content/uploads/2019/09/photo_92647025-1-678x381.jpg', false);
INSERT INTO public.dishes VALUES (224, 'Пицца Морская', 800, 1000, 18, 1, 0, '', 1, 1, 584, 1, 600, 'Пиццы', 'Пиццы', 'https://pizzarini.info/wp-content/uploads/2018/03/pitstsa-morskaya-1-1.jpg', false);
INSERT INTO public.dishes VALUES (225, 'Пицца Карбонара', 540, 1000, 18, 2, 0, '', 1, 1, 581, 1, 420, 'Пиццы', 'Пиццы', 'https://static.1000.menu/img/content/13174/picca-karbonara_1432215044_0_max.jpg', false);
INSERT INTO public.dishes VALUES (226, 'Пицца Карначина', 590, 1000, 18, 3, 0, '', 1, 1, 895, 1, 540, 'Пиццы', 'Пиццы', 'https://cdnn21.img.ria.ru/images/98976/61/989766135_0:100:2000:1233_1920x0_80_0_0_4a3e7af4d4ab43307a68343c059cc57d.jpg', false);
INSERT INTO public.dishes VALUES (227, 'Пицца 4 сезона', 590, 1, 18, 4, 0, '', 1, 1, 952, 1, 580, 'Универсальное', 'Пицца', 'https://www.citypizza.ru/upload/img/shop/pizza/BORT/resize/4-%D1%81%D0%B5%D0%B7%D0%BE%D0%BD%D0%B0-%D0%B1%D0%BE%D1%80%D1%82-listThumb.jpg', false);
INSERT INTO public.dishes VALUES (228, 'Ролл с угрем', 180, 1000, 18, 0, 1, '', 1, 1, 469, 1, 120, 'Пиццы', 'Роллы', 'https://lh3.googleusercontent.com/proxy/uqHy2GczztN6r4P_eH2lySHzJWY57YsPW-vhEuU9GF46-i2_mf8qBESx4VDgPk4Vck5v7HTfIzpS13PG3puLRnbGfCMw-z4qYPFkKjiNMLiabJhV-7R9ncW31eb2X9gCBeywUfBvapgC9v8KGkNYQakwrbPtaiMZb8__LAvccU9bbP7NoV2zFMDuvM8ZDr9FXqP5wlQHHqDOngrSSISAbFkg8I2uIeBsYggTPQayCdO6D5NhwNf9I8w6-ZLnpLRvTScL5Hg4ifyw_7ARlpDmw2yfx0Ypjw', false);
INSERT INTO public.dishes VALUES (229, 'Дыхание дракона', 360, 1000, 18, 1, 1, '', 1, 1, 356, 1, 120, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011355_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (230, 'Горячий краб', 290, 1000, 18, 2, 1, '', 1, 1, 378, 1, 320, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011356_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (231, 'Запеченный угорь', 330, 1000, 18, 3, 1, '', 1, 1, 654, 1, 200, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011359_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (232, 'Сок в ассортименте', 100, 1000, 18, 0, 2, '', 1, 1, 102, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664250_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (233, 'Coca-cola', 65, 1000, 18, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664245_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (234, 'Fanta', 60, 1000, 18, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082269_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (235, 'Sprite', 65, 1000, 18, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082268_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (236, 'Ролл с лососем', 219, 1000, 19, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (237, 'Ролл с огурцом', 149, 1000, 19, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (238, 'Крабик Hot запеченный мини ролл', 189, 1000, 19, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (239, 'Лосось унаги', 319, 1000, 19, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (240, 'Ролл Аяши запеченный', 259, 1000, 19, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (241, 'Ролл Горячий лосось запеченный', 369, 1, 19, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (242, 'Вегетарианский Вок', 289, 1000, 19, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (243, 'Классический Вок', 349, 1000, 19, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (533, 'Картошка Фри', 60, 1000, 41, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (244, 'Сытный Вок', 369, 1000, 19, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (245, 'По-Китайский Вок', 349, 1000, 19, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (246, 'BonAqua', 79, 1000, 19, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (247, 'Coca-cola', 99, 1000, 19, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (248, 'Морс из клюквы', 99, 1000, 19, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (249, 'Морс из черной смородины', 99, 1000, 19, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (250, 'Ролл с лососем', 219, 1000, 20, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (251, 'Ролл с огурцом', 149, 1000, 20, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (252, 'Крабик Hot запеченный мини ролл', 189, 1000, 20, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (253, 'Лосось унаги', 319, 1000, 20, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (254, 'Ролл Аяши запеченный', 259, 1000, 20, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (255, 'Ролл Горячий лосось запеченный', 369, 1, 20, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (256, 'Вегетарианский Вок', 289, 1000, 20, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (257, 'Классический Вок', 349, 1000, 20, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (258, 'Сытный Вок', 369, 1000, 20, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (259, 'По-Китайский Вок', 349, 1000, 20, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (260, 'BonAqua', 79, 1000, 20, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (261, 'Coca-cola', 99, 1000, 20, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (262, 'Морс из клюквы', 99, 1000, 20, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (263, 'Морс из черной смородины', 99, 1000, 20, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (264, 'Шаурма домашняя', 199, 1, 21, 2, 0, 'У этого блюда есть абсолютно всё', 1, 1, 602, 1, 330, 'Универсальное', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542260_m650_b2b212d56f1bc91abe68636af87f1e185e719dfdd01409b7888c3c2884e71d04.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (559, 'Sprite', 65, 1000, 43, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (265, 'Шаурма с тунцом', 229, 1000, 21, 0, 0, '', 1, 1, 697, 1, 260, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542767_m650_9321049bda74a7ef4ad30c7d93ccaf8b8ef4fceb3ce09e67b57e9bf5d42bc897.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (266, 'Бургер с куриной котлетой', 229, 1000, 21, 1, 0, '', 1, 1, 407, 1, 200, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/335894906_m650_e5d32a976fbbc466608c4a9625ec1d73e9034b37d6e9611f2bc4a5eaa427e36f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (267, 'Американо', 150, 1000, 21, 0, 1, '', 1, 1, 458, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263353_m650_f4265fc18ab291c9b77d274364519df8c59d3c398facdc47eaf993cc74f32994.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (268, 'Капучино', 130, 1000, 21, 1, 1, '', 1, 1, 927, 1, 200, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (269, 'Латте', 170, 1000, 21, 2, 1, '', 1, 1, 907, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490932_m650_cb85b5a866a3afaa2c1f08b52b425f6e78fb0891ee033676524d6d90b4c6b2fd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (270, 'Фильтр-кофе', 150, 1000, 21, 3, 1, '', 1, 1, 207, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263356_m650_750ecfdcd43338312299910b4a7777c0466586c170bdef71faa0be4dce816437.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (271, 'Блин Ватрушка', 165, 1000, 21, 0, 2, '', 1, 1, 342, 1, 196, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485249_m650_861f7ff88a9073dfcb8fec2d28d8bc72ea710f7caa343417ae6667a1e516630d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (272, 'Блин с сахаром', 87, 1000, 21, 1, 2, '', 1, 1, 335, 1, 124, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485888_m650_ebfd220955c6f29d73ae8df3ed98ed9a406101b3cf030df5b2573959addaad3d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (273, 'Блин с шоколадным кремом', 151, 1000, 21, 2, 2, '', 1, 1, 327, 1, 166, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485270_m650_36ced98a972f785c015b545212a58d23b3cce201603a556b5c998dbd53d4a8ce.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (274, 'Блин с вишневым вареньем', 93, 1000, 21, 3, 2, '', 1, 1, 292, 1, 126, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485286_m650_7334048751f2aa703a75266d1503ac44e8dd2929a429858ef468c59982d5a45e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (275, 'Шаурма домашняя', 199, 1, 22, 2, 0, 'У этого блюда есть абсолютно всё', 1, 1, 602, 1, 330, 'Универсальное', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542260_m650_b2b212d56f1bc91abe68636af87f1e185e719dfdd01409b7888c3c2884e71d04.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (276, 'Шаурма с тунцом', 229, 1000, 22, 0, 0, '', 1, 1, 697, 1, 260, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542767_m650_9321049bda74a7ef4ad30c7d93ccaf8b8ef4fceb3ce09e67b57e9bf5d42bc897.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (277, 'Бургер с куриной котлетой', 229, 1000, 22, 1, 0, '', 1, 1, 407, 1, 200, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/335894906_m650_e5d32a976fbbc466608c4a9625ec1d73e9034b37d6e9611f2bc4a5eaa427e36f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (278, 'Американо', 150, 1000, 22, 0, 1, '', 1, 1, 458, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263353_m650_f4265fc18ab291c9b77d274364519df8c59d3c398facdc47eaf993cc74f32994.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (279, 'Капучино', 130, 1000, 22, 1, 1, '', 1, 1, 927, 1, 200, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (280, 'Латте', 170, 1000, 22, 2, 1, '', 1, 1, 907, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490932_m650_cb85b5a866a3afaa2c1f08b52b425f6e78fb0891ee033676524d6d90b4c6b2fd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (281, 'Фильтр-кофе', 150, 1000, 22, 3, 1, '', 1, 1, 207, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263356_m650_750ecfdcd43338312299910b4a7777c0466586c170bdef71faa0be4dce816437.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (282, 'Блин Ватрушка', 165, 1000, 22, 0, 2, '', 1, 1, 342, 1, 196, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485249_m650_861f7ff88a9073dfcb8fec2d28d8bc72ea710f7caa343417ae6667a1e516630d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (283, 'Блин с сахаром', 87, 1000, 22, 1, 2, '', 1, 1, 335, 1, 124, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485888_m650_ebfd220955c6f29d73ae8df3ed98ed9a406101b3cf030df5b2573959addaad3d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (284, 'Блин с шоколадным кремом', 151, 1000, 22, 2, 2, '', 1, 1, 327, 1, 166, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485270_m650_36ced98a972f785c015b545212a58d23b3cce201603a556b5c998dbd53d4a8ce.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (285, 'Блин с вишневым вареньем', 93, 1000, 22, 3, 2, '', 1, 1, 292, 1, 126, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485286_m650_7334048751f2aa703a75266d1503ac44e8dd2929a429858ef468c59982d5a45e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (286, 'Ролл с лососем', 219, 1000, 23, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (287, 'Ролл с огурцом', 149, 1000, 23, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (288, 'Крабик Hot запеченный мини ролл', 189, 1000, 23, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (289, 'Лосось унаги', 319, 1000, 23, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (290, 'Ролл Аяши запеченный', 259, 1000, 23, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (291, 'Ролл Горячий лосось запеченный', 369, 1, 23, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (292, 'Вегетарианский Вок', 289, 1000, 23, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (293, 'Классический Вок', 349, 1000, 23, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (294, 'Сытный Вок', 369, 1000, 23, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (295, 'По-Китайский Вок', 349, 1000, 23, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (296, 'BonAqua', 79, 1000, 23, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (297, 'Coca-cola', 99, 1000, 23, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (298, 'Морс из клюквы', 99, 1000, 23, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (299, 'Морс из черной смородины', 99, 1000, 23, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (300, 'Тако', 60, 1000, 24, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (301, 'Пряник', 70, 1000, 24, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (302, 'Чёрный бургер', 139, 1000, 24, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (303, 'Кофе', 149, 1000, 24, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (304, 'Coca-cola', 65, 1000, 24, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (305, 'Fanta', 60, 1000, 24, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (306, 'Sprite', 65, 1000, 24, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (307, 'Картошка Фри', 60, 1000, 24, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (308, 'Картошка по деревенски', 60, 1000, 24, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (309, 'Блюдо со стейком', 256, 1000, 24, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (310, 'Сёмга', 800, 1, 24, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (311, 'Цезарь', 5, 1, 24, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (312, 'Бефстроганов', 1000, 1000, 24, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (313, 'Греческий', 230, 150, 25, 4, 0, '', 1, 1, 520, 1, 5, 'Универсальное', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/15838/328171056_m650_cfcbfe9ab6ebf66ca620792736fc485d3c0a91072b9143258e3c902a05717650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (314, 'Салат Крабовичок', 170, 1000, 25, 0, 0, '', 1, 1, 979, 1, 150, 'Пиццы', 'Салаты', 'https://www.povarenok.ru/data/cache/2013oct/16/04/532964_47569-710x550x.jpg', false);
INSERT INTO public.dishes VALUES (315, 'Теплый салат с говядиной', 370, 1000, 25, 1, 0, '', 1, 1, 926, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (316, 'Салат Бригантина', 350, 1000, 25, 2, 0, '', 1, 1, 901, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (317, 'Цезарь с креветками', 400, 1000, 25, 3, 0, '', 1, 1, 640, 1, 250, 'Пиццы', 'Салаты', 'https://2d-recept.com/wp-content/uploads/2019/05/salat-cezar-s-krevetkami-foto.jpg', false);
INSERT INTO public.dishes VALUES (318, 'Салат Винегрет', 360, 1, 25, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 270, 1, 250, 'Универсальное', 'Салаты', 'https://upload.wikimedia.org/wikipedia/commons/a/a7/Vinegret_cleaned.jpg', false);
INSERT INTO public.dishes VALUES (319, 'Борщ', 250, 1000, 25, 0, 1, '', 1, 1, 873, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128261_m650_c4d2c43bb3f55ec4004d87459556a79120b53b6e2528c352a439cd282ae22f0e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (320, 'Сборная солянка мясная', 350, 1000, 25, 1, 1, '', 1, 1, 335, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128262_m650_2c8a69834bd188549d9dad91e69845b4a67c58a2ca50f19d755ca9a191cd1299.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (321, 'Суп-лапша домашняя', 200, 1000, 25, 2, 1, '', 1, 1, 621, 1, 300, 'Пиццы', 'Супы', 'https://www.maggi.ru/data/images/recept/img640x500/recept_4495_czgj.jpg', false);
INSERT INTO public.dishes VALUES (322, 'Картофельный суп', 200, 1000, 25, 3, 1, '', 1, 1, 743, 1, 300, 'Пиццы', 'Супы', 'https://s1.eda.ru/StaticContent/Photos/120131084848/120213174933/p_O.jpg', false);
INSERT INTO public.dishes VALUES (323, 'Синнабон', 120, 1000, 25, 0, 2, '', 1, 1, 777, 1, 150, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334172939_m650_02814baf09423cd3bd412593eb3dfab2ecd2316537045b96989be181002b2ae9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (324, 'Трайфл Сникерс', 250, 1000, 25, 1, 2, '', 1, 1, 535, 1, 180, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333386700_m650_6cb3cf5e3c893441037dc08ab87054615e947637dcbaf3e6849c642aa1aab7c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (325, 'Кейк попс', 50, 1000, 25, 2, 2, '', 1, 1, 901, 1, 50, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334173008_m650_7c7a29300b7b1a11be154554e2120641a75e49ba543c884cc5ebdb060f682918.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (326, 'Шоколадный фондан', 180, 1000, 25, 3, 2, '', 1, 1, 213, 1, 150, 'Универсальное', 'Десерты', 'https://www.gastronom.ru/binfiles/images/20200121/bf504e3d.jpg', false);
INSERT INTO public.dishes VALUES (327, 'Греческий', 230, 150, 26, 4, 0, '', 1, 1, 520, 1, 5, 'Универсальное', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/15838/328171056_m650_cfcbfe9ab6ebf66ca620792736fc485d3c0a91072b9143258e3c902a05717650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (328, 'Салат Крабовичок', 170, 1000, 26, 0, 0, '', 1, 1, 979, 1, 150, 'Пиццы', 'Салаты', 'https://www.povarenok.ru/data/cache/2013oct/16/04/532964_47569-710x550x.jpg', false);
INSERT INTO public.dishes VALUES (329, 'Теплый салат с говядиной', 370, 1000, 26, 1, 0, '', 1, 1, 926, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (330, 'Салат Бригантина', 350, 1000, 26, 2, 0, '', 1, 1, 901, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (331, 'Цезарь с креветками', 400, 1000, 26, 3, 0, '', 1, 1, 640, 1, 250, 'Пиццы', 'Салаты', 'https://2d-recept.com/wp-content/uploads/2019/05/salat-cezar-s-krevetkami-foto.jpg', false);
INSERT INTO public.dishes VALUES (332, 'Салат Винегрет', 360, 1, 26, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 270, 1, 250, 'Универсальное', 'Салаты', 'https://upload.wikimedia.org/wikipedia/commons/a/a7/Vinegret_cleaned.jpg', false);
INSERT INTO public.dishes VALUES (333, 'Борщ', 250, 1000, 26, 0, 1, '', 1, 1, 873, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128261_m650_c4d2c43bb3f55ec4004d87459556a79120b53b6e2528c352a439cd282ae22f0e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (334, 'Сборная солянка мясная', 350, 1000, 26, 1, 1, '', 1, 1, 335, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128262_m650_2c8a69834bd188549d9dad91e69845b4a67c58a2ca50f19d755ca9a191cd1299.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (335, 'Суп-лапша домашняя', 200, 1000, 26, 2, 1, '', 1, 1, 621, 1, 300, 'Пиццы', 'Супы', 'https://www.maggi.ru/data/images/recept/img640x500/recept_4495_czgj.jpg', false);
INSERT INTO public.dishes VALUES (336, 'Картофельный суп', 200, 1000, 26, 3, 1, '', 1, 1, 743, 1, 300, 'Пиццы', 'Супы', 'https://s1.eda.ru/StaticContent/Photos/120131084848/120213174933/p_O.jpg', false);
INSERT INTO public.dishes VALUES (337, 'Синнабон', 120, 1000, 26, 0, 2, '', 1, 1, 777, 1, 150, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334172939_m650_02814baf09423cd3bd412593eb3dfab2ecd2316537045b96989be181002b2ae9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (338, 'Трайфл Сникерс', 250, 1000, 26, 1, 2, '', 1, 1, 535, 1, 180, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333386700_m650_6cb3cf5e3c893441037dc08ab87054615e947637dcbaf3e6849c642aa1aab7c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (339, 'Кейк попс', 50, 1000, 26, 2, 2, '', 1, 1, 901, 1, 50, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334173008_m650_7c7a29300b7b1a11be154554e2120641a75e49ba543c884cc5ebdb060f682918.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (340, 'Шоколадный фондан', 180, 1000, 26, 3, 2, '', 1, 1, 213, 1, 150, 'Универсальное', 'Десерты', 'https://www.gastronom.ru/binfiles/images/20200121/bf504e3d.jpg', false);
INSERT INTO public.dishes VALUES (341, 'Рождественский кекс', 350, 1, 27, 4, 0, '', 1, 1, 820, 1, 160, 'Универсальное', 'Путешествие в зиму', 'https://www.gastronom.ru/binfiles/images/20151221/b3c8862d.jpg', false);
INSERT INTO public.dishes VALUES (342, 'Блинчики с икрой', 370, 1000, 27, 0, 0, '', 1, 1, 848, 1, 130, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336390947_m650_60910237c1eb8bd3ef90a2d181c4d679663e37c1358fd4586a0e77c75e217acf.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (343, 'Имбирно-медовый пряник', 180, 1000, 27, 1, 0, '', 1, 1, 928, 1, 80, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391648_m650_ad26bba2487253cb6864ab0b3d1d8f3108088a1de250d4ac93e732148cbd2f5f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (344, 'Какао пряничное', 220, 1000, 27, 2, 0, '', 1, 1, 648, 1, 300, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391796_m650_24b78286ca08ad2c2f1f1ed3bcc700d3435092f87021c4fb31d2f8b74da8af37.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (345, 'Сельдь под шубой', 350, 1000, 27, 3, 0, '', 1, 1, 987, 1, 260, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391001_m650_f1c5402c3ed4959fa51d2b82cc7b3538ee70353e835047281033d8cefa75f18b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (346, 'Капучино', 220, 1, 27, 4, 0, '', 1, 1, 78, 1, 300, 'Универсальное', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (347, 'Ролл Цезарь', 270, 1000, 27, 0, 1, '', 1, 1, 798, 1, 190, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490974_m650_a97f85e487af6dffb2a2af88f7fdc71d2d9ee2354ec24601c56a58c0e9540b71.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (348, 'Ролл Филадельфия', 429, 1000, 27, 1, 1, '', 1, 1, 425, 1, 600, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263431_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (349, 'Ролл с ростбифом', 350, 1000, 27, 2, 1, '', 1, 1, 772, 1, 150, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/15838/334459325_m650_7388a0b3c360c2df62ce1525a0053de791fd4cfb1aa70726fa0ea2724839c8ca.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (350, 'Клаб-сэндвич', 370, 1000, 27, 3, 1, '', 1, 1, 770, 1, 300, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490973_m650_ef5daba5a610ee2eabd16c6b8668295390ff4a214e1ffc78a1cc1bfba492097e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (351, 'Мохито', 330, 1000, 27, 0, 2, '', 1, 1, 56, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490992_m650_e3496f99e11ebf227a739e1b89926dc31cfe30495eeedf9424d7919963dbc982.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (352, 'Лимонад Домашний', 270, 1000, 27, 1, 2, '', 1, 1, 230, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490988_m650_fb2879f40902be792dd258eaee85ee9fa0fcffcf228aae32db7b4a9dfcf24a70.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (353, '7Up', 150, 1000, 27, 2, 2, '', 1, 1, 265, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/324535611_m650_7572828939a42371b8fd488d1b0d998755e8148aa2db61abdabb8c439ffd3b39.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (354, 'Пепси', 150, 1000, 27, 3, 2, '', 1, 1, 255, 1, 250, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263484_m650_9d9f9f789f6abc645d15bae17c5c3084a79df0143183553744f4641c8c7afbbd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (355, 'Шашлык из свиной мякоти', 350, 1000, 28, 4, 0, 'Подаётся с луком и зеленью', 1, 1, 649, 1, 200, 'Универсальное', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128163_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (356, 'Шашлык из стейка семги', 540, 1000, 28, 0, 0, '', 1, 1, 1500, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128173_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (357, 'Шашлык из мякоти баранины', 570, 1000, 28, 1, 0, '', 1, 1, 924, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128168_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (358, 'Шашлык из куриного филе', 330, 1000, 28, 2, 0, '', 1, 1, 1234, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128165_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (359, 'Люля-кебаб из курицы', 230, 1000, 28, 3, 0, '', 1, 1, 996, 1, 180, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128170_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (360, 'Овощи на мангале', 230, 1000, 28, 0, 1, '', 1, 1, 750, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128174_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (361, 'Картофель с салом', 160, 1000, 28, 1, 1, '', 1, 1, 526, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128176_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (362, 'Грибы на мангале', 230, 1000, 28, 2, 1, '', 1, 1, 233, 1, 150, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128175_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (363, 'Лимонад Натахтари дюшес', 110, 1000, 28, 0, 2, '', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128181_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (364, 'Coca-cola', 110, 1000, 28, 1, 2, '', 1, 1, 110, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128183_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (365, 'Лимонад Натахтари тархун', 110, 1000, 28, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128182_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (366, 'Лимонад Натахтари барбарис', 110, 1000, 28, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://fermer-63.ru/upload/iblock/10e/10ed883d25ae6ffc445c96c519394779.jpg', false);
INSERT INTO public.dishes VALUES (367, 'Греческий', 230, 150, 29, 4, 0, '', 1, 1, 520, 1, 5, 'Универсальное', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/15838/328171056_m650_cfcbfe9ab6ebf66ca620792736fc485d3c0a91072b9143258e3c902a05717650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (368, 'Салат Крабовичок', 170, 1000, 29, 0, 0, '', 1, 1, 979, 1, 150, 'Пиццы', 'Салаты', 'https://www.povarenok.ru/data/cache/2013oct/16/04/532964_47569-710x550x.jpg', false);
INSERT INTO public.dishes VALUES (369, 'Теплый салат с говядиной', 370, 1000, 29, 1, 0, '', 1, 1, 926, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (370, 'Салат Бригантина', 350, 1000, 29, 2, 0, '', 1, 1, 901, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (371, 'Цезарь с креветками', 400, 1000, 29, 3, 0, '', 1, 1, 640, 1, 250, 'Пиццы', 'Салаты', 'https://2d-recept.com/wp-content/uploads/2019/05/salat-cezar-s-krevetkami-foto.jpg', false);
INSERT INTO public.dishes VALUES (372, 'Салат Винегрет', 360, 1, 29, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 270, 1, 250, 'Универсальное', 'Салаты', 'https://upload.wikimedia.org/wikipedia/commons/a/a7/Vinegret_cleaned.jpg', false);
INSERT INTO public.dishes VALUES (373, 'Борщ', 250, 1000, 29, 0, 1, '', 1, 1, 873, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128261_m650_c4d2c43bb3f55ec4004d87459556a79120b53b6e2528c352a439cd282ae22f0e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (374, 'Сборная солянка мясная', 350, 1000, 29, 1, 1, '', 1, 1, 335, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128262_m650_2c8a69834bd188549d9dad91e69845b4a67c58a2ca50f19d755ca9a191cd1299.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (375, 'Суп-лапша домашняя', 200, 1000, 29, 2, 1, '', 1, 1, 621, 1, 300, 'Пиццы', 'Супы', 'https://www.maggi.ru/data/images/recept/img640x500/recept_4495_czgj.jpg', false);
INSERT INTO public.dishes VALUES (376, 'Картофельный суп', 200, 1000, 29, 3, 1, '', 1, 1, 743, 1, 300, 'Пиццы', 'Супы', 'https://s1.eda.ru/StaticContent/Photos/120131084848/120213174933/p_O.jpg', false);
INSERT INTO public.dishes VALUES (377, 'Синнабон', 120, 1000, 29, 0, 2, '', 1, 1, 777, 1, 150, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334172939_m650_02814baf09423cd3bd412593eb3dfab2ecd2316537045b96989be181002b2ae9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (378, 'Трайфл Сникерс', 250, 1000, 29, 1, 2, '', 1, 1, 535, 1, 180, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333386700_m650_6cb3cf5e3c893441037dc08ab87054615e947637dcbaf3e6849c642aa1aab7c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (379, 'Кейк попс', 50, 1000, 29, 2, 2, '', 1, 1, 901, 1, 50, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334173008_m650_7c7a29300b7b1a11be154554e2120641a75e49ba543c884cc5ebdb060f682918.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (380, 'Шоколадный фондан', 180, 1000, 29, 3, 2, '', 1, 1, 213, 1, 150, 'Универсальное', 'Десерты', 'https://www.gastronom.ru/binfiles/images/20200121/bf504e3d.jpg', false);
INSERT INTO public.dishes VALUES (381, 'Шаурма домашняя', 199, 1, 30, 2, 0, 'У этого блюда есть абсолютно всё', 1, 1, 602, 1, 330, 'Универсальное', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542260_m650_b2b212d56f1bc91abe68636af87f1e185e719dfdd01409b7888c3c2884e71d04.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (382, 'Шаурма с тунцом', 229, 1000, 30, 0, 0, '', 1, 1, 697, 1, 260, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542767_m650_9321049bda74a7ef4ad30c7d93ccaf8b8ef4fceb3ce09e67b57e9bf5d42bc897.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (383, 'Бургер с куриной котлетой', 229, 1000, 30, 1, 0, '', 1, 1, 407, 1, 200, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/335894906_m650_e5d32a976fbbc466608c4a9625ec1d73e9034b37d6e9611f2bc4a5eaa427e36f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (384, 'Американо', 150, 1000, 30, 0, 1, '', 1, 1, 458, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263353_m650_f4265fc18ab291c9b77d274364519df8c59d3c398facdc47eaf993cc74f32994.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (385, 'Капучино', 130, 1000, 30, 1, 1, '', 1, 1, 927, 1, 200, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (386, 'Латте', 170, 1000, 30, 2, 1, '', 1, 1, 907, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490932_m650_cb85b5a866a3afaa2c1f08b52b425f6e78fb0891ee033676524d6d90b4c6b2fd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (387, 'Фильтр-кофе', 150, 1000, 30, 3, 1, '', 1, 1, 207, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263356_m650_750ecfdcd43338312299910b4a7777c0466586c170bdef71faa0be4dce816437.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (388, 'Блин Ватрушка', 165, 1000, 30, 0, 2, '', 1, 1, 342, 1, 196, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485249_m650_861f7ff88a9073dfcb8fec2d28d8bc72ea710f7caa343417ae6667a1e516630d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (389, 'Блин с сахаром', 87, 1000, 30, 1, 2, '', 1, 1, 335, 1, 124, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485888_m650_ebfd220955c6f29d73ae8df3ed98ed9a406101b3cf030df5b2573959addaad3d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (390, 'Блин с шоколадным кремом', 151, 1000, 30, 2, 2, '', 1, 1, 327, 1, 166, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485270_m650_36ced98a972f785c015b545212a58d23b3cce201603a556b5c998dbd53d4a8ce.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (391, 'Блин с вишневым вареньем', 93, 1000, 30, 3, 2, '', 1, 1, 292, 1, 126, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485286_m650_7334048751f2aa703a75266d1503ac44e8dd2929a429858ef468c59982d5a45e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (392, 'Шаурма домашняя', 199, 1, 31, 2, 0, 'У этого блюда есть абсолютно всё', 1, 1, 602, 1, 330, 'Универсальное', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542260_m650_b2b212d56f1bc91abe68636af87f1e185e719dfdd01409b7888c3c2884e71d04.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (393, 'Шаурма с тунцом', 229, 1000, 31, 0, 0, '', 1, 1, 697, 1, 260, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542767_m650_9321049bda74a7ef4ad30c7d93ccaf8b8ef4fceb3ce09e67b57e9bf5d42bc897.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (394, 'Бургер с куриной котлетой', 229, 1000, 31, 1, 0, '', 1, 1, 407, 1, 200, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/335894906_m650_e5d32a976fbbc466608c4a9625ec1d73e9034b37d6e9611f2bc4a5eaa427e36f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (395, 'Американо', 150, 1000, 31, 0, 1, '', 1, 1, 458, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263353_m650_f4265fc18ab291c9b77d274364519df8c59d3c398facdc47eaf993cc74f32994.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (396, 'Капучино', 130, 1000, 31, 1, 1, '', 1, 1, 927, 1, 200, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (397, 'Латте', 170, 1000, 31, 2, 1, '', 1, 1, 907, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490932_m650_cb85b5a866a3afaa2c1f08b52b425f6e78fb0891ee033676524d6d90b4c6b2fd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (398, 'Фильтр-кофе', 150, 1000, 31, 3, 1, '', 1, 1, 207, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263356_m650_750ecfdcd43338312299910b4a7777c0466586c170bdef71faa0be4dce816437.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (399, 'Блин Ватрушка', 165, 1000, 31, 0, 2, '', 1, 1, 342, 1, 196, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485249_m650_861f7ff88a9073dfcb8fec2d28d8bc72ea710f7caa343417ae6667a1e516630d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (400, 'Блин с сахаром', 87, 1000, 31, 1, 2, '', 1, 1, 335, 1, 124, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485888_m650_ebfd220955c6f29d73ae8df3ed98ed9a406101b3cf030df5b2573959addaad3d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (401, 'Блин с шоколадным кремом', 151, 1000, 31, 2, 2, '', 1, 1, 327, 1, 166, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485270_m650_36ced98a972f785c015b545212a58d23b3cce201603a556b5c998dbd53d4a8ce.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (402, 'Блин с вишневым вареньем', 93, 1000, 31, 3, 2, '', 1, 1, 292, 1, 126, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485286_m650_7334048751f2aa703a75266d1503ac44e8dd2929a429858ef468c59982d5a45e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (403, 'Пицца Пепперони', 470, 1, 32, 4, 0, '', 1, 1, 952, 1, 420, 'Универсальное', 'Пицца', 'https://s1.eda.ru/StaticContent/Photos/120131085053/171027192707/p_O.jpg', false);
INSERT INTO public.dishes VALUES (404, 'Пицца Ассорти', 429, 1000, 32, 0, 0, '', 1, 1, 660, 1, 600, 'Пиццы', 'Пиццы', 'https://sp-ao.shortpixel.ai/client/q_lossless,ret_img/https://cookery.site/wp-content/uploads/2019/09/photo_92647025-1-678x381.jpg', false);
INSERT INTO public.dishes VALUES (405, 'Пицца Морская', 800, 1000, 32, 1, 0, '', 1, 1, 584, 1, 600, 'Пиццы', 'Пиццы', 'https://pizzarini.info/wp-content/uploads/2018/03/pitstsa-morskaya-1-1.jpg', false);
INSERT INTO public.dishes VALUES (406, 'Пицца Карбонара', 540, 1000, 32, 2, 0, '', 1, 1, 581, 1, 420, 'Пиццы', 'Пиццы', 'https://static.1000.menu/img/content/13174/picca-karbonara_1432215044_0_max.jpg', false);
INSERT INTO public.dishes VALUES (407, 'Пицца Карначина', 590, 1000, 32, 3, 0, '', 1, 1, 895, 1, 540, 'Пиццы', 'Пиццы', 'https://cdnn21.img.ria.ru/images/98976/61/989766135_0:100:2000:1233_1920x0_80_0_0_4a3e7af4d4ab43307a68343c059cc57d.jpg', false);
INSERT INTO public.dishes VALUES (408, 'Пицца 4 сезона', 590, 1, 32, 4, 0, '', 1, 1, 952, 1, 580, 'Универсальное', 'Пицца', 'https://www.citypizza.ru/upload/img/shop/pizza/BORT/resize/4-%D1%81%D0%B5%D0%B7%D0%BE%D0%BD%D0%B0-%D0%B1%D0%BE%D1%80%D1%82-listThumb.jpg', false);
INSERT INTO public.dishes VALUES (435, 'Сельдь под шубой', 350, 1000, 34, 3, 0, '', 1, 1, 987, 1, 260, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391001_m650_f1c5402c3ed4959fa51d2b82cc7b3538ee70353e835047281033d8cefa75f18b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (409, 'Ролл с угрем', 180, 1000, 32, 0, 1, '', 1, 1, 469, 1, 120, 'Пиццы', 'Роллы', 'https://lh3.googleusercontent.com/proxy/uqHy2GczztN6r4P_eH2lySHzJWY57YsPW-vhEuU9GF46-i2_mf8qBESx4VDgPk4Vck5v7HTfIzpS13PG3puLRnbGfCMw-z4qYPFkKjiNMLiabJhV-7R9ncW31eb2X9gCBeywUfBvapgC9v8KGkNYQakwrbPtaiMZb8__LAvccU9bbP7NoV2zFMDuvM8ZDr9FXqP5wlQHHqDOngrSSISAbFkg8I2uIeBsYggTPQayCdO6D5NhwNf9I8w6-ZLnpLRvTScL5Hg4ifyw_7ARlpDmw2yfx0Ypjw', false);
INSERT INTO public.dishes VALUES (410, 'Дыхание дракона', 360, 1000, 32, 1, 1, '', 1, 1, 356, 1, 120, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011355_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (411, 'Горячий краб', 290, 1000, 32, 2, 1, '', 1, 1, 378, 1, 320, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011356_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (412, 'Запеченный угорь', 330, 1000, 32, 3, 1, '', 1, 1, 654, 1, 200, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011359_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (413, 'Сок в ассортименте', 100, 1000, 32, 0, 2, '', 1, 1, 102, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664250_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (414, 'Coca-cola', 65, 1000, 32, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664245_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (415, 'Fanta', 60, 1000, 32, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082269_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (416, 'Sprite', 65, 1000, 32, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082268_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (417, 'Пицца Пепперони', 470, 1, 33, 4, 0, '', 1, 1, 952, 1, 420, 'Универсальное', 'Пицца', 'https://s1.eda.ru/StaticContent/Photos/120131085053/171027192707/p_O.jpg', false);
INSERT INTO public.dishes VALUES (418, 'Пицца Ассорти', 429, 1000, 33, 0, 0, '', 1, 1, 660, 1, 600, 'Пиццы', 'Пиццы', 'https://sp-ao.shortpixel.ai/client/q_lossless,ret_img/https://cookery.site/wp-content/uploads/2019/09/photo_92647025-1-678x381.jpg', false);
INSERT INTO public.dishes VALUES (419, 'Пицца Морская', 800, 1000, 33, 1, 0, '', 1, 1, 584, 1, 600, 'Пиццы', 'Пиццы', 'https://pizzarini.info/wp-content/uploads/2018/03/pitstsa-morskaya-1-1.jpg', false);
INSERT INTO public.dishes VALUES (420, 'Пицца Карбонара', 540, 1000, 33, 2, 0, '', 1, 1, 581, 1, 420, 'Пиццы', 'Пиццы', 'https://static.1000.menu/img/content/13174/picca-karbonara_1432215044_0_max.jpg', false);
INSERT INTO public.dishes VALUES (421, 'Пицца Карначина', 590, 1000, 33, 3, 0, '', 1, 1, 895, 1, 540, 'Пиццы', 'Пиццы', 'https://cdnn21.img.ria.ru/images/98976/61/989766135_0:100:2000:1233_1920x0_80_0_0_4a3e7af4d4ab43307a68343c059cc57d.jpg', false);
INSERT INTO public.dishes VALUES (422, 'Пицца 4 сезона', 590, 1, 33, 4, 0, '', 1, 1, 952, 1, 580, 'Универсальное', 'Пицца', 'https://www.citypizza.ru/upload/img/shop/pizza/BORT/resize/4-%D1%81%D0%B5%D0%B7%D0%BE%D0%BD%D0%B0-%D0%B1%D0%BE%D1%80%D1%82-listThumb.jpg', false);
INSERT INTO public.dishes VALUES (423, 'Ролл с угрем', 180, 1000, 33, 0, 1, '', 1, 1, 469, 1, 120, 'Пиццы', 'Роллы', 'https://lh3.googleusercontent.com/proxy/uqHy2GczztN6r4P_eH2lySHzJWY57YsPW-vhEuU9GF46-i2_mf8qBESx4VDgPk4Vck5v7HTfIzpS13PG3puLRnbGfCMw-z4qYPFkKjiNMLiabJhV-7R9ncW31eb2X9gCBeywUfBvapgC9v8KGkNYQakwrbPtaiMZb8__LAvccU9bbP7NoV2zFMDuvM8ZDr9FXqP5wlQHHqDOngrSSISAbFkg8I2uIeBsYggTPQayCdO6D5NhwNf9I8w6-ZLnpLRvTScL5Hg4ifyw_7ARlpDmw2yfx0Ypjw', false);
INSERT INTO public.dishes VALUES (424, 'Дыхание дракона', 360, 1000, 33, 1, 1, '', 1, 1, 356, 1, 120, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011355_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (425, 'Горячий краб', 290, 1000, 33, 2, 1, '', 1, 1, 378, 1, 320, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011356_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (426, 'Запеченный угорь', 330, 1000, 33, 3, 1, '', 1, 1, 654, 1, 200, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011359_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (427, 'Сок в ассортименте', 100, 1000, 33, 0, 2, '', 1, 1, 102, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664250_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (428, 'Coca-cola', 65, 1000, 33, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664245_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (429, 'Fanta', 60, 1000, 33, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082269_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (430, 'Sprite', 65, 1000, 33, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082268_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (431, 'Рождественский кекс', 350, 1, 34, 4, 0, '', 1, 1, 820, 1, 160, 'Универсальное', 'Путешествие в зиму', 'https://www.gastronom.ru/binfiles/images/20151221/b3c8862d.jpg', false);
INSERT INTO public.dishes VALUES (432, 'Блинчики с икрой', 370, 1000, 34, 0, 0, '', 1, 1, 848, 1, 130, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336390947_m650_60910237c1eb8bd3ef90a2d181c4d679663e37c1358fd4586a0e77c75e217acf.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (433, 'Имбирно-медовый пряник', 180, 1000, 34, 1, 0, '', 1, 1, 928, 1, 80, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391648_m650_ad26bba2487253cb6864ab0b3d1d8f3108088a1de250d4ac93e732148cbd2f5f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (434, 'Какао пряничное', 220, 1000, 34, 2, 0, '', 1, 1, 648, 1, 300, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391796_m650_24b78286ca08ad2c2f1f1ed3bcc700d3435092f87021c4fb31d2f8b74da8af37.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (436, 'Капучино', 220, 1, 34, 4, 0, '', 1, 1, 78, 1, 300, 'Универсальное', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (437, 'Ролл Цезарь', 270, 1000, 34, 0, 1, '', 1, 1, 798, 1, 190, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490974_m650_a97f85e487af6dffb2a2af88f7fdc71d2d9ee2354ec24601c56a58c0e9540b71.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (438, 'Ролл Филадельфия', 429, 1000, 34, 1, 1, '', 1, 1, 425, 1, 600, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263431_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (439, 'Ролл с ростбифом', 350, 1000, 34, 2, 1, '', 1, 1, 772, 1, 150, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/15838/334459325_m650_7388a0b3c360c2df62ce1525a0053de791fd4cfb1aa70726fa0ea2724839c8ca.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (440, 'Клаб-сэндвич', 370, 1000, 34, 3, 1, '', 1, 1, 770, 1, 300, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490973_m650_ef5daba5a610ee2eabd16c6b8668295390ff4a214e1ffc78a1cc1bfba492097e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (441, 'Мохито', 330, 1000, 34, 0, 2, '', 1, 1, 56, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490992_m650_e3496f99e11ebf227a739e1b89926dc31cfe30495eeedf9424d7919963dbc982.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (442, 'Лимонад Домашний', 270, 1000, 34, 1, 2, '', 1, 1, 230, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490988_m650_fb2879f40902be792dd258eaee85ee9fa0fcffcf228aae32db7b4a9dfcf24a70.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (443, '7Up', 150, 1000, 34, 2, 2, '', 1, 1, 265, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/324535611_m650_7572828939a42371b8fd488d1b0d998755e8148aa2db61abdabb8c439ffd3b39.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (444, 'Пепси', 150, 1000, 34, 3, 2, '', 1, 1, 255, 1, 250, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263484_m650_9d9f9f789f6abc645d15bae17c5c3084a79df0143183553744f4641c8c7afbbd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (445, 'Шашлык из свиной мякоти', 350, 1000, 35, 4, 0, 'Подаётся с луком и зеленью', 1, 1, 649, 1, 200, 'Универсальное', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128163_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (446, 'Шашлык из стейка семги', 540, 1000, 35, 0, 0, '', 1, 1, 1500, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128173_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (447, 'Шашлык из мякоти баранины', 570, 1000, 35, 1, 0, '', 1, 1, 924, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128168_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (448, 'Шашлык из куриного филе', 330, 1000, 35, 2, 0, '', 1, 1, 1234, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128165_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (449, 'Люля-кебаб из курицы', 230, 1000, 35, 3, 0, '', 1, 1, 996, 1, 180, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128170_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (450, 'Овощи на мангале', 230, 1000, 35, 0, 1, '', 1, 1, 750, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128174_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (451, 'Картофель с салом', 160, 1000, 35, 1, 1, '', 1, 1, 526, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128176_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (452, 'Грибы на мангале', 230, 1000, 35, 2, 1, '', 1, 1, 233, 1, 150, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128175_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (453, 'Лимонад Натахтари дюшес', 110, 1000, 35, 0, 2, '', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128181_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (454, 'Coca-cola', 110, 1000, 35, 1, 2, '', 1, 1, 110, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128183_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (455, 'Лимонад Натахтари тархун', 110, 1000, 35, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128182_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (456, 'Лимонад Натахтари барбарис', 110, 1000, 35, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://fermer-63.ru/upload/iblock/10e/10ed883d25ae6ffc445c96c519394779.jpg', false);
INSERT INTO public.dishes VALUES (457, 'Греческий', 230, 150, 36, 4, 0, '', 1, 1, 520, 1, 5, 'Универсальное', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/15838/328171056_m650_cfcbfe9ab6ebf66ca620792736fc485d3c0a91072b9143258e3c902a05717650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (458, 'Салат Крабовичок', 170, 1000, 36, 0, 0, '', 1, 1, 979, 1, 150, 'Пиццы', 'Салаты', 'https://www.povarenok.ru/data/cache/2013oct/16/04/532964_47569-710x550x.jpg', false);
INSERT INTO public.dishes VALUES (459, 'Теплый салат с говядиной', 370, 1000, 36, 1, 0, '', 1, 1, 926, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (460, 'Салат Бригантина', 350, 1000, 36, 2, 0, '', 1, 1, 901, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (461, 'Цезарь с креветками', 400, 1000, 36, 3, 0, '', 1, 1, 640, 1, 250, 'Пиццы', 'Салаты', 'https://2d-recept.com/wp-content/uploads/2019/05/salat-cezar-s-krevetkami-foto.jpg', false);
INSERT INTO public.dishes VALUES (462, 'Салат Винегрет', 360, 1, 36, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 270, 1, 250, 'Универсальное', 'Салаты', 'https://upload.wikimedia.org/wikipedia/commons/a/a7/Vinegret_cleaned.jpg', false);
INSERT INTO public.dishes VALUES (463, 'Борщ', 250, 1000, 36, 0, 1, '', 1, 1, 873, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128261_m650_c4d2c43bb3f55ec4004d87459556a79120b53b6e2528c352a439cd282ae22f0e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (464, 'Сборная солянка мясная', 350, 1000, 36, 1, 1, '', 1, 1, 335, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128262_m650_2c8a69834bd188549d9dad91e69845b4a67c58a2ca50f19d755ca9a191cd1299.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (465, 'Суп-лапша домашняя', 200, 1000, 36, 2, 1, '', 1, 1, 621, 1, 300, 'Пиццы', 'Супы', 'https://www.maggi.ru/data/images/recept/img640x500/recept_4495_czgj.jpg', false);
INSERT INTO public.dishes VALUES (466, 'Картофельный суп', 200, 1000, 36, 3, 1, '', 1, 1, 743, 1, 300, 'Пиццы', 'Супы', 'https://s1.eda.ru/StaticContent/Photos/120131084848/120213174933/p_O.jpg', false);
INSERT INTO public.dishes VALUES (467, 'Синнабон', 120, 1000, 36, 0, 2, '', 1, 1, 777, 1, 150, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334172939_m650_02814baf09423cd3bd412593eb3dfab2ecd2316537045b96989be181002b2ae9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (468, 'Трайфл Сникерс', 250, 1000, 36, 1, 2, '', 1, 1, 535, 1, 180, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333386700_m650_6cb3cf5e3c893441037dc08ab87054615e947637dcbaf3e6849c642aa1aab7c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (469, 'Кейк попс', 50, 1000, 36, 2, 2, '', 1, 1, 901, 1, 50, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334173008_m650_7c7a29300b7b1a11be154554e2120641a75e49ba543c884cc5ebdb060f682918.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (470, 'Шоколадный фондан', 180, 1000, 36, 3, 2, '', 1, 1, 213, 1, 150, 'Универсальное', 'Десерты', 'https://www.gastronom.ru/binfiles/images/20200121/bf504e3d.jpg', false);
INSERT INTO public.dishes VALUES (471, 'Ролл с лососем', 219, 1000, 37, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (472, 'Ролл с огурцом', 149, 1000, 37, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (473, 'Крабик Hot запеченный мини ролл', 189, 1000, 37, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (474, 'Лосось унаги', 319, 1000, 37, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (475, 'Ролл Аяши запеченный', 259, 1000, 37, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (476, 'Ролл Горячий лосось запеченный', 369, 1, 37, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (477, 'Вегетарианский Вок', 289, 1000, 37, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (478, 'Классический Вок', 349, 1000, 37, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (479, 'Сытный Вок', 369, 1000, 37, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (480, 'По-Китайский Вок', 349, 1000, 37, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (481, 'BonAqua', 79, 1000, 37, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (482, 'Coca-cola', 99, 1000, 37, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (531, 'Fanta', 60, 1000, 41, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (483, 'Морс из клюквы', 99, 1000, 37, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (484, 'Морс из черной смородины', 99, 1000, 37, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (485, 'Рождественский кекс', 350, 1, 38, 4, 0, '', 1, 1, 820, 1, 160, 'Универсальное', 'Путешествие в зиму', 'https://www.gastronom.ru/binfiles/images/20151221/b3c8862d.jpg', false);
INSERT INTO public.dishes VALUES (486, 'Блинчики с икрой', 370, 1000, 38, 0, 0, '', 1, 1, 848, 1, 130, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336390947_m650_60910237c1eb8bd3ef90a2d181c4d679663e37c1358fd4586a0e77c75e217acf.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (487, 'Имбирно-медовый пряник', 180, 1000, 38, 1, 0, '', 1, 1, 928, 1, 80, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391648_m650_ad26bba2487253cb6864ab0b3d1d8f3108088a1de250d4ac93e732148cbd2f5f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (488, 'Какао пряничное', 220, 1000, 38, 2, 0, '', 1, 1, 648, 1, 300, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391796_m650_24b78286ca08ad2c2f1f1ed3bcc700d3435092f87021c4fb31d2f8b74da8af37.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (489, 'Сельдь под шубой', 350, 1000, 38, 3, 0, '', 1, 1, 987, 1, 260, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391001_m650_f1c5402c3ed4959fa51d2b82cc7b3538ee70353e835047281033d8cefa75f18b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (490, 'Капучино', 220, 1, 38, 4, 0, '', 1, 1, 78, 1, 300, 'Универсальное', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (491, 'Ролл Цезарь', 270, 1000, 38, 0, 1, '', 1, 1, 798, 1, 190, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490974_m650_a97f85e487af6dffb2a2af88f7fdc71d2d9ee2354ec24601c56a58c0e9540b71.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (492, 'Ролл Филадельфия', 429, 1000, 38, 1, 1, '', 1, 1, 425, 1, 600, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263431_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (493, 'Ролл с ростбифом', 350, 1000, 38, 2, 1, '', 1, 1, 772, 1, 150, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/15838/334459325_m650_7388a0b3c360c2df62ce1525a0053de791fd4cfb1aa70726fa0ea2724839c8ca.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (494, 'Клаб-сэндвич', 370, 1000, 38, 3, 1, '', 1, 1, 770, 1, 300, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490973_m650_ef5daba5a610ee2eabd16c6b8668295390ff4a214e1ffc78a1cc1bfba492097e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (495, 'Мохито', 330, 1000, 38, 0, 2, '', 1, 1, 56, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490992_m650_e3496f99e11ebf227a739e1b89926dc31cfe30495eeedf9424d7919963dbc982.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (496, 'Лимонад Домашний', 270, 1000, 38, 1, 2, '', 1, 1, 230, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490988_m650_fb2879f40902be792dd258eaee85ee9fa0fcffcf228aae32db7b4a9dfcf24a70.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (497, '7Up', 150, 1000, 38, 2, 2, '', 1, 1, 265, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/324535611_m650_7572828939a42371b8fd488d1b0d998755e8148aa2db61abdabb8c439ffd3b39.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (498, 'Пепси', 150, 1000, 38, 3, 2, '', 1, 1, 255, 1, 250, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263484_m650_9d9f9f789f6abc645d15bae17c5c3084a79df0143183553744f4641c8c7afbbd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (499, 'Ролл с лососем', 219, 1000, 39, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (500, 'Ролл с огурцом', 149, 1000, 39, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (501, 'Крабик Hot запеченный мини ролл', 189, 1000, 39, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (502, 'Лосось унаги', 319, 1000, 39, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (503, 'Ролл Аяши запеченный', 259, 1000, 39, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (504, 'Ролл Горячий лосось запеченный', 369, 1, 39, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (505, 'Вегетарианский Вок', 289, 1000, 39, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (506, 'Классический Вок', 349, 1000, 39, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (507, 'Сытный Вок', 369, 1000, 39, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (508, 'По-Китайский Вок', 349, 1000, 39, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (509, 'BonAqua', 79, 1000, 39, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (510, 'Coca-cola', 99, 1000, 39, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (511, 'Морс из клюквы', 99, 1000, 39, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (512, 'Морс из черной смородины', 99, 1000, 39, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (513, 'Тако', 60, 1000, 40, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (514, 'Пряник', 70, 1000, 40, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (515, 'Чёрный бургер', 139, 1000, 40, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (516, 'Кофе', 149, 1000, 40, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (517, 'Coca-cola', 65, 1000, 40, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (518, 'Fanta', 60, 1000, 40, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (519, 'Sprite', 65, 1000, 40, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (520, 'Картошка Фри', 60, 1000, 40, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (521, 'Картошка по деревенски', 60, 1000, 40, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (522, 'Блюдо со стейком', 256, 1000, 40, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (523, 'Сёмга', 800, 1, 40, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (524, 'Цезарь', 5, 1, 40, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (525, 'Бефстроганов', 1000, 1000, 40, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (526, 'Тако', 60, 1000, 41, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (527, 'Пряник', 70, 1000, 41, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (528, 'Чёрный бургер', 139, 1000, 41, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (529, 'Кофе', 149, 1000, 41, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (530, 'Coca-cola', 65, 1000, 41, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (534, 'Картошка по деревенски', 60, 1000, 41, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (535, 'Блюдо со стейком', 256, 1000, 41, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (536, 'Сёмга', 800, 1, 41, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (537, 'Цезарь', 5, 1, 41, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (538, 'Бефстроганов', 1000, 1000, 41, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (539, 'Рождественский кекс', 350, 1, 42, 4, 0, '', 1, 1, 820, 1, 160, 'Универсальное', 'Путешествие в зиму', 'https://www.gastronom.ru/binfiles/images/20151221/b3c8862d.jpg', false);
INSERT INTO public.dishes VALUES (540, 'Блинчики с икрой', 370, 1000, 42, 0, 0, '', 1, 1, 848, 1, 130, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336390947_m650_60910237c1eb8bd3ef90a2d181c4d679663e37c1358fd4586a0e77c75e217acf.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (541, 'Имбирно-медовый пряник', 180, 1000, 42, 1, 0, '', 1, 1, 928, 1, 80, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391648_m650_ad26bba2487253cb6864ab0b3d1d8f3108088a1de250d4ac93e732148cbd2f5f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (542, 'Какао пряничное', 220, 1000, 42, 2, 0, '', 1, 1, 648, 1, 300, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391796_m650_24b78286ca08ad2c2f1f1ed3bcc700d3435092f87021c4fb31d2f8b74da8af37.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (543, 'Сельдь под шубой', 350, 1000, 42, 3, 0, '', 1, 1, 987, 1, 260, 'Пиццы', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/336391001_m650_f1c5402c3ed4959fa51d2b82cc7b3538ee70353e835047281033d8cefa75f18b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (544, 'Капучино', 220, 1, 42, 4, 0, '', 1, 1, 78, 1, 300, 'Универсальное', 'Путешествие в зиму', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (545, 'Ролл Цезарь', 270, 1000, 42, 0, 1, '', 1, 1, 798, 1, 190, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490974_m650_a97f85e487af6dffb2a2af88f7fdc71d2d9ee2354ec24601c56a58c0e9540b71.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (546, 'Ролл Филадельфия', 429, 1000, 42, 1, 1, '', 1, 1, 425, 1, 600, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263431_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (547, 'Ролл с ростбифом', 350, 1000, 42, 2, 1, '', 1, 1, 772, 1, 150, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/15838/334459325_m650_7388a0b3c360c2df62ce1525a0053de791fd4cfb1aa70726fa0ea2724839c8ca.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (548, 'Клаб-сэндвич', 370, 1000, 42, 3, 1, '', 1, 1, 770, 1, 300, 'Пиццы', 'Сэндвичи', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490973_m650_ef5daba5a610ee2eabd16c6b8668295390ff4a214e1ffc78a1cc1bfba492097e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (549, 'Мохито', 330, 1000, 42, 0, 2, '', 1, 1, 56, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490992_m650_e3496f99e11ebf227a739e1b89926dc31cfe30495eeedf9424d7919963dbc982.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (550, 'Лимонад Домашний', 270, 1000, 42, 1, 2, '', 1, 1, 230, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490988_m650_fb2879f40902be792dd258eaee85ee9fa0fcffcf228aae32db7b4a9dfcf24a70.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (551, '7Up', 150, 1000, 42, 2, 2, '', 1, 1, 265, 1, 250, 'Универсальное', 'Холодные напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/324535611_m650_7572828939a42371b8fd488d1b0d998755e8148aa2db61abdabb8c439ffd3b39.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (552, 'Пепси', 150, 1000, 42, 3, 2, '', 1, 1, 255, 1, 250, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263484_m650_9d9f9f789f6abc645d15bae17c5c3084a79df0143183553744f4641c8c7afbbd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (553, 'Тако', 60, 1000, 43, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (554, 'Пряник', 70, 1000, 43, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (555, 'Чёрный бургер', 139, 1000, 43, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (556, 'Кофе', 149, 1000, 43, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (557, 'Coca-cola', 65, 1000, 43, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (558, 'Fanta', 60, 1000, 43, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (560, 'Картошка Фри', 60, 1000, 43, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (561, 'Картошка по деревенски', 60, 1000, 43, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (562, 'Блюдо со стейком', 256, 1000, 43, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (563, 'Сёмга', 800, 1, 43, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (564, 'Цезарь', 5, 1, 43, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (565, 'Бефстроганов', 1000, 1000, 43, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (566, 'Греческий', 230, 150, 44, 4, 0, '', 1, 1, 520, 1, 5, 'Универсальное', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/15838/328171056_m650_cfcbfe9ab6ebf66ca620792736fc485d3c0a91072b9143258e3c902a05717650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (567, 'Салат Крабовичок', 170, 1000, 44, 0, 0, '', 1, 1, 979, 1, 150, 'Пиццы', 'Салаты', 'https://www.povarenok.ru/data/cache/2013oct/16/04/532964_47569-710x550x.jpg', false);
INSERT INTO public.dishes VALUES (568, 'Теплый салат с говядиной', 370, 1000, 44, 1, 0, '', 1, 1, 926, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (569, 'Салат Бригантина', 350, 1000, 44, 2, 0, '', 1, 1, 901, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (570, 'Цезарь с креветками', 400, 1000, 44, 3, 0, '', 1, 1, 640, 1, 250, 'Пиццы', 'Салаты', 'https://2d-recept.com/wp-content/uploads/2019/05/salat-cezar-s-krevetkami-foto.jpg', false);
INSERT INTO public.dishes VALUES (571, 'Салат Винегрет', 360, 1, 44, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 270, 1, 250, 'Универсальное', 'Салаты', 'https://upload.wikimedia.org/wikipedia/commons/a/a7/Vinegret_cleaned.jpg', false);
INSERT INTO public.dishes VALUES (572, 'Борщ', 250, 1000, 44, 0, 1, '', 1, 1, 873, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128261_m650_c4d2c43bb3f55ec4004d87459556a79120b53b6e2528c352a439cd282ae22f0e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (573, 'Сборная солянка мясная', 350, 1000, 44, 1, 1, '', 1, 1, 335, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128262_m650_2c8a69834bd188549d9dad91e69845b4a67c58a2ca50f19d755ca9a191cd1299.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (574, 'Суп-лапша домашняя', 200, 1000, 44, 2, 1, '', 1, 1, 621, 1, 300, 'Пиццы', 'Супы', 'https://www.maggi.ru/data/images/recept/img640x500/recept_4495_czgj.jpg', false);
INSERT INTO public.dishes VALUES (575, 'Картофельный суп', 200, 1000, 44, 3, 1, '', 1, 1, 743, 1, 300, 'Пиццы', 'Супы', 'https://s1.eda.ru/StaticContent/Photos/120131084848/120213174933/p_O.jpg', false);
INSERT INTO public.dishes VALUES (576, 'Синнабон', 120, 1000, 44, 0, 2, '', 1, 1, 777, 1, 150, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334172939_m650_02814baf09423cd3bd412593eb3dfab2ecd2316537045b96989be181002b2ae9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (577, 'Трайфл Сникерс', 250, 1000, 44, 1, 2, '', 1, 1, 535, 1, 180, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333386700_m650_6cb3cf5e3c893441037dc08ab87054615e947637dcbaf3e6849c642aa1aab7c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (578, 'Кейк попс', 50, 1000, 44, 2, 2, '', 1, 1, 901, 1, 50, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334173008_m650_7c7a29300b7b1a11be154554e2120641a75e49ba543c884cc5ebdb060f682918.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (579, 'Шоколадный фондан', 180, 1000, 44, 3, 2, '', 1, 1, 213, 1, 150, 'Универсальное', 'Десерты', 'https://www.gastronom.ru/binfiles/images/20200121/bf504e3d.jpg', false);
INSERT INTO public.dishes VALUES (580, 'Тако', 60, 1000, 45, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (581, 'Пряник', 70, 1000, 45, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (582, 'Чёрный бургер', 139, 1000, 45, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (583, 'Кофе', 149, 1000, 45, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (584, 'Coca-cola', 65, 1000, 45, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (585, 'Fanta', 60, 1000, 45, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (586, 'Sprite', 65, 1000, 45, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (587, 'Картошка Фри', 60, 1000, 45, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (588, 'Картошка по деревенски', 60, 1000, 45, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (589, 'Блюдо со стейком', 256, 1000, 45, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (590, 'Сёмга', 800, 1, 45, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (591, 'Цезарь', 5, 1, 45, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (592, 'Бефстроганов', 1000, 1000, 45, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (593, 'Тако', 60, 1000, 46, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (594, 'Пряник', 70, 1000, 46, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (595, 'Чёрный бургер', 139, 1000, 46, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (596, 'Кофе', 149, 1000, 46, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (597, 'Coca-cola', 65, 1000, 46, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (598, 'Fanta', 60, 1000, 46, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (599, 'Sprite', 65, 1000, 46, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (600, 'Картошка Фри', 60, 1000, 46, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (601, 'Картошка по деревенски', 60, 1000, 46, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (602, 'Блюдо со стейком', 256, 1000, 46, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (603, 'Сёмга', 800, 1, 46, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (604, 'Цезарь', 5, 1, 46, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (605, 'Бефстроганов', 1000, 1000, 46, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (606, 'Ролл с лососем', 219, 1000, 47, 4, 0, '', 1, 1, 209, 1, 102, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238599_m650_539e946ca3db2be5c8b8c4cfc4cd7e660d8fb2cc656d9989f2e03f7cdcae2f47.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (607, 'Ролл с огурцом', 149, 1000, 47, 0, 0, '', 1, 1, 524, 1, 102, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/321238601_m650_82bba223f554acff5d6d25e34ad2c05fe11061526ae33fe2c82a1277d4c24e59.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (608, 'Крабик Hot запеченный мини ролл', 189, 1000, 47, 1, 0, '', 1, 1, 760, 1, 126, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243615_m650_0b4299dcb2bee58f6a558507b8d399f110ec93bbed2364c5b1e8074dbba76621.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (609, 'Лосось унаги', 319, 1000, 47, 2, 0, '', 1, 1, 967, 1, 271, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/42724/333341614_m650_338bc009c123947c0255f084732dc5d0f20788defdceb1f19466b68e847bcc55.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (610, 'Ролл Аяши запеченный', 259, 1000, 47, 3, 0, '', 1, 1, 599, 1, 179, 'Пиццы', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617533_m650_ef7f0e290bf43bd2c56361ea2f21b4effd8b982af3bea6688f97cc43f26bc3dd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (611, 'Ролл Горячий лосось запеченный', 369, 1, 47, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 567, 1, 259, 'Универсальное', 'Мини-роллы', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617534_m650_5e8c0c911604f1979a58f9f1ff212e2a65d441eac2bb3345ad689fbb355a6e4b.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webpg', false);
INSERT INTO public.dishes VALUES (612, 'Вегетарианский Вок', 289, 1000, 47, 0, 1, '', 1, 1, 828, 1, 370, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617539_m650_6108902391e58c4d936e0771259fa5b3743896743d46308420939712614e79c0.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (613, 'Классический Вок', 349, 1000, 47, 1, 1, '', 1, 1, 512, 1, 415, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617540_m650_9599e2bb080be811e6962884e2ec472d13837dad14510384911025404d8398bc.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (614, 'Сытный Вок', 369, 1000, 47, 2, 1, '', 1, 1, 912, 1, 340, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/58031/325617545_m650_d8d9c479220ab8fef7171d9a681a7fb366691be1ac77d1aa8a69491a2e8bd17d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (615, 'По-Китайский Вок', 349, 1000, 47, 3, 1, '', 1, 1, 904, 1, 320, 'Пиццы', 'Горячее', 'https://www.delivery-club.ru/media/cms/relation_product/47930/329243629_m650_1c44e5dad116160d9b0aaa9a8239520d054b5d575f5524d6e66397ad095d169f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (616, 'BonAqua', 79, 1000, 47, 0, 2, 'Горячий, ароматный вода', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153171_m650_a62beee592242eea4214fb25420b3f171a55ecfac517d3752f174cd989554891.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (617, 'Coca-cola', 99, 1000, 47, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153172_m650_22a87eb2036482a82aa3bffe62a66b10cbb35143a5242a88fe412380cedf5fd9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (618, 'Морс из клюквы', 99, 1000, 47, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153174_m650_eb8d7f742393fe4640a80221d0489fc9d19a3bfeb62e3198f02ae300a688e161.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (619, 'Морс из черной смородины', 99, 1000, 47, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/56933/325153175_m650_39454cea676f729e6fec03ad6f6293a119a40d51fce6ff03a52745ea0d95968f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (620, 'Шашлык из свиной мякоти', 350, 1000, 48, 4, 0, 'Подаётся с луком и зеленью', 1, 1, 649, 1, 200, 'Универсальное', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128163_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (621, 'Шашлык из стейка семги', 540, 1000, 48, 0, 0, '', 1, 1, 1500, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128173_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (622, 'Шашлык из мякоти баранины', 570, 1000, 48, 1, 0, '', 1, 1, 924, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128168_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (623, 'Шашлык из куриного филе', 330, 1000, 48, 2, 0, '', 1, 1, 1234, 1, 200, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128165_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (624, 'Люля-кебаб из курицы', 230, 1000, 48, 3, 0, '', 1, 1, 996, 1, 180, 'Пиццы', 'Блюда на мангале', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128170_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (625, 'Овощи на мангале', 230, 1000, 48, 0, 1, '', 1, 1, 750, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128174_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (626, 'Картофель с салом', 160, 1000, 48, 1, 1, '', 1, 1, 526, 1, 200, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128176_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (627, 'Грибы на мангале', 230, 1000, 48, 2, 1, '', 1, 1, 233, 1, 150, 'Пиццы', 'Гарниры', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128175_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (628, 'Лимонад Натахтари дюшес', 110, 1000, 48, 0, 2, '', 1, 1, 90, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128181_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (629, 'Coca-cola', 110, 1000, 48, 1, 2, '', 1, 1, 110, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128183_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (630, 'Лимонад Натахтари тархун', 110, 1000, 48, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/66145/330128182_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (631, 'Лимонад Натахтари барбарис', 110, 1000, 48, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://fermer-63.ru/upload/iblock/10e/10ed883d25ae6ffc445c96c519394779.jpg', false);
INSERT INTO public.dishes VALUES (632, 'Тако', 60, 1000, 49, 0, 0, 'То, что нужно настоящему мексиканцу', 1, 1, 224, 1, 100, 'Горячее', 'Снеки', 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', false);
INSERT INTO public.dishes VALUES (633, 'Пряник', 70, 1000, 49, 1, 0, 'Очень вкусно с чаем', 1, 1, 126, 1, 70, 'К чаю', 'Снеки', 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', false);
INSERT INTO public.dishes VALUES (634, 'Чёрный бургер', 139, 1000, 49, 2, 0, 'Получен в угольных шахтах', 1, 1, 361, 1, 220, 'Универсальное', 'Снеки', 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', false);
INSERT INTO public.dishes VALUES (635, 'Кофе', 149, 1000, 49, 0, 1, 'Горячий, ароматный кофе', 1, 1, 90, 1, 100, 'Универсальное', 'Напитки', 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', false);
INSERT INTO public.dishes VALUES (636, 'Coca-cola', 65, 1000, 49, 1, 1, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/31776/331602242_m650_70f4e0bf0d2e51499d148a41f604b7c15325a3eec89ebc15447163877174a6d4.jpg', false);
INSERT INTO public.dishes VALUES (637, 'Fanta', 60, 1000, 49, 2, 1, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214204_m650.jpg', false);
INSERT INTO public.dishes VALUES (638, 'Sprite', 65, 1000, 49, 3, 1, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/38500/316214205_m650.jpg', false);
INSERT INTO public.dishes VALUES (639, 'Картошка Фри', 60, 1000, 49, 3, 0, 'Классический картофель фри', 1, 1, 232, 1, 120, 'Универсальное', 'Снеки', 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', false);
INSERT INTO public.dishes VALUES (640, 'Картошка по деревенски', 60, 1000, 49, 4, 0, '', 1, 1, 172, 1, 130, 'Универсальное', 'Снеки', 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', false);
INSERT INTO public.dishes VALUES (641, 'Блюдо со стейком', 256, 1000, 49, 0, 2, 'У этого блюда есть абсолютно всё', 1, 1, 756, 1, 400, 'Универсальное', 'Комбо', 'https://www.islandresortandcasino.com/sites/default/wp/blog/wp-content/uploads/2015/02/DSC_2162.jpg', false);
INSERT INTO public.dishes VALUES (642, 'Сёмга', 800, 1, 49, 1, 2, 'У этого блюда есть абсолютно всё', 1, 1, 700, 1, 5, 'Универсальное', 'Комбо', 'http://www.t-h.ru/photomenu/public/img/c/c6dc9d517433930dddcbb26e4824382c.jpg', false);
INSERT INTO public.dishes VALUES (643, 'Цезарь', 5, 1, 49, 2, 2, 'У этого блюда есть абсолютно всё', 1, 1, 784, 1, 5, 'Универсальное', 'Комбо', 'https://pizza-house18.ru/wp-content/uploads/2021/01/IMG_9061.jpg', false);
INSERT INTO public.dishes VALUES (644, 'Бефстроганов', 1000, 1000, 49, 3, 2, 'С грибами и картофельным пюре', 1, 1, 560, 1, 1000, 'Универсальное', 'Комбо', 'http://eday-cafe.ru/wp-content/uploads/2020/03/Бефстроганов-с-грибами-и-толченым-картофелем.jpg', false);
INSERT INTO public.dishes VALUES (645, 'Греческий', 230, 150, 50, 4, 0, '', 1, 1, 520, 1, 5, 'Универсальное', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/15838/328171056_m650_cfcbfe9ab6ebf66ca620792736fc485d3c0a91072b9143258e3c902a05717650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (646, 'Салат Крабовичок', 170, 1000, 50, 0, 0, '', 1, 1, 979, 1, 150, 'Пиццы', 'Салаты', 'https://www.povarenok.ru/data/cache/2013oct/16/04/532964_47569-710x550x.jpg', false);
INSERT INTO public.dishes VALUES (647, 'Теплый салат с говядиной', 370, 1000, 50, 1, 0, '', 1, 1, 926, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (648, 'Салат Бригантина', 350, 1000, 50, 2, 0, '', 1, 1, 901, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (649, 'Цезарь с креветками', 400, 1000, 50, 3, 0, '', 1, 1, 640, 1, 250, 'Пиццы', 'Салаты', 'https://2d-recept.com/wp-content/uploads/2019/05/salat-cezar-s-krevetkami-foto.jpg', false);
INSERT INTO public.dishes VALUES (650, 'Салат Винегрет', 360, 1, 50, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 270, 1, 250, 'Универсальное', 'Салаты', 'https://upload.wikimedia.org/wikipedia/commons/a/a7/Vinegret_cleaned.jpg', false);
INSERT INTO public.dishes VALUES (651, 'Борщ', 250, 1000, 50, 0, 1, '', 1, 1, 873, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128261_m650_c4d2c43bb3f55ec4004d87459556a79120b53b6e2528c352a439cd282ae22f0e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (652, 'Сборная солянка мясная', 350, 1000, 50, 1, 1, '', 1, 1, 335, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128262_m650_2c8a69834bd188549d9dad91e69845b4a67c58a2ca50f19d755ca9a191cd1299.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (653, 'Суп-лапша домашняя', 200, 1000, 50, 2, 1, '', 1, 1, 621, 1, 300, 'Пиццы', 'Супы', 'https://www.maggi.ru/data/images/recept/img640x500/recept_4495_czgj.jpg', false);
INSERT INTO public.dishes VALUES (654, 'Картофельный суп', 200, 1000, 50, 3, 1, '', 1, 1, 743, 1, 300, 'Пиццы', 'Супы', 'https://s1.eda.ru/StaticContent/Photos/120131084848/120213174933/p_O.jpg', false);
INSERT INTO public.dishes VALUES (655, 'Синнабон', 120, 1000, 50, 0, 2, '', 1, 1, 777, 1, 150, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334172939_m650_02814baf09423cd3bd412593eb3dfab2ecd2316537045b96989be181002b2ae9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (656, 'Трайфл Сникерс', 250, 1000, 50, 1, 2, '', 1, 1, 535, 1, 180, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333386700_m650_6cb3cf5e3c893441037dc08ab87054615e947637dcbaf3e6849c642aa1aab7c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (657, 'Кейк попс', 50, 1000, 50, 2, 2, '', 1, 1, 901, 1, 50, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334173008_m650_7c7a29300b7b1a11be154554e2120641a75e49ba543c884cc5ebdb060f682918.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (658, 'Шоколадный фондан', 180, 1000, 50, 3, 2, '', 1, 1, 213, 1, 150, 'Универсальное', 'Десерты', 'https://www.gastronom.ru/binfiles/images/20200121/bf504e3d.jpg', false);
INSERT INTO public.dishes VALUES (659, 'Шаурма домашняя', 199, 1, 51, 2, 0, 'У этого блюда есть абсолютно всё', 1, 1, 602, 1, 330, 'Универсальное', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542260_m650_b2b212d56f1bc91abe68636af87f1e185e719dfdd01409b7888c3c2884e71d04.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (660, 'Шаурма с тунцом', 229, 1000, 51, 0, 0, '', 1, 1, 697, 1, 260, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/334542767_m650_9321049bda74a7ef4ad30c7d93ccaf8b8ef4fceb3ce09e67b57e9bf5d42bc897.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (661, 'Бургер с куриной котлетой', 229, 1000, 51, 1, 0, '', 1, 1, 407, 1, 200, 'Пиццы', 'Донеры', 'https://www.delivery-club.ru/media/cms/relation_product/67635/335894906_m650_e5d32a976fbbc466608c4a9625ec1d73e9034b37d6e9611f2bc4a5eaa427e36f.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (662, 'Американо', 150, 1000, 51, 0, 1, '', 1, 1, 458, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263353_m650_f4265fc18ab291c9b77d274364519df8c59d3c398facdc47eaf993cc74f32994.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (688, 'Цезарь с креветками', 400, 1000, 53, 3, 0, '', 1, 1, 640, 1, 250, 'Пиццы', 'Салаты', 'https://2d-recept.com/wp-content/uploads/2019/05/salat-cezar-s-krevetkami-foto.jpg', false);
INSERT INTO public.dishes VALUES (663, 'Капучино', 130, 1000, 51, 1, 1, '', 1, 1, 927, 1, 200, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490930_m650_784d4b2dc50340e78aae44f47337985ea28f166ff5a26aa5b32fd7d0e494cc21.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (664, 'Латте', 170, 1000, 51, 2, 1, '', 1, 1, 907, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/333490932_m650_cb85b5a866a3afaa2c1f08b52b425f6e78fb0891ee033676524d6d90b4c6b2fd.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (665, 'Фильтр-кофе', 150, 1000, 51, 3, 1, '', 1, 1, 207, 1, 300, 'Пиццы', 'Кофе', 'https://www.delivery-club.ru/media/cms/relation_product/13698/329263356_m650_750ecfdcd43338312299910b4a7777c0466586c170bdef71faa0be4dce816437.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (666, 'Блин Ватрушка', 165, 1000, 51, 0, 2, '', 1, 1, 342, 1, 196, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485249_m650_861f7ff88a9073dfcb8fec2d28d8bc72ea710f7caa343417ae6667a1e516630d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (667, 'Блин с сахаром', 87, 1000, 51, 1, 2, '', 1, 1, 335, 1, 124, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485888_m650_ebfd220955c6f29d73ae8df3ed98ed9a406101b3cf030df5b2573959addaad3d.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (668, 'Блин с шоколадным кремом', 151, 1000, 51, 2, 2, '', 1, 1, 327, 1, 166, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485270_m650_36ced98a972f785c015b545212a58d23b3cce201603a556b5c998dbd53d4a8ce.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (669, 'Блин с вишневым вареньем', 93, 1000, 51, 3, 2, '', 1, 1, 292, 1, 126, 'Универсальное', 'Блины', 'https://www.delivery-club.ru/media/cms/relation_product/32049/328485286_m650_7334048751f2aa703a75266d1503ac44e8dd2929a429858ef468c59982d5a45e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (670, 'Пицца Пепперони', 470, 1, 52, 4, 0, '', 1, 1, 952, 1, 420, 'Универсальное', 'Пицца', 'https://s1.eda.ru/StaticContent/Photos/120131085053/171027192707/p_O.jpg', false);
INSERT INTO public.dishes VALUES (671, 'Пицца Ассорти', 429, 1000, 52, 0, 0, '', 1, 1, 660, 1, 600, 'Пиццы', 'Пиццы', 'https://sp-ao.shortpixel.ai/client/q_lossless,ret_img/https://cookery.site/wp-content/uploads/2019/09/photo_92647025-1-678x381.jpg', false);
INSERT INTO public.dishes VALUES (672, 'Пицца Морская', 800, 1000, 52, 1, 0, '', 1, 1, 584, 1, 600, 'Пиццы', 'Пиццы', 'https://pizzarini.info/wp-content/uploads/2018/03/pitstsa-morskaya-1-1.jpg', false);
INSERT INTO public.dishes VALUES (673, 'Пицца Карбонара', 540, 1000, 52, 2, 0, '', 1, 1, 581, 1, 420, 'Пиццы', 'Пиццы', 'https://static.1000.menu/img/content/13174/picca-karbonara_1432215044_0_max.jpg', false);
INSERT INTO public.dishes VALUES (674, 'Пицца Карначина', 590, 1000, 52, 3, 0, '', 1, 1, 895, 1, 540, 'Пиццы', 'Пиццы', 'https://cdnn21.img.ria.ru/images/98976/61/989766135_0:100:2000:1233_1920x0_80_0_0_4a3e7af4d4ab43307a68343c059cc57d.jpg', false);
INSERT INTO public.dishes VALUES (675, 'Пицца 4 сезона', 590, 1, 52, 4, 0, '', 1, 1, 952, 1, 580, 'Универсальное', 'Пицца', 'https://www.citypizza.ru/upload/img/shop/pizza/BORT/resize/4-%D1%81%D0%B5%D0%B7%D0%BE%D0%BD%D0%B0-%D0%B1%D0%BE%D1%80%D1%82-listThumb.jpg', false);
INSERT INTO public.dishes VALUES (676, 'Ролл с угрем', 180, 1000, 52, 0, 1, '', 1, 1, 469, 1, 120, 'Пиццы', 'Роллы', 'https://lh3.googleusercontent.com/proxy/uqHy2GczztN6r4P_eH2lySHzJWY57YsPW-vhEuU9GF46-i2_mf8qBESx4VDgPk4Vck5v7HTfIzpS13PG3puLRnbGfCMw-z4qYPFkKjiNMLiabJhV-7R9ncW31eb2X9gCBeywUfBvapgC9v8KGkNYQakwrbPtaiMZb8__LAvccU9bbP7NoV2zFMDuvM8ZDr9FXqP5wlQHHqDOngrSSISAbFkg8I2uIeBsYggTPQayCdO6D5NhwNf9I8w6-ZLnpLRvTScL5Hg4ifyw_7ARlpDmw2yfx0Ypjw', false);
INSERT INTO public.dishes VALUES (677, 'Дыхание дракона', 360, 1000, 52, 1, 1, '', 1, 1, 356, 1, 120, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011355_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (678, 'Горячий краб', 290, 1000, 52, 2, 1, '', 1, 1, 378, 1, 320, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011356_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (679, 'Запеченный угорь', 330, 1000, 52, 3, 1, '', 1, 1, 654, 1, 200, 'Пиццы', 'Роллы', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323011359_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (680, 'Сок в ассортименте', 100, 1000, 52, 0, 2, '', 1, 1, 102, 1, 1000, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664250_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (681, 'Coca-cola', 65, 1000, 52, 1, 2, '', 1, 1, 230, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/52007/323664245_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (682, 'Fanta', 60, 1000, 52, 2, 2, '', 1, 1, 225, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082269_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (683, 'Sprite', 65, 1000, 52, 3, 2, '', 1, 1, 215, 1, 500, 'Универсальное', 'Напитки', 'https://www.delivery-club.ru/media/cms/relation_product/64547/331082268_m650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (684, 'Греческий', 230, 150, 53, 4, 0, '', 1, 1, 520, 1, 5, 'Универсальное', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/15838/328171056_m650_cfcbfe9ab6ebf66ca620792736fc485d3c0a91072b9143258e3c902a05717650.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (685, 'Салат Крабовичок', 170, 1000, 53, 0, 0, '', 1, 1, 979, 1, 150, 'Пиццы', 'Салаты', 'https://www.povarenok.ru/data/cache/2013oct/16/04/532964_47569-710x550x.jpg', false);
INSERT INTO public.dishes VALUES (686, 'Теплый салат с говядиной', 370, 1000, 53, 1, 0, '', 1, 1, 926, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (687, 'Салат Бригантина', 350, 1000, 53, 2, 0, '', 1, 1, 901, 1, 150, 'Пиццы', 'Салаты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128256_m650_6e077ea0614f7ec1c092b0f4c6c912d55e03bda63aca0631901ba57af37dc1c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (689, 'Салат Винегрет', 360, 1, 53, 4, 0, 'У этого блюда есть абсолютно всё', 1, 1, 270, 1, 250, 'Универсальное', 'Салаты', 'https://upload.wikimedia.org/wikipedia/commons/a/a7/Vinegret_cleaned.jpg', false);
INSERT INTO public.dishes VALUES (690, 'Борщ', 250, 1000, 53, 0, 1, '', 1, 1, 873, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128261_m650_c4d2c43bb3f55ec4004d87459556a79120b53b6e2528c352a439cd282ae22f0e.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (691, 'Сборная солянка мясная', 350, 1000, 53, 1, 1, '', 1, 1, 335, 1, 300, 'Пиццы', 'Супы', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333128262_m650_2c8a69834bd188549d9dad91e69845b4a67c58a2ca50f19d755ca9a191cd1299.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (692, 'Суп-лапша домашняя', 200, 1000, 53, 2, 1, '', 1, 1, 621, 1, 300, 'Пиццы', 'Супы', 'https://www.maggi.ru/data/images/recept/img640x500/recept_4495_czgj.jpg', false);
INSERT INTO public.dishes VALUES (693, 'Картофельный суп', 200, 1000, 53, 3, 1, '', 1, 1, 743, 1, 300, 'Пиццы', 'Супы', 'https://s1.eda.ru/StaticContent/Photos/120131084848/120213174933/p_O.jpg', false);
INSERT INTO public.dishes VALUES (694, 'Синнабон', 120, 1000, 53, 0, 2, '', 1, 1, 777, 1, 150, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334172939_m650_02814baf09423cd3bd412593eb3dfab2ecd2316537045b96989be181002b2ae9.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (695, 'Трайфл Сникерс', 250, 1000, 53, 1, 2, '', 1, 1, 535, 1, 180, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/333386700_m650_6cb3cf5e3c893441037dc08ab87054615e947637dcbaf3e6849c642aa1aab7c2.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (696, 'Кейк попс', 50, 1000, 53, 2, 2, '', 1, 1, 901, 1, 50, 'Универсальное', 'Десерты', 'https://www.delivery-club.ru/media/cms/relation_product/71395/334173008_m650_7c7a29300b7b1a11be154554e2120641a75e49ba543c884cc5ebdb060f682918.jpg?resize=fit&width=1300&height=544&gravity=ce&out=webp', false);
INSERT INTO public.dishes VALUES (697, 'Шоколадный фондан', 180, 1000, 53, 3, 2, '', 1, 1, 213, 1, 150, 'Универсальное', 'Десерты', 'https://www.gastronom.ru/binfiles/images/20200121/bf504e3d.jpg', false);


--
--



--
--

INSERT INTO public.favorite_restaurant VALUES (1, 1, 1, 0);


--
--

INSERT INTO public.general_user_info VALUES (1, 'root', 'ca2e080a74ed1590cd141171c20e164d40d058fb45817c7b59f83159d059a6c0', 'salt', '88888888888', 'root@root', '/default/defaultUser.jpg', '2021-12-21 22:12:10.689883', false);


--
--



--
--



--
--



--
--



--
--



--
--



--
--

INSERT INTO public.promocode VALUES (1, 'HL4D4', 1, 1, 'Бесплатно куда угодно', 'Бесплатная доставка', '2021-12-21 22:12:10.689883', '2022-04-28 00:00:00', 'https://buslik.by/buyers/delivery/images/icon-2.png', true, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO public.promocode VALUES (2, 'CMO5S', 2, 1, 'Всем скидку!', 'Скидка 20% от цены 300', '2021-12-21 22:12:10.689883', '2022-04-28 00:00:00', 'http://pes-nv.ru/upload/iblock/eed/eedf46366e565834d5e726873a7ae200.jpg', NULL, NULL, NULL, 300, 20, NULL, NULL, NULL, NULL);
INSERT INTO public.promocode VALUES (3, 'DBL45', 3, 1, 'Бесплатное тако за покупку!', 'Бесплатно блюдо от 300', '2021-12-21 22:12:10.689883', '2022-04-28 00:00:00', 'https://image.freepik.com/free-photo/fresh-taco_144627-38286.jpg', NULL, 300, 1, NULL, NULL, NULL, NULL, NULL, NULL);


--
--

INSERT INTO public.radios VALUES (1, 'МакКомбо', 0, 8);
INSERT INTO public.radios VALUES (2, 'Утреннее комбо', 0, 9);
INSERT INTO public.radios VALUES (3, 'Аппетитное комбо', 0, 10);
INSERT INTO public.radios VALUES (4, 'Универсальное комбо', 0, 11);
INSERT INTO public.radios VALUES (5, 'Картофель', 0, 57);
INSERT INTO public.radios VALUES (6, 'Основное блюдо', 1, 57);
INSERT INTO public.radios VALUES (7, 'Напиток', 0, 58);
INSERT INTO public.radios VALUES (8, 'Напиток', 0, 59);
INSERT INTO public.radios VALUES (9, 'Напиток', 0, 60);
INSERT INTO public.radios VALUES (10, 'Картофель', 0, 84);
INSERT INTO public.radios VALUES (11, 'Основное блюдо', 1, 84);
INSERT INTO public.radios VALUES (12, 'Напиток', 0, 85);
INSERT INTO public.radios VALUES (13, 'Напиток', 0, 86);
INSERT INTO public.radios VALUES (14, 'Напиток', 0, 87);
INSERT INTO public.radios VALUES (15, 'Начинка', 0, 100);
INSERT INTO public.radios VALUES (16, 'Начинка', 0, 126);
INSERT INTO public.radios VALUES (17, 'Начинка', 0, 140);
INSERT INTO public.radios VALUES (18, 'Начинка', 0, 194);
INSERT INTO public.radios VALUES (19, 'Начинка', 0, 236);
INSERT INTO public.radios VALUES (20, 'Начинка', 0, 250);
INSERT INTO public.radios VALUES (21, 'Начинка', 0, 286);
INSERT INTO public.radios VALUES (22, 'Картофель', 0, 309);
INSERT INTO public.radios VALUES (23, 'Основное блюдо', 1, 309);
INSERT INTO public.radios VALUES (24, 'Напиток', 0, 310);
INSERT INTO public.radios VALUES (25, 'Напиток', 0, 311);
INSERT INTO public.radios VALUES (26, 'Напиток', 0, 312);
INSERT INTO public.radios VALUES (27, 'Начинка', 0, 471);
INSERT INTO public.radios VALUES (28, 'Начинка', 0, 499);
INSERT INTO public.radios VALUES (29, 'Картофель', 0, 522);
INSERT INTO public.radios VALUES (30, 'Основное блюдо', 1, 522);
INSERT INTO public.radios VALUES (31, 'Напиток', 0, 523);
INSERT INTO public.radios VALUES (32, 'Напиток', 0, 524);
INSERT INTO public.radios VALUES (33, 'Напиток', 0, 525);
INSERT INTO public.radios VALUES (34, 'Картофель', 0, 535);
INSERT INTO public.radios VALUES (35, 'Основное блюдо', 1, 535);
INSERT INTO public.radios VALUES (36, 'Напиток', 0, 536);
INSERT INTO public.radios VALUES (37, 'Напиток', 0, 537);
INSERT INTO public.radios VALUES (38, 'Напиток', 0, 538);
INSERT INTO public.radios VALUES (39, 'Картофель', 0, 562);
INSERT INTO public.radios VALUES (40, 'Основное блюдо', 1, 562);
INSERT INTO public.radios VALUES (41, 'Напиток', 0, 563);
INSERT INTO public.radios VALUES (42, 'Напиток', 0, 564);
INSERT INTO public.radios VALUES (43, 'Напиток', 0, 565);
INSERT INTO public.radios VALUES (44, 'Картофель', 0, 589);
INSERT INTO public.radios VALUES (45, 'Основное блюдо', 1, 589);
INSERT INTO public.radios VALUES (46, 'Напиток', 0, 590);
INSERT INTO public.radios VALUES (47, 'Напиток', 0, 591);
INSERT INTO public.radios VALUES (48, 'Напиток', 0, 592);
INSERT INTO public.radios VALUES (49, 'Картофель', 0, 602);
INSERT INTO public.radios VALUES (50, 'Основное блюдо', 1, 602);
INSERT INTO public.radios VALUES (51, 'Напиток', 0, 603);
INSERT INTO public.radios VALUES (52, 'Напиток', 0, 604);
INSERT INTO public.radios VALUES (53, 'Напиток', 0, 605);
INSERT INTO public.radios VALUES (54, 'Начинка', 0, 606);
INSERT INTO public.radios VALUES (55, 'Картофель', 0, 641);
INSERT INTO public.radios VALUES (56, 'Основное блюдо', 1, 641);
INSERT INTO public.radios VALUES (57, 'Напиток', 0, 642);
INSERT INTO public.radios VALUES (58, 'Напиток', 0, 643);
INSERT INTO public.radios VALUES (59, 'Напиток', 0, 644);


--
--

INSERT INTO public.restaurant VALUES (30, 1, 'Darbar', 'description', '2021-12-21 22:12:10.689883', false, 'https://sovkusom.ru/wp-content/uploads/blog/v/vrednaya-eda/1.jpg', 0, 0, 24, 54, 'city', 'street', 'house', 100, 5, 1, 1, '''darbar'':1');
INSERT INTO public.restaurant VALUES (1, 1, 'Атмосфера', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/6000027_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 250, 15, 90, 'city', 'street', 'house', 100, 5, 1, 1, '''атмосфера'':1');
INSERT INTO public.restaurant VALUES (2, 1, 'Shokolaat', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/5000052_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 10, 25, 65, 'city', 'street', 'house', 100, 3, 1, 1, '''shokolaat'':1');
INSERT INTO public.restaurant VALUES (3, 1, 'Gordon Biersch', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/44000095_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 15, 35, 40, 'city', 'street', 'house', 100, 4, 1, 1, '''biersch'':2 ''gordon'':1');
INSERT INTO public.restaurant VALUES (4, 1, 'Crepevine', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/6000035_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 350, 22, 30, 'city', 'street', 'house', 100, 2, 1, 1, '''crepevin'':1');
INSERT INTO public.restaurant VALUES (5, 1, 'Creamery', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/27000060_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 250, 10, 55, 'city', 'street', 'house', 100, 1, 1, 1, '''creameri'':1');
INSERT INTO public.restaurant VALUES (6, 1, 'Old Pro', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/5f59d56754805_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 100, 31, 32, 'city', 'street', 'house', 100, 2.5, 1, 1, '''old'':1 ''pro'':2');
INSERT INTO public.restaurant VALUES (7, 1, 'Дом вкуснятины', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/5f4a59b84ad69_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 53, 15, 30, 'city', 'street', 'house', 100, 4.5, 1, 1, '''вкуснятины'':2 ''дом'':1');
INSERT INTO public.restaurant VALUES (8, 1, 'Продуктовая печь', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/26000199_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 220, 45, 60, 'city', 'street', 'house', 100, 5, 1, 1, '''печь'':2 ''продуктовая'':1');
INSERT INTO public.restaurant VALUES (9, 1, 'La Strada', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/5f62243740d71_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 121, 17, 20, 'city', 'street', 'house', 100, 3.4, 1, 1, '''la'':1 ''strada'':2');
INSERT INTO public.restaurant VALUES (10, 1, 'Buca di Beppo', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/1000026_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 150, 15, 45, 'city', 'street', 'house', 100, 2.1, 1, 1, '''beppo'':3 ''buca'':1 ''di'':2');
INSERT INTO public.restaurant VALUES (11, 1, 'Мадам Там', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/19000230_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 200, 20, 30, 'city', 'street', 'house', 100, 1.6, 1, 1, '''мадам'':1 ''там'':2');
INSERT INTO public.restaurant VALUES (12, 1, 'Спрут кафе', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/43000086_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 250, 22, 37, 'city', 'street', 'house', 100, 2.3, 1, 1, '''кафе'':2 ''спрут'':1');
INSERT INTO public.restaurant VALUES (13, 1, 'Bistro Maxine', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/61864cc9d00ea_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 300, 15, 46, 'city', 'street', 'house', 100, 3.1, 1, 1, '''bistro'':1 ''maxin'':2');
INSERT INTO public.restaurant VALUES (14, 1, 'Три сезона', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/26000199_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 350, 16, 50, 'city', 'street', 'house', 100, 2.7, 1, 1, '''сезона'':2 ''три'':1');
INSERT INTO public.restaurant VALUES (15, 1, 'Спокойствие', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/5edb9be4ddeba_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 400, 15, 30, 'city', 'street', 'house', 100, 4.9, 1, 1, '''спокойствие'':1');
INSERT INTO public.restaurant VALUES (16, 1, 'Siam Royal', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/5f22cd5325126_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 450, 25, 44, 'city', 'street', 'house', 100, 3.9, 1, 1, '''royal'':2 ''siam'':1');
INSERT INTO public.restaurant VALUES (17, 1, 'Krung Siam', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/60df0e2fec006_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 0, 13, 55, 'city', 'street', 'house', 100, 2.31, 1, 1, '''krung'':1 ''siam'':2');
INSERT INTO public.restaurant VALUES (18, 1, 'Тайфун', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/61027eda0d4c0_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 0, 20, 56, 'city', 'street', 'house', 100, 1.2, 1, 1, '''тайфун'':1');
INSERT INTO public.restaurant VALUES (19, 1, 'Tamarine', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/1000039_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 0, 10, 44, 'city', 'street', 'house', 100, 3.4, 1, 1, '''tamarin'':1');
INSERT INTO public.restaurant VALUES (20, 1, 'Joya', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/60df0e2fec006_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 0, 19, 33, 'city', 'street', 'house', 100, 2.6, 1, 1, '''joya'':1');
INSERT INTO public.restaurant VALUES (21, 1, 'Колокольчик', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/61780d63510c1_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 499, 16, 47, 'city', 'street', 'house', 100, 4.7, 1, 1, '''колокольчик'':1');
INSERT INTO public.restaurant VALUES (22, 1, 'Evvia', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/60d201e011fd4_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 449, 30, 39, 'city', 'street', 'house', 100, 0.5, 1, 1, '''evvia'':1');
INSERT INTO public.restaurant VALUES (23, 1, 'Кафе 220', 'description', '2021-12-21 22:12:10.689883', false, 'https://mywowo.net/media/images/cache/tokyo_meraviglie_tavola_01_introduzione_jpg_1200_630_cover_85.jpg', 0, 399, 40, 50, 'city', 'street', 'house', 100, 0.8, 1, 1, '''220'':2 ''кафе'':1');
INSERT INTO public.restaurant VALUES (24, 1, 'Кафе Ренессанс', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/48000050_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 349, 11, 55, 'city', 'street', 'house', 100, 0.1, 1, 1, '''кафе'':1 ''ренессанс'':2');
INSERT INTO public.restaurant VALUES (25, 1, 'Kan Zeman', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/2000031_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 299, 40, 45, 'city', 'street', 'house', 100, 0.75, 1, 1, '''kan'':1 ''zeman'':2');
INSERT INTO public.restaurant VALUES (26, 1, 'Кафе Манго', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/25000109_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 249, 30, 40, 'city', 'street', 'house', 100, 1.3, 1, 1, '''кафе'':1 ''манго'':2');
INSERT INTO public.restaurant VALUES (27, 1, 'Балаклава', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/61864cc9d00ea_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 199, 24, 53, 'city', 'street', 'house', 100, 2.34, 1, 1, '''балаклава'':1');
INSERT INTO public.restaurant VALUES (28, 1, 'Иностранный гурман', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.delivery-club.ru/naturmort/5ec2443197ff3_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp', 0, 149, 18, 32, 'city', 'street', 'house', 100, 1.23, 1, 1, '''гурман'':2 ''иностранный'':1');
INSERT INTO public.restaurant VALUES (29, 1, 'Частичка Бангкока', 'description', '2021-12-21 22:12:10.689883', false, 'https://incrussia.ru/wp-content/uploads/2018/10/iStock-694189032.jpg', 0, 99, 19, 50, 'city', 'street', 'house', 100, 4.1, 1, 1, '''бангкока'':2 ''частичка'':1');
INSERT INTO public.restaurant VALUES (31, 1, 'Mantra', 'description', '2021-12-21 22:12:10.689883', false, 'https://naked-science.ru/wp-content/uploads/2020/12/fast-fud-pitstsa-burger-chipsy-lukovye-koltsa-kartofel-fri.jpg', 0, 0, 23, 44, 'city', 'street', 'house', 100, 5, 1, 1, '''mantra'':1');
INSERT INTO public.restaurant VALUES (32, 1, 'Janta', 'description', '2021-12-21 22:12:10.689883', false, 'https://static.tildacdn.com/tild6561-6165-4337-b835-316638666562/20-05-20.jpg', 0, 0, 19, 23, 'city', 'street', 'house', 100, 5, 1, 1, '''janta'':1');
INSERT INTO public.restaurant VALUES (33, 1, 'Hyderabad', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.kamis-pripravy.ru/upload/medialibrary/907/9073bb8cc5579504bd22a62e5c1fe0e0.jpg', 0, 0, 25, 50, 'city', 'street', 'house', 100, 5, 1, 1, '''hyderabad'':1');
INSERT INTO public.restaurant VALUES (34, 1, 'Кофейня Джека', 'description', '2021-12-21 22:12:10.689883', false, 'https://images.aif.by/007/433/e73337ac5677e37f8baa002e41232ed4.jpg', 0, 0, 26, 52, 'city', 'street', 'house', 100, 5, 1, 1, '''джека'':2 ''кофейня'':1');
INSERT INTO public.restaurant VALUES (35, 1, 'Coop кофейня', 'description', '2021-12-21 22:12:10.689883', false, 'https://img.gazeta.ru/files3/829/13377829/Depositphotos_412834214_xl-2015-pic905-895x505-19117.jpg', 0, 0, 15, 45, 'city', 'street', 'house', 100, 5, 1, 1, '''coop'':1 ''кофейня'':2');
INSERT INTO public.restaurant VALUES (36, 1, 'Lytton Coffee', 'description', '2021-12-21 22:12:10.689883', false, 'https://cdnmyslo.ru/Photogallery/99/1d/991dffc2-ea20-483e-9352-88cd8e2aa751_b.jpg', 0, 0, 16, 48, 'city', 'street', 'house', 100, 5, 1, 1, '''coffe'':2 ''lytton'':1');
INSERT INTO public.restaurant VALUES (37, 1, 'Il Fornaio', 'description', '2021-12-21 22:12:10.689883', false, 'https://images.ua.prom.st/3125534192_w600_h600_eda-na-vynos.jpg', 0, 0, 17, 51, 'city', 'street', 'house', 100, 4.5, 1, 1, '''fornaio'':2 ''il'':1');
INSERT INTO public.restaurant VALUES (38, 1, 'Lavanda', 'description', '2021-12-21 22:12:10.689883', false, 'https://incrussia.ru/wp-content/uploads/2020/11/iStock-1175505781.jpg', 0, 0, 18, 54, 'city', 'street', 'house', 100, 3.5, 1, 1, '''lavanda'':1');
INSERT INTO public.restaurant VALUES (39, 1, 'MacArthur', 'description', '2021-12-21 22:12:10.689883', false, 'https://kidpassage.com/images/publications/eda-sankt-peterburge-chto-poprobovat-skolko-stoit/cover_original.jpg', 0, 0, 19, 57, 'city', 'street', 'house', 100, 2.5, 1, 1, '''macarthur'':1');
INSERT INTO public.restaurant VALUES (40, 1, 'Osteria', 'description', '2021-12-21 22:12:10.689883', false, 'https://cdn.fishki.net/upload/post/2017/01/30/2205250/2-1485519719-1.jpg', 0, 399, 20, 34, 'city', 'street', 'house', 100, 1.5, 1, 1, '''osteria'':1');
INSERT INTO public.restaurant VALUES (41, 1, 'Vero', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.learnathome.ru/files/media/food.jpg', 0, 499, 20, 40, 'city', 'street', 'house', 100, 0.5, 1, 1, '''vero'':1');
INSERT INTO public.restaurant VALUES (42, 1, 'Renzo', 'description', '2021-12-21 22:12:10.689883', false, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQTvLJHAw98D_0U8xi8fjAN573FWUX42sltrRp2-CkVtQOKUrmoIBP1XyLO5RE_fITY1KQ&usqp=CAU', 0, 299, 21, 42, 'city', 'street', 'house', 100, 0.4, 1, 1, '''renzo'':1');
INSERT INTO public.restaurant VALUES (43, 1, 'Miyake', 'description', '2021-12-21 22:12:10.689883', false, 'https://interesnyefakty.org/wp-content/uploads/Interesnye-fakty-o-ede-v-raznyh-stranah.jpg', 0, 249, 22, 44, 'city', 'street', 'house', 100, 0.3, 1, 1, '''miyak'':1');
INSERT INTO public.restaurant VALUES (44, 1, 'Tomo', 'description', '2021-12-21 22:12:10.689883', false, 'https://billionnews.ru/timthumb/timthumb.php?src=http://billionnews.ru/uploads/posts/2017-01/thumbs/1485519719_1.jpg&w=940&h=600&zc=1', 0, 199, 23, 46, 'city', 'street', 'house', 100, 0.2, 1, 1, '''tomo'':1');
INSERT INTO public.restaurant VALUES (45, 1, 'Kanpai', 'description', '2021-12-21 22:12:10.689883', false, 'https://gorobzor.ru/content/news/2018/06/chto_iz_edy_poprobovat_v_sochi_image_5b2cf79b7278f1.83210187.jpg', 0, 149, 24, 36, 'city', 'street', 'house', 100, 0.1, 1, 1, '''kanpai'':1');
INSERT INTO public.restaurant VALUES (46, 1, 'Любовь моей жизни', 'description', '2021-12-21 22:12:10.689883', false, 'https://kidpassage.com/images/publications/eda-sohi-hto-poprobovat-skolko-stoit/cover_original.jpg', 0, 266, 30, 45, 'city', 'street', 'house', 100, 5, 1, 1, '''жизни'':3 ''любовь'':1 ''моей'':2');
INSERT INTO public.restaurant VALUES (47, 1, 'Новая пицца', 'description', '2021-12-21 22:12:10.689883', false, 'https://www.oum.ru/upload/iblock/4a6/4a689562637ffe31a94e1770388395f8.jpg', 0, 233, 31, 46, 'city', 'street', 'house', 100, 4, 1, 1, '''новая'':1 ''пицца'':2');
INSERT INTO public.restaurant VALUES (48, 1, 'Калифорнийская кухня', 'description', '2021-12-21 22:12:10.689883', false, 'https://cs1.livemaster.ru/storage/15/98/6a9751d56360234808ec8ac68anj--kukly-i-igrushki-eda-dlya-kukol-eda-dlya-barbi-kukolnaya-eda-.jpg', 0, 150, 23, 32, 'city', 'street', 'house', 100, 3, 1, 1, '''калифорнийская'':1 ''кухня'':2');
INSERT INTO public.restaurant VALUES (49, 1, 'Круглый стол', 'description', '2021-12-21 22:12:10.689883', false, 'https://lh3.googleusercontent.com/proxy/HgfW931vlU8WqU-KdGv8doKW5Re0c1qU6t-EkRfRzehj0c1-eEbSMgbSIZe4e7wVyGOGUNFzGWwaTFZwDkD_bu75cIZm4PhFxJj4WI-S-xXWtwhozr8U', 0, 175, 17, 37, 'city', 'street', 'house', 100, 2, 1, 1, '''круглый'':1 ''стол'':2');
INSERT INTO public.restaurant VALUES (50, 1, 'Любимая шляпа', 'description', '2021-12-21 22:12:10.689883', false, 'https://img.the-village.me/the-village.me/post-cover/-k0NDtajdfoONfacIAqvoA-default.jpg', 0, 250, 16, 36, 'city', 'street', 'house', 100, 1, 1, 1, '''любимая'':1 ''шляпа'':2');
INSERT INTO public.restaurant VALUES (51, 1, 'Garden Fresh', 'description', '2021-12-21 22:12:10.689883', false, 'https://tomato.ua/blog/wp-content/uploads/2019/03/000-39-1-1440x961.jpg', 0, 300, 16, 46, 'city', 'street', 'house', 100, 3.2, 1, 1, '''fresh'':2 ''garden'':1');
INSERT INTO public.restaurant VALUES (52, 1, 'Epi', 'description', '2021-12-21 22:12:10.689883', false, 'https://avatars.mds.yandex.net/get-altay/2960979/2a0000017260a9d9f85eb44d3ab634dd7d7f/XXL', 0, 150, 16, 56, 'city', 'street', 'house', 100, 2.1, 1, 1, '''epi'':1');
INSERT INTO public.restaurant VALUES (53, 1, 'Валентино', 'description', '2021-12-21 22:12:10.689883', false, 'https://i1.wp.com/www.agoda.com/wp-content/uploads/2018/07/Experience-Tokyo_food-and-drink_Featured-image-1200x350_sushi-tray_Tokyo.jpg?fit=1200%2C350&ssl=1', 0, 100, 15, 55, 'city', 'street', 'house', 100, 4.2, 1, 1, '''валентино'':1');


--
--

INSERT INTO public.restaurant_category VALUES (57, 'Бар', 20, 0, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (17, 'Бар', 6, 2, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (9, 'Бар', 4, 0, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (1, 'Кафе', 1, 0, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (2, 'Поп-ап', 1, 1, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (3, 'Хенкальная', 2, 0, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (4, 'Буфеты', 2, 1, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (5, 'Кафе', 2, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (6, 'Кальянная', 3, 0, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (7, 'Суши-бар', 3, 1, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (8, 'Поп-ап', 3, 2, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (10, 'Хенкальная', 4, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (11, 'Буфеты', 4, 2, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (12, 'Виртуальный', 5, 0, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (14, 'Общепит', 5, 2, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (13, 'Хенкальная', 5, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (15, 'Поп-ап', 6, 0, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (65, 'Бар', 22, 2, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (16, 'Суши-бар', 6, 1, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (18, 'Поп-ап', 7, 0, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (19, 'Кальянная', 7, 1, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (20, 'Поп-ап', 7, 2, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (21, 'Кальянная', 8, 0, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (66, 'Бар', 23, 0, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (22, 'Хенкальная', 8, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (23, 'Виртуальный', 8, 2, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (24, 'Виртуальный', 9, 0, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (25, 'Кальянная', 9, 1, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (26, 'Кафе', 9, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (27, 'Буфеты', 10, 0, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (28, 'Кафе', 10, 1, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (29, 'Кафе', 10, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (30, 'Пиццерия', 11, 0, '''пиццерия'':1');
INSERT INTO public.restaurant_category VALUES (31, 'Пиццерия', 11, 1, '''пиццерия'':1');
INSERT INTO public.restaurant_category VALUES (32, 'Бар', 11, 2, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (33, 'Общепит', 12, 0, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (34, 'Буфеты', 12, 1, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (35, 'Буфеты', 12, 2, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (36, 'Суши-бар', 13, 0, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (37, 'Виртуальный', 13, 1, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (38, 'Виртуальный', 13, 2, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (39, 'Буфеты', 14, 0, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (40, 'Виртуальный', 14, 1, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (41, 'Виртуальный', 14, 2, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (42, 'Буфеты', 15, 0, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (43, 'Хенкальная', 15, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (44, 'Поп-ап', 15, 2, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (45, 'Бар', 16, 0, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (46, 'Буфеты', 16, 1, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (47, 'Буфеты', 16, 2, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (48, 'Бар', 17, 0, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (50, 'Бар', 17, 2, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (49, 'Суши-бар', 17, 1, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (51, 'Поп-ап', 18, 0, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (52, 'Пиццерия', 18, 1, '''пиццерия'':1');
INSERT INTO public.restaurant_category VALUES (53, 'Хенкальная', 18, 2, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (54, 'Кафе', 19, 0, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (55, 'Поп-ап', 19, 1, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (56, 'Хенкальная', 19, 2, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (58, 'Кальянная', 20, 1, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (59, 'Пиццерия', 20, 2, '''пиццерия'':1');
INSERT INTO public.restaurant_category VALUES (60, 'Хенкальная', 21, 0, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (61, 'Хенкальная', 21, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (62, 'Виртуальный', 21, 2, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (63, 'Кальянная', 22, 0, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (64, 'Кальянная', 22, 1, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (67, 'Виртуальный', 23, 1, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (68, 'Буфеты', 23, 2, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (69, 'Пиццерия', 24, 0, '''пиццерия'':1');
INSERT INTO public.restaurant_category VALUES (70, 'Хенкальная', 24, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (71, 'Бар', 24, 2, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (72, 'Кальянная', 25, 0, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (73, 'Кафе', 25, 1, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (77, 'Кафе', 26, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (74, 'Кальянная', 25, 2, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (75, 'Кальянная', 26, 0, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (76, 'Общепит', 26, 1, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (78, 'Кафе', 27, 0, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (79, 'Бар', 27, 1, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (80, 'Виртуальный', 27, 2, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (147, 'Хенкальная', 50, 0, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (81, 'Суши-бар', 28, 0, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (82, 'Бар', 28, 1, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (83, 'Кафе', 28, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (84, 'Виртуальный', 29, 0, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (85, 'Буфеты', 29, 1, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (86, 'Кафе', 29, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (87, 'Суши-бар', 30, 0, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (88, 'Общепит', 30, 1, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (89, 'Кафе', 30, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (90, 'Буфеты', 31, 0, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (91, 'Кафе', 31, 1, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (92, 'Общепит', 31, 2, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (93, 'Суши-бар', 32, 0, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (94, 'Кальянная', 32, 1, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (95, 'Кальянная', 32, 2, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (96, 'Поп-ап', 33, 0, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (97, 'Суши-бар', 33, 1, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (98, 'Суши-бар', 33, 2, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (99, 'Пиццерия', 34, 0, '''пиццерия'':1');
INSERT INTO public.restaurant_category VALUES (100, 'Общепит', 34, 1, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (101, 'Суши-бар', 34, 2, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (102, 'Кафе', 35, 0, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (103, 'Общепит', 35, 1, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (104, 'Кальянная', 35, 2, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (105, 'Буфеты', 36, 0, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (106, 'Суши-бар', 36, 1, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (107, 'Кафе', 36, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (108, 'Виртуальный', 37, 0, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (109, 'Бар', 37, 1, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (110, 'Суши-бар', 37, 2, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (111, 'Кафе', 38, 0, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (112, 'Виртуальный', 38, 1, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (113, 'Кальянная', 38, 2, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (114, 'Хенкальная', 39, 0, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (115, 'Хенкальная', 39, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (116, 'Буфеты', 39, 2, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (117, 'Виртуальный', 40, 0, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (118, 'Бар', 40, 1, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (119, 'Общепит', 40, 2, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (120, 'Кальянная', 41, 0, '''кальянная'':1');
INSERT INTO public.restaurant_category VALUES (121, 'Хенкальная', 41, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (122, 'Суши-бар', 41, 2, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (123, 'Кафе', 42, 0, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (124, 'Общепит', 42, 1, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (125, 'Пиццерия', 42, 2, '''пиццерия'':1');
INSERT INTO public.restaurant_category VALUES (126, 'Общепит', 43, 0, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (127, 'Хенкальная', 43, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (128, 'Кафе', 43, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (129, 'Общепит', 44, 0, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (130, 'Виртуальный', 44, 1, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (131, 'Бар', 44, 2, '''бар'':1');
INSERT INTO public.restaurant_category VALUES (132, 'Кафе', 45, 0, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (133, 'Виртуальный', 45, 1, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (134, 'Общепит', 45, 2, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (135, 'Виртуальный', 46, 0, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (136, 'Буфеты', 46, 1, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (137, 'Хенкальная', 46, 2, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (138, 'Общепит', 47, 0, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (139, 'Хенкальная', 47, 1, '''хенкальная'':1');
INSERT INTO public.restaurant_category VALUES (140, 'Поп-ап', 47, 2, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (141, 'Кафе', 48, 0, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (142, 'Суши-бар', 48, 1, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (143, 'Буфеты', 48, 2, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (144, 'Виртуальный', 49, 0, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (145, 'Поп-ап', 49, 1, '''ап'':3 ''поп'':2 ''поп-ап'':1');
INSERT INTO public.restaurant_category VALUES (146, 'Суши-бар', 49, 2, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (148, 'Виртуальный', 50, 1, '''виртуальный'':1');
INSERT INTO public.restaurant_category VALUES (149, 'Буфеты', 50, 2, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (150, 'Общепит', 51, 0, '''общепит'':1');
INSERT INTO public.restaurant_category VALUES (151, 'Буфеты', 51, 1, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (152, 'Кафе', 51, 2, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (153, 'Пиццерия', 52, 0, '''пиццерия'':1');
INSERT INTO public.restaurant_category VALUES (154, 'Суши-бар', 52, 1, '''бар'':3 ''суши'':2 ''суши-бар'':1');
INSERT INTO public.restaurant_category VALUES (155, 'Буфеты', 52, 2, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (156, 'Буфеты', 53, 0, '''буфеты'':1');
INSERT INTO public.restaurant_category VALUES (157, 'Кафе', 53, 1, '''кафе'':1');
INSERT INTO public.restaurant_category VALUES (158, 'Кальянная', 53, 2, '''кальянная'':1');


--
--

INSERT INTO public.review VALUES (1, 1, 0, 'rootrootro', '2021-12-01 18:53:51.550501', 2, false);
INSERT INTO public.review VALUES (2, 1, 0, 'rootrootro', '2021-12-01 18:54:02.753424', 2, false);
INSERT INTO public.review VALUES (3, 1, 0, 'rootrootrp', '2021-12-01 18:54:51.736683', 2, false);
INSERT INTO public.review VALUES (4, 1, 0, 'rootrootrp', '2021-12-01 18:57:26.043285', 2, false);
INSERT INTO public.review VALUES (5, 1, 0, 'rootrootrp', '2021-12-01 18:58:22.53117', 2, false);
INSERT INTO public.review VALUES (6, 1, 0, 'rootrootrp', '2021-12-01 18:58:29.022712', 2, false);
INSERT INTO public.review VALUES (7, 1, 0, 'rootrootrp', '2021-12-01 19:03:04.076647', 2, false);
INSERT INTO public.review VALUES (8, 1, 0, 'rootrootrp', '2021-12-01 19:03:04.113509', 2, false);
INSERT INTO public.review VALUES (9, 1, 0, 'rootrootrp', '2021-12-01 19:03:04.126966', 2, false);
INSERT INTO public.review VALUES (11, 1, 27, 'asdasdasdasd', '2021-12-01 19:06:47.812287', 3, false);
INSERT INTO public.review VALUES (10, 1, 29, '2dasdassssdsaasd', '2021-12-01 19:03:27.551381', 5, false);
INSERT INTO public.review VALUES (12, 1, 0, 'asdasda2sdasd', '2021-12-01 19:12:05.52179', 4, false);


--
--

INSERT INTO public.structure_dishes VALUES (1, 'Кетчуп', 5, 1, 0, 1, 1, 1, 1, 5, false, false);
INSERT INTO public.structure_dishes VALUES (2, 'Горчица', 5, 1, 1, 1, 1, 1, 1, 5, false, false);
INSERT INTO public.structure_dishes VALUES (3, 'Сырные бортики', 5, 4, 0, 1, 1, 1, 1, 5, false, false);
INSERT INTO public.structure_dishes VALUES (4, 'Колбаса', 5, 4, 1, 1, 1, 1, 1, 5, false, false);
INSERT INTO public.structure_dishes VALUES (5, 'Сыр Пармезан', 5, 4, 2, 1, 1, 1, 1, 5, false, false);
INSERT INTO public.structure_dishes VALUES (6, 'Сыр Моцарелла', 5, 4, 3, 1, 1, 1, 1, 5, false, false);
INSERT INTO public.structure_dishes VALUES (7, 'Сахар', 5, 5, 0, 1, 1, 1, 1, 5, false, false);
INSERT INTO public.structure_dishes VALUES (8, 'Кетчап', 5, 1, 2, 1, 1, 1, 1, 5, false, false);
INSERT INTO public.structure_dishes VALUES (9, 'Кетчуп', 5, 48, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (10, 'Горчица', 5, 48, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (11, 'Сахар', 5, 51, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (12, 'Кетчуп', 5, 75, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (13, 'Горчица', 5, 75, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (14, 'Сахар', 5, 78, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (15, 'Лук', 25, 88, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (16, 'Зелень', 20, 88, 1, 1, 1, 1, 8, 100, false, false);
INSERT INTO public.structure_dishes VALUES (17, 'Лимон', 21, 89, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (18, 'Помидор', 15, 93, 0, 1, 1, 1, 1, 100, false, false);
INSERT INTO public.structure_dishes VALUES (19, 'Лук', 25, 114, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (20, 'Зелень', 20, 114, 1, 1, 1, 1, 8, 100, false, false);
INSERT INTO public.structure_dishes VALUES (21, 'Лимон', 21, 115, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (22, 'Помидор', 15, 119, 0, 1, 1, 1, 1, 100, false, false);
INSERT INTO public.structure_dishes VALUES (23, 'Лук', 25, 182, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (24, 'Зелень', 20, 182, 1, 1, 1, 1, 8, 100, false, false);
INSERT INTO public.structure_dishes VALUES (25, 'Лимон', 21, 183, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (26, 'Помидор', 15, 187, 0, 1, 1, 1, 1, 100, false, false);
INSERT INTO public.structure_dishes VALUES (27, 'Кетчуп', 5, 300, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (28, 'Горчица', 5, 300, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (29, 'Сахар', 5, 303, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (30, 'Лук', 25, 355, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (31, 'Зелень', 20, 355, 1, 1, 1, 1, 8, 100, false, false);
INSERT INTO public.structure_dishes VALUES (32, 'Лимон', 21, 356, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (33, 'Помидор', 15, 360, 0, 1, 1, 1, 1, 100, false, false);
INSERT INTO public.structure_dishes VALUES (34, 'Лук', 25, 445, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (35, 'Зелень', 20, 445, 1, 1, 1, 1, 8, 100, false, false);
INSERT INTO public.structure_dishes VALUES (36, 'Лимон', 21, 446, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (37, 'Помидор', 15, 450, 0, 1, 1, 1, 1, 100, false, false);
INSERT INTO public.structure_dishes VALUES (38, 'Кетчуп', 5, 513, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (39, 'Горчица', 5, 513, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (40, 'Сахар', 5, 516, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (41, 'Кетчуп', 5, 526, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (42, 'Горчица', 5, 526, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (43, 'Сахар', 5, 529, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (44, 'Кетчуп', 5, 553, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (45, 'Горчица', 5, 553, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (46, 'Сахар', 5, 556, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (47, 'Кетчуп', 5, 580, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (48, 'Горчица', 5, 580, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (49, 'Сахар', 5, 583, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (50, 'Кетчуп', 5, 593, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (51, 'Горчица', 5, 593, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (52, 'Сахар', 5, 596, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (53, 'Лук', 25, 620, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (54, 'Зелень', 20, 620, 1, 1, 1, 1, 8, 100, false, false);
INSERT INTO public.structure_dishes VALUES (55, 'Лимон', 21, 621, 0, 1, 1, 1, 10, 100, false, false);
INSERT INTO public.structure_dishes VALUES (56, 'Помидор', 15, 625, 0, 1, 1, 1, 1, 100, false, false);
INSERT INTO public.structure_dishes VALUES (57, 'Кетчуп', 5, 632, 0, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (58, 'Горчица', 5, 632, 1, 1, 1, 1, 1, 1, false, false);
INSERT INTO public.structure_dishes VALUES (59, 'Сахар', 5, 635, 0, 1, 1, 1, 1, 1, false, false);


--
--

INSERT INTO public.structure_radios VALUES (1, 'Картофель Фри', 1, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (2, 'Картофель по деревенски', 1, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (3, 'Сырный соус', 2, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (4, 'Чесночный соус', 2, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (5, 'Кисло-сладкий соус', 2, 2, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (6, 'Картофель Фри', 3, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (7, 'Картофель по деревенски', 3, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (8, 'Сырный соус', 4, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (9, 'Чесночный соус', 4, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (10, 'Картофель Фри', 5, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (11, 'Картофель по Деревенски', 5, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (12, 'Стейк', 6, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (13, 'Кофе', 7, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (14, 'Кофе', 8, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (15, 'Кофе', 9, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (16, 'Картофель Фри', 10, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (17, 'Картофель по Деревенски', 10, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (18, 'Стейк', 11, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (19, 'Кофе', 12, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (20, 'Кофе', 13, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (21, 'Кофе', 14, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (22, 'Рис', 15, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (23, 'Нори', 15, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (24, 'Рис', 16, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (25, 'Нори', 16, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (26, 'Рис', 17, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (27, 'Нори', 17, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (28, 'Рис', 18, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (29, 'Нори', 18, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (30, 'Рис', 19, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (31, 'Нори', 19, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (32, 'Рис', 20, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (33, 'Нори', 20, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (34, 'Рис', 21, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (35, 'Нори', 21, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (36, 'Картофель Фри', 22, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (37, 'Картофель по Деревенски', 22, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (38, 'Стейк', 23, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (39, 'Кофе', 24, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (40, 'Кофе', 25, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (41, 'Кофе', 26, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (42, 'Рис', 27, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (43, 'Нори', 27, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (44, 'Рис', 28, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (45, 'Нори', 28, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (46, 'Картофель Фри', 29, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (47, 'Картофель по Деревенски', 29, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (48, 'Стейк', 30, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (49, 'Кофе', 31, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (50, 'Кофе', 32, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (51, 'Кофе', 33, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (52, 'Картофель Фри', 34, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (53, 'Картофель по Деревенски', 34, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (54, 'Стейк', 35, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (55, 'Кофе', 36, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (56, 'Кофе', 37, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (57, 'Кофе', 38, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (58, 'Картофель Фри', 39, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (59, 'Картофель по Деревенски', 39, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (60, 'Стейк', 40, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (61, 'Кофе', 41, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (62, 'Кофе', 42, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (63, 'Кофе', 43, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (64, 'Картофель Фри', 44, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (65, 'Картофель по Деревенски', 44, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (66, 'Стейк', 45, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (67, 'Кофе', 46, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (68, 'Кофе', 47, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (69, 'Кофе', 48, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (70, 'Картофель Фри', 49, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (71, 'Картофель по Деревенски', 49, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (72, 'Стейк', 50, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (73, 'Кофе', 51, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (74, 'Кофе', 52, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (75, 'Кофе', 53, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (76, 'Рис', 54, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (77, 'Нори', 54, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (78, 'Картофель Фри', 55, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (79, 'Картофель по Деревенски', 55, 1, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (80, 'Стейк', 56, 0, 10, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (81, 'Кофе', 57, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (82, 'Кофе', 58, 0, 1, 1, 1, 1);
INSERT INTO public.structure_radios VALUES (83, 'Кофе', 59, 0, 1, 1, 1, 1);


--
--



--
--

SELECT pg_catalog.setval('public.address_user_id_seq', 1, true);


--
--

SELECT pg_catalog.setval('public.card_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.cart_food_id_seq', 295, true);


--
--

SELECT pg_catalog.setval('public.cart_radios_food_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.cart_structure_food_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.cart_user_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.client_id_seq', 1, true);


--
--

SELECT pg_catalog.setval('public.cookie_id_seq', 1, true);


--
--

SELECT pg_catalog.setval('public.courier_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.dishes_id_seq', 697, true);


--
--

SELECT pg_catalog.setval('public.event_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.favorite_restaurant_id_seq', 1, true);


--
--

SELECT pg_catalog.setval('public.general_user_info_id_seq', 1, true);


--
--

SELECT pg_catalog.setval('public.host_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.manager_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.order_list_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.order_radios_list_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.order_structure_list_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.order_user_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.promocode_id_seq', 1, false);


--
--

SELECT pg_catalog.setval('public.radios_id_seq', 59, true);


--
--

SELECT pg_catalog.setval('public.restaurant_category_id_seq', 158, true);


--
--

SELECT pg_catalog.setval('public.restaurant_id_seq', 53, true);


--
--

SELECT pg_catalog.setval('public.review_id_seq', 12, true);


--
--

SELECT pg_catalog.setval('public.structure_dishes_id_seq', 59, true);


--
--

SELECT pg_catalog.setval('public.structure_radios_id_seq', 83, true);


--
--

SELECT pg_catalog.setval('public.worker_id_seq', 1, false);


--
--

ALTER TABLE ONLY public.address_user
    ADD CONSTRAINT address_user_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.card
    ADD CONSTRAINT card_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.cart_food
    ADD CONSTRAINT cart_food_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_client_id_key UNIQUE (client_id);


--
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_client_id_promo_code_key UNIQUE (client_id, promo_code);


--
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_client_id_key UNIQUE (client_id);


--
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.cookie
    ADD CONSTRAINT cookie_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.courier
    ADD CONSTRAINT courier_client_id_key UNIQUE (client_id);


--
--

ALTER TABLE ONLY public.courier
    ADD CONSTRAINT courier_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.dishes
    ADD CONSTRAINT dishes_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.favorite_restaurant
    ADD CONSTRAINT favorite_restaurant_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.general_user_info
    ADD CONSTRAINT general_user_info_email_key UNIQUE (email);


--
--

ALTER TABLE ONLY public.general_user_info
    ADD CONSTRAINT general_user_info_phone_key UNIQUE (phone);


--
--

ALTER TABLE ONLY public.general_user_info
    ADD CONSTRAINT general_user_info_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.host
    ADD CONSTRAINT host_client_id_key UNIQUE (client_id);


--
--

ALTER TABLE ONLY public.host
    ADD CONSTRAINT host_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.manager
    ADD CONSTRAINT manager_client_id_key UNIQUE (client_id);


--
--

ALTER TABLE ONLY public.manager
    ADD CONSTRAINT manager_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.order_list
    ADD CONSTRAINT order_list_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.promocode
    ADD CONSTRAINT promocode_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.radios
    ADD CONSTRAINT radios_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.restaurant_category
    ADD CONSTRAINT restaurant_category_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.restaurant
    ADD CONSTRAINT restaurant_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.review
    ADD CONSTRAINT review_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.structure_dishes
    ADD CONSTRAINT structure_dishes_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.structure_radios
    ADD CONSTRAINT structure_radios_pkey PRIMARY KEY (id);


--
--

ALTER TABLE ONLY public.worker
    ADD CONSTRAINT worker_client_id_key UNIQUE (client_id);


--
--

ALTER TABLE ONLY public.worker
    ADD CONSTRAINT worker_pkey PRIMARY KEY (id);


--
--

CREATE INDEX restaurant_category_fts ON public.restaurant_category USING btree (fts);


--
--

CREATE INDEX restaurant_fts ON public.restaurant USING btree (fts);


--
--

ALTER TABLE ONLY public.address_user
    ADD CONSTRAINT address_user_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.card
    ADD CONSTRAINT card_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_cart_id_fkey FOREIGN KEY (cart_id) REFERENCES public.cart_food(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_radios_fkey FOREIGN KEY (radios) REFERENCES public.structure_radios(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_radios_id_fkey FOREIGN KEY (radios_id) REFERENCES public.radios(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_cart_id_fkey FOREIGN KEY (cart_id) REFERENCES public.cart_food(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_checkbox_fkey FOREIGN KEY (checkbox) REFERENCES public.structure_dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.cookie
    ADD CONSTRAINT cookie_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.courier
    ADD CONSTRAINT courier_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.dishes
    ADD CONSTRAINT dishes_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.favorite_restaurant
    ADD CONSTRAINT favorite_restaurant_client_fkey FOREIGN KEY (client) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.favorite_restaurant
    ADD CONSTRAINT favorite_restaurant_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.host
    ADD CONSTRAINT host_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.manager
    ADD CONSTRAINT manager_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_list
    ADD CONSTRAINT order_list_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_list
    ADD CONSTRAINT order_list_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.order_user(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_list_id_fkey FOREIGN KEY (list_id) REFERENCES public.order_list(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.order_user(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_radios_fkey FOREIGN KEY (radios) REFERENCES public.structure_radios(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_radios_id_fkey FOREIGN KEY (radios_id) REFERENCES public.radios(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_list_id_fkey FOREIGN KEY (list_id) REFERENCES public.order_list(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.order_user(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_structure_food_fkey FOREIGN KEY (structure_food) REFERENCES public.structure_dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_address_id_fkey FOREIGN KEY (address_id) REFERENCES public.address_user(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_courier_id_fkey FOREIGN KEY (courier_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_restaurant_id_fkey FOREIGN KEY (restaurant_id) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.promocode
    ADD CONSTRAINT promocode_free_dish_id_fkey FOREIGN KEY (free_dish_id) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.promocode
    ADD CONSTRAINT promocode_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.radios
    ADD CONSTRAINT radios_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.restaurant_category
    ADD CONSTRAINT restaurant_category_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.restaurant
    ADD CONSTRAINT restaurant_owner_fkey FOREIGN KEY (owner) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.structure_dishes
    ADD CONSTRAINT structure_dishes_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.structure_radios
    ADD CONSTRAINT structure_radios_radios_fkey FOREIGN KEY (radios) REFERENCES public.radios(id) ON DELETE CASCADE;


--
--

ALTER TABLE ONLY public.worker
    ADD CONSTRAINT worker_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
--

GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
--

