CREATE TABLE IF NOT EXISTS nexoapp (
   greetings VARCHAR ( 50 ) UNIQUE NOT NULL,
   username VARCHAR ( 50 ) UNIQUE NOT NULL
);

INSERT INTO nexoapp (greetings, username) VALUES ('Greetings from Postgre, Nexo :)', 'Anton Chernaev');

/* 
Delete row if needed: DELETE FROM nexoapp WHERE username='Anton Chernaev'; 
Login to DB: psql --username postgres -d nexo
*/