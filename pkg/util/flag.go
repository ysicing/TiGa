// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package util

import (
	"bytes"
	"encoding/csv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/ysicing/tiga/internal/types"
)

// thank autok3s for this code
// https://github.com/cnrancher/autok3s/blob/master/pkg/utils/flag.go

// BashCompEnvVarFlag cobra flag's annotation used for bind env to flag.
const BashCompEnvVarFlag = "cobra_annotation_bash_env_var_flag"

// ConvertFlags change  flags to FlagSet, will mark required annotation if possible.
func ConvertFlags(cmd *cobra.Command, fs []types.Flag) *pflag.FlagSet {
	for _, f := range fs {
		if f.Shorthand == "" {
			if cmd.Flags().Lookup(f.Name) == nil {
				pf := cmd.Flags()
				switch t := f.V.(type) {
				case bool:
					pf.BoolVar(f.P.(*bool), f.Name, t, f.Usage)
				case string:
					pf.StringVar(f.P.(*string), f.Name, t, f.Usage)
				case map[string]string:
					pf.StringToStringVar(f.P.(*map[string]string), f.Name, t, f.Usage)
				case []string:
					pf.StringArrayVar(f.P.(*[]string), f.Name, t, f.Usage)
				case types.StringArray:
					pf.Var(newStringArrayValue(t, f.P.(*types.StringArray)), f.Name, f.Usage)
				case int:
					pf.IntVar(f.P.(*int), f.Name, t, f.Usage)
				default:
					continue
				}
				if f.Required {
					_ = cobra.MarkFlagRequired(pf, f.Name)
				}
			}
		} else {
			if cmd.Flags().Lookup(f.Name) == nil {
				pf := cmd.Flags()
				switch t := f.V.(type) {
				case bool:
					pf.BoolVarP(f.P.(*bool), f.Name, f.Shorthand, t, f.Usage)
				case string:
					pf.StringVarP(f.P.(*string), f.Name, f.Shorthand, t, f.Usage)
				case map[string]string:
					pf.StringToStringVarP(f.P.(*map[string]string), f.Name, f.Shorthand, t, f.Usage)
				case []string:
					pf.StringArrayVarP(f.P.(*[]string), f.Name, f.Shorthand, t, f.Usage)
				case types.StringArray:
					pf.VarP(newStringArrayValue(t, f.P.(*types.StringArray)), f.Name, f.Shorthand, f.Usage)
				default:
					continue
				}
				if f.Required {
					_ = cobra.MarkFlagRequired(pf, f.Name)
				}
			}
		}

		if f.EnvVar != "" {
			_ = cmd.Flags().SetAnnotation(f.Name, BashCompEnvVarFlag, []string{f.EnvVar})
		}
	}
	return cmd.Flags()
}

type stringArrayValue struct {
	value   *types.StringArray
	changed bool
}

func newStringArrayValue(val []string, p *types.StringArray) *stringArrayValue {
	ssv := new(stringArrayValue)
	ssv.value = p
	*ssv.value = val
	return ssv
}

func (s *stringArrayValue) Set(val string) error {
	if !s.changed {
		*s.value = []string{val}
		s.changed = true
	} else {
		*s.value = append(*s.value, val)
	}
	return nil
}

func (s *stringArrayValue) Append(val string) error {
	*s.value = append(*s.value, val)
	return nil
}

func (s *stringArrayValue) Replace(val []string) error {
	out := make([]string, len(val))
	copy(out, val)
	*s.value = out
	return nil
}

func (s *stringArrayValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	copy(out, *s.value)
	return out
}

func (s *stringArrayValue) Type() string {
	return "stringArray"
}

func (s *stringArrayValue) String() string {
	str, _ := writeAsCSV(*s.value)
	return "[" + str + "]"
}

func writeAsCSV(ss []string) (string, error) {
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	err := w.Write(ss)
	if err != nil {
		return "", err
	}
	w.Flush()
	return strings.TrimSuffix(b.String(), "\n"), nil
}
