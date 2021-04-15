#!/usr/bin/env bash
# Generate the documentation using tfplugindocs and remove changes to files that shouldn't change

exclude_files=(
  # Timeboard + Screenbaord is deprecated
  'docs/resources/timeboard.md'
  'docs/resources/screenboard.md'
  # There is an issue with the security_monitoring schema that requires the docs to be updated by hand if needed
  'docs/data-sources/security_monitoring_rules.md'
  'docs/resources/security_monitoring_default_rule.md'
  'docs/resources/security_monitoring_rule.md'
)

# Check if manual changes were made to any excluded files and exit
# otherwise these will be lost with `tfplugindocs`
if [ "$(git status --porcelain "${exclude_files[@]}")" ]; then
  echo "Uncommitted changes were detected to the following files. These aren't autogenerated, please commit or stash these changes and try again"
  echo $(git status --porcelain "${exclude_files[@]}")
  exit 1
fi

tfplugindocs


# Remove the changes to files we don't autogenerate
git checkout HEAD -- "${exclude_files[@]}"
