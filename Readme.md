This is the rest api that will call internally grpc kafka_example_java

How to generate a proto 
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
ex.  protoc -I=proto --go_out=proto proto/Greet.proto 

protoc -I=proto --go_out=plugins=grpc:proto proto/greet.proto