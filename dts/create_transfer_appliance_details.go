// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Transfer Service API
//
// Data Transfer Service API Specification
//

package dts

import (
	"github.com/oracle/oci-go-sdk/v35/common"
)

// CreateTransferApplianceDetails The representation of CreateTransferApplianceDetails
type CreateTransferApplianceDetails struct {
	CustomerShippingAddress *ShippingAddress `mandatory:"false" json:"customerShippingAddress"`
}

func (m CreateTransferApplianceDetails) String() string {
	return common.PointerString(m)
}
