package api

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	m "github.com/AntonyIS/go-grpc-crud/server/domain"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

type movieHandler interface {
	CreateMovie(*gin.Context)
	GetMovie(*gin.Context)
	GetMovies(*gin.Context)
	UpdateMovie(*gin.Context)
	DeleteMovie(*gin.Context)
}

type handler struct{}

func gRPCClient() (pb.MovieServiceClient, error) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pb.NewMovieServiceClient(conn)

	return client, nil
}

func NewHandler() movieHandler {
	return handler{}
}

func (handler) CreateMovie(ctx *gin.Context) {
	movie := m.Movie{}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := gRPCClient()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}

	res, err := client.CreateMovie(ctx, &pb.MovieRequest{
		Name:        movie.Name,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Image:       movie.Image,
	})
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}
	movie.ID = res.Id
	ctx.JSON(http.StatusOK, movie)
}

func (handler) GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	client, err := gRPCClient()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err),
		})
		return
	}
	res, err := client.GetMovie(ctx, &pb.MovieID{Id: id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"errors": "Movie not found",
		})
		return
	}
	movie := make(map[string]string)
	movie["id"] = res.GetId()
	movie["name"] = res.GetName()
	movie["description"] = res.GetDescription()
	movie["release_date"] = res.GetReleaseDate()
	movie["image"] = res.GetReleaseDate()
	ctx.JSON(http.StatusOK, movie)
}

func (handler) GetMovies(ctx *gin.Context) {
	client, err := gRPCClient()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}
	res, err := client.GetMovies(ctx, &pb.EmptyRequest{})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Movies)

}

func (handler) UpdateMovie(ctx *gin.Context) {
	movie := m.Movie{}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := gRPCClient()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}

	res, err := client.UpdateMovie(ctx, &pb.MovieRequest{
		Id:          movie.ID,
		Name:        movie.Name,
		Description: movie.Description,
		ReleaseDate: movie.ReleaseDate,
		Image:       movie.Image,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}
	movie.ID = res.Id
	ctx.JSON(http.StatusOK, movie)
}

func (handler) DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	client, err := gRPCClient()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}
	res, err := client.DeleteMovie(ctx, &pb.MovieID{Id: id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Error: %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": res.GetMessage(),
	})
}
