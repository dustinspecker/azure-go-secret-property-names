#!/bin/bash
set -e

# 1. find all json files that are not examples
# 2. find secret properties in the json files
# 3. convert the secret properties to lowercase
# 4. sort alphabetically
# 5. remove duplicates
secret_property_names=$( \
  find ./azure-rest-api-specs/specification/ \
    -name "*.json" \
    -not -path '*/examples/*' \
    -print0 \
  | xargs -0 --max-args 1 --max-procs 16 node find-secret-property-names.js\
  | tr '[:upper:]' '[:lower:]' \
  | sort \
  | uniq \
)

mkdir -p ./secretnames

tab=$'\t'

# Create a Go file that contains all secret property names.
{
  echo "package secretnames"
  echo ""
  echo "// GetAll returns all Azure Go SDK secret property names"
  echo "func GetAll() []string {"
  echo -e "${tab}return []string{"
  for secret_property_name in $secret_property_names; do
    echo "${tab}${tab}\"$secret_property_name\","
  done
  echo "${tab}}"
  echo "}"
} > ./secretnames/secretnames_all.go
