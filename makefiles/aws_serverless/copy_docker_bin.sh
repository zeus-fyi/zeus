#!/bin/sh

output=$(docker run ${{ env.SERVERLESS_LATEST }} /usr/bin/ethereumsignbls 2>/dev/null) || true
echo "$output" > main
zip main.zip main
