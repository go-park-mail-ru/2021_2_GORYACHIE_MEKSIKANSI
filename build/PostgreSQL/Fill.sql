INSERT INTO general_user_info (name, email, phone, password, salt)
VALUES ('root','root@root', 88888888888,'ca2e080a74ed1590cd141171c20e164d40d058fb45817c7b59f83159d059a6c0', 'salt');

INSERT INTO client (client_id, date_birthday) VALUES (1, NOW());

INSERT INTO restaurant (owner, name, description, price_delivery, min_delivery_time, max_delivery_time, city,
                        street, house, floor, rating, latitude, longitude) VALUES
      (1, 'Cheesecake Factory', 'description', 250, 15, 90, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'Shokolaat', 'description', 10, 25, 65, 'city', 'street', 'house', 100, 3, 1, 1),
      (1, 'Gordon Biersch', 'description', 15, 35, 40, 'city', 'street', 'house', 100, 4, 1, 1),
      (1, 'Crepevine', 'description', 350, 22, 30, 'city', 'street', 'house', 100, 2, 1, 1),
      (1, 'Creamery', 'description', 250, 10, 55, 'city', 'street', 'house', 100, 1, 1, 1),
      (1, 'Old Pro', 'description', 100, 31, 30, 'city', 'street', 'house', 100, 2.5, 1, 1),
      (1, 'House of Bagels', 'description', 53, 15, 30, 'city', 'street', 'house', 100, 4.5, 1, 1),
      (1, 'The Prolific Oven', 'description', 220, 45, 60, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'La Strada', 'description', 121, 17, 20, 'city', 'street', 'house', 100, 3.4, 1, 1),
      (1, 'Buca di Beppo', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.1, 1, 1),
      (1, 'Madame Tam', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.6, 1, 1),
      (1, 'Sprout Cafe', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.3, 1, 1),
      (1, 'Bistro Maxine', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.1, 1, 1),
      (1, 'Three Seasons', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.7, 1, 1),
      (1, 'Reposado', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.9, 1, 1),
      (1, 'Siam Royal', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.9, 1, 1),
      (1, 'Krung Siam', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.31, 1, 1),
      (1, 'Thaiphoon', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.2, 1, 1),
      (1, 'Tamarine', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.4, 1, 1),
      (1, 'Joya', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.6, 1, 1),
      (1, 'Jing Jing', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.7, 1, 1),
      (1, 'Evvia Estiatorio', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.5, 1, 1),
      (1, 'Cafe 220', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.8, 1, 1),
      (1, 'Cafe Renaissance', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.1, 1, 1),
      (1, 'Kan Zeman', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.75, 1, 1),
      (1, 'Mango Caribbean Cafe', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.3, 1, 1),
      (1, 'Baklava', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.34, 1, 1),
      (1, 'Mandarin Gourmet', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.23, 1, 1),
      (1, 'Bangkok Cuisine', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.1, 1, 1),
      (1, 'Darbar Indian Cuisine', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'Mantra', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'Janta', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'Hyderabad House', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'Starbucks', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'Coupa Cafe', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'Lytton Coffee Company', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'Il Fornaio', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.5, 1, 1),
      (1, 'Lavanda', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.5, 1, 1),
      (1, 'MacArthur Park', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.5, 1, 1),
      (1, 'Osteria', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.5, 1, 1),
      (1, 'Vero', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.5, 1, 1),
      (1, 'Cafe Renzo', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.4, 1, 1),
      (1, 'Miyake', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.3, 1, 1),
      (1, 'Sushi Tomo', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.2, 1, 1),
      (1, 'Kanpai', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.1, 1, 1),
      (1, 'Pizza My Heart', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 1, 1),
      (1, 'New York Pizza', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4, 1, 1),
      (1, 'California Pizza Kitchen', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3, 1, 1),
      (1, 'Round Table', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2, 1, 1),
      (1, 'Loving Hut', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1, 1, 1),
      (1, 'Garden Fresh', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.2, 1, 1),
      (1, 'Cafe Epi', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.1, 1, 1),
      (1, 'Tai Pan', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.2, 1, 1)
;

UPDATE restaurant
SET
    avatar = (
            array[
            'https://nypost.com/wp-content/uploads/sites/2/2020/05/mcdonalds-feature.jpg?quality=90&strip=all',
            'https://top-samyh.ru/assets/components/phpthumbof/cache/alias4.8ec1866f31359b98f52e8c1d06fc6bcb.1eabef41af95455b89f319c7ba68f516.jpg',
            'https://kod.ru/content/images/2020/11/KFC_has_unveiled_a_new-27a410c3f16231bf8ac7977a566697c4.png',
            'https://avatars.mds.yandex.net/get-zen_doc/51081/pub_5f96d7f1bc35081b5203ce74_5f96dd8924d0d15a6614547a/scale_1200',
            'https://pokatim.ru/uploads/posts/2020-08/1598606636_rpwswmq3ullhobrcmp5evrhuifhl6x5k0nnt8dda.jpeg',
            'https://ligabiznesa.ru/wp-content/uploads/2020/01/ris.-1.-logotip-dodo-picca.jpg',
            'https://skidka-na-prazdnik.ru/wp-content/uploads/2020/06/xxl.jpg',
            'https://avatars.mds.yandex.net/get-zen_doc/1878571/pub_5d1f02fb24e56600ad2b65d5_5d1f036ff221ef00adfa7d8f/scale_1200'
                ]
        ) [floor(random() * 8 + 1)]
;

INSERT INTO cookie (client_id, session_id, date_life, csrf_token) VALUES (1, '1', NOW(), '');

INSERT INTO dishes (name, cost, restaurant, description, protein, falt, kilocalorie, carbohydrates, category_dishes, category_restaurant, count, weight, avatar) VALUES
    ('Тако', 60, 1, '', 1, 1, 224, 1, 'Горячее', 'Снеки', 1000, 1, 'https://pbs.twimg.com/media/DtAriH3U8AAD3jV.jpg'),
    ('Пряник', 60, 1, '', 1, 1, 126, 1, 'К чаю', 'К чаю', 1000, 1, 'https://s3.amazonaws.com/images.ecwid.com/images/38011115/2135671623.jpg'),
    ('Чёрный бургер', 60, 1, '', 1, 1, 361, 1, 'горячее', 'Снеки', 1000, 1, 'https://meat-pepper.ru/image/cache/catalog/products/burgers/kotlety-dly-burgerov-black-angus-3-800x667.jpg'),
    ('Пицца Ассорти', 60, 1, '', 1, 1, 1024, 1, 'горячее', 'Пиццы', 1000, 1, 'https://www.koolinar.ru/all_image/recipes/156/156543/recipe_7b4bb7f7-1d42-428a-bb0a-3db8df03093a.jpg'),
    ('Кофе', 60, 1, '', 1, 1, 90, 1, 'горячее', 'Напитки', 1000, 1, 'https://traveltimes.ru/wp-content/uploads/2021/08/kofe-caska-penka-scaled.jpg'),
    ('Картошка Фри', 60, 1, '', 1, 1, 232, 1, 'горячее', 'Снеки', 1000, 1, 'https://вести35.рф/images/2020/07/06/5f2775ffddc94d76a57605479b3f02e0.jpg'),
    ('Картошка по деревенски', 60, 1, '', 1, 1, 172, 1, 'Горячее', 'Снеки', 1000, 1, 'https://cherkessk.crazybrothers.ru/wp-content/uploads/Kartofel-po-derevenski.jpg'),
    ('МакКомбо', 256, 1, '', 1, 1, 5036, 1, 'Горячее', 'Комбо', 1000, 1, 'https://www.eatthis.com/wp-content/uploads/sites/4/2019/05/mcdonalds-fries-food-lights.jpg'),
    ('Утреннее комбо', 99, 1, '', 1, 1, 4708, 1, 'Горячее', 'Комбо', 1000, 1, 'https://imageproxy.ru/img/crop/1380x920/https/xn--h1ame.xn--80adxhks/storage/app/uploads/public/5e2/700/f07/5e2700f079c4c587329799.jpg'),
    ('Аппетитное комбо', 150, 1, '', 1, 1, 3575, 1, 'Горячее', 'Комбо', 1000, 1, 'https://www.iphones.ru/wp-content/plugins/wonderm00ns-simple-facebook-open-graph-tags/fbimg.php?img=https%3A%2F%2Fwww.iphones.ru%2Fwp-content%2Fuploads%2F2018%2F08%2FBurgerN.jpg'),
    ('Универсальное комбо', 100, 1, '', 1, 1, 1500, 1, 'Горячее', 'Комбо', 1000, 1, 'https://eda.yandex.ru/images/3667559/9724883e03ae48c2b6a1e28c5b9ea111-680x500.jpeg')
    ;

INSERT INTO structure_dishes (name, food, cost, protein, falt, carbohydrates, kilocalorie, count_element) VALUES
    ('Кетчуп', 1, 5, 1, 1, 1, 1, 5),
    ('Горчица', 1, 5, 1, 1, 1, 1, 5),
    ('Сырные бортики', 4, 5, 1, 1, 1, 1, 5),
    ('Колбаса', 4, 5, 1, 1, 1, 1, 5),
    ('Сыр Пармезан', 4, 5, 1, 1, 1, 1, 5),
    ('Сыр Моцарелла', 4, 5, 1, 1, 1, 1, 5),
    ('Сахар', 5, 5, 1, 1, 1, 1, 5),
    ('Кетчуп', 1, 5, 1, 1, 1, 1, 5)
    ;

INSERT INTO radios (name, food) VALUES
    ('МакКомбо', 8),
    ('Утреннее комбо', 9),
    ('Аппетитное комбо', 10),
    ('Универсальное комбо', 11)
    ;

INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie) VALUES
    ('Картофель Фри', 1, 1, 1, 1, 1),
    ('Картофель по деревенски', 1, 1, 1, 1, 1),
    ('Сырный соус', 2, 1, 1, 1, 1),
    ('Чесночный соус', 2, 1, 1, 1, 1),
    ('Кисло-сладкий соус', 2, 1, 1, 1, 1)
    ;

INSERT INTO dishes (name, cost, restaurant, description, protein, falt, kilocalorie, carbohydrates, category_dishes, category_restaurant, count, weight)
SELECT 'name', random() * (500 - 10) + 10, random() * (53 - 1 - 1) + 1 + 1, 'description', 1, 1, random() * (1000 - 1) + 1, 1, 'text', 'text', random() * (101 - 1) - 1, random() * (51 - 1) + 1 FROM generate_series(1, 1000);
UPDATE dishes
SET
    name = (
            array[
                'Бургер',
            'Двойной бургер',
            'Картофель',
            'Сибас',
            'Икра',
            'Белый хлеб',
            'Молоко',
            'Кефир',
            'Печенье',
            'Арахис',
            'Колбаса',
            'Сыр',
            'Пицца Пеппрони',
            'Пицца Ассорти',
            'Пицца Четыре сыра',
            'Кока-кола',
            'Фанта',
            'Спрайт',
            'Чёрый чай',
            'Кофе',
            'Кофе (негорячий)',
            'Тако',
            'БигМак',
            'Стрипсы',
            'Мороженое',
            'Красный рак'
                ]
        ) [floor(random() * 26 + 1)],
    category_dishes = (
            array[
                'Суп',
                'Вода',
                'Почти кола',
                'Арбузы',
                'Что-то похожее на еду'
                ]
           ) [floor(random() * 5 + 1)],
    category_restaurant = (
            array[
                'Вкусно',
                'Напитки',
                'Пиццы',
                'Суши',
                'Десерт',
                'Закуски'
                ]
           ) [floor(random() * 6 + 1)],
    description = (
            array[
                'Большое комбо. При изменении комплектации цена комбо может измениться.',
                'Вкуснейшее произведение кулинарии, которое отправит вас в гастрономический оргазм',
                'Самое вкусное блюдо, после которого вы захотите ещё в 0.5 раз больше!',
                'Красивейшее оформление блюдо подчерквиает мастерство нашего повара - мистера Пеннивайз',
                'Думаешь это вкусно? Правильно думаешь! Бери его и получишь его же!'
                ]
           ) [floor(random() * 5 + 1)],
    avatar = (
            array[
            'https://s.fishki.net/upload/users/2019/11/29/1518141/85ae61606d59ee2b2701adeed30a1d71.png',
            'https://pbs.twimg.com/media/EBGtDaiVAAAv8Nh.jpg',
            'https://avatars.mds.yandex.net/i?id=8190d1c6f0a87d9dcc258676e69b5018-2431862-images-thumbs&n=13',
            'https://avatars.mds.yandex.net/get-altay/1031166/2a00000162001941c99c64ec1cf8a9fa2edd/XXL',
            'https://moscow-restaurants.ru/netcat_files/38/26/1391/IMG_9985.jpeg_800.jpg',
            'https://www.iphones.ru/wp-content/plugins/wonderm00ns-simple-facebook-open-graph-tags/fbimg.php?img=https%3A%2F%2Fwww.iphones.ru%2Fwp-content%2Fuploads%2F2018%2F08%2FBurgerN.jpg',
            'https://imageproxy.ru/img/crop/1380x920/https/xn--h1ame.xn--80adxhks/storage/app/uploads/public/5e2/700/f07/5e2700f079c4c587329799.jpg',
            'https://islam.ru/sites/default/files/img/2017/news/90-main-image-w1480.jpg',
            'https://www.eatthis.com/wp-content/uploads/sites/4/2019/05/mcdonalds-fries-food-lights.jpg',
            'https://s1.1zoom.ru/big7/529/Holidays_Christmas_399885.jpg',
            'https://pbs.twimg.com/media/EVy0CxpXYAAZrD4.jpg:large',
                'https://sun9-51.userapi.com/impf/c636323/v636323212/36406/KDUTCfKG8b0.jpg?size=867x1080&quality=96&sign=fe39ffdc81aff59212e58cd37348f5e9&c_uniq_tag=oOwnCg7Lulpo9y6f4R_lk1fuUkM3QkuTNRk0ekN92Vk&type=album'
            ]
            ) [floor(random() * 12 + 1)]
WHERE id > 12
;

INSERT INTO restaurant_category (restaurant, category)
SELECT random() * (53 - 1) + 1, 'name' FROM generate_series(1, 350);
UPDATE restaurant_category
SET
    category = (
            array[
                'Хороший',
                'Лучший',
                'Единственный',
                'Суши-бар',
                'Кальянная',
                'Пиццерия',
                'Бар',
                'Хенкальная',
                'Общепит',
                'Кафе',
                'Буфеты',
                'Поп-ап',
                'Виртуальный'
                ]
        ) [floor(random() * 13 + 1)]
;

INSERT INTO radios (name, food)
SELECT 'name', random() * (1000 - 1) + 12 FROM generate_series(1, 1000);
UPDATE radios
SET
    name = (
            array[
                'Комплимент от повара',
                'Бургер',
                'Напиток',
                'Секретный ингредиент',
                'Что-то к чаю',
                'Снеки'
                ]
        ) [floor(random() * 6 + 1)]
WHERE id > 4
;

INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie)
SELECT 'name', random() * (1000 - 1) + 1 + 3, 1, 1, 1, 1 FROM generate_series(1, 1000);
UPDATE structure_radios
SET
    name = (
                array[
                    'Кола-ванила',
                    'Картофель фри',
                    'Коктейль',
                    'Вкусный бургер',
                    'Яблоки',
                    'Семга',
                    'Форель'
                    ]
        ) [floor(random() * 7 + 1)]
WHERE id > 9
;

INSERT INTO structure_dishes (name, food, cost, protein, falt, carbohydrates, kilocalorie, count_element)
SELECT 'name', random() * (1000 - 1) + 12, random() * (25 - 1) + 1, 1, 1, 1, 1, random() * (5 - 1) + 1 FROM generate_series(1, 2000);
UPDATE structure_dishes
SET
    name = (
                array[
                    'Сахар',
                    'Соль',
                    'Кетчуп',
                    'Лук',
                    'Морковь',
                    'Котлета',
                    'Огурец',
                    'Помидор'
                    ]
        ) [floor(random() * 8 + 1)]
WHERE id > 8
;

INSERT INTO address_user (client_id, alias, comment, city, street, house, floor, flat, porch, intercom, latitude, longitude)
VALUES (1, 'Мой дом', 'Есть злая собака', 'Москва', 'Вязов', 2, 5, 28, 2, '28K', 500, 500);

INSERT INTO promocode (restaurant, name, end_date)
VALUES (1, 'Бесплатно всё', NOW());