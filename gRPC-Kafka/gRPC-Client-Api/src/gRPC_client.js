var PROTO_PATH = __dirname + '/proto/games.proto';

var parseArgs = require('minimist');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
  });

var games_proto = grpc.loadPackageDefinition(packageDefinition).games;

var argv = parseArgs(process.argv.slice(2), {
  string: 'target'
});

var target;
if (argv.target) {
  target = argv.target;
} else {
  target = process.env.GRPC_SERVER_HOST + ':'+ process.env.GRPC_SERVER_PORT; 
}
var client = new games_proto.IngressGame(target, grpc.credentials.createInsecure());

module.exports = client;
