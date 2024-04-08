CREATE TABLE IF NOT EXISTS catalog (
    reg_num VARCHAR(255) PRIMARY KEY,
    mark VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    owner_name VARCHAR(255) NOT NULL,
    owner_surname VARCHAR(255) NOT NULL,
    owner_patronymic VARCHAR(255)
);
