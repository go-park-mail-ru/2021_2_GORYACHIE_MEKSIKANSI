INSERT INTO general_user_info (name, email, phone, password, salt)
VALUES ('root','root@root', 88888888888,'ca2e080a74ed1590cd141171c20e164d40d058fb45817c7b59f83159d059a6c0', 'salt');

INSERT INTO client (client_id, date_birthday) VALUES (1, NOW());

INSERT INTO restaurant (owner, name, description, price_delivery, min_delivery_time, max_delivery_time, city,
                        street, house, floor, rating, location) VALUES
      (1, 'Cheesecake Factory', 'description', 100, 15, 90, 'city', 'street', 'house', 100, 5, 'location'),
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

INSERT INTO restaurant_category (restaurant, category) VALUES (1, 'category');
INSERT INTO restaurant_category (restaurant, category) VALUES (1, 'text');

INSERT INTO dishes (name, cost, restaurant, description, protein, falt, kilocalorie, carbohydrates, category_dishes, category_restaurant) VALUES
      ('name', 0, 1, 'descr', 1, 1, 1, 1, 'cat_dis', 'cat_rest');

INSERT INTO dishes (name, cost, restaurant, description, protein, falt, kilocalorie, carbohydrates, category_dishes, category_restaurant) VALUES
      ('name2', 0, 1, 'descr2', 2, 2, 2, 2, 'cat_dis2', 'cat_rest2');

INSERT INTO radios (name, food) VALUES ('name', 1);
INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie) VALUES ('name', 1, 1, 1, 1, 1);
INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie) VALUES ('name2', 1, 1, 1, 1, 1);
INSERT INTO radios (name, food) VALUES ('name2', 1);
INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie) VALUES ('name', 2, 1, 1, 1, 1);
INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie) VALUES ('name2', 2, 1, 1, 1, 1);
INSERT INTO structure_dishes (name, food, cost, protein, falt, carbohydrates, kilocalorie, count_element) VALUES ('name', 1, 1, 1, 1, 1, 1, 1);
INSERT INTO structure_dishes (name, food, cost, protein, falt, carbohydrates, kilocalorie, count_element) VALUES ('name2', 1, 2, 2, 2, 2, 2, 2);
