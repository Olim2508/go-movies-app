-- +goose Up
-- +goose StatementBegin
CREATE TABLE genres (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL
);

ALTER TABLE movies
    ADD COLUMN genre_id INT,
ADD CONSTRAINT fk_genre
    FOREIGN KEY (genre_id)
    REFERENCES genres(id)
    ON DELETE SET NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE movies
DROP CONSTRAINT fk_genre,
DROP COLUMN genre_id;

DROP TABLE IF EXISTS genres;
-- +goose StatementEnd
