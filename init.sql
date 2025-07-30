CREATE TABLE ru_en (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(50) UNIQUE,
                       translation VARCHAR(255)
);

INSERT INTO ru_en (title, translation) VALUES
                                           ('Привет', 'Hello'),
                                           ('Мир', 'World'),
                                           ('Книга', 'Book'),
                                           ('Стол', 'Table'),
                                           ('Яблоко', 'Apple'),
                                           ('Солнце', 'Sun'),
                                           ('Вода', 'Water'),
                                           ('Дом', 'House'),
                                           ('Кот', 'Cat'),
                                           ('Собака', 'Dog'),
                                           ('Человек', 'Human'),
                                           ('Школа', 'School'),
                                           ('Машина', 'Car'),
                                           ('Окно', 'Window'),
                                           ('Ручка', 'Pen');

CREATE TABLE IF NOT EXISTS reports (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

