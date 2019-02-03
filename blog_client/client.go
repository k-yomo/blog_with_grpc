package main

import (
	"context"
	"fmt"
	"github.com/k-yomo/blog_with_grpc/blogpb"
	"google.golang.org/grpc"
	"io"
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

	// create blog
	fmt.Println("Createing the blog")
	blog := &blogpb.Blog{
		AuthorId: "k-yomo",
		Title: "My first blog",
		Content: "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error while creating blog: %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)

	// read Blog
	fmt.Println("Reading the blog")

	_, err = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5c54b08ebe020342da2b13a8"})
	if err != nil {
		fmt.Printf("Error happend while reading: %v\n", err)
	}

	readBlogRes, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: createBlogRes.GetBlog().GetId()})
	if err != nil {
		fmt.Printf("Error happend while reading: %v\n", err)
	}

	fmt.Printf("Blog was read: %v", readBlogRes)

	// update Blog
	fmt.Println("Createing the blog")
	updatedBlog := &blogpb.Blog{
		Id: createBlogRes.GetBlog().GetId(),
		AuthorId: "abc",
		Title: "My first blog(updated)",
		Content: "Content of the first blog(updated)",
	}
	updateBlogRes, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: updatedBlog})
	if err != nil {
		log.Fatalf("Unexpected error while creating blog: %v", err)
	}
	fmt.Printf("Blog has been updated: %v\n", updateBlogRes)

	// delete Blog
	deleteRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: createBlogRes.GetBlog().GetId()})

	if err != nil {
		fmt.Printf("Error happened while deleting: %v\n", err)
	}

	fmt.Printf("Blog has been deleted %v\n", deleteRes)

	// list Blogs
	stream , err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})

	if err != nil {
		log.Fatalf("Error while calling ListBlog RPC: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}

