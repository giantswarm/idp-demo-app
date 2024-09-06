#!/bin/bash

if [[ "$#" -ne 1 ]]; then
  echo "Usage: $0 <organization/repo_to_copy_to>"
  exit 1
fi

if ! command -v "gh" &>/dev/null; then
  echo "Please install gh"
  exit 1
fi

if [[ -z "$GITHUB_TOKEN" ]]; then
  echo "\$GITHUB_TOKEN is not set"
  exit 1
fi

target_repo="$1"
project_name="${target_repo##*/}"
source_dir=$(pwd)
cd /tmp || exit 2
if [[ -d "${project_name}" ]]; then
  echo "Please cleanup (or just remove) /tmp/${project_name}"
  exit 3
fi
gh repo clone "$target_repo"
cd "${project_name}" || exit 2
git config user.email "developer@demotechinc.com"
git config user.name "Platform Developer"
git checkout -b new-app-feature
cp "$source_dir"/go{.mod,.sum} .
cp -r "$source_dir"/templates/ .
cp "$source_dir"/*.go .
cp helm/values.* helm/
cp helm/templates/*.yaml helm/templates/
sed -i "s/idp-demo-app/${project_name}/g" helm/templates/*.yaml
git add -A
git commit -am "First release of the new app"
git push --set-upstream origin new-app-feature
gh pr create --title "New functionality of the app" --body "This is a new release of the app" --base main
cd "$source_dir" || exit 2
