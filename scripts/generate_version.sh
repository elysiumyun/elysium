#!/bin/bash --posix

VER=$(cat .version)
UPT=$(date +"%Y%m%d")

WS=$(dirname $0)
TMPFILE=`mktemp /tmp/elysium.XXXXXXXXXXXX`

echo "Geterate Version Const ..."

cat > $TMPFILE << EOF
package version

const verStr string = "${VER}"
const uptStr string = "${UPT}"
EOF

cat $TMPFILE > ${WS}/../internal/version/const.go
