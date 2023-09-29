// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package types

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Flag struct {
	Name      string
	Shorthand string
	Usage     string
	Required  bool
	EnvVar    string
	P         interface{}
	V         interface{}
}

// StringArray gorm custom string array flag type.
type StringArray []string

// Scan gorm Scan implement.
func (a *StringArray) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case string:
		if v != "" {
			*a = strings.Split(v, ",")
		}
	default:
		return fmt.Errorf("failed to scan array value %v", value)
	}
	return nil
}

// Value gorm Value implement.
func (a StringArray) Value() (driver.Value, error) {
	if len(a) == 0 {
		return nil, nil
	}
	return strings.Join(a, ","), nil
}

// GormDataType returns gorm data type.
func (a StringArray) GormDataType() string {
	return "string"
}

func (a StringArray) Contains(target string) bool {
	for _, content := range a {
		if target == content {
			return true
		}
	}
	return false
}
