package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "gorpc/proto"
	"io"
)

type Client struct {
	client pb.DemoClient
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := Client{
		client: pb.NewDemoClient(conn),
	}
	return &client, nil
}

func (c *Client) Version() (string, error) {
	reply, err := c.client.ServerVersion(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		return "", err
	}

	return reply.Version, nil
}

func (c *Client) Watch(id string) error {
	reply, err := c.client.Watch(context.Background(), &pb.VideoRequest{
		Id: id,
	})
	if err != nil {
		return err
	}

	for {
		block, err := reply.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(string(block.Block))
	}

	return nil
}
