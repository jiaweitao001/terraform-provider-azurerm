package clusters

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

type ClusterProperties struct {
	CustomResourceName *string           `json:"customResourceName,omitempty"`
	DatastoreIds       *[]string         `json:"datastoreIds,omitempty"`
	InventoryItemId    *string           `json:"inventoryItemId,omitempty"`
	MoName             *string           `json:"moName,omitempty"`
	MoRefId            *string           `json:"moRefId,omitempty"`
	NetworkIds         *[]string         `json:"networkIds,omitempty"`
	ProvisioningState  *string           `json:"provisioningState,omitempty"`
	Statuses           *[]ResourceStatus `json:"statuses,omitempty"`
	Uuid               *string           `json:"uuid,omitempty"`
	VCenterId          *string           `json:"vCenterId,omitempty"`
}
