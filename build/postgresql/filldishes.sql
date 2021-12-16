INSERT INTO dishes (name, cost, restaurant, description, protein, falt, kilocalorie, carbohydrates, category_dishes, category_restaurant, count, weight, avatar, place_category, place) VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
    RETURNING id;

INSERT INTO structure_dishes (name, food, cost, protein, falt, carbohydrates, kilocalorie, count_element, place) VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, $9)
;

INSERT INTO radios (name, food, place) VALUES
    ($1, $2, $3)
    RETURNING id
;

INSERT INTO structure_radios (name, radios, protein, falt, carbohydrates, kilocalorie, place) VALUES
    ($1, $2, $3, $4, $5, $6, $7)
;
