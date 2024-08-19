DROP TABLE IF EXISTS Drawings;
DROP TABLE IF EXISTS Collections;
DROP TABLE IF EXISTS Users;

CREATE TABLE Collections(
  Name TEXT UNIQUE, 
  ID TEXT PRIMARY KEY
);

CREATE TABLE Drawings(
  Name TEXT,
  CollectionID TEXT, 
  ID TEXT PRIMARY KEY,
  Data TEXT, 
  FOREIGN KEY(CollectionID) REFERENCES Collections(ID),
  UNIQUE(Name, CollectionID)
);

CREATE TABLE Users(
  ID TEXT PRIMARY KEY,
  FirstName TEXT,
  LastName TEXT,
  Email TEXT UNIQUE,
  PasswordHash TEXT NOT NULL
);