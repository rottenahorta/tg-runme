package repo

var initUsersDB = `CREATE TABLE IF NOT EXISTS users(
    id SERIAL NOT NULL UNIQUE,
    chatid int NOT NULL UNIQUE,
    zptoken VARCHAR(255),
)`