CREATE TABLE user
(
    id    BIGSERIAL PRIMARY KEY NOT NULL,
    email VARCHAR(255)          NOT NULL,
    age   BIGINT                NOT NULL
);

INSERT INTO user (email, age)
VALUES ('user1@example.com', 25),
       ('user2@example.com', 32),
       ('user3@example.com', 40),
       ('user4@example.com', 28),
       ('user5@example.com', 37),
       ('user6@example.com', 19),
       ('user7@example.com', 42),
       ('user8@example.com', 31),
       ('user9@example.com', 26),
       ('user10@example.com', 35);
