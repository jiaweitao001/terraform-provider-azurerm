package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OsProfile struct {
	AdminPassword      *string `json:"adminPassword,omitempty"`
	AdminUsername      *string `json:"adminUsername,omitempty"`
	ComputerName       *string `json:"computerName,omitempty"`
	OsName             *string `json:"osName,omitempty"`
	OsType             *OsType `json:"osType,omitempty"`
	ToolsRunningStatus *string `json:"toolsRunningStatus,omitempty"`
	ToolsVersion       *string `json:"toolsVersion,omitempty"`
	ToolsVersionStatus *string `json:"toolsVersionStatus,omitempty"`
}
