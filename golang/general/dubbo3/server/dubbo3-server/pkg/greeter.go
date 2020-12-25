/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"context"
	"fmt"
	"github.com/apache/dubbo-samples/golang/general/dubbo3/protobuf/dubbo3"
)

type GreeterProvider struct {
	*dubbo3.Dubbo3GreeterProviderBase
}

func NewGreeterProvider() *GreeterProvider {
	return &GreeterProvider{
		Dubbo3GreeterProviderBase: &dubbo3.Dubbo3GreeterProviderBase{},
	}
}

// Dubbo3SayHello2 is a unary-grpc-grpc-client rpc example
func (s *GreeterProvider) Dubbo3SayHello2(ctx context.Context, in *dubbo3.Dubbo3HelloRequest) (*dubbo3.Dubbo3HelloReply, error) {
	fmt.Println("######### get server request name :" + in.Myname)
	return &dubbo3.Dubbo3HelloReply{Msg: "Hello " + in.Myname}, nil
}

// Dubbo3SayHello is a server rpc exmple
func (g *GreeterProvider) Dubbo3SayHello(svr dubbo3.Dubbo3Greeter_Dubbo3SayHelloServer) error {
	c, err := svr.Recv()
	if err != nil {
		return err
	}
	fmt.Println("server server recv 1 = ", c)
	c2, err := svr.Recv()
	if err != nil {
		return err
	}
	fmt.Println("server server recv 2 = ", c2)
	c3, err := svr.Recv()
	fmt.Println("server server recv 3 = ", c3)

	svr.Send(&dubbo3.Dubbo3HelloReply{
		Msg: c.Myname + c2.Myname,
	})
	fmt.Println("server server send 1 = ", c.Myname+c2.Myname)
	svr.Send(&dubbo3.Dubbo3HelloReply{
		Msg: c3.Myname,
	})
	fmt.Println("server server send 2 = ", c3.Myname)
	return nil
}

func (g *GreeterProvider) Reference() string {
	return "GrpcGreeterImpl"
}
