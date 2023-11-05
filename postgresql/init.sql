CREATE DATABASE IF NOT EXISTS workflow;
USE workflow;

CREATE TABLE IF NOT EXISTS process(
    Key BIGINT PRIMARY KEY,
    BpmnProcessId VARCHAR(50) NOT NULL,
    Version INT NOT NULL,
    Resource TEXT NOT NULL,
    Timestamp BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS process_instance(
    Key BIGINT PRIMARY KEY,
    PartitionID BIGINT NOT NULL,
    ProcessDefinitionKey BIGINT NOT NULL,
    BpmnProcessId VARCHAR(50) NOT NULL,
    Version INT NOT NULL,
    ParentProcessInstanceKey BIGINT NOT NULL,
    ParentElementInstanceKey BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS variable(
    PartitionID BIGINT NOT NULL,
    Position BIGINT NOT NULL,
    Name VARCHAR(50) NOT NULL,
    Value VARCHAR(50) NOT NULL,
    ProcessInstanceKey REFERENCES process_instance (key),
    ScopeKey BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS job(
    key BIGINT PRIMARY KEY,
    Timestamp BIGINT NOT NULL,
    ProcessInstanceKey REFERENCES process_instance (key) NOT NULL,
    ElementInstanceKey BIGINT NOT NULL,
    JobType VARCHAR(50),
    Worker VARCHAR(50),
    Retries INT
);

CREATE TABLE IF NOT EXISTS incident(
    Key BIGINT PRIMARY KEY,
    BpmnProcessId VARCHAR(50) NOT NULL,
    ProcessInstanceKey REFERENCES process_instance (key) NOT NULL,
    ElementInstanceKey BIGINT NOT NULL,
    JobKey BIGINT NOT NULL;
    ErrorType VARCHAR(50) NOT NULL,
    ErrorMessage TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS message(
    Key BIGINT PRIMARY KEY,
    Name VARCHAR(50) NOT NULL,
    CorrelationKey VARCHAR(50) NOT NULL,
    MessageId VARCHAR(50) NOT NULL,
    Timestamp BIGINT NOT NULL
);