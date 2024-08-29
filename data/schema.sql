DROP TABLE IF EXISTS DrawingAccessLog;
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
  FirstName  TEXT NOT NULL,
  LastName  TEXT NOT NULL,
  Email TEXT UNIQUE,
  PasswordHash TEXT NOT NULL
);

CREATE TABLE DrawingAccessLog(
  ID TEXT PRIMARY KEY,
  UserID TEXT NOT NULL,
  DrawingID TEXT NOT NULL,
  AccessedAt INTEGER
  FOREIGN KEY(UserID) REFERENCES Users(ID),
  FOREIGN KEY(DrawingID) REFERENCES Drawings(ID)
);