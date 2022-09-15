-- +goose Up
-- +goose StatementBegin
CREATE TABLE spendings (
    id int NOT NULL PRIMARY KEY IDENTITY(1,1),
    userId int not null,
    value decimal(10,2) NOT NULL,
    createdAt datetimeoffset NOT NULL,
    updatedAt datetimeoffset,

    foreign key (userId) references users(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE spendings
-- +goose StatementEnd
