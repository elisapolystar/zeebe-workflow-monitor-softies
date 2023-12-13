*** Settings ***
Documentation    Deploy instance data to backend through zbctl and query it from the websocket.
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
Variables        ../../variables/websocket_messages.yml
Test Timeout     1 minutes
Suite Setup      Suite Setup
Test Setup       Test Setup
Test Teardown    Test Teardown


*** Variables ***
${WEBSOCKET} =        DEFAULT WEBSOCKET CONNECTION
${DEPLOYED_DATA} =    DATA DEPLOYED IN SUITE SETUP


*** Test Cases ***
All instances message is of correct type
    ${all_processes_msg} =     data_requests.Request Raw All Instances
    Message Type Should Be     ${all_processes_msg}    all-instances


Return instances with correct ids
    ${instances} =    data_requests.Request All Instances
    ${received_instance_ids} =   message_utils.Get All Items With Key From List Of Dicts    ${instances}    ProcessInstanceKey
    ${correct_instance_ids} =    zbctl.Get Deployed Instance Ids
    ${current_received_instance_ids} =    message_utils.Get Latest Instances    ${received_instance_ids}    ${correct_instance_ids}
    data_validation.Lists Should Contain Exactly Same Items    ${current_received_instance_ids}    ${correct_instance_ids}


Return instances with correct process names
    ${instances} =    data_requests.Request All Instances
    ${received_instance_names} =   message_utils.Get All Items With Key From List Of Dicts    ${instances}    BpmnProcessId
    ${correct_instances} =    zbctl.Get Names Of Deployed Instances
    ${correct_instance_names} =    Collections.Convert To List    ${correct_instances}
    ${current_received_instance_ids} =    message_utils.Get Latest Instances    ${received_instance_names}    ${correct_instance_names}
    data_validation.Lists Should Contain Exactly Same Items    ${current_received_instance_ids}    ${correct_instance_names}


Return instances with correct process ids
    ${instances} =    data_requests.Request All Instances
    ${received_instance_ids} =   message_utils.Get All Items With Key From List Of Dicts    ${instances}    ProcessDefinitionKey
    ${correct_instances} =      zbctl.Get Process Ids Of Deployed Instances
    ${correct_process_ids} =    Collections.Convert To List    ${correct_instances}
    ${current_received_process_ids} =    message_utils.Get Latest Instances    ${received_instance_ids}    ${correct_process_ids}
    data_validation.Lists Should Contain Exactly Same Items    ${current_received_process_ids}    ${correct_process_ids}


Single instance message is of correct type
    ${instance_ids} =    zbctl.Get Deployed Instance Ids
    ${instance_msg} =    data_requests.Request Raw Instance By Id    ${instance_ids[0]}
    Message Type Should Be    ${instance_msg}    instance


Return the correct process id when requested by instance id
    ${deployed_instance_ids} =    zbctl.Get Deployed Instance Ids
    FOR    ${instance_id}    IN    @{deployed_instance_ids}
        ${received_message} =       data_requests.Request Instance By Id    ${instance_id}
        ${received_process} =       Collections.Get From Dictionary    ${received_message}    process
        ${received_process_id} =    Collections.Get From Dictionary    ${received_process}    key
        ${correct_process_id} =     zbctl.Get Process Id With Instance Id    ${instance_id}
        BuiltIn.Should Be Equal As Strings    ${correct_process_id}    ${received_process_id}
    END


Return the correct process name when requested by instance id
    ${deployed_instance_ids} =    zbctl.Get Deployed Instance Ids
    FOR    ${instance_id}    IN    @{deployed_instance_ids}
        ${received_message} =       data_requests.Request Instance By Id    ${instance_id}
        ${received_process} =       Collections.Get From Dictionary    ${received_message}    process
        ${received_process_name} =    Collections.Get From Dictionary    ${received_process}    bpmnProcessId
        ${correct_process_name} =     zbctl.Get Process Name Of Deployed Instance    ${instance_id}
        BuiltIn.Should Be Equal As Strings    ${received_process_name}    ${correct_process_name}
    END


Return the correct bpmn when requested by instance id
    ${deployed_instance_ids} =    zbctl.Get Deployed Instance Ids
    FOR    ${instance_id}    IN    @{deployed_instance_ids}
        ${received_message} =          data_requests.Request Instance By Id    ${instance_id}
        ${received_process} =          Collections.Get From Dictionary    ${received_message}    process
        ${received_instance_bpmn} =    message_utils.Get BPMN From Message    ${received_process}
        ${process_name} =              zbctl.Get Process Name Of Deployed Instance    ${instance_id}
        ${correct_instance_bpmn} =     message_utils.Get Contents Of A BPMN File    ${process_name}
        BuiltIn.Should Be Equal As Strings    ${received_instance_bpmn}    ${correct_instance_bpmn}
    END


Return the correct fields in instance message
    ${deployed_instance_ids} =    zbctl.Get Deployed Instance Ids
    FOR    ${instance_id}    IN    @{deployed_instance_ids}
        ${received_message} =       data_requests.Request Instance By Id    ${instance_id}
        ${correct_keys} =         Collections.Get From Dictionary    ${expected_fields}    instance
        data_validation.Dictionary Should Have The Keys    ${received_message}    ${correct_keys}
    END


Return an error when requesting an instance that does not exist
    ${received_message} =    data_requests.Request Raw Instance By Id    123456789123456798
    data_validation.Message Type Should Be    ${received_message}    instance
    ${correct_error_message} =    Collections.Get From Dictionary    ${response}    instance-not-found
    ${received_error_message} =   message_utils.Get Message Data As Dictionary    ${received_message}
    BuiltIn.Should Be Equal As Strings    ${received_error_message}    ${correct_error_message}


*** Keywords ***
Suite Setup
    zbctl.Deploy All


Test Setup
    ${WEBSOCKET} =    websocket.Init Websocket Connection    ${BACKEND_WS_URL}
    BuiltIn.Set Suite Variable    ${WEBSOCKET}


Test Teardown
    websocket.Close Websocket Connection    ${WEBSOCKET}
