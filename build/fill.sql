
-- INSERT INTO users (login,user_name, password, role) VALUES ('user1', 'Mark','password1','0');
-- INSERT INTO users (login,user_name, password, role) VALUES ('user2','Michael', 'password2','0');
-- INSERT INTO users (login,user_name, password, role) VALUES ('user3','Vanessa', 'password3','0');
-- INSERT INTO users (login,user_name, password, role) VALUES ('user4','Emily', 'moder','2');
--
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Сатурн', 'Сатурн - шестая планета Солнечной системы, наиболее известная благодаря своим кольцам из льда и камней, которые делают ее уникальной среди других планет. Сатурн также является газовым гигантом с многочисленными спутниками, включая крупнейший - Титан. Несмотря на то, что Сатурн находится на значительном расстоянии от Земли, его потрясающая красота и тайны привлекают учёных и астрономов.', 3389.5, 1200000000, 107, 'http://127.0.0.1:9000/amsflights/saturn.jpg','Планета',false);
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Марс', 'Марс - четвёртая планета от Солнца и ближайшая к Земле внешняя планета. Он известен своим красноватым оттенком, который обусловлен наличием оксида железа на его поверхности. Марс также имеет атмосферу и полярные капюшоны, а исследование этой планеты помогает ученым лучше понять процессы, протекающие на Земле.', 3389.5, 55000000, 37, 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/36/Mars_Valles_Marineris_EDIT.jpg/274px-Mars_Valles_Marineris_EDIT.jpg','Планета',false);
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Луна', 'Луна - естественный спутник Земли, являющийся единственным небесным телом, на котором человек уже побывал. Она имеет покрытую кратерами поверхность и орбитирует вокруг Земли, повышая красоту ночного неба.', 1737.1, 384400, 16.6, 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/83/Moon_nearside_LRO_color_mosaic.png/300px-Moon_nearside_LRO_color_mosaic.png','Спутник',false);


-- INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Тест', 'Тест - естественный спутник Земли, являющийся единственным небесным телом, на котором человек уже побывал. Она имеет покрытую кратерами поверхность и орбитирует вокруг Земли, повышая красоту ночного неба.', 1737.1, 384400, 16.6, 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/83/Moon_nearside_LRO_color_mosaic.png/300px-Moon_nearside_LRO_color_mosaic.png','Спутник',false);

-- INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id, user_login) VALUES (1,'2020-01-01','2022-01-01', '0001-01-01', 'создан', 'AMS123', 1,4,'user1');
-- INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id, user_login) VALUES (2,'2021-02-07','2022-02-01', '0001-01-01', 'в работе', 'AMS456', 2,4,'user2');
-- INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id, user_login) VALUES (3,'2021-04-02','2022-03-01', '0001-01-01', 'в работе', 'AMS789', 3,4,'user3');
-- INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id, user_login) VALUES (4,'2021-05-02','2022-04-01', '0001-01-01', 'в работе', 'AMS789', 4,4,'user4');
--
--
-- INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (1, 1, 1);
-- INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (2, 2, 1);
-- INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (3, 3, 1);
-- INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (3, 1, 1);
-- INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (3, 2, 3);
-- INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (3, 3, 2);




-- ALTER TABLE planets_requests DROP COLUMN created_at;
-- ALTER TABLE planets_requests DROP COLUMN updated_at;
-- ALTER TABLE planets_requests DROP COLUMN deleted_at;
-- ALTER TABLE planets_requests DROP COLUMN id;

-- ALTER TABLE users DROP COLUMN created_at;
-- ALTER TABLE users DROP COLUMN updated_at;
-- ALTER TABLE users DROP COLUMN deleted_at;

-- ALTER TABLE flight_requests DROP COLUMN created_at;
-- ALTER TABLE flight_requests DROP COLUMN updated_at;
-- ALTER TABLE flight_requests DROP COLUMN deleted_at;

-- ALTER TABLE planets DROP COLUMN created_at;
-- ALTER TABLE planets DROP COLUMN updated_at;
-- ALTER TABLE planets DROP COLUMN deleted_at;