package service

import (
	"context"
	"errors"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type PostServiceImpl struct {
	repository.PostRepository
	repository.CommentRepository
	repository.RewardRepository
	CommentService
	UserService
}

func (s *PostServiceImpl) FindByTopic(topicName string) []web.PostBySearchResponse {
	posts, err := s.PostRepository.FindByTopic(topicName)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	postsResponse := []web.PostBySearchResponse{}
	for _, post := range posts {
		postResponse := helper.ToPostBySearchResponse(post)
		post.CommentCount, _ = s.PostRepository.GetTotalCommentByPostId(post.Id)
		postsResponse = append(postsResponse, postResponse)
	}

	return postsResponse
}

func (s *PostServiceImpl) FindByUserId(userId int) []web.PostBySearchResponse {
	posts, err := s.PostRepository.FindByUserId(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	postsResponse := []web.PostBySearchResponse{}
	for _, post := range posts {
		postResponse := helper.ToPostBySearchResponse(post)
		post.CommentCount, _ = s.PostRepository.GetTotalCommentByPostId(post.Id)
		postsResponse = append(postsResponse, postResponse)
	}

	return postsResponse
}

func (s *PostServiceImpl) FindAll() []web.PostBySearchResponse {
	posts, err := s.PostRepository.FindAll()
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	postsResponse := []web.PostBySearchResponse{}
	for _, post := range posts {
		if post.CommentReply {
			comments, err := s.CommentRepository.FindByPostId(post.Id)
			if err != nil {
				panic(helper.NewNotFoundError(err.Error()))
			}
			post.Comments = comments
		}
		postResponse := helper.ToPostBySearchResponse(post)
		post.CommentCount, _ = s.PostRepository.GetTotalCommentByPostId(post.Id)
		postsResponse = append(postsResponse, postResponse)
	}
	return postsResponse
}

func (s *PostServiceImpl) FindByKeyword(query string) ([]web.PostBySearchResponse, error) {
	if query == "" {
		return nil, errors.New("user did not submit a valid query")
	}

	query = strings.ToLower(strings.TrimSpace(query))
	results := []domain.Post{}

	if !strings.Contains(query, " ") {
		posts, err := s.PostRepository.FindByKeywords(query, 20)
		if err != nil {
			return nil, err
		}
		results = append(results, posts...)
	} else {
		splitQuery := strings.Fields(query)
		for _, keyword := range splitQuery {
			posts, err := s.PostRepository.FindByKeywords(keyword, 20)
			if err != nil {
				return nil, err
			}
			results = append(results, posts...)
		}
	}

	postsResponse := []web.PostBySearchResponse{}
	for _, post := range results {
		postResponse := helper.ToPostBySearchResponse(post)
		post.CommentCount, _ = s.PostRepository.GetTotalCommentByPostId(post.Id)
		postsResponse = append(postsResponse, postResponse)
	}

	if len(results) == 0 {
		return nil, errors.New("post not found ")
	}

	return postsResponse, nil
}

func (s *PostServiceImpl) FindById(postId int) web.PostResponse {
	findById, err := s.PostRepository.FindById(postId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	postResponse := helper.ToPostResponse(findById)
	if postResponse.CommentReply {
		postResponse.Comments = s.CommentService.FindByPostId(postResponse.Id)
	}
	postResponse.CommentCount, _ = s.PostRepository.GetTotalCommentByPostId(postResponse.Id)
	return postResponse
}

func (s *PostServiceImpl) Create(request web.PostCreateInput) web.PostBySearchResponse {
	post := domain.Post{
		UserId:      request.UserId,
		Title:       request.Title,
		Slug:        request.Slug,
		Description: request.Description,
		Content:     request.Content,
		Topic:       request.Topic,
		Keyworks:    request.Keyworks,
	}

	post.CommentReply = false

	if post.UserId == 0 {
		panic(helper.NewUnauthorizedError("You're not an user"))
	}

	user := s.UserService.FindById(post.UserId)
	post.ProfileImageAlt = user.Avatar
	post.ProfileImageName = user.LastName

	save := s.PostRepository.Save(post)

	auth := helper.AuthGenerator(helper.Client)
	add, err := helper.Manage.AddPost(auth, big.NewInt(int64(save.Id)), common.HexToAddress(user.Address))
	if err != nil {
		helper.PanicIfError(err)
	}
	log.Printf("add successfull", add)

	getPost, err := helper.Manage.GetPosts(&bind.CallOpts{}, big.NewInt(int64(save.Id)))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("getPost", getPost)

	return helper.ToPostBySearchResponse(save)
}

func (s *PostServiceImpl) LikePost(ctx context.Context, userId int, postId int) (int, bool, error) {
	didUserLike, err := helper.RedisCli.SIsMember(ctx, "user:"+strconv.Itoa(userId)+":liked_posts", strconv.Itoa(postId)).Result()
	if err != nil {
		return 0, false, err
	}

	findById, err := s.PostRepository.FindById(postId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if didUserLike {
		if err := helper.RedisCli.SRem(ctx, "user:"+strconv.Itoa(userId)+":liked_posts", strconv.Itoa(postId)).Err(); err != nil {
			return 0, false, err
		}
		findById.Likes--
	} else {
		if err := helper.RedisCli.SAdd(ctx, "user:"+strconv.Itoa(userId)+":liked_posts", strconv.Itoa(postId)).Err(); err != nil {
			return 0, false, err
		}
		findById.Likes++
	}

	// Cập nhật thông tin bài viết vào cơ sở dữ liệu
	postUpdate := s.PostRepository.Update(findById)

	// if didUserLike && (post.Likes == 5 || post.Likes == 10 || post.Likes == 25) {
	// 	notification := models.Notification{
	// 		Message: "Your post, " + post.Title + ", has " + strconv.Itoa(post.Likes) + " likes",
	// 		IsReply: false,
	// 	}
	// 	postCreator, err := s.UserRepo.FindByID(post.User)
	// 	if err != nil {
	// 		return 0, false, err
	// 	}
	// 	postCreator.Notifications = append(postCreator.Notifications, notification)
	// 	if err := s.UserRepo.Update(postCreator); err != nil {
	// 		return 0, false, err
	// 	}
	// }

	return postUpdate.Likes, !didUserLike, nil
}

func (s *PostServiceImpl) DisLikePost(ctx context.Context, userId int, postId int) (int, bool, error) {
	didUserDislike, err := helper.RedisCli.SIsMember(ctx, "user:"+strconv.Itoa(userId)+":dis_likes", strconv.Itoa(postId)).Result()
	if err != nil {
		return 0, false, err
	}

	findById, err := s.PostRepository.FindById(postId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if didUserDislike {
		if err := helper.RedisCli.SRem(ctx, "user:"+strconv.Itoa(userId)+":dis_likes", strconv.Itoa(postId)).Err(); err != nil {
			return 0, false, err
		}
		findById.Dislikes--
	} else {
		if err := helper.RedisCli.SAdd(ctx, "user:"+strconv.Itoa(userId)+":dis_likes", strconv.Itoa(postId)).Err(); err != nil {
			return 0, false, err
		}
		findById.Dislikes++
	}
	postUpdate := s.PostRepository.Update(findById)

	return postUpdate.Dislikes, !didUserDislike, nil
}

func CalculatePoints(likes, dislikes, commCount int) int {
	return likes - dislikes + 3*commCount
}

func (s *PostServiceImpl) RewardPost(postId int, point int) {
	eduManageAddress := common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")
	balance1, err := helper.Token.BalanceOf(&bind.CallOpts{}, eduManageAddress)
	if err != nil {
		log.Fatal(err)
	}

	auth := helper.AuthGenerator(helper.Client)
	add, err := helper.Manage.CheckAndTransferRewardPost(auth, big.NewInt(int64(postId)), big.NewInt(int64(point)))
	if err != nil {
		helper.PanicIfError(err)
	}
	log.Printf("post, reward", add)

	balance2, err := helper.Token.BalanceOf(&bind.CallOpts{}, eduManageAddress)
	if err != nil {
		log.Fatal(err)
	}

	post, err := s.PostRepository.FindById(postId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	txHash := add.Hash().Hex()
	rewardDetail := domain.RewardHistory{}
	rewardDetail.RewardAddress = txHash
	rewardDetail.UserId = post.UserId
	rewardDetail.RewardType = 1
	balance1Int64 := balance1.Int64()
	balance2Int64 := balance2.Int64()
	rewardDetail.CountReward = (int(balance1Int64) - int(balance2Int64)) / 100000000
	rewardDetail.RewardAt = time.Now()

	save := s.RewardRepository.Save(rewardDetail)
	log.Printf("post, reward", save)

}

// ProcessPosts processes all posts for points calculation and rewards
func (s *PostServiceImpl) ProcessPosts() bool {
	posts, err := s.PostRepository.FindAll()
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	for _, post := range posts {
		post.CommentCount, _ = s.PostRepository.GetTotalCommentByPostId(post.Id)
		post.Points = CalculatePoints(post.Likes, post.Dislikes, int(post.CommentCount))
		s.PostRepository.Update(post)
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("post not enough to reward %d: %v\n", post.Id, r)
				}
			}()
			s.RewardPost(post.Id, post.Points)
		}()
	}
	return true
}

func NewPostService(postRepository repository.PostRepository, commentRepository repository.CommentRepository, commentService CommentService, userService UserService, rewardRepository repository.RewardRepository) PostService {
	return &PostServiceImpl{PostRepository: postRepository, CommentRepository: commentRepository, CommentService: commentService, UserService: userService, RewardRepository: rewardRepository}
}
