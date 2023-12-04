*** Settings ***
Documentation    Deploy process data to backend through zbctl and query it from the websocket.
...              Assumes that the backend and zeebe-broker are running.
Resource         ../../resources/zbctl.resource
Resource         ../../resources/websocket.resource
Resource         ../../resources/data_validation.resource
Resource         ../../resources/message_utils.resource
Resource         ../../resources/data_requests.resource
Library          Collections
Library          OperatingSystem
Variables        ../../variables/global.yml
Variables        ../../variables/bpmn.yml
Test Timeout     1 minutes
Suite Setup      Suite Setup
Test Setup       Test Setup
Test Teardown    Test Teardown


*** Variables ***
${WEBSOCKET} =        DEFAULT WEBSOCKET CONNECTION
${DEPLOYED_DATA} =    DATA DEPLOYED IN SUITE SETUP


*** Test Cases ***
All processes message is of correct type
    ${all_processes_msg} =       data_requests.Request Raw All Processes
    Message Type Should Be    ${all_processes_msg}    all-processes


Return correct amount of processes
    ${processes} =    data_requests.Request All Processes
    ${deployed_process_count} =    zbctl.Get Count Of Any Deployed Element    process
    BuiltIn.Length Should Be    ${processes}    ${deployed_process_count}


Return processes with correct names
    ${processes} =    data_requests.Request All Processes
    ${received_process_names} =   message_utils.Get All Items With Key From List Of Dicts    ${processes}    bpmnProcessId
    ${correct_process_names} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    process
    data_validation.Lists Should Contain Exactly Same Items    ${received_process_names}    ${correct_process_names}


Return processes with correct ids
    ${processes} =    data_requests.Request All Processes
    ${received_process_ids} =   message_utils.Get All Items With Key From List Of Dicts    ${processes}    key
    ${correct_process_ids} =    zbctl.Get Deployed Process Ids
    data_validation.Lists Should Contain Exactly Same Items    ${received_process_ids}    ${correct_process_ids}


Single process message is of correct type
    ${deployed_process_ids} =    zbctl.Get Deployed Process Ids
    ${single_process_msg} =      data_requests.Request Raw Process By Id    ${deployed_process_ids[0]}
    Message Type Should Be    ${single_process_msg}    process


Return the correct process id when requested by id
    ${deployed_process_ids} =    zbctl.Get Deployed Process Ids
    FOR    ${process_id}    IN    @{deployed_process_ids}
        ${process} =                  data_requests.Request Process By Id    ${process_id}
        ${received_process_id} =      Collections.Get From Dictionary    ${process}    key
        BuiltIn.Should Be Equal As Strings    ${received_process_id}    ${process_id}
    END


Return the correct process name when requested by id
    ${deployed_process_ids} =    zbctl.Get Deployed Process Ids
    FOR    ${process_id}    IN    @{deployed_process_ids}
        ${process} =                  data_requests.Request Process By Id    ${process_id}
        ${received_process_name} =    Collections.Get From Dictionary    ${process}    bpmnProcessId
        ${correct_process_name} =     zbctl.Get Name Of Deployed Process    ${process_id}
        BuiltIn.Should Be Equal As Strings    ${received_process_name}    ${correct_process_name}
    END


Return the correct process bpmn when requested by id
    ${deployed_process_ids} =    zbctl.Get Deployed Process Ids
    FOR    ${process_id}    IN    @{deployed_process_ids}
        ${process} =                  data_requests.Request Process By Id    ${process_id}
        ${process_name} =             zbctl.Get Name Of Deployed Process    ${process_id}
        ${received_process_bpmn} =    message_utils.Get BPMN From Message    ${process}
        ${correct_process_bpmn} =     message_utils.Get Contents Of A BPMN File    ${process_name}
        BuiltIn.Should Be Equal As Strings    ${received_process_bpmn}    ${correct_process_bpmn}
    END


Return an error when requesting a non-existing process by id
    ${process} =    data_requests.Request Raw Process By Id    123456789123456789
    Message Type Should Be    ${process}    process
    ${correct_error_message} =    Collections.Get From Dictionary    ${response}    process-not-found
    ${received_error_message} =   message_utils.Get Message Data As Dictionary    ${process}
    BuiltIn.Should Be Equal As Strings    ${received_error_message}    ${correct_error_message}


*** Keywords ***
Suite Setup
    zbctl.Deploy All


Test Setup
    ${WEBSOCKET} =    websocket.Init Websocket Connection    ${BACKEND_WS_URL}
    BuiltIn.Set Suite Variable    ${WEBSOCKET}


Test Teardown
    websocket.Close Websocket Connection    ${WEBSOCKET}
