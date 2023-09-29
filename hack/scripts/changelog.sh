#!/usr/bin/env sh
# Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
# Use of this source code is covered by the following dual licenses:
# (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
# (2) Affero General Public License 3.0 (AGPL 3.0)
# License that can be found in the LICENSE file.

set -eu

echo '# Changelog'
echo

tag=
git tag -l 'v*' | sort -rV | while read last; do
  if [ "$tag" != "" ]; then
    echo "## $(git for-each-ref --format='%(refname:strip=2) (%(creatordate:short))' refs/tags/${tag})"
    echo
    git_log='git --no-pager log --no-merges --invert-grep --grep=^\(build\|ci\|docs\|test\|chore\):'
	  $git_log --format=' * [%h](https://github.com/ysicing/tiga/commit/%H) %s' $last..$tag
	  echo
	  echo "### Contributors"
	  echo
	  $git_log --format=' * %an'  $last..$tag | sort -u
	  echo
  fi
  tag=$last
done
