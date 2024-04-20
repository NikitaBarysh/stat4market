Запуск docker-compose up 

Задание 1 

В файле sql.txt

Задание 2

Вставка тестовых данных реализована в методе fakeData

Получение по заданному eventType и временному диапазону
    get  localhost:8080/get/{from}--{to}/eventType
    Example: localhost:8080/get/2015-02-01--2025-02-01/login P.s между датой два дефиса (--)

Задание 3
    post  localhost:8080/api/event