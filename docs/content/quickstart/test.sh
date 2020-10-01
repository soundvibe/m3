#!/bin/bash
curl -X "POST" "http://localhost:7201/api/v1/query_range?
  query=third_avenue&
  start=$(date -v -45S "+%s")&
  end=$(date "+%s")&
  step=5s" | jq .
