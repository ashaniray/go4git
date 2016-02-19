#!/bin/bash

if [ $# -ne 1 ]
then
	echo "Usage: testShowPack <pack_file>"
	exit -1
fi

PACK_FILE=$1

echo "diff -b <(showpack -v ${PACK_FILE}) <(git verify-pack -v ${PACK_FILE})"
diff -b <(showpack -v ${PACK_FILE}) <(git verify-pack -v ${PACK_FILE})
if [ $? -eq 0 ]
then
	echo "Success"
	exit 0
else
	echo "FAILED..."
	exit -1
fi

