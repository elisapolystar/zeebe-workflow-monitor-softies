*** Settings ***
Library      Collections
Library      JSONLibrary
Library      String
Library      OperatingSystem
Library      ./utils/conversions.py
Variables    ../variables/bpmn.yml


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


Get Contents Of A BPMN File
    [Documentation]   Return the contents of a BPMN file.
    [Arguments]    ${process_name}
    ${process_instance} =    Collections.Get From Dictionary    ${bpmn}    ${process_name}
    ${bpmn_filepath} =    Collections.Get From Dictionary    ${process_instance}    filepath
    ${bpmn_contents} =    OperatingSystem.Get File    ${bpmn_filepath}
    [Return]    ${bpmn_contents}


Decode Base64 To String
    [Documentation]    Decode a base64 string. E.g. a Process BPMN.
    [Arguments]    ${base64}
    ${bytes} =     conversions.Base64 To Bytes    ${base64}
    ${string} =    conversions.Bytes To String    ${bytes}
    [Return]    ${string}


Get BPMN From Message
    [Documentation]    Get the BPMN from a message of string-type JSON.
    [Arguments]    ${message}
    ${bpmn_encoded} =    Collections.Get From Dictionary    ${message}    resource
    ${bpmn} =            Decode Base64 To String    ${bpmn_encoded}
    [Return]    ${bpmn}


Get Latest Instances
    [Documentation]    Get x latest instaces from ${received_instances}, where x is the length of ${deployed_instances}.
    [Arguments]    ${received_instances}    ${deployed_instances}
    ${length} =    Get Length    ${deployed_instances}
    ${latest_instances} =    Collections.Get Slice From List    ${received_instances}    end=${length}
    [Return]    ${latest_instances}
