#!/bin/bash

REPO_BRANCH=
DA_API_ENABLE=0

while getopts "b:d" opt; do
  case $opt in
    b)
        REPO_BRANCH=$OPTARG
        echo "REPO_BRANCH :REPO_BRANCH" ;;
    d)
        DA_API_ENABLE=1
        echo "DA_API_ENABLE :DA_API_ENABLE" ;;
    *)
        echo "$0: invalid option -$OPTARG" >&2
		    echo "Usage: $0 [-b repo_branch] [-d <enable da api>]" >&2
		    exit
		    ;;
  esac
done

CURRENT_PATH="$( cd "$(dirname "$0")" && pwd -P )"
echo "Script dir: $CURRENT_PATH"

cd $CURRENT_PATH && cd ..
CURRENT_PATH="$(pwd -P)"
echo "Program dir: $CURRENT_PATH"

echo "===>Start to pull repo..."
CURRENT_BRANCH="$( git branch | grep '\*' | sed 's/\* //' )"
echo "[Current branch] $CURRENT_BRANCH"
git fetch
git reset --hard origin/$CURRENT_BRANCH
if [ -n "$REPO_BRANCH" ]; then
  echo "[Start to switch branch] $REPO_BRANCH"
  git fetch origin $REPO_BRANCH
  git checkout $REPO_BRANCH || exit 1
fi
echo "<===Repo pulled successfully."

echo "===>Start to build image..."
docker-compose build --no-cache || exit 1
echo "<===Images built successfully. "

echo "===>Start to launch services..."
docker-compose up --force-recreate -d sync
docker-compose up --force-recreate -d stat
echo "DA_API_ENABLE: $DA_API_ENABLE"
if [ "$DA_API_ENABLE" -ne 1 ]; then
  docker-compose up --force-recreate -d api
else
  docker-compose up --force-recreate -d da_api
fi
echo "<===Services launched successfully."