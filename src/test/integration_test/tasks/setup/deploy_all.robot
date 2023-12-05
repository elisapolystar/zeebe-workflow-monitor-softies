*** Settings ***
Documentation    Deploy all data defined in 'variables/bpmn.yml' 'to-deploy' dict: ${to-deploy}
Resource         ../../resources/zbctl.resource
Variables        ../../variables/bpmn.yml
Test Timeout     1 minutes


*** Variables ***
${DEPLOYED_DATA} =    DATA DEPLOYED BY THE TASK


*** Tasks ***
Deploy all
    zbctl.Deploy All
    Log To Console    \nDeployed data:\n${DEPLOYED_DATA}\n
    Log To Console    \nDetails:\n${bpmn}\n
