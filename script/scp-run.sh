#!/bin/bash

SSH_KEY_File=
LOGIN_USER=
LOGIN_SERVER_HOST=

while getopts "i:u:h:" opt; do
  case $opt in
    i)
        SSH_KEY_File=$OPTARG
        echo "SSH_KEY_File :$SSH_KEY_File" ;;
    u)
        LOGIN_USER=$OPTARG
        echo "LOGIN_USER :$LOGIN_USER" ;;
    h)
        LOGIN_SERVER_HOST=$OPTARG
        echo "LOGIN_SERVER_HOST :$LOGIN_SERVER_HOST" ;;
    *)
        echo "$0: invalid option -$OPTARG" >&2
		    echo "Usage: $0 [-i ssh_key_file] [-u login_user] [-h login_server_host]" >&2
		    exit
		    ;;
  esac
done

ssh -i $SSH_KEY_File $LOGIN_USER@$LOGIN_SERVER_HOST << bash
source /etc/profile
cd ~/0g-storage-scan/script
./run.sh
bash