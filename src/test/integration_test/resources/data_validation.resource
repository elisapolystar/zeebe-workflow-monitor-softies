*** Settings ***
Library      Collections


*** Keywords ***
Lists Should Contain Exactly Same Items
    [Documentation]    Two lists should contain exactly same items, in any order.
    [Arguments]    ${list1}    ${list2}
    List Lengths Should Be Equal    ${list1}    ${list2}
    FOR    ${item}    IN    @{list1}
        Collections.List Should Contain Value    ${list2}    ${item}
    END


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
