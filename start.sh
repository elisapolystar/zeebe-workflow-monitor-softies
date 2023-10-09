#!/bin/bash

docker compose up -d postgres
docker compose up -d --force-recreate kafka

k_started=""
while [[ $k_started == "" ]]; do
    k_started=$(docker compose logs kafka 2>&1 | grep 'started (kafka.server.KafkaServer)')
done

while read topic; do docker exec -t zeebe-workflow-monitor-softies-kafka-1 /opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server kafka:9092 \
    --replication-factor 1 --partitions 1 --if-not-exists --topic "$topic"; done < topics

docker compose up -d zeebe
docker compose up -d postgres
docker compose up -d backend
docker compose up -d frontend
