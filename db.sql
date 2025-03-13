CREATE ROLE match_rw;

CREATE USER match_user WITH PASSWORD 'swifty';
GRANT match_rw TO match_user;

CREATE DATABASE match_db WITH OWNER = match_rw;
CREATE SCHEMA match_schema;

CREATE TABLE match_schema.user (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    email VARCHAR
);


GRANT USAGE ON SCHEMA match_schema to match_rw;
GRANT SELECT, INSERT, UPDATE, DELETE ON match_schema.user TO match_rw;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA match_schema TO match_rw;



-- INSERT INTO match_schema.user (name, email) VALUES('Martin', 'martinmail@arizona.edu');
-- SELECT * FROM match_schema.user;