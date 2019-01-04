rm -rf ByteGraph
if [ `uname` = Linux ];
then
  echo "thrift linux"
  exit $?
fi
if [ `uname` = Darwin ];
then
  echo "thrift darwin"
  ./thrift-darwin -r --out . --gen go:package_prefix=github.com\/lemonwx\/read-signed-num\/idl\/,thrift_import=code.byted.org/gopkg/thrift t.thrift
  exit $?
fi
