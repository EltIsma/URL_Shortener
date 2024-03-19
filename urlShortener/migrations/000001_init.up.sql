CREATE TABLE urls(
    id bigserial PRIMARY KEY,
    url varchar NOT NULL,
    short_url varchar(15) UNIQUE
);
