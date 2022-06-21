BASEDIR=$(dirname "$0")
echo $BASEDIR
mkdir -p $BASEDIR/../client/src/types/generated
mkdir -p $BASEDIR/../proto/cosmos/base/query/v1beta1
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.42.6/proto/cosmos/base/query/v1beta1/pagination.proto -o $BASEDIR/../proto/cosmos/base/query/v1beta1/pagination.proto
mkdir -p $BASEDIR/../proto/google/api
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.42.6/third_party/proto/google/api/annotations.proto -o $BASEDIR/../proto/google/api/annotations.proto
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.42.6/third_party/proto/google/api/http.proto -o $BASEDIR/../proto/google/api/http.proto
ls $BASEDIR/../proto/checkers | xargs -I {} ./node_modules/protoc/protoc/bin/protoc \
    --plugin="$BASEDIR/node_modules/.bin/protoc-gen-ts_proto" \
    --ts_proto_out="$BASEDIR/../client/src/types/generated" \
    --proto_path="$BASEDIR/../proto" \
    --ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
    checkers/{}