-- +goose Up
CREATE TABLE users (
    id int NOT NULL PRIMARY KEY IDENTITY(1,1),
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    createdAt datetimeoffset NOT NULL,
    updatedAt datetimeoffset,
    deletedAt datetimeoffset
);

-- +goose Down
DROP TABLE users;
