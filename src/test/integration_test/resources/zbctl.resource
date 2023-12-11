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
    [Documentation]    Creates a single BPMN instance by the name of the process and variables type.
    ...
    ...                The bpmn_name is the key in the 'bpmn.yml' file's bpmn dict.
    ...                The variables_type is the key in the 'bpmn.yml' file's variables dict inside the process.
    [Arguments]    ${bpmn_name}    ${variables_type}
    ${bpmn_instance} =  Collections.Get From Dictionary  ${bpmn}  ${bpmn_name}
    ${process_id} =  Collections.Get From Dictionary  ${bpmn_instance}  process-id
    ${all_variables} =  Collections.Get From Dictionary  ${bpmn_instance}  variables
    ${variables} =  Collections.Get From Dictionary  ${all_variables}  ${variables_type}
    ${res} =  Process.Run Process  zbctl  --insecure  create  instance  ${process_id}  --variables  ${variables}
    Log  ${res.stdout}
    Should Be Equal As Integers  ${res.rc}  0  msg="Creating instance failed"


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


Set Process Id
    [Documentation]    Sets the process id to the dynamic version of the 'bpmn.yml' file's bpmn dict.
    [Arguments]    ${bpmn_name}    ${process_id}
    ${bpmn_instance} =  Collections.Get From Dictionary     ${bpmn}    ${bpmn_name}
    ${process_id} =    Convert To Integer    ${process_id}
    Collections.Set To Dictionary    ${bpmn_instance}    process-id    ${process_id}
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
    FOR    ${instance_name}    IN    @{deployed_instances}
        ${instance_instance} =    Collections.Get From Dictionary    ${bpmn}    ${instance_name}
        ${instance_id} =    Collections.Get From Dictionary    ${instance_instance}    process-id
        Append To List    ${ids}    ${instance_id}
    END
    ${clean_list} =    Collections.Remove Duplicates    ${ids}
    [Return]    ${clean_list}


Get Name Of Deployed Instance
    [Documentation]   Return the name of a deployed instance by id.
    [Arguments]    ${instance_id}
    FOR    ${instance_name}    IN    @{DEPLOYED_DATA}[instance]
        ${instance_instance} =    Collections.Get From Dictionary    ${bpmn}    ${instance_name}
        ${instance_id_from_instance} =    Collections.Get From Dictionary    ${instance_instance}    process-id
        IF    ${instance_id_from_instance} == ${instance_id}
            ${found_instance_name} =    BuiltIn.Set Variable    ${instance_name}
        END
    END
    BuiltIn.Variable Should Exist    ${found_instance_name}    msg=No deployed instance with id ${instance_id} found.
    [Return]    ${found_instance_name}


Get Single Instance Id Of Deployed Instance
    [Documentation]   Return a random instance id from the list of deployed instances.
    ${instance_ids} =    Get Deployed Instance Ids
    ${random_instance_id} =    Get From List    ${instance_ids}    0
    [Return]    ${random_instance_id}
