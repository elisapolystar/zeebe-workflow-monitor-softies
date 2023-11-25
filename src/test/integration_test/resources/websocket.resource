*** Settings ***
Library      WebSocketClient
Library      Collections
Library      JSONLibrary
Variables    ../variables/websocket_messages.yml


*** Keywords ***
Init Websocket Connection
    [Arguments]    ${address}
    ${websocket} =    WebSocketClient.Connect    ${address}
    Log    ${websocket}
    [Return]    ${websocket}


Close Websocket Connection
    [Arguments]    ${websocket}
    WebSocketClient.Close    ${websocket}


Request From Websocket
    [Documentation]    Send message with the given name to the websocket and return the response.
    ...
    ...                Messages are defined in the websocket_messages.yml file.
    ...                The message name is the key of the message in the 'query' dictionary.
    ...                The response message is assumed to be string-type JSON and returned as a dictionary.
    [Arguments]    ${websocket}    ${message_name}
    Send To Websocket    ${websocket}    ${message_name}
    ${message} =    Receive From Websocket    ${websocket}
    [Return]    ${message}


Send To Websocket
    [Documentation]    Send message with the given name to the websocket.
    ...
    ...                Messages are defined in the websocket_messages.yml file.
    ...                The message name is the key of the message in the 'query' dictionary.
    [Arguments]    ${websocket}    ${message_name}
    ${message} =    Collections.Get From Dictionary    ${query}    ${message_name}
    WebSocketClient.Send    ${websocket}    ${message_name}


Receive From Websocket
    [Documentation]    Receive and return a single message from the websocket.
    ...
    ...                The message is assumed to be string-type JSON and returned as a dictionary.
    [Arguments]    ${websocket}
    ${message} =    WebSocketClient.Recv    ${websocket}
    ${json} =    JSONLibrary.Convert String To Json    ${message}
    ${dictionary} =    Collections.Convert To Dictionary    ${json}
    Log    ${dictionary}
    [Return]    ${dictionary}

