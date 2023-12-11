# Zeebe Workflow Monitor Integration Test Asset

## Setting up the environment

1. Create virtual environment with `python3 -m venv .venv`

1. Activate the virtual environment with `source .venv/bin/activate`

1. Install the Python dependencies with `python3 -m pip install -r requirements.txt`

## Deploying data to zeebe

This task is useful to deploy test data into the program for manual testing. The automatic test cases include this step so if you are running those you can skip to that part.

1. The tasks assume that the services have been started by running the docker-compose from the root of this project.

1. Deploy all data from [the variable file](./variables/bpmn.yml)'s to-deploy dict with `robot --outputdir output/ tasks/setup/deploy_all.robot`

To empty the data from the program:

1. Stop the running containers.

1. Run `docker system prune` to delete all data from docker.

1. Run the docker-compose up again.

## Running the tests

1. The tests assume that the services have been started by running the docker-compose from the root of this project.

1. Run all tests with `robot --outputdir output/ test_suites/`
