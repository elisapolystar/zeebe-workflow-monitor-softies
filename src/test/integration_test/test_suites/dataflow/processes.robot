*** Settings ***
Documentation    Deploy process data to backend through zbctl and query it from the websocket.
...              Assumes that the backend and zeebe-broker are running.
Resource         ../../resources/zbctl.resource
Resource         ../../resources/websocket.resource
Resource         ../../resources/data_validation.resource
Resource         ../../resources/message_utils.resource
Library          Collections
Variables        ../../variables/global.yml
Variables        ../../variables/bpmn.yml
Test Timeout     2 minutes
Suite Setup      Suite Setup
Test Setup       Test Setup
Test Teardown    Test Teardown


*** Variables ***
${WEBSOCKET} =        DEFAULT WEBSOCKET CONNECTION
${DEPLOYED_DATA} =    DATA DEPLOYED IN SUITE SETUP


*** Test Cases ***
Return correct amount of processes
    ${processes} =    Request All Processes
    ${deployed_process_count} =    Get Count Of Deployed Instance    process
    BuiltIn.Length Should Be    ${processes}    ${deployed_process_count}


Return processes with correct names
    ${processes} =    Request All Processes
    ${received_process_names} =   message_utils.Get All Items With Key From List Of Dicts    ${processes}    bpmnProcessId
    ${correct_process_names} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    process
    data_validation.Lists Should Contain Exactly Same Items    ${received_process_names}    ${correct_process_names}


Return processes with correct ids
    ${processes} =    Request All Processes
    ${received_process_ids} =   message_utils.Get All Items With Key From List Of Dicts    ${processes}    key
    ${correct_process_ids} =    Get Deployed Process Ids
    data_validation.Lists Should Contain Exactly Same Items    ${received_process_ids}    ${correct_process_ids}


*** Keywords ***
Suite Setup
    ${DEPLOYED_DATA} =    zbctl.Deploy All
    BuiltIn.Set Suite Variable    ${DEPLOYED_DATA}


Test Setup
    ${WEBSOCKET} =    websocket.Init Websocket Connection    ${BACKEND_WS_URL}
    BuiltIn.Set Suite Variable    ${WEBSOCKET}


Test Teardown
    websocket.Close Websocket Connection    ${WEBSOCKET}


Request All Processes
    [Documentation]   Request and return all processes from backend.
    ${reply} =        websocket.Request From Websocket    ${WEBSOCKET}    all_processes
    ${processes} =    message_utils.Get Message Data As List    ${reply}
    Log    ${processes}
    [Return]    ${processes}


Get Count Of Deployed Instance
    [Documentation]   Return the amount of deployed instances of a given type.
    ...
    ...               Type here is the name of the type of instance, e.g. process, instance, etc.
    [Arguments]    ${type}
    ${deployed_instances} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    ${type}
    ${instance_count} =    Get Length    ${deployed_instances}
    [Return]    ${instance_count}


Get Deployed Process Ids
    [Documentation]   Return the ids of all processes deployed during Suite Setup.
    ${ids} =    Create List
    ${deployed_processes} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    process
    FOR    ${process_name}    IN    @{deployed_processes}
        ${process_instance} =    Collections.Get From Dictionary    ${bpmn}    ${process_name}
        ${process_id} =    Collections.Get From Dictionary    ${process_instance}    process-id
        Append To List    ${ids}    ${process_id}
    END
    [Return]    ${ids}
