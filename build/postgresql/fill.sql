INSERT INTO general_user_info (name, email, phone, password, salt)
VALUES ('root','root@root', 88888888888,'ca2e080a74ed1590cd141171c20e164d40d058fb45817c7b59f83159d059a6c0', 'salt');

INSERT INTO client (client_id, date_birthday) VALUES (1, NOW());

INSERT INTO restaurant (owner, name, description, price_delivery, min_delivery_time, max_delivery_time, city,
                        street, house, floor, rating, latitude, longitude, avatar) VALUES
      (1, 'Атмосфера', 'description', 250, 15, 90, 'city', 'street', 'house', 100, 5, 1, 1, 'https://img01.rl0.ru/afisha/o/www.afisha.ru/uploads/images/0/c5/0c562e8d522a4e11998692015d184789.jpg'),
      (1, 'Shokolaat', 'description', 10, 25, 65, 'city', 'street', 'house', 100, 3, 1, 1, 'https://media-cdn.tripadvisor.com/media/photo-s/1a/20/47/aa/roof-gastro-bar.jpg'),
      (1, 'Gordon Biersch', 'description', 15, 35, 40, 'city', 'street', 'house', 100, 4, 1, 1, 'https://e-kazan.ru/upload/redactor/images/6e53fa627dd67f3e799d4b840506af55.jpg'),
      (1, 'Crepevine', 'description', 350, 22, 30, 'city', 'street', 'house', 100, 2, 1, 1, 'https://avatars.mds.yandex.net/get-altay/1981910/2a0000016f2eb47cc2b7d68bdc0e5f647348/XXL'),
      (1, 'Creamery', 'description', 250, 10, 55, 'city', 'street', 'house', 100, 1, 1, 1, 'https://avatars.mds.yandex.net/get-altay/1583511/2a0000016f2eb4101852518618c8bb2471ec/XXL'),
      (1, 'Old Pro', 'description', 100, 31, 32, 'city', 'street', 'house', 100, 2.5, 1, 1, 'https://media-cdn.tripadvisor.com/media/photo-s/11/e7/e2/97/view-from-terrace.jpg'),
      (1, 'Дом вкуснятины', 'description', 53, 15, 30, 'city', 'street', 'house', 100, 4.5, 1, 1, 'https://roomester.ru/wp-content/uploads/2019/06/dizajn-restorana1.jpg'),
      (1, 'Продуктовая печь', 'description', 220, 45, 60, 'city', 'street', 'house', 100, 5, 1, 1, 'https://komod-kiev.com/wp-content/uploads/rest-catch-5-1201x800.jpg'),
      (1, 'La Strada', 'description', 121, 17, 20, 'city', 'street', 'house', 100, 3.4, 1, 1, 'https://his.ua/img/articles/P2zlzJUplx.jpg'),
      (1, 'Buca di Beppo', 'description', 150, 15, 45, 'city', 'street', 'house', 100, 2.1, 1, 1, 'https://tomato.ua/image/resize/storage/restaurants/5a75b217e26d5d00431d8ecc/15433380355bfd7833789d8_5bfd782fbcbe68.58579386.jpeg?w=1200&h=1200'),
      (1, 'Мадам Там', 'description', 200, 20, 30, 'city', 'street', 'house', 100, 1.6, 1, 1, 'https://images.restoclub.ru/uploads/article/3/0/d/1/30d10ce7078246f6d9c4487641da10ff_w828_h552--big.jpg'),
      (1, 'Спрут кафе', 'description', 250, 22, 37, 'city', 'street', 'house', 100, 2.3, 1, 1, 'https://peterburg.guide/wp-content/uploads/2019/07/a9LrzH0m.png'),
      (1, 'Bistro Maxine', 'description', 300, 15, 46, 'city', 'street', 'house', 100, 3.1, 1, 1, 'https://www.sobaka.ru/images/image/00/93/54/41/_normal.jpg'),
      (1, 'Три сезона', 'description', 350, 16, 50, 'city', 'street', 'house', 100, 2.7, 1, 1, 'https://100dorog.ru/upload/contents/432/2016/6924462-TOP10_Luchshikh_priv01502302.jpg?width=795&scale=both'),
      (1, 'Спокойствие', 'description', 400, 15, 30, 'city', 'street', 'house', 100, 4.9, 1, 1, 'http://voyagist.ru/wp-content/uploads/2016/12/restoran-evropa-v-sankt-peterburge.jpg'),
      (1, 'Siam Royal', 'description', 450, 25, 44, 'city', 'street', 'house', 100, 3.9, 1, 1, 'https://www.1000ideas.ru/upload/iblock/f55/nPm4YlFbb94.jpg'),
      (1, 'Krung Siam', 'description', 0, 13, 55, 'city', 'street', 'house', 100, 2.31, 1, 1, 'https://img04.rl0.ru/afisha/o/www.afisha.ru/uploads/images/f/90/f90b076437f349aebd57d8a15cf48286.jpg'),
      (1, 'Тайфун', 'description', 0, 20, 56, 'city', 'street', 'house', 100, 1.2, 1, 1, 'https://media-cdn.tripadvisor.com/media/photo-s/12/3c/b5/14/getlstd-property-photo.jpg'),
      (1, 'Tamarine', 'description', 0, 10, 44, 'city', 'street', 'house', 100, 3.4, 1, 1, 'https://b1.m24.ru/c/1168064.jpg'),
      (1, 'Joya', 'description', 0, 19, 33, 'city', 'street', 'house', 100, 2.6, 1, 1, 'https://avatars.mds.yandex.net/get-altay/2776464/2a00000170e89d7d4ffaea7485fe8c72c1e5/XXL'),
      (1, 'Колокольчик', 'description', 499, 16, 47, 'city', 'street', 'house', 100, 4.7, 1, 1, 'https://media-cdn.tripadvisor.com/media/photo-s/19/a3/45/67/caption.jpg'),
      (1, 'Evvia', 'description', 449, 30, 39, 'city', 'street', 'house', 100, 0.5, 1, 1, 'https://media-cdn.tripadvisor.com/media/photo-s/10/38/3b/d0/caption.jpg'),
      (1, 'Кафе 220', 'description', 399, 40, 50, 'city', 'street', 'house', 100, 0.8, 1, 1, 'https://avatars.mds.yandex.net/get-altay/1871013/2a0000016d970de1ac02fd1bab0d585c3d38/XXL'),
      (1, 'Кафе Ренессанс', 'description', 349, 11, 55, 'city', 'street', 'house', 100, 0.1, 1, 1, 'https://salon.ru/storage/thumbs/gallery/525/524912/835_3500_s219.jpg'),
      (1, 'Kan Zeman', 'description', 299, 40, 45, 'city', 'street', 'house', 100, 0.75, 1, 1, 'https://restolife.kz/upload/information_system_5/2/2/9/item_22980/information_items_property_27178.jpg'),
      (1, 'Кафе Манго', 'description', 249, 30, 40, 'city', 'street', 'house', 100, 1.3, 1, 1, 'https://aroma-profi.ru/upload/medialibrary/shop/12po.jpg'),
      (1, 'Балаклава', 'description', 199, 24, 53, 'city', 'street', 'house', 100, 2.34, 1, 1, 'https://www.bankfax.ru/files/_thumbs/resize/files/images/news/2021/10/2021102105_0x700_1634795582.jpg'),
      (1, 'Иностранный гурман', 'description', 149, 18, 32, 'city', 'street', 'house', 100, 1.23, 1, 1, 'https://img.the-village.ru/pOGVpe0KPe1QKnX5TSS1Oq9gJRZsVa6-0-DHmUQbbTw/rs:fill:940:630/q:88/plain/2020/08/31/KKS_0520_1597225612117_rkqgE5c4Zfv.jpg'),
      (1, 'Частичка Бангкока', 'description', 99, 19, 50, 'city', 'street', 'house', 100, 4.1, 1, 1, 'https://robb.report/upload/custom/87a/87aaad790f174c8e07962279f6df7597.jpg'),
      (1, 'Darbar', 'description', 0, 24, 54, 'city', 'street', 'house', 100, 5, 1, 1, 'https://media-cdn.tripadvisor.com/media/photo-s/17/c0/08/c7/caption.jpg'),
      (1, 'Mantra', 'description', 0, 23, 44, 'city', 'street', 'house', 100, 5, 1, 1, 'https://vipdivani.ru/upload/images/restoran-ampir.jpg'),
      (1, 'Janta', 'description', 0, 19, 23, 'city', 'street', 'house', 100, 5, 1, 1, 'https://cdnimg.rg.ru/img/content/181/83/43/01_d_850.jpg'),
      (1, 'Hyderabad', 'description', 0, 25, 50, 'city', 'street', 'house', 100, 5, 1, 1, 'https://cdn.the-village.ru/the-village.ru/post_image-image/EXs8hyK8HWMLPGN8a-YK7w.jpg'),
      (1, 'Кофейня Джека', 'description', 0, 26, 52, 'city', 'street', 'house', 100, 5, 1, 1, 'https://static.ngs.ru/news/2020/99/preview/9fb93df42dde394bf24fab1529978ca81454b591a_599_399_c.jpg'),
      (1, 'Coop кофейня', 'description', 0, 15, 45, 'city', 'street', 'house', 100, 5, 1, 1, 'https://bering-spb.ru/assets/cache_image/photo/main/about/021A6262_1920x0_96c.jpg'),
      (1, 'Lytton Coffee', 'description', 0, 16, 48, 'city', 'street', 'house', 100, 5, 1, 1, 'https://mziurirest.ru/wp-content/uploads/2021/03/DSC_6799-1.jpg'),
      (1, 'Il Fornaio', 'description', 0, 17, 51, 'city', 'street', 'house', 100, 4.5, 1, 1, 'https://brsg.ru/assets/gallery/591.jpg'),
      (1, 'Lavanda', 'description', 0, 18, 54, 'city', 'street', 'house', 100, 3.5, 1, 1, 'https://n1s1.elle.ru/6a/5e/b4/6a5eb4c6039a86a9e0944d83d18c3bd5/728x546_1_249b28fed25c22fdcbacf0314406ba13@1880x1409_0xac120003_7646892161567585889.jpg'),
      (1, 'MacArthur', 'description', 0, 19, 57, 'city', 'street', 'house', 100, 2.5, 1, 1, 'https://topspb.tv/media/768x432/news_covers/vlcsnap-2021-08-02-19h55m22s330e73moay.png'),
      (1, 'Osteria', 'description', 399, 20, 34, 'city', 'street', 'house', 100, 1.5, 1, 1, 'https://lh3.googleusercontent.com/proxy/GdiOjjIJPf-yc-GdAcYy8CS1UUfP95EVYa4cqvioR-RvaqY3HL70wwNdJjl6qthJh3ZXz84H8PNs9KhX6S12HxG8uCABTklxN4g1ffT99nPxrO38heSo3qawvJH8BrfTMgQSBpAh_hPGAE3aHBdipQ'),
      (1, 'Vero', 'description', 499, 20, 40, 'city', 'street', 'house', 100, 0.5, 1, 1, 'https://roomester.ru/wp-content/uploads/2019/06/dizajn-restorana3.jpg'),
      (1, 'Renzo', 'description', 299, 21, 42, 'city', 'street', 'house', 100, 0.4, 1, 1, 'https://drivenew.ru/upload/resize_cache/iblock/1dc/944_629_2/1dc8467631871f72f9e5cdf0f2f185f0.jpg'),
      (1, 'Miyake', 'description', 249, 22, 44, 'city', 'street', 'house', 100, 0.3, 1, 1, 'https://мойбизнес.рф/upload/iblock/fc7/027e015f43ec06fa514ef7554bbde3c8[1].png'),
      (1, 'Tomo', 'description', 199, 23, 46, 'city', 'street', 'house', 100, 0.2, 1, 1, 'https://s3.eu-central-1.amazonaws.com/cdn-eu.jowi.club/jowi.club/ckeditor_assets/pictures/296/content_7.jpg'),
      (1, 'Kanpai', 'description', 149, 24, 36, 'city', 'street', 'house', 100, 0.1, 1, 1, 'https://piteronline.tv/images/2019/07/09/kvartira-kosti-krojtsa_dc4ac.webp'),
      (1, 'Любовь моей жизни', 'description', 266, 30, 45, 'city', 'street', 'house', 100, 5, 1, 1, 'https://restoran-kuvshin.ru/wp-content/uploads/2016/07/3int.jpg'),
      (1, 'Новая пицца', 'description', 233, 31, 46, 'city', 'street', 'house', 100, 4, 1, 1, 'https://n1s1.hsmedia.ru/6f/b5/a0/6fb5a0f016591b1b99b3d595d9543354/728x485_1_43154b801b639f404bcf90a90c8c1023@2880x1920_0xac120003_6030516551597998814.jpeg'),
      (1, 'Калифорнийская кухня', 'description', 150, 23, 32, 'city', 'street', 'house', 100, 3, 1, 1, 'https://restoran-kuvshin.ru/wp-content/uploads/wp-responsive-images-thumbnail-slider/Kuvshin-145cc97388478b4.jpg'),
      (1, 'Круглый стол', 'description', 175, 17, 37, 'city', 'street', 'house', 100, 2, 1, 1, 'https://cdnn21.img.ria.ru/images/144726/45/1447264552_0:498:5616:3657_1920x0_80_0_0_66aa2d824e34eb8bc06049fea22942fd.jpg'),
      (1, 'Любимая шляпа', 'description', 250, 16, 36, 'city', 'street', 'house', 100, 1, 1, 1, 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/a6/%D0%98%D0%BD%D1%82%D0%B5%D1%80%D1%8C%D0%B5%D1%80_%D1%80%D0%B5%D1%81%D1%82%D0%BE%D1%80%D0%B0%D0%BD%D0%B0_White_Rabbit.jpg/1200px-%D0%98%D0%BD%D1%82%D0%B5%D1%80%D1%8C%D0%B5%D1%80_%D1%80%D0%B5%D1%81%D1%82%D0%BE%D1%80%D0%B0%D0%BD%D0%B0_White_Rabbit.jpg'),
      (1, 'Garden Fresh', 'description', 300, 16, 46, 'city', 'street', 'house', 100, 3.2, 1, 1, 'https://n1s1.elle.ru/c1/41/cd/c141cdcca9b54270b6a51f86f5072cff/728x485_1_d5081d5daee74fa9dd6c6924b9d5d1c1@1880x1253_0xac120003_12637632411579102325.jpg'),
      (1, 'Epi', 'description', 150, 16, 56, 'city', 'street', 'house', 100, 2.1, 1, 1, 'https://a-a-ah-ru.s3.amazonaws.com/uploads/items/9414/24911/large_Panoramnye-restorany-Moskvy-TOP-luchshih.jpeg'),
      (1, 'Валентино', 'description', 100, 15, 55, 'city', 'street', 'house', 100, 4.2, 1, 1, 'https://www.gastronom.ru/binfiles/images/20170913/b0d48eb7.jpg')
;

UPDATE restaurant
SET
    fts = to_tsvector(name)
;

INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES (1, '1', NOW(), '');

-- INSERT INTO dishes (name, cost, restaurant, description, protein, falt, kilocalorie, carbohydrates, category_dishes, category_restaurant, count, weight, avatar, place_category, place) VALUES
--     ('Тако', 60, 1, '', 1, 1, 224, 1, 'Горячее', 'Снеки', 1000, 1, 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg', 0, 0),
--     ('Пряник', 60, 1, '', 1, 1, 126, 1, 'К чаю', 'К чаю', 1000, 1, 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg', 1, 0),
--     ('Чёрный бургер', 60, 1, '', 1, 1, 361, 1, 'горячее', 'Снеки', 1000, 1, 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg', 0, 1),
--     ('Пицца Ассорти', 60, 1, '', 1, 1, 1024, 1, 'горячее', 'Пиццы', 1000, 1, 'https://www.koolinar.ru/all_image/recipes/156/156543/recipe_7b4bb7f7-1d42-428a-bb0a-3db8df03093a.jpg', 2, 0),
--     ('Кофе', 60, 1, '', 1, 1, 90, 1, 'горячее', 'Напитки', 1000, 1, 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg', 3, 0),
--     ('Картошка Фри', 60, 1, '', 1, 1, 232, 1, 'горячее', 'Снеки', 1000, 1, 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg', 0, 2),
--     ('Картошка по деревенски', 60, 1, '', 1, 1, 172, 1, 'Горячее', 'Снеки', 1000, 1, 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg', 0, 3),
--     ('МакКомбо', 256, 1, '', 1, 1, 5036, 1, 'Горячее', 'Комбо', 1000, 1, 'https://www.eatthis.com/wp-content/uploads/sites/4/2019/05/mcdonalds-fries-food-lights.jpg', 4, 0),
--     ('Утреннее комбо', 99, 1, '', 1, 1, 4708, 1, 'Горячее', 'Комбо', 1000, 1, 'https://imageproxy.ru/img/crop/1380x920/https/xn--h1ame.xn--80adxhks/storage/app/uploads/public/5e2/700/f07/5e2700f079c4c587329799.jpg', 4, 1),
--     ('Аппетитное комбо', 150, 1, '', 1, 1, 3575, 1, 'Горячее', 'Комбо', 1000, 1, 'https://www.iphones.ru/wp-content/plugins/wonderm00ns-simple-facebook-open-graph-tags/fbimg.php?img=https%3A%2F%2Fwww.iphones.ru%2Fwp-content%2Fuploads%2F2018%2F08%2FBurgerN.jpg', 4, 2),
--     ('Универсальное комбо', 100, 1, '', 1, 1, 1500, 1, 'Горячее', 'Комбо', 1000, 1, 'https://eda.yandex.ru/images/3667559/9724883e03ae48c2b6a1e28c5b9ea111-680x500.jpeg', 4, 3)
--     ;

-- INSERT INTO structure_dishes (name, food, cost, protein, falt, carbohydrates, kilocalorie, count_element, place) VALUES
--     ('Кетчуп', 1, 5, 1, 1, 1, 1, 5, 0),
--     ('Горчица', 1, 5, 1, 1, 1, 1, 5, 1),
--     ('Сырные бортики', 4, 5, 1, 1, 1, 1, 5, 0),
--     ('Колбаса', 4, 5, 1, 1, 1, 1, 5, 1),
--     ('Сыр Пармезан', 4, 5, 1, 1, 1, 1, 5, 2),
--     ('Сыр Моцарелла', 4, 5, 1, 1, 1, 1, 5, 3),
--     ('Сахар', 5, 5, 1, 1, 1, 1, 5, 0),
--     ('Кетчап', 1, 5, 1, 1, 1, 1, 5, 2)
--     ;

-- INSERT INTO radios (name, food, place) VALUES
--     ('МакКомбо', 8, 0),
--     ('Утреннее комбо', 9, 0),
--     ('Аппетитное комбо', 10, 0),
--     ('Универсальное комбо', 11, 0)
--     ;

-- INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie, place) VALUES
--     ('Картофель Фри', 1, 1, 1, 1, 1, 0),
--     ('Картофель по деревенски', 1, 1, 1, 1, 1, 1),
--     ('Сырный соус', 2, 1, 1, 1, 1, 0),
--     ('Чесночный соус', 2, 1, 1, 1, 1, 1),
--     ('Кисло-сладкий соус', 2, 1, 1, 1, 1, 2),
--     ('Картофель Фри', 3, 1, 1, 1, 1, 0),
--     ('Картофель по деревенски', 3, 1, 1, 1, 1, 1),
--     ('Сырный соус', 4, 1, 1, 1, 1, 0),
--     ('Чесночный соус', 4, 1, 1, 1, 1, 1)
--     ;

-- INSERT INTO restaurant_category (restaurant, category, place)
-- VALUES
-- (1, 'Кафе', 0),
-- (1, 'Поп-ап', 1)
-- ;

UPDATE restaurant_category
    SET fts = to_tsvector(category);


INSERT INTO address_user (client_id, alias, comment, city, street, house, floor, flat, porch, intercom, latitude, longitude)
VALUES (1, 'Мой дом', 'Есть злая собака', 'Москва', 'Вязов', 2, 5, 28, 2, '28K', 500, 500);

INSERT INTO promocode (restaurant, name, end_date)
VALUES (1, 'Бесплатно всё', NOW());

INSERT INTO favorite_restaurant (restaurant, client)
VALUES (1, 1);
