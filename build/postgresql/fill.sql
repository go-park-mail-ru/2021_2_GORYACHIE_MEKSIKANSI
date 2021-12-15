INSERT INTO general_user_info (name, email, phone, password, salt)
VALUES ('root','root@root', 88888888888,'ca2e080a74ed1590cd141171c20e164d40d058fb45817c7b59f83159d059a6c0', 'salt');

INSERT INTO client (client_id, date_birthday) VALUES (1, NOW());

INSERT INTO restaurant (owner, name, description, price_delivery, min_delivery_time, max_delivery_time, city,
                        street, house, floor, rating, latitude, longitude, avatar) VALUES
      (1, 'Атмосфера', 'description', 250, 15, 90, 'city', 'street', 'house', 100, 5, 1, 1, 'https://www.delivery-club.ru/naturmort/6000027_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Shokolaat', 'description', 10, 25, 65, 'city', 'street', 'house', 100, 3, 1, 1, 'https://www.delivery-club.ru/naturmort/5000052_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Gordon Biersch', 'description', 15, 35, 40, 'city', 'street', 'house', 100, 4, 1, 1, 'https://www.delivery-club.ru/naturmort/44000095_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Crepevine', 'description', 350, 22, 30, 'city', 'street', 'house', 100, 2, 1, 1, 'https://www.delivery-club.ru/naturmort/6000035_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Creamery', 'description', 250, 10, 55, 'city', 'street', 'house', 100, 1, 1, 1, 'https://www.delivery-club.ru/naturmort/27000060_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Old Pro', 'description', 100, 31, 32, 'city', 'street', 'house', 100, 2.5, 1, 1, 'https://www.delivery-club.ru/naturmort/5f59d56754805_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Дом вкуснятины', 'description', 53, 15, 30, 'city', 'street', 'house', 100, 4.5, 1, 1, 'https://www.delivery-club.ru/naturmort/5f4a59b84ad69_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Продуктовая печь', 'description', 220, 45, 60, 'city', 'street', 'house', 100, 5, 1, 1, 'https://www.delivery-club.ru/naturmort/26000199_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'La Strada', 'description', 121, 17, 20, 'city', 'street', 'house', 100, 3.4, 1, 1, 'https://www.delivery-club.ru/naturmort/5f62243740d71_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Buca di Beppo', 'description', 150, 15, 45, 'city', 'street', 'house', 100, 2.1, 1, 1, 'https://www.delivery-club.ru/naturmort/1000026_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Мадам Там', 'description', 200, 20, 30, 'city', 'street', 'house', 100, 1.6, 1, 1, 'https://www.delivery-club.ru/naturmort/19000230_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Спрут кафе', 'description', 250, 22, 37, 'city', 'street', 'house', 100, 2.3, 1, 1, 'https://www.delivery-club.ru/naturmort/43000086_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Bistro Maxine', 'description', 300, 15, 46, 'city', 'street', 'house', 100, 3.1, 1, 1, 'https://www.delivery-club.ru/naturmort/61864cc9d00ea_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Три сезона', 'description', 350, 16, 50, 'city', 'street', 'house', 100, 2.7, 1, 1, 'https://www.delivery-club.ru/naturmort/26000199_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Спокойствие', 'description', 400, 15, 30, 'city', 'street', 'house', 100, 4.9, 1, 1, 'https://www.delivery-club.ru/naturmort/5edb9be4ddeba_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Siam Royal', 'description', 450, 25, 44, 'city', 'street', 'house', 100, 3.9, 1, 1, 'https://www.delivery-club.ru/naturmort/5f22cd5325126_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Krung Siam', 'description', 0, 13, 55, 'city', 'street', 'house', 100, 2.31, 1, 1, 'https://www.delivery-club.ru/naturmort/60df0e2fec006_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Тайфун', 'description', 0, 20, 56, 'city', 'street', 'house', 100, 1.2, 1, 1, 'https://www.delivery-club.ru/naturmort/61027eda0d4c0_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Tamarine', 'description', 0, 10, 44, 'city', 'street', 'house', 100, 3.4, 1, 1, 'https://www.delivery-club.ru/naturmort/1000039_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Joya', 'description', 0, 19, 33, 'city', 'street', 'house', 100, 2.6, 1, 1, 'https://www.delivery-club.ru/naturmort/60df0e2fec006_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Колокольчик', 'description', 499, 16, 47, 'city', 'street', 'house', 100, 4.7, 1, 1, 'https://www.delivery-club.ru/naturmort/61780d63510c1_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Evvia', 'description', 449, 30, 39, 'city', 'street', 'house', 100, 0.5, 1, 1, 'https://www.delivery-club.ru/naturmort/60d201e011fd4_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Кафе 220', 'description', 399, 40, 50, 'city', 'street', 'house', 100, 0.8, 1, 1, 'https://mywowo.net/media/images/cache/tokyo_meraviglie_tavola_01_introduzione_jpg_1200_630_cover_85.jpg'),
      (1, 'Кафе Ренессанс', 'description', 349, 11, 55, 'city', 'street', 'house', 100, 0.1, 1, 1, 'https://www.delivery-club.ru/naturmort/48000050_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Kan Zeman', 'description', 299, 40, 45, 'city', 'street', 'house', 100, 0.75, 1, 1, 'https://www.delivery-club.ru/naturmort/2000031_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Кафе Манго', 'description', 249, 30, 40, 'city', 'street', 'house', 100, 1.3, 1, 1, 'https://www.delivery-club.ru/naturmort/25000109_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Балаклава', 'description', 199, 24, 53, 'city', 'street', 'house', 100, 2.34, 1, 1, 'https://www.delivery-club.ru/naturmort/61864cc9d00ea_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Иностранный гурман', 'description', 149, 18, 32, 'city', 'street', 'house', 100, 1.23, 1, 1, 'https://www.delivery-club.ru/naturmort/5ec2443197ff3_480x300.jpg?resize=fill&width=960&height=960&gravity=ce&out=webp'),
      (1, 'Частичка Бангкока', 'description', 99, 19, 50, 'city', 'street', 'house', 100, 4.1, 1, 1, 'https://incrussia.ru/wp-content/uploads/2018/10/iStock-694189032.jpg'),
      (1, 'Darbar', 'description', 0, 24, 54, 'city', 'street', 'house', 100, 5, 1, 1, 'https://sovkusom.ru/wp-content/uploads/blog/v/vrednaya-eda/1.jpg'),
      (1, 'Mantra', 'description', 0, 23, 44, 'city', 'street', 'house', 100, 5, 1, 1, 'https://naked-science.ru/wp-content/uploads/2020/12/fast-fud-pitstsa-burger-chipsy-lukovye-koltsa-kartofel-fri.jpg'),
      (1, 'Janta', 'description', 0, 19, 23, 'city', 'street', 'house', 100, 5, 1, 1, 'https://static.tildacdn.com/tild6561-6165-4337-b835-316638666562/20-05-20.jpg'),
      (1, 'Hyderabad', 'description', 0, 25, 50, 'city', 'street', 'house', 100, 5, 1, 1, 'https://www.kamis-pripravy.ru/upload/medialibrary/907/9073bb8cc5579504bd22a62e5c1fe0e0.jpg'),
      (1, 'Кофейня Джека', 'description', 0, 26, 52, 'city', 'street', 'house', 100, 5, 1, 1, 'https://images.aif.by/007/433/e73337ac5677e37f8baa002e41232ed4.jpg'),
      (1, 'Coop кофейня', 'description', 0, 15, 45, 'city', 'street', 'house', 100, 5, 1, 1, 'https://img.gazeta.ru/files3/829/13377829/Depositphotos_412834214_xl-2015-pic905-895x505-19117.jpg'),
      (1, 'Lytton Coffee', 'description', 0, 16, 48, 'city', 'street', 'house', 100, 5, 1, 1, 'https://cdnmyslo.ru/Photogallery/99/1d/991dffc2-ea20-483e-9352-88cd8e2aa751_b.jpg'),
      (1, 'Il Fornaio', 'description', 0, 17, 51, 'city', 'street', 'house', 100, 4.5, 1, 1, 'https://images.ua.prom.st/3125534192_w600_h600_eda-na-vynos.jpg'),
      (1, 'Lavanda', 'description', 0, 18, 54, 'city', 'street', 'house', 100, 3.5, 1, 1, 'https://incrussia.ru/wp-content/uploads/2020/11/iStock-1175505781.jpg'),
      (1, 'MacArthur', 'description', 0, 19, 57, 'city', 'street', 'house', 100, 2.5, 1, 1, 'https://kidpassage.com/images/publications/eda-sankt-peterburge-chto-poprobovat-skolko-stoit/cover_original.jpg'),
      (1, 'Osteria', 'description', 399, 20, 34, 'city', 'street', 'house', 100, 1.5, 1, 1, 'https://cdn.fishki.net/upload/post/2017/01/30/2205250/2-1485519719-1.jpg'),
      (1, 'Vero', 'description', 499, 20, 40, 'city', 'street', 'house', 100, 0.5, 1, 1, 'https://www.learnathome.ru/files/media/food.jpg'),
      (1, 'Renzo', 'description', 299, 21, 42, 'city', 'street', 'house', 100, 0.4, 1, 1, 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQTvLJHAw98D_0U8xi8fjAN573FWUX42sltrRp2-CkVtQOKUrmoIBP1XyLO5RE_fITY1KQ&usqp=CAU'),
      (1, 'Miyake', 'description', 249, 22, 44, 'city', 'street', 'house', 100, 0.3, 1, 1, 'https://interesnyefakty.org/wp-content/uploads/Interesnye-fakty-o-ede-v-raznyh-stranah.jpg'),
      (1, 'Tomo', 'description', 199, 23, 46, 'city', 'street', 'house', 100, 0.2, 1, 1, 'https://billionnews.ru/timthumb/timthumb.php?src=http://billionnews.ru/uploads/posts/2017-01/thumbs/1485519719_1.jpg&w=940&h=600&zc=1'),
      (1, 'Kanpai', 'description', 149, 24, 36, 'city', 'street', 'house', 100, 0.1, 1, 1, 'https://gorobzor.ru/content/news/2018/06/chto_iz_edy_poprobovat_v_sochi_image_5b2cf79b7278f1.83210187.jpg'),
      (1, 'Любовь моей жизни', 'description', 266, 30, 45, 'city', 'street', 'house', 100, 5, 1, 1, 'https://kidpassage.com/images/publications/eda-sohi-hto-poprobovat-skolko-stoit/cover_original.jpg'),
      (1, 'Новая пицца', 'description', 233, 31, 46, 'city', 'street', 'house', 100, 4, 1, 1, 'https://www.oum.ru/upload/iblock/4a6/4a689562637ffe31a94e1770388395f8.jpg'),
      (1, 'Калифорнийская кухня', 'description', 150, 23, 32, 'city', 'street', 'house', 100, 3, 1, 1, 'https://cs1.livemaster.ru/storage/15/98/6a9751d56360234808ec8ac68anj--kukly-i-igrushki-eda-dlya-kukol-eda-dlya-barbi-kukolnaya-eda-.jpg'),
      (1, 'Круглый стол', 'description', 175, 17, 37, 'city', 'street', 'house', 100, 2, 1, 1, 'https://lh3.googleusercontent.com/proxy/HgfW931vlU8WqU-KdGv8doKW5Re0c1qU6t-EkRfRzehj0c1-eEbSMgbSIZe4e7wVyGOGUNFzGWwaTFZwDkD_bu75cIZm4PhFxJj4WI-S-xXWtwhozr8U'),
      (1, 'Любимая шляпа', 'description', 250, 16, 36, 'city', 'street', 'house', 100, 1, 1, 1, 'https://img.the-village.me/the-village.me/post-cover/-k0NDtajdfoONfacIAqvoA-default.jpg'),
      (1, 'Garden Fresh', 'description', 300, 16, 46, 'city', 'street', 'house', 100, 3.2, 1, 1, 'https://tomato.ua/blog/wp-content/uploads/2019/03/000-39-1-1440x961.jpg'),
      (1, 'Epi', 'description', 150, 16, 56, 'city', 'street', 'house', 100, 2.1, 1, 1, 'https://avatars.mds.yandex.net/get-altay/2960979/2a0000017260a9d9f85eb44d3ab634dd7d7f/XXL'),
      (1, 'Валентино', 'description', 100, 15, 55, 'city', 'street', 'house', 100, 4.2, 1, 1, 'https://i1.wp.com/www.agoda.com/wp-content/uploads/2018/07/Experience-Tokyo_food-and-drink_Featured-image-1200x350_sushi-tray_Tokyo.jpg?fit=1200%2C350&ssl=1')
;

UPDATE restaurant
SET
    fts = to_tsvector(name)
;

INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES (1, '1', NOW(), '');

INSERT INTO dishes (name, cost, restaurant, description, protein, falt, kilocalorie, carbohydrates, category_dishes, category_restaurant, count, weight, avatar, place_category, place) VALUES
    ('Тако', 60, 1, '', 1, 1, 224, 1, 'Горячее', 'Снеки', 1000, 1, 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', 0, 0),
    ('Пряник', 60, 1, '', 1, 1, 126, 1, 'К чаю', 'К чаю', 1000, 1, 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', 1, 0),
    ('Чёрный бургер', 60, 1, '', 1, 1, 361, 1, 'горячее', 'Снеки', 1000, 1, 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', 0, 1),
    ('Пицца Ассорти', 60, 1, '', 1, 1, 1024, 1, 'горячее', 'Пиццы', 1000, 1, 'https://www.koolinar.ru/all_image/recipes/156/156543/recipe_7b4bb7f7-1d42-428a-bb0a-3db8df03093a.jpg', 2, 0),
    ('Кофе', 60, 1, '', 1, 1, 90, 1, 'горячее', 'Напитки', 1000, 1, 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', 3, 0),
    ('Картошка Фри', 60, 1, '', 1, 1, 232, 1, 'горячее', 'Снеки', 1000, 1, 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', 0, 2),
    ('Картошка по деревенски', 60, 1, '', 1, 1, 172, 1, 'Горячее', 'Снеки', 1000, 1, 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', 0, 3),
    ('МакКомбо', 256, 1, '', 1, 1, 5036, 1, 'Горячее', 'Комбо', 1000, 1, 'https://www.eatthis.com/wp-content/uploads/sites/4/2019/05/mcdonalds-fries-food-lights.jpg', 4, 0),
    ('Утреннее комбо', 99, 1, '', 1, 1, 4708, 1, 'Горячее', 'Комбо', 1000, 1, 'https://imageproxy.ru/img/crop/1380x920/https/xn--h1ame.xn--80adxhks/storage/app/uploads/public/5e2/700/f07/5e2700f079c4c587329799.jpg', 4, 1),
    ('Аппетитное комбо', 150, 1, '', 1, 1, 3575, 1, 'Горячее', 'Комбо', 1000, 1, 'https://www.iphones.ru/wp-content/plugins/wonderm00ns-simple-facebook-open-graph-tags/fbimg.php?img=https%3A%2F%2Fwww.iphones.ru%2Fwp-content%2Fuploads%2F2018%2F08%2FBurgerN.jpg', 4, 2),
    ('Универсальное комбо', 100, 1, '', 1, 1, 1500, 1, 'Горячее', 'Комбо', 1000, 1, 'https://eda.yandex.ru/images/3667559/9724883e03ae48c2b6a1e28c5b9ea111-680x500.jpeg', 4, 3)
    ;

INSERT INTO structure_dishes (name, food, cost, protein, falt, carbohydrates, kilocalorie, count_element, place) VALUES
    ('Кетчуп', 1, 5, 1, 1, 1, 1, 5, 0),
    ('Горчица', 1, 5, 1, 1, 1, 1, 5, 1),
    ('Сырные бортики', 4, 5, 1, 1, 1, 1, 5, 0),
    ('Колбаса', 4, 5, 1, 1, 1, 1, 5, 1),
    ('Сыр Пармезан', 4, 5, 1, 1, 1, 1, 5, 2),
    ('Сыр Моцарелла', 4, 5, 1, 1, 1, 1, 5, 3),
    ('Сахар', 5, 5, 1, 1, 1, 1, 5, 0),
    ('Кетчап', 1, 5, 1, 1, 1, 1, 5, 2)
    ;

INSERT INTO radios (name, food, place) VALUES
    ('МакКомбо', 8, 0),
    ('Утреннее комбо', 9, 0),
    ('Аппетитное комбо', 10, 0),
    ('Универсальное комбо', 11, 0)
    ;

INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie, place) VALUES
    ('Картофель Фри', 1, 1, 1, 1, 1, 0),
    ('Картофель по деревенски', 1, 1, 1, 1, 1, 1),
    ('Сырный соус', 2, 1, 1, 1, 1, 0),
    ('Чесночный соус', 2, 1, 1, 1, 1, 1),
    ('Кисло-сладкий соус', 2, 1, 1, 1, 1, 2),
    ('Картофель Фри', 3, 1, 1, 1, 1, 0),
    ('Картофель по деревенски', 3, 1, 1, 1, 1, 1),
    ('Сырный соус', 4, 1, 1, 1, 1, 0),
    ('Чесночный соус', 4, 1, 1, 1, 1, 1)
    ;

INSERT INTO restaurant_category (restaurant, category, place)
VALUES
(1, 'Кафе', 0),
(1, 'Поп-ап', 1)
;

UPDATE restaurant_category
    SET fts = to_tsvector(category);


INSERT INTO address_user (client_id, alias, comment, city, street, house, floor, flat, porch, intercom, latitude, longitude)
VALUES (1, 'Мой дом', 'Есть злая собака', 'Москва', 'Вязов', 2, 5, 28, 2, '28K', 500, 500);

INSERT INTO promocode (id, code, type, restaurant, name, description, end_date, avatar, free_delivery)
VALUES (1, 'Double Time', 1, 1, 'Тук-тук', 'Кто там?', TIMESTAMP '2022-04-28', 'url/photo/', true);

INSERT INTO favorite_restaurant (restaurant, client, position)
VALUES (1, 1, 0);
