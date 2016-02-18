#!/bin/bash

INDEX_FILE="sample.idx"

echo "diff <(showindex  ${INDEX_FILE}) <(git show-index < ${INDEX_FILE})"
diff <(showindex  ${INDEX_FILE}) <(git show-index < ${INDEX_FILE})
if [ $? -eq 0 ]
then
	echo "Success"
	exit 0
else
	echo "FAILED..."
	exit -1
fi

