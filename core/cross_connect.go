// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// CrossConnect. For use with Oracle Cloud Infrastructure FastConnect. A cross-connect represents a
// physical connection between an existing network and Oracle. Customers who are colocated
// with Oracle in a FastConnect location create and use cross-connects. For more
// information, see [FastConnect Overview]({{DOC_SERVER_URL}}/Content/Network/Concepts/fastconnect.htm).
// Oracle recommends you create each cross-connect in a
// CrossConnectGroup so you can use link aggregation
// with the connection.
// **Note:** If you're a provider who is setting up a physical connection to Oracle so customers
// can use FastConnect over the connection, be aware that your connection is modeled the
// same way as a colocated customer's (with `CrossConnect` and `CrossConnectGroup` objects, and so on).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies]({{DOC_SERVER_URL}}/Content/Identity/Concepts/policygetstarted.htm).
type CrossConnect struct {

	// The OCID of the compartment containing the cross-connect group.
	CompartmentID *string `mandatory:"false" json:"compartmentId,omitempty"`

	// The OCID of the cross-connect group this cross-connect belongs to (if any).
	CrossConnectGroupID *string `mandatory:"false" json:"crossConnectGroupId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The cross-connect's Oracle ID (OCID).
	ID *string `mandatory:"false" json:"id,omitempty"`

	// The cross-connect's current state.
	LifecycleState CrossConnectLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The name of the FastConnect location where this cross-connect is installed.
	LocationName *string `mandatory:"false" json:"locationName,omitempty"`

	// A string identifying the meet-me room port for this cross-connect.
	PortName *string `mandatory:"false" json:"portName,omitempty"`

	// The port speed for this cross-connect.
	// Example: `10 Gbps`
	PortSpeedShapeName *string `mandatory:"false" json:"portSpeedShapeName,omitempty"`

	// The date and time the cross-connect was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated,omitempty"`
}

func (model CrossConnect) String() string {
	return common.PointerString(model)
}

type CrossConnectLifecycleStateEnum string

const (
	CROSS_CONNECT_LIFECYCLE_STATE_PENDING_CUSTOMER CrossConnectLifecycleStateEnum = "PENDING_CUSTOMER"
	CROSS_CONNECT_LIFECYCLE_STATE_PROVISIONING     CrossConnectLifecycleStateEnum = "PROVISIONING"
	CROSS_CONNECT_LIFECYCLE_STATE_PROVISIONED      CrossConnectLifecycleStateEnum = "PROVISIONED"
	CROSS_CONNECT_LIFECYCLE_STATE_INACTIVE         CrossConnectLifecycleStateEnum = "INACTIVE"
	CROSS_CONNECT_LIFECYCLE_STATE_TERMINATING      CrossConnectLifecycleStateEnum = "TERMINATING"
	CROSS_CONNECT_LIFECYCLE_STATE_TERMINATED       CrossConnectLifecycleStateEnum = "TERMINATED"
	CROSS_CONNECT_LIFECYCLE_STATE_UNKNOWN          CrossConnectLifecycleStateEnum = "UNKNOWN"
)

var mapping_crossconnect_lifecycleState = map[string]CrossConnectLifecycleStateEnum{
	"PENDING_CUSTOMER": CROSS_CONNECT_LIFECYCLE_STATE_PENDING_CUSTOMER,
	"PROVISIONING":     CROSS_CONNECT_LIFECYCLE_STATE_PROVISIONING,
	"PROVISIONED":      CROSS_CONNECT_LIFECYCLE_STATE_PROVISIONED,
	"INACTIVE":         CROSS_CONNECT_LIFECYCLE_STATE_INACTIVE,
	"TERMINATING":      CROSS_CONNECT_LIFECYCLE_STATE_TERMINATING,
	"TERMINATED":       CROSS_CONNECT_LIFECYCLE_STATE_TERMINATED,
	"UNKNOWN":          CROSS_CONNECT_LIFECYCLE_STATE_UNKNOWN,
}

func GetCrossConnectLifecycleStateEnumValues() []CrossConnectLifecycleStateEnum {
	values := make([]CrossConnectLifecycleStateEnum, 0)
	for _, v := range mapping_crossconnect_lifecycleState {
		if v != CROSS_CONNECT_LIFECYCLE_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}