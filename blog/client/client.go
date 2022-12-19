package main

import (
	"context"
	"fmt"
	blogpb "grpc/blogpb"
	"log"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close() // Maybe this should be in a separate function and the error handled?

	c := blogpb.NewBlogServiceClient(cc)

	// create Blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Stephane",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v ", err)
	}
	fmt.Printf("Blog has been created: %v \n", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	//readblog
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "saudhuiadhoaid"})
	if err2 != nil {
		fmt.Printf("Error happend while reading: %v\n", err2)
	}

	readblogreq := &blogpb.ReadBlogRequest{BlogId: blogID}
	res, err := c.ReadBlog(context.Background(), readblogreq)
	if err != nil {
		fmt.Printf("Error happend while reading \n: %v", err)
	}
	fmt.Printf("blog read: %v", res)
}
