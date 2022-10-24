package guestagents

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

type GuestAgentProperties struct {
	Credentials        *GuestCredential        `json:"credentials,omitempty"`
	CustomResourceName *string                 `json:"customResourceName,omitempty"`
	HttpProxyConfig    *HttpProxyConfiguration `json:"httpProxyConfig,omitempty"`
	ProvisioningAction *ProvisioningAction     `json:"provisioningAction,omitempty"`
	ProvisioningState  *string                 `json:"provisioningState,omitempty"`
	Status             *string                 `json:"status,omitempty"`
	Statuses           *[]ResourceStatus       `json:"statuses,omitempty"`
	Uuid               *string                 `json:"uuid,omitempty"`
}
