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
    ...                The response message is assumed to be string-type JSON and returned as a dictionary.
    [Arguments]    ${websocket}    ${message}
    WebSocketClient.Send    ${websocket}    ${message}
    ${message} =    Receive From Websocket    ${websocket}
    [Return]    ${message}


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

