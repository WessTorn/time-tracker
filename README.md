# time-tracker

## swagger

https://wesstorn.github.io/time-tracker/

## Библиотеки

go get -u github.com/gin-gonic/gin

go get github.com/lib/pq

go get github.com/sirupsen/logrus

go get github.com/joho/godotenv

go get github.com/swaggo/swag

go get -u github.com/swaggo/gin-swagger

go get -u github.com/swaggo/files

## config.env
| cvar               |     value      |           description           |
|:-------------------|:--------------:|:-------------------------------:|
| DB_ADDRESS         | localhost |              address               |
| DB_PORT         | 5432 |              port               |
| DB_USER            | postgres |              user               |
| DB_PASSWORD        | root |              pass               |
| DB_DATABASE            | tz_iul |            database             |
| HOST_URL           | localhost:8080 |              Адрес              |
| LOG_LEVEL          | debug | Уровень лога `info` или `debug` |
| EXTERNAL_API_URL           | http://localhost:8081/info |              Адрес внешнего апи              |
