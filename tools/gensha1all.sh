if [ $# -ne 2 ]; then
	echo "Usage: $0 <object_type> <hash>."
	echo '<object_type> is one of "tree", "blob", "commit"'
	exit -1
fi

OBJ_TYPE=$1
HASH=$2
echo =========Data=========
git cat-file $OBJ_TYPE $HASH
echo
echo =======================
echo Length: $(git cat-file $OBJ_TYPE $HASH | wc -c)
echo =========Compute sha1 for========
(printf "$OBJ_TYPE %s\0" $(git cat-file $OBJ_TYPE $HASH | wc -c); git cat-file $OBJ_TYPE $HASH) 
echo
echo ==============================

echo -n "Sha1: "
(printf "$OBJ_TYPE %s\0" $(git cat-file $OBJ_TYPE $HASH | wc -c); git cat-file $OBJ_TYPE $HASH) | sha1sum
