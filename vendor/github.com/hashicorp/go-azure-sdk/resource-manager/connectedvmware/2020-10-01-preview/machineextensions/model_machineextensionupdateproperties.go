package machineextensions

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

type MachineExtensionUpdateProperties struct {
	AutoUpgradeMinorVersion *bool        `json:"autoUpgradeMinorVersion,omitempty"`
	ForceUpdateTag          *string      `json:"forceUpdateTag,omitempty"`
	ProtectedSettings       *interface{} `json:"protectedSettings,omitempty"`
	Publisher               *string      `json:"publisher,omitempty"`
	Settings                *interface{} `json:"settings,omitempty"`
	Type                    *string      `json:"type,omitempty"`
	TypeHandlerVersion      *string      `json:"typeHandlerVersion,omitempty"`
}
