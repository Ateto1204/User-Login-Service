CREATE TABLE IF NOT EXISTS `user` (
        `email` VARCHAR(100) PRIMARY KEY,
        `name` VARCHAR(100),
        `pwd` VARCHAR(100) NOT NULL
);