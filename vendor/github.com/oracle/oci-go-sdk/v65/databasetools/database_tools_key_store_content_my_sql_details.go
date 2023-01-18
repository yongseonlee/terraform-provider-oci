// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseToolsKeyStoreContentMySqlDetails The key store content.
type DatabaseToolsKeyStoreContentMySqlDetails interface {
}

type databasetoolskeystorecontentmysqldetails struct {
	JsonData  []byte
	ValueType string `json:"valueType"`
}

// UnmarshalJSON unmarshals json
func (m *databasetoolskeystorecontentmysqldetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdatabasetoolskeystorecontentmysqldetails databasetoolskeystorecontentmysqldetails
	s := struct {
		Model Unmarshalerdatabasetoolskeystorecontentmysqldetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ValueType = s.Model.ValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *databasetoolskeystorecontentmysqldetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ValueType {
	case "SECRETID":
		mm := DatabaseToolsKeyStoreContentSecretIdMySqlDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m databasetoolskeystorecontentmysqldetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m databasetoolskeystorecontentmysqldetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum Enum with underlying type: string
type DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum string

// Set of constants representing the allowable values for DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum
const (
	DatabaseToolsKeyStoreContentMySqlDetailsValueTypeSecretid DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum = "SECRETID"
)

var mappingDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum = map[string]DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum{
	"SECRETID": DatabaseToolsKeyStoreContentMySqlDetailsValueTypeSecretid,
}

var mappingDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnumLowerCase = map[string]DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum{
	"secretid": DatabaseToolsKeyStoreContentMySqlDetailsValueTypeSecretid,
}

// GetDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnumValues Enumerates the set of values for DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum
func GetDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnumValues() []DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum {
	values := make([]DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum, 0)
	for _, v := range mappingDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnumStringValues Enumerates the set of values in String for DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum
func GetDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnumStringValues() []string {
	return []string{
		"SECRETID",
	}
}

// GetMappingDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum(val string) (DatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnum, bool) {
	enum, ok := mappingDatabaseToolsKeyStoreContentMySqlDetailsValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
