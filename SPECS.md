# Specifications

This application instantiates **5** containers via a `docker-compose.yml` file for orchestration.

## Container A - MongoDB

* Runs a `MongoDB` instance which stores the data used for the application.

## Container B - DBMS/DBAdmin

* Runs a `MongoExpress` instance which works as a `DBMS` and allow us to access to the mongo collections in a more graphical way.

## Container C - API Scraper/Fetcher/Consumer

* This container will execute a scraper to fetch the data from the <https://vocadb.net/> web service.

## Container D - FrontEnd

* Runs the frontend GUI.

## Container E - Message Queue Broker

* This container will work as a mediator from the `D` container and the `A` container, so every `POST` request performed from the frontend will pass throught the `RabbitMQ` broker, and then it will be send to the `MongoDB` for its storage
