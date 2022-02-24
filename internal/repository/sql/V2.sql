insert into public.players (id,name,email,mobile) values
('8ef8067b-0480-4b07-a5ff-1a526c2dab7f','player1','player1@gmail.com','12356'),
('383c0598-4a21-42e1-b592-3c5b00bec66d','player2','player2@gmail.com','12356'),
('00bed94d-009f-4e19-bc74-8df934ea3bdf','player3','player3@gmail.com','12356'),
('b1749ac4-58b9-4d88-87e0-a6ab88ac989d','player4','player4@gmail.com','12356'),
('1d2d0ca6-7728-4b5f-8e30-60ab89a273f2','player5','player5@gmail.com','12356'),
('c99122e3-892d-4dfe-b161-89283b9044e7','player6','player6@gmail.com','12356'),
('ebc1effa-681a-449e-9d37-3947e7ca4786','player7','player7@gmail.com','12356'),
('c11eee9c-0847-40a6-8a7f-1308302ae6e6','player8','player8@gmail.com','12356'),
('4c496290-9a69-46bc-8a89-f8d72ed32adb','player9','player9@gmail.com','12356'),
('c8f40618-c751-4f45-98b8-954096e63d3b','player10','player10@gmail.com','12356');

insert into public.games (id,name) values
('e2d8038b-caef-4d91-934f-509d35fe1b4e','game1'),
('36d15764-ea97-4ead-b6c3-db555c15d8b1','game2'),
('6747e6d9-e00f-4674-ae11-9e959324b4c1','game3'),
('b1a37c7f-e7da-436b-a145-dd61560acadf','game4'),
('4b79ce0f-5703-4b51-a090-0df36af213ec','game5'),
('5c55e3c0-2957-4e0e-ab30-5f7e81ed7e3f','game6'),
('0232a5ad-b52b-48aa-8142-8052d7f9fde5','game7'),
('0722276f-dc67-4ade-b37e-fa04ad446610','game8'),
('5d23a6e8-36cd-453a-90eb-da8a9155f97b','game9'),
('e0413b6f-394b-4fa1-9de7-0edd2fc82e2f','game10');

insert into public.scores (id,game_id,player_id,score,region) values
('b7e00d03-b77f-41c4-add6-1269c5b55a40','e2d8038b-caef-4d91-934f-509d35fe1b4e','8ef8067b-0480-4b07-a5ff-1a526c2dab7f',7654,'US'),
('01ddd910-e932-466c-b43e-e9314295c3b6','e2d8038b-caef-4d91-934f-509d35fe1b4e','8ef8067b-0480-4b07-a5ff-1a526c2dab7f',98765,'US'),
('177d3249-03d7-4855-8f04-c18eb7e268e6','e2d8038b-caef-4d91-934f-509d35fe1b4e','8ef8067b-0480-4b07-a5ff-1a526c2dab7f',876,'US'),
('68cc519a-c6aa-40c6-8271-ab8ebd500cea','e2d8038b-caef-4d91-934f-509d35fe1b4e','383c0598-4a21-42e1-b592-3c5b00bec66d',765456,'US'),
('44f4a656-e58a-4667-b9eb-49dfd8f551df','e2d8038b-caef-4d91-934f-509d35fe1b4e','00bed94d-009f-4e19-bc74-8df934ea3bdf',4567,'US'),
('faaee9ba-35ae-4c19-a40b-a1ac0afcd25d','e2d8038b-caef-4d91-934f-509d35fe1b4e','b1749ac4-58b9-4d88-87e0-a6ab88ac989d',0987621,'US'),
('a5fe2334-8908-43cc-aae5-0deb541b09c2','e2d8038b-caef-4d91-934f-509d35fe1b4e','1d2d0ca6-7728-4b5f-8e30-60ab89a273f2',123,'US'),
('3ed1f44f-cff9-4c7e-969d-c05ff675c733','e2d8038b-caef-4d91-934f-509d35fe1b4e','c99122e3-892d-4dfe-b161-89283b9044e7',987654,'US'),
('37ef5bb0-ed8f-4dd4-a3f0-c034687f570f','e2d8038b-caef-4d91-934f-509d35fe1b4e','ebc1effa-681a-449e-9d37-3947e7ca4786',99998,'US'),
('a9f55f4d-10a6-4ba2-845e-726ddd430cba','e2d8038b-caef-4d91-934f-509d35fe1b4e','c11eee9c-0847-40a6-8a7f-1308302ae6e6',999999,'europe'),
('dfb2584c-7f61-4685-9a25-5842c9323a44','e2d8038b-caef-4d91-934f-509d35fe1b4e','383c0598-4a21-42e1-b592-3c5b00bec66d',99998,'US');
