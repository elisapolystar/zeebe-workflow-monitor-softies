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
Test Timeout     1 minutes
# Suite Setup      Suite Setup
Test Setup       Test Setup
Test Teardown    Test Teardown


*** Variables ***
${WEBSOCKET} =        DEFAULT WEBSOCKET CONNECTION
${DEPLOYED_DATA} =    DATA DEPLOYED IN SUITE SETUP


*** Test Cases ***
# Return correct amount of instances
#     ${instances} =    data_requests.Request All Instances
#     ${deployed_instance_count} =    zbctl.Get Count Of Any Deployed Element    instance
#     BuiltIn.Length Should Be    ${instances}    ${deployed_instance_count}


# Return instances with correct names
#     ${instances} =    data_requests.Request All Instances
#     ${received_instance_names} =   message_utils.Get All Items With Key From List Of Dicts    ${instances}    BpmnProcessId
#     ${correct_instances} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    instance
#     ${correct_instance_names} =    Collections.Convert To List    ${correct_instances}
#     ${correct_instance_count} =    BuiltIn.Get Length    ${correct_instance_names}
#     ${received_instance_names} =    Collections.Get Slice From List    ${received_instance_names}    end=${correct_instance_count}
#     data_validation.Lists Should Contain Exactly Same Items    ${received_instance_names}    ${correct_instance_names}


# Return instances with correct ids
#     ${instances} =    data_requests.Request All Instances
#     ${received_instance_ids} =   message_utils.Get All Items With Key From List Of Dicts    ${instances}    key
#     ${correct_instance_ids} =    zbctl.Get Deployed Instance Ids
#     data_validation.Lists Should Contain Exactly Same Items    ${received_instance_ids}    ${correct_instance_ids}


# Return the correct instance id when requested by id
#     ${deployed_instance_ids} =    zbctl.Get Deployed Instance Ids
#     FOR    ${instance_id}    IN    @{deployed_instance_ids}
#         ${instance} =                  data_requests.Request Instance By Id    ${instance_id}
#         ${received_instance_id} =      Collections.Get From Dictionary    ${instance}    key
#         BuiltIn.Should Be Equal As Strings    ${received_instance_id}    ${instance_id}
#     END


# Return the correct instance name when requested by id
#     ${deployed_instance_ids} =    zbctl.Get Deployed Instance Ids
#     FOR    ${instance_id}    IN    @{deployed_instance_ids}
#         ${instance} =                  data_requests.Request Instance By Id    ${instance_id}
#         ${received_instance_name} =    Collections.Get From Dictionary    ${instance}    bpmnProcessId
#         ${correct_instance_name} =     zbctl.Get Name Of Deployed Instance    ${instance_id}
#         BuiltIn.Should Be Equal As Strings    ${received_instance_name}    ${correct_instance_name}
#     END


# Return the correct instance bpmn when requested by id
#     ${deployed_instance_ids} =    zbctl.Get Deployed Instance Ids
#     FOR    ${instance_id}    IN    @{deployed_instance_ids}
#         ${instance} =                  data_requests.Request Instance By Id    ${instance_id}
#         ${instance_name} =             zbctl.Get Name Of Deployed Instance    ${instance_id}
#         ${received_instance_bpmn} =    message_utils.Get BPMN From Message    ${instance}
#         ${correct_instance_bpmn} =     message_utils.Get Contents Of A BPMN File    ${instance_name}
#         BuiltIn.Should Be Equal As Strings    ${received_instance_bpmn}    ${correct_instance_bpmn}
#     END


# Return nothing when requesting a non-existing instance by id
#     BuiltIn.Run Keyword And Expect Error    timeout    Request Instance By Id    123456789123456789


DEBUG REMOVE THIS
    Request All Instances
    Request Instance By Id    2251799813685257


*** Keywords ***
Suite Setup
    zbctl.Deploy All


Test Setup
    ${WEBSOCKET} =    websocket.Init Websocket Connection    ${BACKEND_WS_URL}
    BuiltIn.Set Suite Variable    ${WEBSOCKET}


Test Teardown
    websocket.Close Websocket Connection    ${WEBSOCKET}
