
INSERT INTO users (login, password, is_admin) VALUES ('user1', 'password1',false);
INSERT INTO users (login, password, is_admin) VALUES ('user2', 'password2',false);
INSERT INTO users (login, password, is_admin) VALUES ('user3', 'password3',false);
INSERT INTO users (login, password, is_admin) VALUES ('user4', 'moder',true);

INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Сатурн', 'Сатурн - шестая планета Солнечной системы, наиболее известная благодаря своим кольцам из льда и камней, которые делают ее уникальной среди других планет. Сатурн также является газовым гигантом с многочисленными спутниками, включая крупнейший - Титан. Несмотря на то, что Сатурн находится на значительном расстоянии от Земли, его потрясающая красота и тайны привлекают учёных и астрономов.', 3389.5, 1200000000, 107, 'http://172.20.0.5:9000/amsflights/saturn.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=HOI8P8C0YJVTC9AR2JMA%2F20231206%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231206T074509Z&X-Amz-Expires=604799&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJIT0k4UDhDMFlKVlRDOUFSMkpNQSIsImV4cCI6MTcwMTg1MTI0MSwicGFyZW50IjoibWluaW8ifQ.47r8eZVXwLXvGfK5bjaICX0HLyKbcnlBxiTlTLpP0cZhbaGmrA-fEvk40K24aeOSLYaeXXO8MhpjHZaIxRaMew&X-Amz-SignedHeaders=host&versionId=12ab384c-8992-4535-b6fd-a19e8030838c&X-Amz-Signature=afdd02cc15d6f0e8c87b0d36831fb899465e8a26676c8e9b233227db80c8e9bd','Планета',false);
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Марс', 'Марс - четвёртая планета от Солнца и ближайшая к Земле внешняя планета. Он известен своим красноватым оттенком, который обусловлен наличием оксида железа на его поверхности. Марс также имеет атмосферу и полярные капюшоны, а исследование этой планеты помогает ученым лучше понять процессы, протекающие на Земле.', 3389.5, 55000000, 37, 'http://172.20.0.5:9000/amsflights/mars.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=HOI8P8C0YJVTC9AR2JMA%2F20231206%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231206T073304Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJIT0k4UDhDMFlKVlRDOUFSMkpNQSIsImV4cCI6MTcwMTg1MTI0MSwicGFyZW50IjoibWluaW8ifQ.47r8eZVXwLXvGfK5bjaICX0HLyKbcnlBxiTlTLpP0cZhbaGmrA-fEvk40K24aeOSLYaeXXO8MhpjHZaIxRaMew&X-Amz-SignedHeaders=host&versionId=622c7aa6-32c6-4ae6-9f69-97b5232bc574&X-Amz-Signature=f6095e04ea159823dc5686ce8a586f017af57dd874d9ea1b57dad471e87d8620','Планета',false);
INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Луна', 'Луна - естественный спутник Земли, являющийся единственным небесным телом, на котором человек уже побывал. Она имеет покрытую кратерами поверхность и орбитирует вокруг Земли, повышая красоту ночного неба.', 1737.1, 384400, 16.6, 'http://172.20.0.5:9000/amsflights/moon.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=HOI8P8C0YJVTC9AR2JMA%2F20231206%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231206T073135Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJIT0k4UDhDMFlKVlRDOUFSMkpNQSIsImV4cCI6MTcwMTg1MTI0MSwicGFyZW50IjoibWluaW8ifQ.47r8eZVXwLXvGfK5bjaICX0HLyKbcnlBxiTlTLpP0cZhbaGmrA-fEvk40K24aeOSLYaeXXO8MhpjHZaIxRaMew&X-Amz-SignedHeaders=host&versionId=56364c60-bbca-44ef-9bef-45b9e81d3d84&X-Amz-Signature=69184d96c39803076a86f2a1dcbb390ef5578d35bdbb80dfdda7ee05644343a2','Спутник',false);


INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id, user_login) VALUES (1,'2020-01-01','2022-01-01', '0001-01-01', 'создан', 'AMS123', 1,4,'user1');
INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id, user_login) VALUES (2,'2021-02-07','2022-02-01', '0001-01-01', 'в работе', 'AMS456', 2,4,'user2');
INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id, user_login) VALUES (3,'2021-04-02','2022-03-01', '0001-01-01', 'в работе', 'AMS789', 3,4,'user3');
INSERT INTO flight_requests (id,date_create,date_formation, date_completion, status, AMS, user_id,moder_id, user_login) VALUES (4,'2021-05-02','2022-04-01', '0001-01-01', 'в работе', 'AMS789', 4,4,'user4');


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



-- export const mockPlanets: IPlanet[] = [
--     {id: 1, name: 'Сатурн', radius: 3389.5,distance:1200000000,gravity:107, is_delete:false, description: 'Сатурн - шестая планета Солнечной системы, наиболее известная благодаря своим кольцам из льда и камней, которые делают ее уникальной среди других планет. Сатурн также является газовым гигантом с многочисленными спутниками, включая крупнейший - Титан. Несмотря на то, что Сатурн находится на значительном расстоянии от Земли, его потрясающая красота и тайны привлекают учёных и астрономов.', image_url: 'http://172.18.0.6:9000/amsflights/saturn.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=66V5TUU991OFSUBCI48Y%2F20231014%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231014T214233Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiI2NlY1VFVVOTkxT0ZTVUJDSTQ4WSIsImV4cCI6MTY5NzMyMzEwMCwicGFyZW50IjoibWluaW8ifQ.FtOVDZLuOj5NQllP9lei-HIwEYi409lWj3di4arS1_bC9j6snsftvdbXveqoY_XB_mCrvdqpHPutIxxqvCSsWw&X-Amz-SignedHeaders=host&versionId=12ab384c-8992-4535-b6fd-a19e8030838c&X-Amz-Signature=2ccdc4647f9c2cff661c062aa757d070f65225f0a0009928eac436c6ae2a5ca7'},
--     {id: 2, name: 'Марс', radius: 3389.5,distance:55000000,gravity:37, is_delete:false, description: 'Марс - четвёртая планета от Солнца и ближайшая к Земле внешняя планета. Он известен своим красноватым оттенком, который обусловлен наличием оксида железа на его поверхности. Марс также имеет атмосферу и полярные капюшоны, а исследование этой планеты помогает ученым лучше понять процессы, протекающие на Земле.', image_url: 'http://172.18.0.6:9000/amsflights/mars.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=66V5TUU991OFSUBCI48Y%2F20231014%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231014T214219Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiI2NlY1VFVVOTkxT0ZTVUJDSTQ4WSIsImV4cCI6MTY5NzMyMzEwMCwicGFyZW50IjoibWluaW8ifQ.FtOVDZLuOj5NQllP9lei-HIwEYi409lWj3di4arS1_bC9j6snsftvdbXveqoY_XB_mCrvdqpHPutIxxqvCSsWw&X-Amz-SignedHeaders=host&versionId=622c7aa6-32c6-4ae6-9f69-97b5232bc574&X-Amz-Signature=335d2e64b03b3ede11b7128e046107e3fcbe39b8711cdba81fd3c463f426a09d'},
--     {id: 3, name: 'Луна', radius: 1737.1,distance:384400,gravity:16.6, is_delete:false, description: 'Луна - естественный спутник Земли, являющийся единственным небесным телом, на котором человек уже побывал. Она имеет покрытую кратерами поверхность и орбитирует вокруг Земли, повышая красоту ночного неба.', image_url: 'http://172.18.0.6:9000/amsflights/moon.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=66V5TUU991OFSUBCI48Y%2F20231014%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231014T214254Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiI2NlY1VFVVOTkxT0ZTVUJDSTQ4WSIsImV4cCI6MTY5NzMyMzEwMCwicGFyZW50IjoibWluaW8ifQ.FtOVDZLuOj5NQllP9lei-HIwEYi409lWj3di4arS1_bC9j6snsftvdbXveqoY_XB_mCrvdqpHPutIxxqvCSsWw&X-Amz-SignedHeaders=host&versionId=56364c60-bbca-44ef-9bef-45b9e81d3d84&X-Amz-Signature=172ec073126bcbca01eca312af045dd4ab51db01322690f68456fb4eaacf5523'},
-- ]
--
--
-- INSERT INTO planets (name, description,radius, distance, gravity, image, type,is_delete) VALUES ('Test', 'Test', 111, 222, 333, 'http://172.20.0.5:9000/amsflights/moon.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=HOI8P8C0YJVTC9AR2JMA%2F20231206%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20231206T073135Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiJIT0k4UDhDMFlKVlRDOUFSMkpNQSIsImV4cCI6MTcwMTg1MTI0MSwicGFyZW50IjoibWluaW8ifQ.47r8eZVXwLXvGfK5bjaICX0HLyKbcnlBxiTlTLpP0cZhbaGmrA-fEvk40K24aeOSLYaeXXO8MhpjHZaIxRaMew&X-Amz-SignedHeaders=host&versionId=56364c60-bbca-44ef-9bef-45b9e81d3d84&X-Amz-Signature=69184d96c39803076a86f2a1dcbb390ef5578d35bdbb80dfdda7ee05644343a2','Test',false);
