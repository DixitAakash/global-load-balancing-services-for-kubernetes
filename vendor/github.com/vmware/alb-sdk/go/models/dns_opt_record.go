// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// DNSOptRecord Dns opt record
// swagger:model DnsOptRecord
type DNSOptRecord struct {

	// Flag indicating client is DNSSEC aware. Field introduced in 17.1.1.
	DnssecOk *bool `json:"dnssec_ok,omitempty"`

	// EDNS options. Field introduced in 17.1.1.
	Options []*DNSEdnsOption `json:"options,omitempty"`

	// Client requestor's UDP payload size. Field introduced in 17.1.1.
	UDPPayloadSize *int32 `json:"udp_payload_size,omitempty"`

	// EDNS version. Field introduced in 17.1.1.
	Version *int32 `json:"version,omitempty"`
}
