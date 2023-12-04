*** Settings ***
Library      Collections
Library      JSONLibrary


*** Keywords ***
Get Message Data As Dictionary
    [Documentation]    Get the data from a message of string-type JSON as a dictionary.
    [Arguments]    ${message}
    ${json} =          Get Message Data As JSON    ${message}
    ${dictionary} =    Collections.Convert To Dictionary    ${json}
    Log    ${dictionary}
    [Return]    ${dictionary}


Get Message Data As List
    [Documentation]    Get the data from a message of string-type JSON as a list.
    [Arguments]    ${message}
    ${json} =    Get Message Data As JSON    ${message}
    ${list} =    Collections.Convert To List    ${json}
    Log    ${list}
    [Return]    ${list}


Get Message Data As JSON
    [Documentation]    Get the data from a message of string-type JSON as JSON.
    [Arguments]    ${message}
    ${data} =    Collections.Get From Dictionary    ${message}    data
    ${json} =    JSONLibrary.Convert String To Json    ${data}
    Log    ${json}
    [Return]    ${json}


Get All Items With Key From List Of Dicts
    [Documentation]    Iterate through a list of dictionaries and return a list of all first instances of a value in the dicts with the given key.
    [Arguments]    ${list}    ${key}
    ${items} =    Create List
    FOR    ${item}    IN    @{list}
        ${value} =    Collections.Get From Dictionary    ${item}    ${key}
        Append To List    ${items}    ${value}
    END
    [Return]    ${items}

