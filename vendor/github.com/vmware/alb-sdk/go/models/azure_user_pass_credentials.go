// Copyright 2021 VMware, Inc.
// SPDX-License-Identifier: Apache License 2.0
package models

// This file is auto-generated.

// AzureUserPassCredentials azure user pass credentials
// swagger:model AzureUserPassCredentials
type AzureUserPassCredentials struct {

	// Password for Azure subscription. Required only if username is provided. Field introduced in 17.2.1.
	Password *string `json:"password,omitempty"`

	// Tenant or the active directory associated with the subscription. Required for user name password authentication. Field introduced in 17.2.1.
	TenantName *string `json:"tenant_name,omitempty"`

	// Username for Azure subscription. Required only for username password based authentication. Field introduced in 17.2.1.
	Username *string `json:"username,omitempty"`
}
