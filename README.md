
# Test task for shop.by - "Currency Service"

## Task


>Напишите сервис валютчика:
>
>Раз в сутки должны собираться данные всех курсов валют nbrb по отношению к бел.рублю, https://api.nbrb.by/exrates/rates?periodicity=0;
>
>Реализовывать Http API сервер: 
>1) метод получения всех записей собранных вами;
>2) метод получения записи за выбранный день;
>Технологии: Mysql, Golang

## Requirements

- Linux
- [Docker](https://www.docker.com/)

## Demo

![Demo](/docs/demo.gif)

## About solution

When run the application, an HTTP server is started on local port `8080` and a cron job for fetching currency is set to run every midnight.

The received exchange rates are stored in the `rates` table in MySQL.

The following addresses: `/rates` and `/rates/:date` can be used to retrieve the records.

Minimal API test coverage. Test data is populated using fixtures.

Docker is used for ease of deployment.

See more developer command in [Taskfile.yml](/Taskfile.yml).

## Implementation details

 - [Gin](https://gin-gonic.com/) - for easy setup HTTP server and routings.
 - [Cron](https://github.com/robfig/cron) - for cron jobs.
 - [Gorm](https://gorm.io/) - ORM for works with DB.
 - [Testify](https://github.com/stretchr/testify) and [go-testfixtures](github.com/go-testfixtures/testfixtures) - for testing.
 - [Task](https://github.com/go-task/task) - for commands alias.

## Running the Project

1. Clone the repository:
    ```bash
    git clone https://github.com/abelapko/shop-by-test.git currency-service
    cd currency-service
    ```

2. Install tool for run commands
    ```bash
    sudo wget https://github.com/go-task/task/releases/download/v3.40.1/task_linux_amd64.tar.gz -O /tmp/task_linux_amd64.tar.gz && sudo tar -xz -C /usr/local/bin -f /tmp/task_linux_amd64.tar.gz && sudo rm -f /tmp/task_linux_amd64.tar.gz
    ```

3. Set up environment variables:
    ```bash
    task set-env
    ```

4. Init and run app:
    ```bash
    task app-init
    ```    

5. Force fetching currency rates:
    ```bash
    task fetch
    ```

6.  Use the server:
    - Go to [http://localhost:8080/rates](http://localhost:8080/rates) to get all currency rates.
    - For getting the rate for a specific day: [http://localhost:8080/rates/2025-01-05](http://localhost:8080/rates/2025-01-05).
    
7. Run tests:
   ```bash
   task test
   ```