# Zeebe Workflow Monitor Integration Test Asset

## Setting up the environment

1. Create virtual environment with `python3 -m venv .venv`

1. Activate the virtual environment with `source .venv/bin/activate`

1. Install the Python dependencies with `python3 -m pip install -r requirements.txt`

## Running the tests

1. The tests assume that the services have been started by running the docker-compose from the root of this project.

1. Run all tests with `robot --outputdir output/ test_suites/`
