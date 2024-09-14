#!/bin/bash

REPO_BRANCH=

while getopts "b:" opt; do
  case $opt in
    b)
        REPO_BRANCH=$OPTARG
        echo "REPO_BRANCH :REPO_BRANCH" ;;
    *)
        echo "$0: invalid option -$OPTARG" >&2
		    echo "Usage: $0 [-b repo_branch]" >&2
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
if [ -n "$REPO_BRANCH" ]; then
  echo "[Start to switch branch] $REPO_BRANCH"
  git checkout $REPO_BRANCH || exit 1
fi
git pull || exit 1
echo "<===Repo is pulled successfully."

echo "===>Start to build image..."
docker-compose build --no-cache || exit 1
echo "<===Images are built successfully. "

echo "===>Start to launch services..."
docker-compose up -d
echo "<===Services are launched successfully."