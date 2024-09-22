-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS directors (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           age INT NOT NULL
);
CREATE TABLE IF NOT EXISTS movies (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(255) NOT NULL,
                        rating INT NOT NULL,
                        director_id INT NOT NULL,
                        CONSTRAINT fk_director
                            FOREIGN KEY (director_id)
                                REFERENCES directors(id)
                                ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS directors;
-- +goose StatementEnd
