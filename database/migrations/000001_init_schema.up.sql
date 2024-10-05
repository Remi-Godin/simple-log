CREATE TABLE users(
    UserId SERIAL NOT NULL PRIMARY KEY,
    FirstName VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    PasswordHash VARCHAR(255) NOT NULL
);

CREATE TABLE logbooks(
    LogbookId SERIAL NOT NULL PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    OwnedBy INTEGER NOT NULL REFERENCES users(UserId) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE entries(
    EntryId SERIAL NOT NULL PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Description VARCHAR(5000) NOT NULL,
    CreatedOn TIMESTAMP DEFAULT NOW() NOT NULL,
    CreatedBy INTEGER NOT NULL REFERENCES users(UserId) ON DELETE CASCADE ON UPDATE CASCADE,
    LogbookId INTEGER NOT NULL REFERENCES logbooks(LogbookId) ON DELETE CASCADE ON UPDATE CASCADE
);
