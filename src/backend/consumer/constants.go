package main

// for topic message parsing
const intent_str = "\"intent\":"

// Broker
const broker = "kafka:9093"

// Topics
const (
	zeebe                                  = "zeebe"
	zeebe_deployment                       = "zeebe-deployment"
	zeebe_deploy_distribution              = "zeebe-deploy-distribution"
	zeebe_error                            = "zeebe-error"
	zeebe_incident                         = "zeebe-incident"
	zeebe_job_batch                        = "zeebe-job-batch"
	zeebe_job                              = "zeebe-job"
	zeebe_message                          = "zeebe-message"
	zeebe_message_subscription             = "zeebe-message-subscription"
	zeebe_message_subscription_start_event = "zeebe-message-subscription-start-event"
	zeebe_process                          = "zeebe-process"
	zeebe_process_event                    = "zeebe-process-event"
	zeebe_process_instance                 = "zeebe-process-instance"
	zeebe_process_instance_result          = "zeebe-process-instance-result"
	zeebe_process_message_subscription     = "zeebe-process-message-subscription"
	zeebe_timer                            = "zeebe-timer"
	zeebe_variable                         = "zeebe-variable"
)

// Types of messages frontend and backend use for communication
const (
	empty           = ""
	process_string  = "process"
	all_processes   = "all-processes"
	instance_string = "instance"
	all_instances   = "all-instances"
	incident_string = "incident"
	all_incidents   = "all-incidents"
)
