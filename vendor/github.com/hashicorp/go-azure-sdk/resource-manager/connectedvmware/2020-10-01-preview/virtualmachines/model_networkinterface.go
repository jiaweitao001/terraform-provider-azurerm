package virtualmachines

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

type NetworkInterface struct {
	DeviceKey      *int64             `json:"deviceKey,omitempty"`
	IpAddresses    *[]string          `json:"ipAddresses,omitempty"`
	IpSettings     *NicIPSettings     `json:"ipSettings,omitempty"`
	Label          *string            `json:"label,omitempty"`
	MacAddress     *string            `json:"macAddress,omitempty"`
	Name           *string            `json:"name,omitempty"`
	NetworkId      *string            `json:"networkId,omitempty"`
	NetworkMoName  *string            `json:"networkMoName,omitempty"`
	NetworkMoRefId *string            `json:"networkMoRefId,omitempty"`
	NicType        *NICType           `json:"nicType,omitempty"`
	PowerOnBoot    *PowerOnBootOption `json:"powerOnBoot,omitempty"`
}
