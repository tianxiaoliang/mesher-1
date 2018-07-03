diff -u <(echo -n) <(find . -name "*.go" -print0 | xargs --null goconst)
if [ $? == 0 ]; then
	echo "No goConst problem"
	exit 0
else
	echo "Has goConst Problem"
	exit 1
fi
