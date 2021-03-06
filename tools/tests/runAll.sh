#!/bin/bash

testShowIndex.sh sample.idx
testShowPack.sh sample.pack

# bug01.pack cased incorrect decomprssing
# of data.
testShowIndex.sh bug01.idx
testShowPack.sh bug01.pack 

# bug02.pack due to incorrect reading of ref_delta
testShowIndex.sh bug02.idx
testShowPack.sh bug02.pack 

# bug03.pack due to incorrect reading of ref_delta
# when length = 0
testShowIndex.sh bug03.idx
testShowPack.sh bug03.pack 

# bug04.pack due to single object entry
testShowIndex.sh bug04.idx
testShowPack.sh bug04.pack 

for PACK_FILE in $(find ~ -name "*.pack" 2>/dev/null)
do
	echo -n 
# Optionally run the following
	testShowPack.sh ${PACK_FILE}
done

