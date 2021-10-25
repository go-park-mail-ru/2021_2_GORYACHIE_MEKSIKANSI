INSERT INTO general_user_info (name, email, phone, password, salt)
VALUES ('root','root@root', 88888888888,'ca2e080a74ed1590cd141171c20e164d40d058fb45817c7b59f83159d059a6c0', 'salt');

INSERT INTO client (client_id, date_birthday) VALUES (1, NOW());

INSERT INTO restaurant (owner, name, description, price_delivery, min_delivery_time, max_delivery_time, city,
                        street, house, floor, rating, location) VALUES
      (1, 'Cheesecake Factory', 'description', 250, 15, 90, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'Shokolaat', 'description', 10, 25, 65, 'city', 'street', 'house', 100, 3, 'location'),
      (1, 'Gordon Biersch', 'description', 15, 35, 40, 'city', 'street', 'house', 100, 4, 'location'),
      (1, 'Crepevine', 'description', 350, 22, 30, 'city', 'street', 'house', 100, 2, 'location'),
      (1, 'Creamery', 'description', 250, 10, 55, 'city', 'street', 'house', 100, 1, 'location'),
      (1, 'Old Pro', 'description', 100, 31, 30, 'city', 'street', 'house', 100, 2.5, 'location'),
      (1, 'House of Bagels', 'description', 53, 15, 30, 'city', 'street', 'house', 100, 4.5, 'location'),
      (1, 'The Prolific Oven', 'description', 220, 45, 60, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'La Strada', 'description', 121, 17, 20, 'city', 'street', 'house', 100, 3.4, 'location'),
      (1, 'Buca di Beppo', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.1, 'location'),
      (1, 'Madame Tam', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.6, 'location'),
      (1, 'Sprout Cafe', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.3, 'location'),
      (1, 'Bistro Maxine', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.1, 'location'),
      (1, 'Three Seasons', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.7, 'location'),
      (1, 'Reposado', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.9, 'location'),
      (1, 'Siam Royal', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.9, 'location'),
      (1, 'Krung Siam', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.31, 'location'),
      (1, 'Thaiphoon', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.2, 'location'),
      (1, 'Tamarine', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.4, 'location'),
      (1, 'Joya', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.6, 'location'),
      (1, 'Jing Jing', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.7, 'location'),
      (1, 'Evvia Estiatorio', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.5, 'location'),
      (1, 'Cafe 220', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.8, 'location'),
      (1, 'Cafe Renaissance', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.1, 'location'),
      (1, 'Kan Zeman', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.75, 'location'),
      (1, 'Mango Caribbean Cafe', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.3, 'location'),
      (1, 'Baklava', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.34, 'location'),
      (1, 'Mandarin Gourmet', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.23, 'location'),
      (1, 'Bangkok Cuisine', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.1, 'location'),
      (1, 'Darbar Indian Cuisine', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'Mantra', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'Janta', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'Hyderabad House', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'Starbucks', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'Coupa Cafe', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'Lytton Coffee Company', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'Il Fornaio', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.5, 'location'),
      (1, 'Lavanda', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.5, 'location'),
      (1, 'MacArthur Park', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.5, 'location'),
      (1, 'Osteria', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1.5, 'location'),
      (1, 'Vero', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.5, 'location'),
      (1, 'Cafe Renzo', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.4, 'location'),
      (1, 'Miyake', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.3, 'location'),
      (1, 'Sushi Tomo', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.2, 'location'),
      (1, 'Kanpai', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 0.1, 'location'),
      (1, 'Pizza My Heart', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 5, 'location'),
      (1, 'New York Pizza', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4, 'location'),
      (1, 'California Pizza Kitchen', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3, 'location'),
      (1, 'Round Table', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2, 'location'),
      (1, 'Loving Hut', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 1, 'location'),
      (1, 'Garden Fresh', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 3.2, 'location'),
      (1, 'Cafe Epi', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 2.1, 'location'),
      (1, 'Tai Pan', 'description', 0, 15, 30, 'city', 'street', 'house', 100, 4.2, 'location')
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

INSERT INTO dishes (name, cost, restaurant, description, protein, falt, kilocalorie, carbohydrates, category_dishes, category_restaurant, count)
SELECT 'name', random() * (500 - 10) + 10, random() * (53 - 1) + 1, 'description', 1, 1, 1, 1, 'text', 'text', random() * (51 - 1) - 1 FROM generate_series(1, 1000);
UPDATE dishes
SET
    name = (
            array[
                'Бургер',
            'Двойной бургер',
            'Тройной бургер',
            'Четверной бургер',
            'Пятерной бургер',
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
            'Тако'
                ]
        ) [floor(random() * 25 + 1)],
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
                'Что-то похожее на еду',
                'Пиццы',
                'Суши',
                'Десерт'
                ]
           ) [floor(random() * 6 + 1)],
    description = (
            array[
                'Описание Бургера',
                'Описание Вкусной еды',
                'Описание Данного предмета',
                'Описание Четверного бургера',
                'Описание Пятерного бургера'
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
            'https://worldofmeat.ru/wp-content/uploads/2020/09/64777fa0b34f8985bdbfbaa242baccdd.jpg'
            ]
            ) [floor(random() * 8 + 1)]
;

INSERT INTO restaurant_category (restaurant, category)
SELECT random() * (53 - 1) + 1, 'name' FROM generate_series(1, 500);
UPDATE restaurant_category
SET
    category = (
            array[
                'Хороший',
                'Лучший',
                'Единственный',
                'Кака',
                'Лучший, т.к. единственный',
                'Хз',
                'Призрак',
                'Суши-бар',
                'Кальянная',
                'Пиццерия',
                'Бар',
                'Хенкальная',
                'Общепит',
                'Тестовый'
                ]
        ) [floor(random() * 14 + 1)]
;

INSERT INTO radios (name, food)
SELECT 'name', random() * (1000 - 1) + 1 FROM generate_series(1, 1000);
UPDATE radios
SET
    name = (
            array[
                'Комплимент от повара',
                'Бургер',
                'Напиток',
                'Секретный ингредиент',
                'Что-то к чаю',
                'Снеки',
                'В одном из них яд. Угадай где)'
                ]
        ) [floor(random() * 7 + 1)]
;

INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie)
SELECT 'name', random() * (1000 - 1) + 1, 1, 1, 1, 1 FROM generate_series(1, 1000);
UPDATE structure_radios
SET
    name = (
                array[
                    'Кола-ванила',
                    'Картофель фри',
                    'Коктейль',
                    'Вкусный бургер',
                    'Волос повара',
                    'Яблоки',
                    'Кошка',
                    'Кака'
                    ]
        ) [floor(random() * 8 + 1)]
;

INSERT INTO structure_dishes (name, food, cost, protein, falt, carbohydrates, kilocalorie, count_element)
SELECT 'name', random() * (1000 - 1) + 1, random() * (25 - 1) + 1, 1, 1, 1, 1, random() * (5 - 1) + 1 FROM generate_series(1, 2000);
UPDATE structure_dishes
SET
    name = (
                array[
                    'Сахар',
                    'Соль',
                    'Кетчап',
                    'Лук',
                    'Морковь',
                    'Котлета',
                    'Огурец',
                    'Помидор'
                    ]
        ) [floor(random() * 8 + 1)]
;

INSERT INTO cart (client_id, food, count_food, number_item, restaurant_id) VALUES (1, 1, 1, 1, 1);
INSERT INTO cart_structure_food (client_id, checkbox) VALUES (1, 1);
INSERT INTO cart_radios_food (client_id, radios_id, radios) VALUES (1, 1, 1);
