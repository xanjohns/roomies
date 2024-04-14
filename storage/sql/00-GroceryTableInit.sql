CREATE TABLE IF NOT EXISTS GroceryList(
    listID int unique NOT NULL,
    itemID int unique NOT NULL,
    groupID int NOT NULL,
    listItem varchar(255) NULL,
    addedByID varchar(255) NULL,
    recurring boolean NULL,
    recurringDate datetime(3) NULL,
    timestamp Date NOT NULL,
    PRIMARY KEY(listID)
);

CREATE TABLE IF NOT EXISTS Users(
    userID varchar(255) NOT NULL,
    groupID varchar(255) NOT NULL,
    firstName varchar(255) NOT NULL,
    lastName varchar(255) NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL,
    PRIMARY KEY(userID)
);
