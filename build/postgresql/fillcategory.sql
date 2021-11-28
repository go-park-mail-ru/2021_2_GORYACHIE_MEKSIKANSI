INSERT INTO restaurant_category (restaurant, category, place)
VALUES
    ($1, $2, $3)
;

UPDATE restaurant_category
SET fts = to_tsvector(category);
