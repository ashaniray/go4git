#!/bin/bash

INDEX_FILE="sample.pack"

echo "diff -b <(showpack -v ${INDEX_FILE}) <(git verify-pack -v ${INDEX_FILE})"
diff -b <(showpack -v ${INDEX_FILE}) <(git verify-pack -v ${INDEX_FILE})
if [ $? -eq 0 ]
then
	echo "Success"
	exit 0
else
	echo "FAILED..."
	exit -1
fi

