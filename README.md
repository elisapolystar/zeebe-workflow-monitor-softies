# Team Softikset – Zeebe Workflow Monitor

## About

Zeebe Workflow Monitor is an application designed for the purpose of visualization of the workflows and monitoring their execution paths on the Zeebe. 
The application imports the data from Zeebe through Zeebe Kafka Exporter. The application subscribes to the Kafka topics, consumes the data, parses it, and stores it into a database. The data can then be displayed in the React frontend.   

## Prerequisites

* Docker Engine version 18.02.0 or greater

Also recommended are tools to communicate with Zeebe.

## How to run

On root folder, type `docker compose up` to launch. After the environment is up, you may view to monitor on your web browser at `localhost:3000`.


TUNI University Software Project - Elisa
