/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	pb "github.com/apache/dubbo-samples/golang/general/dubbo3/protobuf/grpc"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:20000"
	//address     = "localhost:50051"
)

func main() {
	// Set up a connection to the grpc-client.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDubbo3GreeterClient(conn)

	// Contact the grpc-client and print out its response.
	name := "jifeng"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for i := 0; i < 100; i++ {
		r, err := c.Dubbo3SayHello(context.Background())
		if err != nil {
			fmt.Println("say hello err:", err)
		}

		if err := r.Send(&pb.Dubbo3HelloRequest{Myname: name}); err != nil {
			fmt.Println("say hello err:", err)
		}
		if err := r.Send(&pb.Dubbo3HelloRequest{Myname: name}); err != nil {
			fmt.Println("say hello err:", err)
		}
		//time.Sleep(time.Second * 10)
		if err := r.Send(&pb.Dubbo3HelloRequest{Myname: name}); err != nil {
			fmt.Println("say hello err:", err)
		}
		rsp := &pb.Dubbo3HelloReply{}
		if err := r.RecvMsg(rsp); err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Printf("firstSend Got rsp = %+v\n", rsp)
		rsp = &pb.Dubbo3HelloReply{}
		if err := r.RecvMsg(rsp); err != nil {
			fmt.Println("err = ", err)
		}
		fmt.Printf("firstSend Got rsp = %+v\n", rsp)

	}

}
