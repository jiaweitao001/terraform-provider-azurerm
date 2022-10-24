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

type VirtualDisk struct {
	ControllerKey *int64    `json:"controllerKey,omitempty"`
	DeviceKey     *int64    `json:"deviceKey,omitempty"`
	DeviceName    *string   `json:"deviceName,omitempty"`
	DiskMode      *DiskMode `json:"diskMode,omitempty"`
	DiskObjectId  *string   `json:"diskObjectId,omitempty"`
	DiskSizeGB    *int64    `json:"diskSizeGB,omitempty"`
	DiskType      *DiskType `json:"diskType,omitempty"`
	Label         *string   `json:"label,omitempty"`
	Name          *string   `json:"name,omitempty"`
	UnitNumber    *int64    `json:"unitNumber,omitempty"`
}
