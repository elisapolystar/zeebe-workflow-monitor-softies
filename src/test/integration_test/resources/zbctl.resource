*** Settings ***
Library      Process
Library      Collections
Library      JSONLibrary
Variables    ../variables/bpmn.yml


*** Variables ***
${DEPLOYED_DATA} =    DATA DEPLOYED BY KEYWORD 'Deploy All'


*** Keywords ***
Deploy All
    [Documentation]    Deploys all BPMN processes, subprocesses, instances and workers.
    ...
    ...                'All' here means all the items that are defined in the 'bpmn.yml' file's to-deploy dict.
    ...                Returns that dictionary containing all the deployed items.
    ${processes} =    Collections.Get From Dictionary    ${to-deploy}    process
    ${sub_processes} =    Collections.Get From Dictionary    ${to-deploy}    sub-process
    ${instances} =    Collections.Get From Dictionary    ${to-deploy}    instance
    ${workers} =    Collections.Get From Dictionary    ${to-deploy}    worker
    FOR    ${process}    IN    @{processes}
        Deploy BPMN Process    ${process}
    END
    FOR    ${sub_process}    IN    @{sub_processes}
        Deploy BPMN SubProcess    ${sub_process}
    END
    FOR    ${process}    ${variable_type_list}    IN    &{instances}
        FOR    ${variable_type}    IN    @{variable_type_list}
            Create Instance    ${process}    ${variable_type}
            Sleep    1s
        END
    END
    FOR    ${worker}    IN    @{workers}
        Create Worker    ${worker}    create
    END
    BuiltIn.Set Suite Variable    ${DEPLOYED_DATA}    ${to-deploy}


Deploy BPMN Process
    [Documentation]    Deploys a single BPMN process by its name and saves its id.
    ...
    ...                The bpmn_name is the key in the 'bpmn.yml' file's bpmn dict.
    ...                The process_id is saved in the dynamic version of the 'bpmn.yml' file's bpmn dict.
    [Arguments]    ${bpmn_name}
    ${bpmn_instance} =  Collections.Get From Dictionary  ${bpmn}  ${bpmn_name}
    ${bpmn_path} =  Collections.Get From Dictionary  ${bpmn_instance}  filepath
    ${res} =  Process.Run Process  zbctl  --insecure  deploy  ${bpmn_path}
    Log  ${res.stdout}
    Should Be Equal As Integers  ${res.rc}  0  msg="Deploying BPMN process failed"
    ${process_id} =  Get Process Id From Deployment  ${res.stdout}
    Set Process Id  ${bpmn_name}  ${process_id}


Deploy BPMN SubProcess
    [Documentation]    Deploys a single BPMN subProcess by its name.
    ...
    ...                The bpmn_name is the key in the 'bpmn.yml' file's bpmn dict.
    [Arguments]    ${bpmn_name}
    ${bpmn_instance} =  Collections.Get From Dictionary  ${bpmn}  ${bpmn_name}
    ${bpmn_id} =  Collections.Get From Dictionary  ${bpmn_instance}  process-id
    ${res} =  Process.Run Process  zbctl  --insecure  create  instance  ${bpmn_id}
    Log  ${res.stdout}
    Should Be Equal As Integers  ${res.rc}  0  msg="Deploying BPMN SubProcess failed"


Create Instance
    [Documentation]    Creates a single BPMN instance by the name of the process and instance type.
    ...
    ...                The bpmn_name is the key in the 'bpmn.yml' file's bpmn dict.
    ...                The instance_type is the key in the 'bpmn.yml' file's instances dict inside the process.
    [Arguments]    ${bpmn_name}    ${instance_type}
    ${bpmn_instance} =  Collections.Get From Dictionary  ${bpmn}  ${bpmn_name}
    ${process_id} =  Collections.Get From Dictionary  ${bpmn_instance}  process-id
    ${all_instances} =  Collections.Get From Dictionary  ${bpmn_instance}  instances
    ${instance} =  Collections.Get From Dictionary  ${all_instances}  ${instance_type}
    ${instance_variables} =  Collections.Get From Dictionary  ${instance}  variables
    ${res} =  Process.Run Process  zbctl  --insecure  create  instance  ${process_id}  --variables  ${instance_variables}
    Log  ${res.stdout}
    Should Be Equal As Integers  ${res.rc}  0  msg="Creating instance failed"
    ${instance_id} =  Get Instance Id From Deployment  ${res.stdout}
    Set Instance Id    ${instance}    ${instance_id}


Create Worker
    [Documentation]    Creates a single BPMN worker by the name of the worker and handler.
    ...
    ...                Creating a worker never returns so the return value cannot be checked.
    ...                The worker_name is a string of a name defined in the parent BPMN process.
    ...                These names are visible in the 'bpmn.yml' file's bpmn dict.
    [Arguments]    ${worker_name}    ${handler_name}
    ${res} =  Process.Start Process  zbctl  --insecure  create  worker  ${worker_name}  --handler  ${handler_name}
    Log  ${res}


Publish Message
    [Documentation]    Publishes a single BPMN message.
    ...
    ...                correlation_key is the id of the process instance that is linked to this message.
    [Arguments]    ${message}    ${correlation_key}
    ${res} =  Process.Run Process  zbctl  --insecure  publish  message  ${message}  \-\-correlationKey\=${correlation_key}
    Log  ${res.stdout}
    Should Be Equal As Integers  ${res.rc}  0  msg="Publishing message failed"


Get Process Id From Deployment
    [Documentation]    Gets the process id from the stdout of the deployment.
    [Arguments]    ${stdout}
    ${parent_json} =    JSONLibrary.Convert String To Json    ${stdout}
    ${parent_dict} =    Collections.Convert To Dictionary    ${parent_json}
    ${sublist} =        Collections.Get From Dictionary    ${parent_dict}    processes
    ${subdict} =        Collections.Get From List    ${sublist}    0
    ${key} =            Collections.Get From Dictionary    ${subdict}    processDefinitionKey
    [Return]    ${key}


Get Instance Id From Deployment
    [Documentation]    Gets the instance id from the stdout of the deployment.
    [Arguments]    ${stdout}
    ${json} =    JSONLibrary.Convert String To Json    ${stdout}
    ${dict} =    Collections.Convert To Dictionary    ${json}
    ${key} =     Collections.Get From Dictionary    ${dict}    processInstanceKey
    [Return]    ${key}


Set Process Id
    [Documentation]    Sets the process id to the dynamic version of the 'bpmn.yml' file's bpmn dict.
    [Arguments]    ${bpmn_name}    ${process_id}
    ${bpmn_instance} =  Collections.Get From Dictionary     ${bpmn}    ${bpmn_name}
    ${process_id} =    Convert To Integer    ${process_id}
    Collections.Set To Dictionary    ${bpmn_instance}    process-id    ${process_id}
    Log    ${bpmn}


Set Instance Id
    [Documentation]    Sets the instance id to the dynamic version of the 'bpmn.yml' file's bpmn dict.
    [Arguments]    ${instance_entry}    ${instance_id}
    ${instance_id} =    Convert To Integer    ${instance_id}
    Collections.Set To Dictionary    ${instance_entry}    instance-id    ${instance_id}
    Log    ${bpmn}


Get Name Of Deployed Process
    [Documentation]   Return the name of a deployed process by id.
    [Arguments]    ${process_id}
    FOR    ${process_name}    IN    @{DEPLOYED_DATA}[process]
        ${process_instance} =    Collections.Get From Dictionary    ${bpmn}    ${process_name}
        ${process_id_from_instance} =    Collections.Get From Dictionary    ${process_instance}    process-id
        IF    ${process_id_from_instance} == ${process_id}
            ${found_process_name} =    BuiltIn.Set Variable    ${process_name}
        END
    END
    BuiltIn.Variable Should Exist    ${found_process_name}    msg=No deployed process with id ${process_id} found.
    [Return]    ${found_process_name}


Get Process Id With Instance Id
    [Documentation]    Return the process id of a deployed instance by id.
    [Arguments]    ${instance_id}
    FOR    ${process_name}    IN    @{DEPLOYED_DATA}[process]
        ${process_instance} =    Collections.Get From Dictionary    ${bpmn}    ${process_name}
        Builtin.Continue For Loop If    ${process_instance}[instances] == ${None}
        ${process_id} =    Collections.Get From Dictionary    ${process_instance}    process-id
        ${all_instances_of_process} =    Collections.Get From Dictionary    ${process_instance}    instances
        FOR    ${instance_type}    IN    @{all_instances_of_process}
            ${instance_data} =    Collections.Get From Dictionary    ${all_instances_of_process}    ${instance_type}
            ${instance_id_from_instance} =    Collections.Get From Dictionary    ${instance_data}    instance-id
            IF    ${instance_id_from_instance} == ${instance_id}
                ${found_process_id} =    BuiltIn.Set Variable    ${process_id}
            END
        END
    END
    BuiltIn.Variable Should Exist    ${found_process_id}    msg=No deployed instance with id ${instance_id} found.
    [Return]    ${found_process_id}

Get Count Of Any Deployed Element
    [Documentation]   Return the amount of deployed instances of a given type.
    ...
    ...               Type here is the name of the type of instance, e.g. process, instance, etc.
    [Arguments]    ${type}
    ${deployed_instances} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    ${type}
    ${no_duplicates} =     Collections.Remove Duplicates    ${deployed_instances}
    ${instance_count} =    Get Length    ${no_duplicates}
    [Return]    ${instance_count}


Get Count Of Deployed Instances
    [Documentation]   Return the amount of deployed instances of a given process.
    [Arguments]    ${process_name}
    ${deployed_instances} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    instance
    ${no_duplicates} =     Collections.Remove Duplicates    ${deployed_instances}
    ${instance_count} =    Get Length    ${no_duplicates}
    [Return]    ${instance_count}


Get Deployed Process Ids
    [Documentation]   Return the ids of all processes deployed during Suite Setup.
    ${ids} =    Create List
    ${deployed_processes} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    process
    FOR    ${process_name}    IN    @{deployed_processes}
        ${process_instance} =    Collections.Get From Dictionary    ${bpmn}    ${process_name}
        ${process_id} =    Collections.Get From Dictionary    ${process_instance}    process-id
        Append To List    ${ids}    ${process_id}
    END
    ${clean_list} =    Collections.Remove Duplicates    ${ids}
    [Return]    ${clean_list}


Get Single Process Id Of Deployed Process
    [Documentation]   Return a random process id from the list of deployed processes.
    ${process_ids} =    Get Deployed Process Ids
    ${random_process_id} =    Get From List    ${process_ids}    0
    [Return]    ${random_process_id}


Get Deployed Instance Ids
    [Documentation]   Return the ids of all instances deployed during Suite Setup.
    ${ids} =    Create List
    ${deployed_instances} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    instance
    FOR    ${process_name}  ${instance_list}    IN    &{deployed_instances}
        ${process_data} =    Collections.Get From Dictionary    ${bpmn}    ${process_name}
        Builtin.Continue For Loop If    ${process_data}[instances] == ${None}
        ${all_instances_of_process} =    Collections.Get From Dictionary    ${process_data}    instances
        FOR    ${instance_type}    IN    @{instance_list}
            ${instance_data} =    Collections.Get From Dictionary    ${all_instances_of_process}    ${instance_type}
            ${instance_id} =    Collections.Get From Dictionary    ${instance_data}    instance-id
            Append To List    ${ids}    ${instance_id}
        END
    END
    ${clean_list} =    Collections.Remove Duplicates    ${ids}
    [Return]    ${clean_list}


Get Names Of Deployed Instances
    [Documentation]   Return the names of all instances deployed during Suite Setup.
    ...
    ...               The names are repeated for the amount of instances deployed for that process.
    ${names} =    Create List
    ${deployed_instances} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    instance
    FOR    ${process_name}  ${instance_list}    IN    &{deployed_instances}
        FOR    ${deployed_instance}    IN    @{instance_list}
            Collections.Append To List    ${names}    ${process_name}
        END
    END
    [Return]    ${names}


Get Process Ids Of Deployed Instances
    [Documentation]   Return the ids of all processes of instances deployed during Suite Setup.
    ${ids} =    Create List
    ${deployed_instances} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    instance
    FOR    ${process_name}  ${instance_list}    IN    &{deployed_instances}
        ${process_data} =    Collections.Get From Dictionary    ${bpmn}    ${process_name}
        Builtin.Continue For Loop If    ${process_data}[instances] == ${None}
        ${process_id} =      Collections.Get From Dictionary    ${process_data}    process-id
        ${all_instances_of_process} =    Collections.Get From Dictionary    ${process_data}    instances
        FOR    ${instance_type}    IN    @{instance_list}
            Append To List    ${ids}    ${process_id}
        END
    END
    [Return]    ${ids}


Get Process Name Of Deployed Instance
    [Documentation]   Return the name of a deployed instance by id.
    [Arguments]    ${instance_id}
    ${process_id} =    Get Process Id With Instance Id    ${instance_id}
    ${found_process_name} =    Get Name Of Deployed Process    ${process_id}
    [Return]    ${found_process_name}


Get Instance With Id
    [Documentation]   Return the instance dict of a deployed instance by id.
    [Arguments]    ${instance_id}
    ${deployed_instances} =    Collections.Get From Dictionary    ${DEPLOYED_DATA}    instance
    FOR    ${process_name}  ${instance_list}    IN    &{deployed_instances}
        ${process_data} =    Collections.Get From Dictionary    ${bpmn}    ${process_name}
        Builtin.Continue For Loop If    ${process_data}[instances] == ${None}
        ${all_instances_of_process} =    Collections.Get From Dictionary    ${process_data}    instances
        FOR    ${instance_type}    IN    @{instance_list}
            ${instance_data} =    Collections.Get From Dictionary    ${all_instances_of_process}    ${instance_type}
            ${instance_id_from_instance} =    Collections.Get From Dictionary    ${instance_data}    instance-id
            IF    ${instance_id_from_instance} == ${instance_id}
                ${found_instance_data} =    BuiltIn.Set Variable    ${instance_data}
            END
        END
    END
    BuiltIn.Variable Should Exist    ${found_instance_data}    msg=No deployed instance with id ${instance_id} found.
    [Return]    ${found_instance_data}


Get Amount Of Elements In Instance
    [Documentation]   Return the amount of elements in an instance.
    [Arguments]    ${instance_id}
    ${instance_data} =    Get Instance With Id    ${instance_id}
    ${amount_of_elements} =    Collections.Get From Dictionary    ${instance_data}    elements
    [Return]    ${amount_of_elements}


Get Amount Of Variables In Instance
    [Documentation]   Return the amount of variables in an instance.
    [Arguments]    ${instance_id}
    ${instance_data} =    Get Instance With Id    ${instance_id}
    ${amount_of_variables} =    Collections.Get From Dictionary    ${instance_data}    message-variables
    [Return]    ${amount_of_variables}


Get Amount Of Timers In Instance
    [Documentation]   Return the amount of timers in an instance.
    [Arguments]    ${instance_id}
    ${instance_data} =    Get Instance With Id    ${instance_id}
    ${amount_of_timers} =    Collections.Get From Dictionary    ${instance_data}    timers
    [Return]    ${amount_of_timers}


Get Amount Of Incidents In Instance
    [Documentation]   Return the amount of incidents in an instance.
    [Arguments]    ${instance_id}
    ${instance_data} =    Get Instance With Id    ${instance_id}
    ${amount_of_incidents} =    Collections.Get From Dictionary    ${instance_data}    incidents
    [Return]    ${amount_of_incidents}


Get Single Instance Id Of Deployed Instance
    [Documentation]   Return a random instance id from the list of deployed instances.
    ${instance_ids} =    Get Deployed Instance Ids
    ${random_instance_id} =    Get From List    ${instance_ids}    0
    [Return]    ${random_instance_id}
