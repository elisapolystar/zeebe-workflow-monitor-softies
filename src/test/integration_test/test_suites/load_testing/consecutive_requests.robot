*** Settings ***
Documentation    Load test the backend with a large number of consequtive requests.
Resource         ../../resources/zbctl.resource
Resource         ../../resources/websocket.resource
Resource         ../../resources/data_validation.resource
Resource         ../../resources/message_utils.resource
Resource         ../../resources/data_requests.resource
Library          Collections
Library          OperatingSystem
Variables        ../../variables/global.yml
Variables        ../../variables/bpmn.yml
Test Timeout     5 minutes
Suite Setup      Suite Setup
Test Setup       Test Setup
Test Teardown    Test Teardown


*** Variables ***
# Deploying data takes significantly longer than requesting it.
# ~10 deployment iterations / minute
# ~2500 request iterations / minute
${DEPLOYMENT_ITERATIONS} =       ${20}
${REQUEST_ITERATIONS} =       ${500}
${WEBSOCKET} =        DEFAULT WEBSOCKET CONNECTION
${DEPLOYED_DATA} =    DATA DEPLOYED IN SUITE SETUP


*** Test Cases ***
Request all processes
    FOR    ${_}    IN RANGE    ${REQUEST_ITERATIONS}
        ${processes} =    data_requests.Request All Processes
        ${received_process_ids} =   message_utils.Get All Items With Key From List Of Dicts    ${processes}    key
        ${correct_process_ids} =    zbctl.Get Deployed Process Ids
        data_validation.Lists Should Contain Exactly Same Items    ${received_process_ids}    ${correct_process_ids}
    END


Request a single process
    ${process_id} =    zbctl.Get Single Process Id Of Deployed Process
    FOR    ${_}    IN RANGE    ${REQUEST_ITERATIONS}
        ${process} =                  data_requests.Request Process By Id    ${process_id}
        ${received_process_id} =      Collections.Get From Dictionary    ${process}    key
        BuiltIn.Should Be Equal As Strings    ${received_process_id}    ${process_id}
    END


Deploy All Data
    FOR    ${_}    IN RANGE    ${DEPLOYMENT_ITERATIONS}
        zbctl.Deploy All
    END


*** Keywords ***
Suite Setup
    zbctl.Deploy All


Test Setup
    ${WEBSOCKET} =    websocket.Init Websocket Connection    ${BACKEND_WS_URL}
    BuiltIn.Set Suite Variable    ${WEBSOCKET}


Test Teardown
    websocket.Close Websocket Connection    ${WEBSOCKET}
