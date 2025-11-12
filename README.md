<!--
SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
Copyright 2019 free5GC.org

SPDX-License-Identifier: Apache-2.0
-->
[![Go Report Card](https://goreportcard.com/badge/github.com/omec-project/pcf)](https://goreportcard.com/report/github.com/omec-project/pcf)

# pcf

The Policy Control Function (PCF) supports unified policy framework to govern
network behaviour, provides policy rules to Control Plane function(s) to enforce
them and Accesses subscription information relevant for policy decisions in a
Unified Data Repository (UDR)

The reference 3GPP specification for PCF are as follows
- PCC framework Specification 23.503,
- Session Management Policy Control - Specification 29.512
- Policy and Charging Control signalling flows and QoS parameter mapping (Spec 29.513)

## Repository Structure

Below is a high-level view of the repository and its main components:
```
.
├── ampolicy                    # Implements APIs and routers related to Access and Mobility Policies.
│   ├── api_default.go
│   └── routers.go
├── bdtpolicy                   # Contains controllers for Background Data Transfer (BDT) policies, including creation and update of individual BDT policies.
│   ├── api_bdt_policies_collection_routers.go
│   ├── api_individual_bdt_policy_document_routers.go
│   └── routers.go
├── callback                    # Defines notification handlers (Subscribe/Notify) between the PCF and other Network Functions (NFs).
│   ├── api_nf_subscribe_notify.go
│   └── router.go
├── consumer                    # Includes logic for NF discovery and communication (NF Discovery, NF Management).
│   ├── communication.go
│   ├── nf_discovery.go
│   ├── nf_management.go
│   └── search_nf_service.go
├── context                     # Manages the internal PCF context, including UE information, active sessions, and policies.
│   ├── context.go
│   └── ue.go
├── Dockerfile
├── Dockerfile.fast
├── docs                        # Contains documentation assets and diagrams, such as PCF flow diagrams.
│   └── images
│       ├── Flow-Diagram.drawio
│       ├── Flow-Diagram.drawio.license
│       ├── README-PCF.png
│       └── README-PCF.png.license
├── factory                     # Defines configuration files (pcfcfg.yaml) and initialization logic for PCF setup.
│   ├── config.go 
│   ├── factory.go
│   ├── pcfcfg_with_custom_webui_url.yaml
│   ├── pcfcfg.yaml
│   └── pcf_config_test.go
├── go.mod
├── go.mod.license
├── go.sum
├── go.sum.license
├── httpcallback                # Implements HTTP notification handlers (e.g., AMF status change, SM Policy Notify).
│   ├── amf_status_change_notify.go
│   ├── router.go
│   └── sm_policy_notify.go
├── internal                    # Handles internal event and notification logic, including policy update or termination messages.
│   └── notifyevent
│       ├── dispatcher.go
│       ├── listener.go
│       ├── send_smpolicy_termination.go
│       └── send_smpolicy_update.go
├── LICENSES
│   └── Apache-2.0.txt
├── logger                      # Centralized logging configuration for the PCF.
│   └── logger.go
├── Makefile                    # Provide build, testing, and deployment automation for the PCF.
├── metrics                     # Logic for telemetry collection and metric exporting.
│   └── telemetry.go
├── nfregistration              # Implements PCF registration procedures with the Network Repository Function (NRF).
│   ├── nf_registration.go
│   └── nf_registration_test.go
├── NOTICE.txt
├── oam                         # APIs related to Operations, Administration, and Maintenance (OAM) — for example, retrieving policy data.
│   ├── api_get_am_policy.go  
│   └── routers.go
├── pcf.go                      # Main entry point for the PCF service.
├── pcftests                    # Unit and integration tests, including example configuration files (pcfcfg.yaml).
│   ├── pcfcfg.yaml
│   ├── pcf_test.go
│   └── test_input_data.go
├── policyauthorization         # Controllers and APIs for Application Policy Authorization (App Sessions and Event Subscriptions).
│   ├── api_application_sessions_collection.go
│   ├── api_events_subscription_document.go
│   ├── api_individual_application_session_context_document.go
│   └── routers.go
├── polling                     # Implements PCC policy and QoS configuration handling based on IMSI polling.
│   ├── imsi_qos_config.go
│   ├── imsi_qos_config_test.go
│   ├── pcc_policy_config.go
│   ├── pcc_policy_config_test.go
│   ├── polling.go
│   └── polling_test.go
├── producer                    # Core backend logic for generating and applying policy decisions (SM, BDT, OAM, etc.).
│   ├── ampolicy.go
│   ├── bdtpolicy.go
│   ├── callback.go
│   ├── oam.go
│   ├── policyauthorization.go
│   ├── smpolicy.go
│   └── smpolicy_test.go
├── README.md
├── service                     # Defines the service initialization logic (main PCF server startup).
│   └── init.go
├── smpolicy                    # APIs and routers for Session Management Policy Control.
│   ├── api_default.go
│   └── routers.go
├── Taskfile.yml                # Provide build, testing, and deployment automation for the PCF.
├── uepolicy                    # APIs and routers for User Equipment Policies, defining UE-specific rules.
│   ├── api_default.go
│   └── routers.go
├── util                        # Utility functions for conversions, PCC rules, and general PCF helpers.
│   ├── convert.go
│   ├── init_context.go
│   ├── pcc_rule.go
│   └── pcf_util.go
├── VERSION
└── VERSION.license

25 directories, 78 files

```


## Configuration and Deployment

**Docker**

To build the container image:
```
task mod-start
task build
task docker-build-fast
```

**Kubernetes**

The standard deployment uses Helm charts from the Aether project. The version of the Chart can be found in the OnRamp repository in the `vars/main.yml` file.


## Quick Navigation

| Type             | File / Directory                                                               | Description                                                            |
| ---------------- | ------------------------------------------------------------------------------ | ---------------------------------------------------------------------- |
| Core          | [`pcf.go`](./pcf.go)                                                           | Main entry point of the PCF service. Initializes routes and modules.   |
| Configuration | [`factory/pcfcfg.yaml`](./factory/pcfcfg.yaml)                                 | Main configuration file for the PCF service.                           |
| Context       | [`context/context.go`](./context/context.go)                                   | Defines the internal PCF context data structures and state management. |
| AM Policy     | [`ampolicy/api_default.go`](./ampolicy/api_default.go)                         | Implementation of Access and Mobility Policy APIs.                     |
| SM Policy     | [`smpolicy/api_default.go`](./smpolicy/api_default.go)                         | Implementation of Session Management Policy Control APIs.              |
| UE Policy     | [`uepolicy/api_default.go`](./uepolicy/api_default.go)                         | Defines policies specific to individual User Equipment (UEs).          |
| BDT Policy    | [`bdtpolicy/routers.go`](./bdtpolicy/routers.go)                               | Routers for Background Data Transfer (BDT) policy handling.            |
| Callbacks     | [`callback/api_nf_subscribe_notify.go`](./callback/api_nf_subscribe_notify.go) | Notification handlers for inter-NF communication.                      |
| NF Management | [`consumer/nf_management.go`](./consumer/nf_management.go)                     | NF registration and discovery logic with the NRF.                      |
| Metrics       | [`metrics/telemetry.go`](./metrics/telemetry.go)                               | Telemetry collection and metric export.                                |
| Utilities     | [`util/pcc_rule.go`](./util/pcc_rule.go)                                       | Helper functions for PCC rule management.                              |
| Tests         | [`pcftests/pcf_test.go`](./pcftests/pcf_test.go)                               | Integration and validation tests for the PCF.                          |
| License       | [`LICENSES/Apache-2.0.txt`](./LICENSES/Apache-2.0.txt)                         | Main project license file.                                             |



## PCF Block Diagram
![PCF Block Diagram](/docs/images/README-PCF.png)

## Supported Features
- PCF provides Access and Mobility Management related policies to the AMF
Subscription Data retrieval and AM Policy management
- PCF provides Session Management Policy Control Service to the SMF Subscription
Data retrieval and SM Policy management
- Policy Control Function (PCF) shall support interactions with the access and
mobility policy enforcement in the AMF, through service-based interfaces

## Upcoming Changes in PCF
- Process configuration received from Configuration Service and prepare PCC
Rules, Session Rules and Qos Flows
- Send PCC Rules, Session Rules to SMF when a SMF Creates Policy subscriber
- Send notification towards SMF PDU Session when PCF detects any changes in
Subscriber’s Rule/Qos information.
- Dedicated QoS flows addition and removal through APIs



Compliance of the 5G Network functions can be found at [5G Compliance](https://docs.sd-core.opennetworking.org/main/overview/3gpp-compliance-5g.html)

## Reach out to us thorugh

1. #sdcore-dev channel in [ONF Community Slack](https://onf-community.slack.com/)
2. Raise Github issues
