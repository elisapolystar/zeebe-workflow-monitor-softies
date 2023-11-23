CREATE USER test PASSWORD 'password';
CREATE DATABASE workflow OWNER test;
GRANT CONNECT ON DATABASE workflow TO test;
\c workflow;

CREATE TABLE IF NOT EXISTS process (
    Key BIGINT PRIMARY KEY,
    BpmnProcessId VARCHAR(50) NOT NULL,
    Version INT NOT NULL,
    Resource TEXT NOT NULL,
    Timestamp BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS process_instance (
    ProcessInstanceKey BIGINT PRIMARY KEY,
    PartitionID BIGINT NOT NULL,
    ProcessDefinitionKey BIGINT NOT NULL,
    BpmnProcessId VARCHAR(50) NOT NULL,
    Version INT NOT NULL,
    Timestamp BIGINT NOT NULL,
    Active BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS variable (
    PartitionID BIGINT NOT NULL,
    Position BIGINT NOT NULL,
    Name VARCHAR(50) NOT NULL,
    Value VARCHAR(50) NOT NULL,
    ProcessInstanceKey REFERENCES process_instance (ProcessInstanceKey),
    ScopeKey BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS job (
    Key BIGINT PRIMARY KEY,
    Timestamp BIGINT NOT NULL,
    ProcessInstanceKey REFERENCES process_instance (ProcessInstanceKey),
    ElementInstanceKey BIGINT NOT NULL,
    JobType VARCHAR(50),
    Worker VARCHAR(50),
    Retries INT
);

CREATE TABLE IF NOT EXISTS incident (
    Key BIGINT PRIMARY KEY,
    BpmnProcessId VARCHAR(50) NOT NULL,
    ProcessInstanceKey REFERENCES process_instance (ProcessInstanceKey),
    ElementInstanceKey BIGINT NOT NULL,
    JobKey BIGINT NOT NULL;
    ErrorType VARCHAR(50) NOT NULL,
    ErrorMessage TEXT NOT NULL,
    Timestamp BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS message (
    Key BIGINT PRIMARY KEY,
    Name VARCHAR(50) NOT NULL,
    CorrelationKey VARCHAR(50) NOT NULL,
    MessageId VARCHAR(50) NOT NULL,
    Timestamp BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS timer (
    Key BIGINT PRIMARY KEY,
    Timestamp BIGINT NOT NULL,
    ProcessDefinitionKey BIGINT NOT NULL,
    ProcessInstanceKey BIGINT NOT NULL,
    ElementInstanceKey BIGINT NOT NULL,
    TargetElementId VARCHAR(50) NOT NULL,
    Duedate BIGINT NOT NULL,
    Repetitions BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS element (
    Key BIGINT PRIMARY KEY,
    ProcessInstanceKey BIGINT NOT NULL,
    ProcessDefinitionKey BIGINT NOT NULL,
    BpmnProcessId VARCHAR(50) NOT NULL,
    ElementId VARCHAR(50) NOT NULL,
    BpmnElementType VARCHAR(50) NOT NULL,
    Intent VARCHAR(50) NOT NULL
);