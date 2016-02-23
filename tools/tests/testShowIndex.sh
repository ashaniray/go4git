#!/bin/bash

if [ $# -ne 1 ]
then
	echo "Usage: testShowIndex <index_file>"
	exit -1
fi

INDEX_FILE=$1

echo "diff <(showindex  ${INDEX_FILE}) <(git show-index < ${INDEX_FILE})"
diff <(showindex  ${INDEX_FILE}) <(git show-index < ${INDEX_FILE}) > /dev/null
if [ $? -eq 0 ]
then
	echo "Success"
	exit 0
else
	echo "FAILED..."
	exit -1
fi

