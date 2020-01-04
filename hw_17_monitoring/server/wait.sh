#!/bin/sh
# wait-for-postgres.sh

RETRIES=5

until psql -h ${REG_SERVICE_DB_HOST} -U ${REG_SERVICE_DB_USER} -d ${REG_SERVICE_DB_NAME} -c "select 1" > /dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
  echo "Waiting for postgres server, $((RETRIES--)) remaining attempts..."
  sleep 2
done

while psql -h ${REG_SERVICE_DB_HOST} -U ${REG_SERVICE_DB_USER} -d ${REG_SERVICE_DB_NAME} -c "select 1";
do
  echo "Waiting for postgres server..."
  sleep 2;
 done

done psql -h ${REG_SERVICE_DB_HOST} -U ${REG_SERVICE_DB_USER} -d ${REG_SERVICE_DB_NAME} -c "select 1" > /dev/null 2>&1 || [ $RETRIES -eq 0 ]; do
  echo "Waiting for postgres server, $((RETRIES--)) remaining attempts..."
  sleep 2
done
#
#
#set -e
#
#host="${1}"
#echo "NOW - HOST"
#echo host
#echo "${REG_SERVICE_LOG_LEVEL}   long This work"
#echo "${LOG_LEVEL} short"