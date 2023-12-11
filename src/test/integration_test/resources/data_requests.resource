*** Settings ***
Documentation     Abstarcts request made from websocket.
Resource         ../resources/websocket.resource
Resource         ../resources/message_utils.resource
Library          Collections
Library          OperatingSystem
Variables        ../variables/global.yml
Variables        ../variables/bpmn.yml
Variables        ../variables/websocket_messages.yml


*** Variables ***
${WEBSOCKET} =        DEFAULT WEBSOCKET CONNECTION
${DEPLOYED_DATA} =    DATA DEPLOYED IN SUITE SETUP


*** Keywords ***
Request All Processes
    [Documentation]   Request and return all processes from backend.
    ${message} =      Collections.Get From Dictionary    ${query}    all-processes
    ${reply} =        websocket.Request From Websocket    ${WEBSOCKET}    ${message}
    ${processes} =    message_utils.Get Message Data As List    ${reply}
    Log    ${processes}
    [Return]    ${processes}


Request Process By Id
    [Documentation]   Request and return a process from backend by id.
    [Arguments]    ${process_id}
    ${message} =         Collections.Get From Dictionary    ${query}    process-by-id
    ${full_message} =    BuiltIn.Replace Variables    ${message}
    ${reply} =           websocket.Request From Websocket    ${WEBSOCKET}    ${full_message}
    ${process} =         message_utils.Get Message Data As Dictionary    ${reply}
    Log    ${process}
    [Return]    ${process}


Request Raw All Processes
    [Documentation]   Request and return all processes from backend as unparsed message.
    ${message} =      Collections.Get From Dictionary    ${query}    all-processes
    ${reply} =        websocket.Request From Websocket    ${WEBSOCKET}    ${message}
    Log    ${reply}
    [Return]    ${reply}


Request Raw Process By Id
    [Documentation]   Request and return a process from backend by id as unparsed message.
    [Arguments]    ${process_id}
    ${message} =         Collections.Get From Dictionary    ${query}    process-by-id
    ${full_message} =    BuiltIn.Replace Variables    ${message}
    ${reply} =           websocket.Request From Websocket    ${WEBSOCKET}    ${full_message}
    Log    ${reply}
    [Return]    ${reply}


Request Raw All Instances
    [Documentation]   Request and return all instances from backend as unparsed message.
    ${message} =      Collections.Get From Dictionary    ${query}    all-instances
    ${reply} =        websocket.Request From Websocket    ${WEBSOCKET}    ${message}
    Log    ${reply}
    [Return]    ${reply}


Request Raw Instance By Id
    [Documentation]   Request and return an instance from backend by id as unparsed message.
    [Arguments]    ${instance_id}
    ${message} =         Collections.Get From Dictionary    ${query}    instance-by-id
    ${full_message} =    BuiltIn.Replace Variables    ${message}
    ${reply} =           websocket.Request From Websocket    ${WEBSOCKET}    ${full_message}
    Log    ${reply}
    [Return]    ${reply}


Request All Instances
    [Documentation]   Request and return all instances from backend.
    ${message} =      Collections.Get From Dictionary    ${query}    all-instances
    ${reply} =        websocket.Request From Websocket    ${WEBSOCKET}    ${message}
    ${instances} =    message_utils.Get Message Data As List    ${reply}
    Log    ${instances}
    [Return]    ${instances}


Request Instance By Id
    [Documentation]   Request and return an instance from backend by id.
    [Arguments]    ${instance_id}
    ${message} =         Collections.Get From Dictionary    ${query}    instance-by-id
    ${full_message} =    BuiltIn.Replace Variables    ${message}
    ${reply} =           websocket.Request From Websocket    ${WEBSOCKET}    ${full_message}
    ${instance} =        message_utils.Get Message Data As Dictionary    ${reply}
    Log    ${instance}
    [Return]    ${instance}
