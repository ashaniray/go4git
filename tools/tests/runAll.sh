#!/bin/bash

testShowIndex.sh sample.idx
testShowPack.sh sample.pack

# bug01.pack cased incorrect decomprssing
# of data.
testShowIndex.sh bug01.idx
testShowPack.sh bug01.pack 

