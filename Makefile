
all: proto

proto:
	protoc -I events events/events.proto --go_out=plugins=grpc:events
