package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"
)

type MyCSIDriver struct{}

func (*MyCSIDriver) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	fmt.Println("CreateVolume called")
	// Implement your logic to create a volume
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			Id:           "vol-123",
			CapacityBytes: 1024 * 1024 * 1024, // 1 GB
		},
	}, nil
}

func (*MyCSIDriver) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	fmt.Println("DeleteVolume called")
	// Implement your logic to delete a volume
	return &csi.DeleteVolumeResponse{}, nil
}

func (*MyCSIDriver) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	fmt.Println("ControllerPublishVolume called")
	// Implement your logic to publish a volume
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (*MyCSIDriver) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	fmt.Println("NodePublishVolume called")
	// Implement your logic to publish a volume on a node
	return &csi.NodePublishVolumeResponse{}, nil
}

func (*MyCSIDriver) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	fmt.Println("NodeUnpublishVolume called")
	// Implement your logic to unpublish a volume from a node
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func main() {
	var (
		endpoint = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	)
	flag.Parse()

	listener, err := net.Listen("unix", *endpoint)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	server := grpc.NewServer()
	csi.RegisterIdentityServer(server, &MyCSIDriver{})
	csi.RegisterControllerServer(server, &MyCSIDriver{})
	csi.RegisterNodeServer(server, &MyCSIDriver{})

	fmt.Printf("Listening for CSI requests at %s\n", *endpoint)
	err = server.Serve(listener)
	if err != nil {
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}
