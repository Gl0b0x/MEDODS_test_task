CREATE table users (
    uuid UUID PRIMARY KEY,
    email varchar(50) not null,
    ip inet,
    refresh_token varchar(200)
);

INSERT into users (uuid, email, ip) values
('7ad7c029-1ae4-411f-b96d-6ea566866311', 'vanyakyz@mail.ru', '45.135.180.220'),
('fef1022a-5952-4237-96d4-c12ec113839f', 'ivankuz3@yandex.ru', '213.87.144.70'),
('0184b4a5-4391-446c-a7c1-0449336ed3b1', 'nepisatmne80@gmail.com', '95.220.56.245');