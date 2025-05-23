package devcenters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CatalogItemSyncEnableStatus string

const (
	CatalogItemSyncEnableStatusDisabled CatalogItemSyncEnableStatus = "Disabled"
	CatalogItemSyncEnableStatusEnabled  CatalogItemSyncEnableStatus = "Enabled"
)

func PossibleValuesForCatalogItemSyncEnableStatus() []string {
	return []string{
		string(CatalogItemSyncEnableStatusDisabled),
		string(CatalogItemSyncEnableStatusEnabled),
	}
}

func (s *CatalogItemSyncEnableStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCatalogItemSyncEnableStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCatalogItemSyncEnableStatus(input string) (*CatalogItemSyncEnableStatus, error) {
	vals := map[string]CatalogItemSyncEnableStatus{
		"disabled": CatalogItemSyncEnableStatusDisabled,
		"enabled":  CatalogItemSyncEnableStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CatalogItemSyncEnableStatus(input)
	return &out, nil
}

type IdentityType string

const (
	IdentityTypeDelegatedResourceIdentity IdentityType = "delegatedResourceIdentity"
	IdentityTypeSystemAssignedIdentity    IdentityType = "systemAssignedIdentity"
	IdentityTypeUserAssignedIdentity      IdentityType = "userAssignedIdentity"
)

func PossibleValuesForIdentityType() []string {
	return []string{
		string(IdentityTypeDelegatedResourceIdentity),
		string(IdentityTypeSystemAssignedIdentity),
		string(IdentityTypeUserAssignedIdentity),
	}
}

func (s *IdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityType(input string) (*IdentityType, error) {
	vals := map[string]IdentityType{
		"delegatedresourceidentity": IdentityTypeDelegatedResourceIdentity,
		"systemassignedidentity":    IdentityTypeSystemAssignedIdentity,
		"userassignedidentity":      IdentityTypeUserAssignedIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityType(input)
	return &out, nil
}

type InstallAzureMonitorAgentEnableStatus string

const (
	InstallAzureMonitorAgentEnableStatusDisabled InstallAzureMonitorAgentEnableStatus = "Disabled"
	InstallAzureMonitorAgentEnableStatusEnabled  InstallAzureMonitorAgentEnableStatus = "Enabled"
)

func PossibleValuesForInstallAzureMonitorAgentEnableStatus() []string {
	return []string{
		string(InstallAzureMonitorAgentEnableStatusDisabled),
		string(InstallAzureMonitorAgentEnableStatusEnabled),
	}
}

func (s *InstallAzureMonitorAgentEnableStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstallAzureMonitorAgentEnableStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstallAzureMonitorAgentEnableStatus(input string) (*InstallAzureMonitorAgentEnableStatus, error) {
	vals := map[string]InstallAzureMonitorAgentEnableStatus{
		"disabled": InstallAzureMonitorAgentEnableStatusDisabled,
		"enabled":  InstallAzureMonitorAgentEnableStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstallAzureMonitorAgentEnableStatus(input)
	return &out, nil
}

type MicrosoftHostedNetworkEnableStatus string

const (
	MicrosoftHostedNetworkEnableStatusDisabled MicrosoftHostedNetworkEnableStatus = "Disabled"
	MicrosoftHostedNetworkEnableStatusEnabled  MicrosoftHostedNetworkEnableStatus = "Enabled"
)

func PossibleValuesForMicrosoftHostedNetworkEnableStatus() []string {
	return []string{
		string(MicrosoftHostedNetworkEnableStatusDisabled),
		string(MicrosoftHostedNetworkEnableStatusEnabled),
	}
}

func (s *MicrosoftHostedNetworkEnableStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftHostedNetworkEnableStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftHostedNetworkEnableStatus(input string) (*MicrosoftHostedNetworkEnableStatus, error) {
	vals := map[string]MicrosoftHostedNetworkEnableStatus{
		"disabled": MicrosoftHostedNetworkEnableStatusDisabled,
		"enabled":  MicrosoftHostedNetworkEnableStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftHostedNetworkEnableStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateAccepted                  ProvisioningState = "Accepted"
	ProvisioningStateCanceled                  ProvisioningState = "Canceled"
	ProvisioningStateCreated                   ProvisioningState = "Created"
	ProvisioningStateCreating                  ProvisioningState = "Creating"
	ProvisioningStateDeleted                   ProvisioningState = "Deleted"
	ProvisioningStateDeleting                  ProvisioningState = "Deleting"
	ProvisioningStateFailed                    ProvisioningState = "Failed"
	ProvisioningStateMovingResources           ProvisioningState = "MovingResources"
	ProvisioningStateNotSpecified              ProvisioningState = "NotSpecified"
	ProvisioningStateRolloutInProgress         ProvisioningState = "RolloutInProgress"
	ProvisioningStateRunning                   ProvisioningState = "Running"
	ProvisioningStateStorageProvisioningFailed ProvisioningState = "StorageProvisioningFailed"
	ProvisioningStateSucceeded                 ProvisioningState = "Succeeded"
	ProvisioningStateTransientFailure          ProvisioningState = "TransientFailure"
	ProvisioningStateUpdated                   ProvisioningState = "Updated"
	ProvisioningStateUpdating                  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreated),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMovingResources),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStateRolloutInProgress),
		string(ProvisioningStateRunning),
		string(ProvisioningStateStorageProvisioningFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateTransientFailure),
		string(ProvisioningStateUpdated),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":                  ProvisioningStateAccepted,
		"canceled":                  ProvisioningStateCanceled,
		"created":                   ProvisioningStateCreated,
		"creating":                  ProvisioningStateCreating,
		"deleted":                   ProvisioningStateDeleted,
		"deleting":                  ProvisioningStateDeleting,
		"failed":                    ProvisioningStateFailed,
		"movingresources":           ProvisioningStateMovingResources,
		"notspecified":              ProvisioningStateNotSpecified,
		"rolloutinprogress":         ProvisioningStateRolloutInProgress,
		"running":                   ProvisioningStateRunning,
		"storageprovisioningfailed": ProvisioningStateStorageProvisioningFailed,
		"succeeded":                 ProvisioningStateSucceeded,
		"transientfailure":          ProvisioningStateTransientFailure,
		"updated":                   ProvisioningStateUpdated,
		"updating":                  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
