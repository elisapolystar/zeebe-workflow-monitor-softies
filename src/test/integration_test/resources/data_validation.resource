*** Settings ***
Library      Collections


*** Keywords ***
Lists Should Contain Exactly Same Items
    [Documentation]    Two lists should contain exactly same items, in any order.
    [Arguments]    ${list1}    ${list2}
    List Lengths Should Be Equal    ${list1}    ${list2}
    Collections.Sort List    ${list1}
    Collections.Sort List    ${list2}
    Collections.Lists Should Be Equal    ${list1}    ${list2}


List Lengths Should Be Equal
    [Arguments]    ${list1}    ${list2}
    ${len1} =    Builtin.Get Length    ${list1}
    ${len2} =    Builtin.Get Length    ${list2}
    Builtin.Should Be Equal As Integers    ${len1}    ${len2}


Message Type Should Be
    [Documentation]    Message should be of given type.
    [Arguments]    ${message}    ${type}
    ${actual_type} =    Collections.Get From Dictionary    ${message}    type
    Builtin.Should Be Equal    ${actual_type}    ${type}


Get Length Or 0
    [Documentation]    Get length of list or 0 if list is None.
    [Arguments]    ${list}
    ${length} =    Run Keyword If    ${list}    Builtin.Get Length    ${list}    ELSE    Set Variable    0
    [Return]    ${length}


Dictionary Should Have The Keys
    [Documentation]    Dictionary should have the given keys.
    [Arguments]    ${dictionary}    ${keys}
    ${actual_keys} =    Collections.Get Dictionary Keys    ${dictionary}
    Lists Should Contain Exactly Same Items    ${actual_keys}    ${keys}
