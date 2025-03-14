CREATE ROLE match_rw;

CREATE USER match_user WITH PASSWORD 'swifty';
GRANT match_rw TO match_user;

CREATE DATABASE match_db WITH OWNER = match_rw;

\connect match_db;
CREATE SCHEMA  IF NOT EXISTS match_schema;

CREATE TABLE  IF NOT EXISTS match_schema.user (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    email VARCHAR
);


GRANT USAGE ON SCHEMA match_schema TO match_rw;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA match_schema TO match_rw;
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA match_schema TO match_rw;



-- INSERT INTO match_schema.user (name, email) VALUES('Martin', 'martinmail@arizona.edu');
-- SELECT * FROM match_schema.user;