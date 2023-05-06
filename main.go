package main

import (
	"log"

	"github.com/achmadrizkin/go_social_media_API/allproducts"
	"github.com/achmadrizkin/go_social_media_API/comment"
	"github.com/achmadrizkin/go_social_media_API/explore"
	"github.com/achmadrizkin/go_social_media_API/handler"
	"github.com/achmadrizkin/go_social_media_API/like"
	"github.com/achmadrizkin/go_social_media_API/reels"
	"github.com/achmadrizkin/go_social_media_API/user"
	userfollower "github.com/achmadrizkin/go_social_media_API/user_follower"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// connect to database
	// PLEASE CREATE go-ecommerce database first.
	dsn := "root:meru@1221@tcp(127.0.0.1:3306)/goecommerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	db.AutoMigrate(&user.User{}, &explore.Explore{}, &reels.Reels{}, &allproducts.AllProduct{}, &comment.Comment{}, &userfollower.UserFollower{}, &like.Like{})

	// API Versioning
	v1 := r.Group("/v1")

	// USER
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewBookHandler(userService)

	v1.GET("/users/:name_user", userHandler.GetUserByName)
	v1.GET("/users/email/:email_user", userHandler.GetUserByEmail)
	v1.GET("/users/id/:id", userHandler.GetUserById)
	v1.POST("/users/a/:email_user", userHandler.CreateIfNotExistOrUpdateIfExist)
	v1.POST("/users/a/:email_user/:id", userHandler.UpdateUserPost)
	v1.POST("/users/following/:email_user/:id", userHandler.UpdateUserFollowing)
	v1.PUT("/users/id/:id", userHandler.UpdateUser)
	v1.POST("/users/following/a/:email_user")

	// EXPLORE
	exploreRepository := explore.NewRepository(db)
	exploreService := explore.NewService(exploreRepository)
	exploreHandler := handler.NewExploreHandler(exploreService)

	v1.POST("/explore/a/post", exploreHandler.PostExploreHandler)
	v1.GET("/explore/:user", exploreHandler.GetExploreNotUserAndOrderByLike)
	v1.GET("/explore/user/:email", exploreHandler.GetExploreByEmailAndOrderByCreateAt)
	v1.GET("/explore/home/:email", exploreHandler.GetUserFollowingPost)

	// REELS
	reelsRepository := reels.NewRepository(db)
	reelsService := reels.NewService(reelsRepository)
	reelsHandler := handler.NewReelsHandler(reelsService)

	v1.POST("/reels", reelsHandler.PostReelsHandler)
	v1.GET("/reels/e/:user", reelsHandler.GetReelsNotUserAndOrderByLike)

	// ALL PRODUCTS
	allProductRepository := allproducts.NewRepository(db)
	allProductService := allproducts.NewService(allProductRepository)
	allProductHandler := handler.NewAllProductHandler(allProductService)

	v1.POST("/products", allProductHandler.PostBooksHandler)
	v1.GET("/products", allProductHandler.GetBooksList)
	v1.GET("/products/:id", allProductHandler.GetBookById)
	v1.GET("/products/category/:category", allProductHandler.GetBookByCategory)
	v1.GET("/products/user/:email_user", allProductHandler.GetBookByUser)
	v1.PUT("/products/:id", allProductHandler.UpdateBook)
	v1.DELETE("/products/:id", allProductHandler.DeleteBook)

	// GET BY name_product, and email
	v1.GET("/products/a/:name_product/:email", allProductHandler.GetBookByProductNameAndEmail)

	// search ALL PRODUCTS
	v1.GET("/products/s/:name_product", allProductHandler.GetProductListByName)

	// COMMENT
	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	v1.POST("/comment/a", commentHandler.PostCommentHandler)
	v1.GET("/comment/:postUser/:toId", commentHandler.GetToUserAndPostUser)

	// user followers
	userFollowersRepository := userfollower.NewRepository(db)
	userFollowersService := userfollower.NewService(userFollowersRepository)
	userFollowersHandler := handler.NewUserFollowerHandler(userFollowersService)

	v1.GET("/user/followers/:user", userFollowersHandler.GetJoinUserToUserFollowing)
	v1.GET("/user/followers/e/:email", userFollowersHandler.GetByEmail)
	v1.POST("/user/a/followers", userFollowersHandler.PostUserFollowerHandler)

	// LIKE
	likeRepository := like.NewRepository(db)
	likeService := like.NewService(likeRepository)
	likeHandler := handler.NewLikeHandler(likeService)

	v1.POST("/post/like", likeHandler.PostLikeHandler)
	v1.GET("/post/like/:id/:email_user", likeHandler.GetLikeByUser)

	r.Run(":3000")
}
