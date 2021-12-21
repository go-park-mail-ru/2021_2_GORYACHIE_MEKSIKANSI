--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4 (Ubuntu 13.4-1.pgdg20.04+1)
-- Dumped by pg_dump version 13.4 (Ubuntu 13.4-1.pgdg20.04+1)

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
-- Name: address_user; Type: TABLE; Schema: public; Owner: root
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
-- Name: address_user_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: address_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.address_user_id_seq OWNED BY public.address_user.id;


--
-- Name: card; Type: TABLE; Schema: public; Owner: root
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
-- Name: card_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: card_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.card_id_seq OWNED BY public.card.id;


--
-- Name: cart_food; Type: TABLE; Schema: public; Owner: root
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
-- Name: cart_food_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: cart_food_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cart_food_id_seq OWNED BY public.cart_food.id;


--
-- Name: cart_radios_food; Type: TABLE; Schema: public; Owner: root
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
-- Name: cart_radios_food_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: cart_radios_food_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cart_radios_food_id_seq OWNED BY public.cart_radios_food.id;


--
-- Name: cart_structure_food; Type: TABLE; Schema: public; Owner: root
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
-- Name: cart_structure_food_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: cart_structure_food_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cart_structure_food_id_seq OWNED BY public.cart_structure_food.id;


--
-- Name: cart_user; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.cart_user (
    id integer NOT NULL,
    client_id integer,
    promo_code text,
    restaurant integer
);


ALTER TABLE public.cart_user OWNER TO root;

--
-- Name: cart_user_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: cart_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cart_user_id_seq OWNED BY public.cart_user.id;


--
-- Name: client; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.client (
    id integer NOT NULL,
    client_id integer,
    date_birthday timestamp without time zone
);


ALTER TABLE public.client OWNER TO root;

--
-- Name: client_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: client_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.client_id_seq OWNED BY public.client.id;


--
-- Name: cookie; Type: TABLE; Schema: public; Owner: root
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
-- Name: cookie_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: cookie_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.cookie_id_seq OWNED BY public.cookie.id;


--
-- Name: courier; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.courier (
    id integer NOT NULL,
    client_id integer
);


ALTER TABLE public.courier OWNER TO root;

--
-- Name: courier_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: courier_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.courier_id_seq OWNED BY public.courier.id;


--
-- Name: dishes; Type: TABLE; Schema: public; Owner: root
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
-- Name: dishes_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: dishes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.dishes_id_seq OWNED BY public.dishes.id;


--
-- Name: event; Type: TABLE; Schema: public; Owner: root
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
-- Name: event_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: event_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.event_id_seq OWNED BY public.event.id;


--
-- Name: favorite_restaurant; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.favorite_restaurant (
    id integer NOT NULL,
    client integer,
    restaurant integer,
    "position" integer
);


ALTER TABLE public.favorite_restaurant OWNER TO root;

--
-- Name: favorite_restaurant_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: favorite_restaurant_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.favorite_restaurant_id_seq OWNED BY public.favorite_restaurant.id;


--
-- Name: general_user_info; Type: TABLE; Schema: public; Owner: root
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
-- Name: general_user_info_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: general_user_info_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.general_user_info_id_seq OWNED BY public.general_user_info.id;


--
-- Name: host; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.host (
    id integer NOT NULL,
    client_id integer
);


ALTER TABLE public.host OWNER TO root;

--
-- Name: host_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: host_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.host_id_seq OWNED BY public.host.id;


--
-- Name: manager; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.manager (
    id integer NOT NULL,
    client_id integer
);


ALTER TABLE public.manager OWNER TO root;

--
-- Name: manager_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: manager_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.manager_id_seq OWNED BY public.manager.id;


--
-- Name: order_list; Type: TABLE; Schema: public; Owner: root
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
-- Name: order_list_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: order_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.order_list_id_seq OWNED BY public.order_list.id;


--
-- Name: order_radios_list; Type: TABLE; Schema: public; Owner: root
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
-- Name: order_radios_list_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: order_radios_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.order_radios_list_id_seq OWNED BY public.order_radios_list.id;


--
-- Name: order_structure_list; Type: TABLE; Schema: public; Owner: root
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
-- Name: order_structure_list_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: order_structure_list_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.order_structure_list_id_seq OWNED BY public.order_structure_list.id;


--
-- Name: order_user; Type: TABLE; Schema: public; Owner: root
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
-- Name: order_user_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: order_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.order_user_id_seq OWNED BY public.order_user.id;


--
-- Name: promocode; Type: TABLE; Schema: public; Owner: root
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
    time_for_sale_start time without time zone,
    time_for_sale_finish time without time zone,
    sale_in_time_percent integer,
    sale_in_time_amount integer
);


ALTER TABLE public.promocode OWNER TO root;

--
-- Name: promocode_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: promocode_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.promocode_id_seq OWNED BY public.promocode.id;


--
-- Name: radios; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.radios (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    place integer,
    food integer
);


ALTER TABLE public.radios OWNER TO root;

--
-- Name: radios_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: radios_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.radios_id_seq OWNED BY public.radios.id;


--
-- Name: restaurant; Type: TABLE; Schema: public; Owner: root
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
-- Name: restaurant_category; Type: TABLE; Schema: public; Owner: root
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
-- Name: restaurant_category_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: restaurant_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.restaurant_category_id_seq OWNED BY public.restaurant_category.id;


--
-- Name: restaurant_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: restaurant_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.restaurant_id_seq OWNED BY public.restaurant.id;


--
-- Name: review; Type: TABLE; Schema: public; Owner: root
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
-- Name: review_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: review_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.review_id_seq OWNED BY public.review.id;


--
-- Name: structure_dishes; Type: TABLE; Schema: public; Owner: root
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
-- Name: structure_dishes_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: structure_dishes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.structure_dishes_id_seq OWNED BY public.structure_dishes.id;


--
-- Name: structure_radios; Type: TABLE; Schema: public; Owner: root
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
-- Name: structure_radios_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: structure_radios_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.structure_radios_id_seq OWNED BY public.structure_radios.id;


--
-- Name: worker; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.worker (
    id integer NOT NULL,
    client_id integer
);


ALTER TABLE public.worker OWNER TO root;

--
-- Name: worker_id_seq; Type: SEQUENCE; Schema: public; Owner: root
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
-- Name: worker_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.worker_id_seq OWNED BY public.worker.id;


--
-- Name: address_user id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.address_user ALTER COLUMN id SET DEFAULT nextval('public.address_user_id_seq'::regclass);


--
-- Name: card id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.card ALTER COLUMN id SET DEFAULT nextval('public.card_id_seq'::regclass);


--
-- Name: cart_food id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_food ALTER COLUMN id SET DEFAULT nextval('public.cart_food_id_seq'::regclass);


--
-- Name: cart_radios_food id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_radios_food ALTER COLUMN id SET DEFAULT nextval('public.cart_radios_food_id_seq'::regclass);


--
-- Name: cart_structure_food id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_structure_food ALTER COLUMN id SET DEFAULT nextval('public.cart_structure_food_id_seq'::regclass);


--
-- Name: cart_user id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_user ALTER COLUMN id SET DEFAULT nextval('public.cart_user_id_seq'::regclass);


--
-- Name: client id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.client ALTER COLUMN id SET DEFAULT nextval('public.client_id_seq'::regclass);


--
-- Name: cookie id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cookie ALTER COLUMN id SET DEFAULT nextval('public.cookie_id_seq'::regclass);


--
-- Name: courier id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.courier ALTER COLUMN id SET DEFAULT nextval('public.courier_id_seq'::regclass);


--
-- Name: dishes id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.dishes ALTER COLUMN id SET DEFAULT nextval('public.dishes_id_seq'::regclass);


--
-- Name: event id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.event ALTER COLUMN id SET DEFAULT nextval('public.event_id_seq'::regclass);


--
-- Name: favorite_restaurant id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.favorite_restaurant ALTER COLUMN id SET DEFAULT nextval('public.favorite_restaurant_id_seq'::regclass);


--
-- Name: general_user_info id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.general_user_info ALTER COLUMN id SET DEFAULT nextval('public.general_user_info_id_seq'::regclass);


--
-- Name: host id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.host ALTER COLUMN id SET DEFAULT nextval('public.host_id_seq'::regclass);


--
-- Name: manager id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.manager ALTER COLUMN id SET DEFAULT nextval('public.manager_id_seq'::regclass);


--
-- Name: order_list id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_list ALTER COLUMN id SET DEFAULT nextval('public.order_list_id_seq'::regclass);


--
-- Name: order_radios_list id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_radios_list ALTER COLUMN id SET DEFAULT nextval('public.order_radios_list_id_seq'::regclass);


--
-- Name: order_structure_list id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_structure_list ALTER COLUMN id SET DEFAULT nextval('public.order_structure_list_id_seq'::regclass);


--
-- Name: order_user id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_user ALTER COLUMN id SET DEFAULT nextval('public.order_user_id_seq'::regclass);


--
-- Name: promocode id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.promocode ALTER COLUMN id SET DEFAULT nextval('public.promocode_id_seq'::regclass);


--
-- Name: radios id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.radios ALTER COLUMN id SET DEFAULT nextval('public.radios_id_seq'::regclass);


--
-- Name: restaurant id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.restaurant ALTER COLUMN id SET DEFAULT nextval('public.restaurant_id_seq'::regclass);


--
-- Name: restaurant_category id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.restaurant_category ALTER COLUMN id SET DEFAULT nextval('public.restaurant_category_id_seq'::regclass);


--
-- Name: review id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.review ALTER COLUMN id SET DEFAULT nextval('public.review_id_seq'::regclass);


--
-- Name: structure_dishes id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.structure_dishes ALTER COLUMN id SET DEFAULT nextval('public.structure_dishes_id_seq'::regclass);


--
-- Name: structure_radios id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.structure_radios ALTER COLUMN id SET DEFAULT nextval('public.structure_radios_id_seq'::regclass);


--
-- Name: worker id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.worker ALTER COLUMN id SET DEFAULT nextval('public.worker_id_seq'::regclass);


--
-- Data for Name: address_user; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: card; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: cart_food; Type: TABLE DATA; Schema: public; Owner: root
--

INSERT INTO public.cart_food VALUES (295, 0, 1, 1, 1, 1, 1);


--
-- Data for Name: cart_radios_food; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: cart_structure_food; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: cart_user; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: client; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: cookie; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: courier; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: dishes; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: event; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: favorite_restaurant; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: general_user_info; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: host; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: manager; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: order_list; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: order_radios_list; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: order_structure_list; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: order_user; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: promocode; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: radios; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: restaurant; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: restaurant_category; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: review; Type: TABLE DATA; Schema: public; Owner: root
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
-- Data for Name: structure_dishes; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: structure_radios; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Data for Name: worker; Type: TABLE DATA; Schema: public; Owner: root
--



--
-- Name: address_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.address_user_id_seq', 1, false);


--
-- Name: card_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.card_id_seq', 1, false);


--
-- Name: cart_food_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.cart_food_id_seq', 295, true);


SELECT pg_catalog.setval('public.cart_radios_food_id_seq', 1, false);

SELECT pg_catalog.setval('public.cart_structure_food_id_seq', 1, false);

SELECT pg_catalog.setval('public.cart_user_id_seq', 1, false);



SELECT pg_catalog.setval('public.client_id_seq', 1, false);



SELECT pg_catalog.setval('public.cookie_id_seq', 1, false);


SELECT pg_catalog.setval('public.courier_id_seq', 1, false);


SELECT pg_catalog.setval('public.dishes_id_seq', 1, false);

SELECT pg_catalog.setval('public.event_id_seq', 1, false);


SELECT pg_catalog.setval('public.favorite_restaurant_id_seq', 1, false);


SELECT pg_catalog.setval('public.general_user_info_id_seq', 1, false);



SELECT pg_catalog.setval('public.host_id_seq', 1, false);



SELECT pg_catalog.setval('public.manager_id_seq', 1, false);


SELECT pg_catalog.setval('public.order_list_id_seq', 1, false);


SELECT pg_catalog.setval('public.order_radios_list_id_seq', 1, false);



SELECT pg_catalog.setval('public.order_structure_list_id_seq', 1, false);


SELECT pg_catalog.setval('public.order_user_id_seq', 1, false);


--
-- Name: promocode_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.promocode_id_seq', 1, false);


--
-- Name: radios_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.radios_id_seq', 1, false);


--
-- Name: restaurant_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.restaurant_category_id_seq', 1, false);


--
-- Name: restaurant_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.restaurant_id_seq', 1, false);


--
-- Name: review_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.review_id_seq', 12, true);


--
-- Name: structure_dishes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.structure_dishes_id_seq', 1, false);


--
-- Name: structure_radios_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.structure_radios_id_seq', 1, false);


--
-- Name: worker_id_seq; Type: SEQUENCE SET; Schema: public; Owner: root
--

SELECT pg_catalog.setval('public.worker_id_seq', 1, false);


--
-- Name: address_user address_user_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.address_user
    ADD CONSTRAINT address_user_pkey PRIMARY KEY (id);


--
-- Name: card card_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.card
    ADD CONSTRAINT card_pkey PRIMARY KEY (id);


--
-- Name: cart_food cart_food_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_food
    ADD CONSTRAINT cart_food_pkey PRIMARY KEY (id);


--
-- Name: cart_radios_food cart_radios_food_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_pkey PRIMARY KEY (id);


--
-- Name: cart_structure_food cart_structure_food_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_pkey PRIMARY KEY (id);


--
-- Name: cart_user cart_user_client_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_client_id_key UNIQUE (client_id);


--
-- Name: cart_user cart_user_client_id_promo_code_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_client_id_promo_code_key UNIQUE (client_id, promo_code);


--
-- Name: cart_user cart_user_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_pkey PRIMARY KEY (id);


--
-- Name: client client_client_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_client_id_key UNIQUE (client_id);


--
-- Name: client client_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_pkey PRIMARY KEY (id);


--
-- Name: cookie cookie_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cookie
    ADD CONSTRAINT cookie_pkey PRIMARY KEY (id);


--
-- Name: courier courier_client_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.courier
    ADD CONSTRAINT courier_client_id_key UNIQUE (client_id);


--
-- Name: courier courier_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.courier
    ADD CONSTRAINT courier_pkey PRIMARY KEY (id);


--
-- Name: dishes dishes_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.dishes
    ADD CONSTRAINT dishes_pkey PRIMARY KEY (id);


--
-- Name: event event_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_pkey PRIMARY KEY (id);


--
-- Name: favorite_restaurant favorite_restaurant_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.favorite_restaurant
    ADD CONSTRAINT favorite_restaurant_pkey PRIMARY KEY (id);


--
-- Name: general_user_info general_user_info_email_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.general_user_info
    ADD CONSTRAINT general_user_info_email_key UNIQUE (email);


--
-- Name: general_user_info general_user_info_phone_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.general_user_info
    ADD CONSTRAINT general_user_info_phone_key UNIQUE (phone);


--
-- Name: general_user_info general_user_info_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.general_user_info
    ADD CONSTRAINT general_user_info_pkey PRIMARY KEY (id);


--
-- Name: host host_client_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.host
    ADD CONSTRAINT host_client_id_key UNIQUE (client_id);


--
-- Name: host host_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.host
    ADD CONSTRAINT host_pkey PRIMARY KEY (id);


--
-- Name: manager manager_client_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.manager
    ADD CONSTRAINT manager_client_id_key UNIQUE (client_id);


--
-- Name: manager manager_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.manager
    ADD CONSTRAINT manager_pkey PRIMARY KEY (id);


--
-- Name: order_list order_list_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_list
    ADD CONSTRAINT order_list_pkey PRIMARY KEY (id);


--
-- Name: order_radios_list order_radios_list_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_pkey PRIMARY KEY (id);


--
-- Name: order_structure_list order_structure_list_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_pkey PRIMARY KEY (id);


--
-- Name: order_user order_user_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_pkey PRIMARY KEY (id);


--
-- Name: promocode promocode_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.promocode
    ADD CONSTRAINT promocode_pkey PRIMARY KEY (id);


--
-- Name: radios radios_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.radios
    ADD CONSTRAINT radios_pkey PRIMARY KEY (id);


--
-- Name: restaurant_category restaurant_category_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.restaurant_category
    ADD CONSTRAINT restaurant_category_pkey PRIMARY KEY (id);


--
-- Name: restaurant restaurant_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.restaurant
    ADD CONSTRAINT restaurant_pkey PRIMARY KEY (id);


--
-- Name: review review_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.review
    ADD CONSTRAINT review_pkey PRIMARY KEY (id);


--
-- Name: structure_dishes structure_dishes_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.structure_dishes
    ADD CONSTRAINT structure_dishes_pkey PRIMARY KEY (id);


--
-- Name: structure_radios structure_radios_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.structure_radios
    ADD CONSTRAINT structure_radios_pkey PRIMARY KEY (id);


--
-- Name: worker worker_client_id_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.worker
    ADD CONSTRAINT worker_client_id_key UNIQUE (client_id);


--
-- Name: worker worker_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.worker
    ADD CONSTRAINT worker_pkey PRIMARY KEY (id);


--
-- Name: restaurant_category_fts; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX restaurant_category_fts ON public.restaurant_category USING btree (fts);


--
-- Name: restaurant_fts; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX restaurant_fts ON public.restaurant USING btree (fts);


--
-- Name: address_user address_user_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.address_user
    ADD CONSTRAINT address_user_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: card card_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.card
    ADD CONSTRAINT card_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: cart_radios_food cart_radios_food_cart_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_cart_id_fkey FOREIGN KEY (cart_id) REFERENCES public.cart_food(id) ON DELETE CASCADE;


--
-- Name: cart_radios_food cart_radios_food_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: cart_radios_food cart_radios_food_food_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
-- Name: cart_radios_food cart_radios_food_radios_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_radios_fkey FOREIGN KEY (radios) REFERENCES public.structure_radios(id) ON DELETE CASCADE;


--
-- Name: cart_radios_food cart_radios_food_radios_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_radios_food
    ADD CONSTRAINT cart_radios_food_radios_id_fkey FOREIGN KEY (radios_id) REFERENCES public.radios(id) ON DELETE CASCADE;


--
-- Name: cart_structure_food cart_structure_food_cart_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_cart_id_fkey FOREIGN KEY (cart_id) REFERENCES public.cart_food(id) ON DELETE CASCADE;


--
-- Name: cart_structure_food cart_structure_food_checkbox_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_checkbox_fkey FOREIGN KEY (checkbox) REFERENCES public.structure_dishes(id) ON DELETE CASCADE;


--
-- Name: cart_structure_food cart_structure_food_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: cart_structure_food cart_structure_food_food_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_structure_food
    ADD CONSTRAINT cart_structure_food_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
-- Name: cart_user cart_user_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: cart_user cart_user_restaurant_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cart_user
    ADD CONSTRAINT cart_user_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
-- Name: client client_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.client
    ADD CONSTRAINT client_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: cookie cookie_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.cookie
    ADD CONSTRAINT cookie_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: courier courier_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.courier
    ADD CONSTRAINT courier_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: dishes dishes_restaurant_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.dishes
    ADD CONSTRAINT dishes_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
-- Name: event event_restaurant_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.event
    ADD CONSTRAINT event_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
-- Name: favorite_restaurant favorite_restaurant_client_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.favorite_restaurant
    ADD CONSTRAINT favorite_restaurant_client_fkey FOREIGN KEY (client) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: favorite_restaurant favorite_restaurant_restaurant_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.favorite_restaurant
    ADD CONSTRAINT favorite_restaurant_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


--
-- Name: host host_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.host
    ADD CONSTRAINT host_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: manager manager_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.manager
    ADD CONSTRAINT manager_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


--
-- Name: order_list order_list_food_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_list
    ADD CONSTRAINT order_list_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


--
-- Name: order_list order_list_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_list
    ADD CONSTRAINT order_list_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.order_user(id) ON DELETE CASCADE;


--
-- Name: order_radios_list order_radios_list_food_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_list_id_fkey FOREIGN KEY (list_id) REFERENCES public.order_list(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.order_user(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_radios_fkey FOREIGN KEY (radios) REFERENCES public.structure_radios(id) ON DELETE CASCADE;



ALTER TABLE ONLY public.order_radios_list
    ADD CONSTRAINT order_radios_list_radios_id_fkey FOREIGN KEY (radios_id) REFERENCES public.radios(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_list_id_fkey FOREIGN KEY (list_id) REFERENCES public.order_list(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.order_user(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.order_structure_list
    ADD CONSTRAINT order_structure_list_structure_food_fkey FOREIGN KEY (structure_food) REFERENCES public.structure_dishes(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_address_id_fkey FOREIGN KEY (address_id) REFERENCES public.address_user(id) ON DELETE CASCADE;



ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;

ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_courier_id_fkey FOREIGN KEY (courier_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.order_user
    ADD CONSTRAINT order_user_restaurant_id_fkey FOREIGN KEY (restaurant_id) REFERENCES public.restaurant(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.promocode
    ADD CONSTRAINT promocode_free_dish_id_fkey FOREIGN KEY (free_dish_id) REFERENCES public.dishes(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.promocode
    ADD CONSTRAINT promocode_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.radios
    ADD CONSTRAINT radios_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.restaurant_category
    ADD CONSTRAINT restaurant_category_restaurant_fkey FOREIGN KEY (restaurant) REFERENCES public.restaurant(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.restaurant
    ADD CONSTRAINT restaurant_owner_fkey FOREIGN KEY (owner) REFERENCES public.general_user_info(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.structure_dishes
    ADD CONSTRAINT structure_dishes_food_fkey FOREIGN KEY (food) REFERENCES public.dishes(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.structure_radios
    ADD CONSTRAINT structure_radios_radios_fkey FOREIGN KEY (radios) REFERENCES public.radios(id) ON DELETE CASCADE;


ALTER TABLE ONLY public.worker
    ADD CONSTRAINT worker_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.general_user_info(id) ON DELETE CASCADE;

