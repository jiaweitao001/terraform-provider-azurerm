package virtualmachinetemplates

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/edgezones"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NicIPSettings struct {
	AllocationMethod    *IPAddressAllocationMethod `json:"allocationMethod,omitempty"`
	DnsServers          *[]string                  `json:"dnsServers,omitempty"`
	Gateway             *[]string                  `json:"gateway,omitempty"`
	IpAddress           *string                    `json:"ipAddress,omitempty"`
	IpAddressInfo       *[]NicIPAddressSettings    `json:"ipAddressInfo,omitempty"`
	PrimaryWinsServer   *string                    `json:"primaryWinsServer,omitempty"`
	SecondaryWinsServer *string                    `json:"secondaryWinsServer,omitempty"`
	SubnetMask          *string                    `json:"subnetMask,omitempty"`
}
