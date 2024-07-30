-- CREATE DATABASE go-simple-rpg-api;

CREATE TABLE Player (
    ID UUID PRIMARY KEY,
    Nickname VARCHAR(255) NOT NULL,
    Life INT NOT NULL,
    Attack INT NOT NULL
);

CREATE TABLE Enemy (
    ID UUID PRIMARY KEY,
    Nickname VARCHAR(255) NOT NULL,
    Life INT NOT NULL,
    Attack INT NOT NULL,
    Defense INT NOT NULL
);

CREATE TABLE Battle (
    ID UUID PRIMARY KEY,
    PlayerID UUID NOT NULL,
    EnemyID UUID NOT NULL,
    PlayerName VARCHAR(255) NOT NULL,
    EnemyName VARCHAR(255) NOT NULL,
    Result VARCHAR(50) NOT NULL,
    FOREIGN KEY (PlayerID) REFERENCES Player(ID),
    FOREIGN KEY (EnemyID) REFERENCES Enemy(ID)
);





























