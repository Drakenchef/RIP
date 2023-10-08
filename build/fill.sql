
INSERT INTO users (login, password, is_admin) VALUES ('user1', 'password1',false);
INSERT INTO users (login, password, is_admin) VALUES ('user2', 'password2',false);
INSERT INTO users (login, password, is_admin) VALUES ('user3', 'moder',true);

INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Сатурн', 'Сатурн - шестая планета Солнечной системы, наиболее известная благодаря своим кольцам из льда и камней, которые делают ее уникальной среди других планет. Сатурн также является газовым гигантом с многочисленными спутниками, включая крупнейший - Титан. Несмотря на то, что Сатурн находится на значительном расстоянии от Земли, его потрясающая красота и тайны привлекают учёных и астрономов.', 3389.5, 1200000000, 107, 'saturn.jpg','Планета',false);
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Марс', 'Марс - четвёртая планета от Солнца и ближайшая к Земле внешняя планета. Он известен своим красноватым оттенком, который обусловлен наличием оксида железа на его поверхности. Марс также имеет атмосферу и полярные капюшоны, а исследование этой планеты помогает ученым лучше понять процессы, протекающие на Земле.', 3389.5, 55000000, 37, 'mars.jpg','Планета',false);
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Луна', 'Луна - естественный спутник Земли, являющийся единственным небесным телом, на котором человек уже побывал. Она имеет покрытую кратерами поверхность и орбитирует вокруг Земли, повышая красоту ночного неба.', 1737.1, 384400, 16.6, 'moon.jpg','Спутник',false);

INSERT INTO flight_requests (date_create,date_formation, date_completion, status, AMS, user_id,moder_id) VALUES ('2020-01-01','2022-01-01', '2022-01-10', 'существует', 'AMS123', 1,3);
INSERT INTO flight_requests (date_create,date_formation, date_completion, status, AMS, user_id,moder_id) VALUES ('2021-02-07','2022-02-01', '2022-02-10', 'существует', 'AMS456', 2,3);
INSERT INTO flight_requests (date_create,date_formation, date_completion, status, AMS, user_id,moder_id) VALUES ('2021-04-02','2022-03-01', '2022-03-10', 'существует', 'AMS789', 2,3);

INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (1, 1, 1);
INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (2, 2, 1);
INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (2, 1, 2);
INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (3, 1, 1);
INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (3, 2, 3);
INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (3, 3, 2);




-- ALTER TABLE planets_requests DROP COLUMN created_at;
-- ALTER TABLE planets_requests DROP COLUMN updated_at;
-- ALTER TABLE planets_requests DROP COLUMN deleted_at;
ALTER TABLE planets_requests DROP COLUMN id;

-- ALTER TABLE users DROP COLUMN created_at;
-- ALTER TABLE users DROP COLUMN updated_at;
-- ALTER TABLE users DROP COLUMN deleted_at;

-- ALTER TABLE flight_requests DROP COLUMN created_at;
-- ALTER TABLE flight_requests DROP COLUMN updated_at;
-- ALTER TABLE flight_requests DROP COLUMN deleted_at;

-- ALTER TABLE planets DROP COLUMN created_at;
-- ALTER TABLE planets DROP COLUMN updated_at;
-- ALTER TABLE planets DROP COLUMN deleted_at;