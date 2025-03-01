// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// PortMatch port match
// swagger:model PortMatch
type PortMatch struct {

	// Criterion to use for port matching the HTTP request. Enum options - IS_IN, IS_NOT_IN.
	// Required: true
	MatchCriteria *string `json:"match_criteria"`

	// Listening TCP port(s). Allowed values are 1-65535. Minimum of 1 items required.
	Ports []int64 `json:"ports,omitempty,omitempty"`
}
