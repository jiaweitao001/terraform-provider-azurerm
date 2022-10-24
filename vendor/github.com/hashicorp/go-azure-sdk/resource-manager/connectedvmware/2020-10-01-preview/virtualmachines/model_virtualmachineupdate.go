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

type VirtualMachineUpdate struct {
	Identity   *identity.SystemAssigned        `json:"identity,omitempty"`
	Properties *VirtualMachineUpdateProperties `json:"properties,omitempty"`
	Tags       *map[string]string              `json:"tags,omitempty"`
}
