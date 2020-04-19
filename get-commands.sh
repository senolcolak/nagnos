#!/bin/bash
curl 'http://127.0.0.1:5000/commandz' | jq -r '.commands[][]' >.commandz
