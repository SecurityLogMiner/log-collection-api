#!/bin/bash
token="$1"
curl --request GET --url http:/localhost:6060/api/messages/private --header 'authorization: Bearer ${token}'
echo ""