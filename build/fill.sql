
INSERT INTO users (login, password, is_admin) VALUES ('user1', 'password1',false);
INSERT INTO users (login, password, is_admin) VALUES ('user2', 'password2',false);
INSERT INTO users (login, password, is_admin) VALUES ('user3', 'password3',false);
INSERT INTO users (login, password, is_admin) VALUES ('user4', 'moder',true);

INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Сатурн', 'Сатурн - шестая планета Солнечной системы, наиболее известная благодаря своим кольцам из льда и камней, которые делают ее уникальной среди других планет. Сатурн также является газовым гигантом с многочисленными спутниками, включая крупнейший - Титан. Несмотря на то, что Сатурн находится на значительном расстоянии от Земли, его потрясающая красота и тайны привлекают учёных и астрономов.', 3389.5, 1200000000, 107, 'http://172.18.0.6:9000/amsflights/saturn.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=66V5TUU991OFSUBCI48Y%2F20231014%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231014T214233Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiI2NlY1VFVVOTkxT0ZTVUJDSTQ4WSIsImV4cCI6MTY5NzMyMzEwMCwicGFyZW50IjoibWluaW8ifQ.FtOVDZLuOj5NQllP9lei-HIwEYi409lWj3di4arS1_bC9j6snsftvdbXveqoY_XB_mCrvdqpHPutIxxqvCSsWw&X-Amz-SignedHeaders=host&versionId=12ab384c-8992-4535-b6fd-a19e8030838c&X-Amz-Signature=2ccdc4647f9c2cff661c062aa757d070f65225f0a0009928eac436c6ae2a5ca7','Планета',false);
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Марс', 'Марс - четвёртая планета от Солнца и ближайшая к Земле внешняя планета. Он известен своим красноватым оттенком, который обусловлен наличием оксида железа на его поверхности. Марс также имеет атмосферу и полярные капюшоны, а исследование этой планеты помогает ученым лучше понять процессы, протекающие на Земле.', 3389.5, 55000000, 37, 'http://172.18.0.6:9000/amsflights/mars.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=66V5TUU991OFSUBCI48Y%2F20231014%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231014T214219Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiI2NlY1VFVVOTkxT0ZTVUJDSTQ4WSIsImV4cCI6MTY5NzMyMzEwMCwicGFyZW50IjoibWluaW8ifQ.FtOVDZLuOj5NQllP9lei-HIwEYi409lWj3di4arS1_bC9j6snsftvdbXveqoY_XB_mCrvdqpHPutIxxqvCSsWw&X-Amz-SignedHeaders=host&versionId=622c7aa6-32c6-4ae6-9f69-97b5232bc574&X-Amz-Signature=335d2e64b03b3ede11b7128e046107e3fcbe39b8711cdba81fd3c463f426a09d','Планета',false);
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Луна', 'Луна - естественный спутник Земли, являющийся единственным небесным телом, на котором человек уже побывал. Она имеет покрытую кратерами поверхность и орбитирует вокруг Земли, повышая красоту ночного неба.', 1737.1, 384400, 16.6, 'http://172.18.0.6:9000/amsflights/moon.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=66V5TUU991OFSUBCI48Y%2F20231014%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231014T214254Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiI2NlY1VFVVOTkxT0ZTVUJDSTQ4WSIsImV4cCI6MTY5NzMyMzEwMCwicGFyZW50IjoibWluaW8ifQ.FtOVDZLuOj5NQllP9lei-HIwEYi409lWj3di4arS1_bC9j6snsftvdbXveqoY_XB_mCrvdqpHPutIxxqvCSsWw&X-Amz-SignedHeaders=host&versionId=56364c60-bbca-44ef-9bef-45b9e81d3d84&X-Amz-Signature=172ec073126bcbca01eca312af045dd4ab51db01322690f68456fb4eaacf5523','Спутник',false);

INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id) VALUES (1,'2020-01-01','2022-01-01', '2022-01-10', 'создан', 'AMS123', 1,4);
INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id) VALUES (2,'2021-02-07','2022-02-01', '2022-02-10', 'создан', 'AMS456', 2,4);
INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id) VALUES (3,'2021-04-02','2022-03-01', '2022-03-10', 'создан', 'AMS789', 3,4);
INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id) VALUES (4,'2021-05-02','2022-04-01', '2022-04-10', 'в работе', 'AMS789', 4,4);


INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (1, 1, 1);
INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (2, 2, 1);
INSERT INTO planets_requests (fr_id, planet_id, flight_number) VALUES (3, 3, 1);
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