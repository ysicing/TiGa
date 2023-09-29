// Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
// Use of this source code is covered by the following dual licenses:
// (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
// (2) Affero General Public License 3.0 (AGPL 3.0)
// License that can be found in the LICENSE file.

package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	logpkg "github.com/ysicing/tiga/pkg/log"
	"github.com/ysicing/tiga/pkg/util/systemutil"

	"github.com/spf13/cobra"
	"github.com/ysicing/tiga/common"
)

type ListOptions struct {
	NameOnly    bool
	Verifier    PathVerifier
	PluginPaths []string
}

func (o *ListOptions) Complete(cmd *cobra.Command) error {
	o.Verifier = &CommandOverrideVerifier{
		root:        cmd.Root(),
		seenPlugins: make(map[string]string),
	}

	o.PluginPaths = filepath.SplitList(systemutil.GetOSPath())
	return nil
}

func (o *ListOptions) Run() error {
	log := logpkg.GetInstance()
	plugins := o.ListPlugins()

	if len(plugins) == 0 {
		return fmt.Errorf("error: unable to find any tiga plugins in your PATH")
	}

	log.Info("The following compatible plugins are available:")

	for _, pluginPath := range plugins {
		if err := o.Verifier.Verify(pluginPath); err != nil {
			if o.NameOnly {
				fmt.Printf("%s  %s\n", filepath.Base(pluginPath), err)
			} else {
				fmt.Printf("%s  %s\n", pluginPath, err)
			}
		} else {
			if o.NameOnly {
				fmt.Printf("%s\n", filepath.Base(pluginPath))
			} else {
				fmt.Printf("%s\n", pluginPath)
			}
		}
	}

	return nil
}

// ListPlugins returns list of plugin paths.
func (o *ListOptions) ListPlugins() []string {
	log := logpkg.GetInstance()
	var plugins []string

	for _, dir := range uniquePathsList(o.PluginPaths) {
		if len(strings.TrimSpace(dir)) == 0 {
			continue
		}

		files, err := os.ReadDir(dir)
		if err != nil {
			log.Debugf("unable to read directory %q from your path: %v. skipping...", dir, err)
			continue
		}

		for _, f := range files {
			if f.IsDir() {
				continue
			}
			if !hasValidPrefix(f.Name(), common.ValidPrefixes) {
				continue
			}

			plugins = append(plugins, filepath.Join(dir, f.Name()))
		}
	}

	return plugins
}

// PathVerifier receives a path and determines if it is valid or not
type PathVerifier interface {
	// Verify determines if a given path is valid
	Verify(path string) error
}

type CommandOverrideVerifier struct {
	root        *cobra.Command
	seenPlugins map[string]string
}

// Verify implements PathVerifier and determines if a given path
// is valid depending on whether or not it overwrites an existing
// kubectl command path, or a previously seen plugin.
func (v *CommandOverrideVerifier) Verify(path string) error {
	if v.root == nil {
		return fmt.Errorf("unable to verify path with nil root")
	}

	// extract the plugin binary name
	segs := strings.Split(path, "/")
	fullBinName := segs[len(segs)-1]
	binName := strings.Join(strings.Split(fullBinName, "-")[1:], "-")

	cmdPath := strings.Split(binName, "-")
	if len(cmdPath) > 1 {
		// the first argument is always "kubectl" for a plugin binary
		cmdPath = cmdPath[1:]
	}

	logpkg.GetBaseInstance().Debugf("cmdPath: %v, fullBinName: %v , binName: %v", cmdPath, fullBinName, binName)

	if isExec, err := isExecutable(path); err == nil && !isExec {
		return fmt.Errorf("warning: %s identified as a tiga plugin, but it is not executable", path)
	} else if err != nil {
		return fmt.Errorf("error: unable to identify %s as an executable file: %v", path, err)
	}

	if existingPath, ok := v.seenPlugins[binName]; ok {
		return fmt.Errorf("warning: %s is overshadowed by a similarly named plugin: %s", path, existingPath)
	} else {
		v.seenPlugins[binName] = path
	}

	if cmd, _, err := v.root.Find(cmdPath); err == nil {
		return fmt.Errorf("warning: %s overwrites existing command: %q", binName, cmd.CommandPath())
	}

	return nil
}

func isExecutable(fullPath string) (bool, error) {
	info, err := os.Stat(fullPath)
	if err != nil {
		return false, err
	}

	if runtime.GOOS == "windows" {
		fileExt := strings.ToLower(filepath.Ext(fullPath))

		switch fileExt {
		case ".bat", ".cmd", ".com", ".exe", ".ps1":
			return true, nil
		}
		return false, nil
	}

	if m := info.Mode(); !m.IsDir() && m&0111 != 0 {
		return true, nil
	}

	return false, nil
}

// uniquePathsList deduplicates a given slice of strings without
// sorting or otherwise altering its order in any way.
func uniquePathsList(paths []string) []string {
	seen := map[string]bool{}
	newPaths := []string{}
	for _, p := range paths {
		if seen[p] {
			continue
		}
		seen[p] = true
		newPaths = append(newPaths, p)
	}
	return newPaths
}

func hasValidPrefix(filepath string, validPrefixes []string) bool {
	for _, prefix := range validPrefixes {
		if !strings.HasPrefix(filepath, prefix+"-") {
			continue
		}
		return true
	}
	return false
}
