package main

import (
	"context"
	"fmt"
	"github.com/k-yomo/blog_with_grpc/blogpb"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	fmt.Println("Blog Client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Createing the blog")
	blog := &blogpb.Blog{
		AuthorId: "k-yomo",
		Title: "My first blog",
		Content: "Content of the first blog",
	}
	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error while creating blog: %v", err)
	}
	fmt.Printf("Blog has been created: %v", res)
}

