package service

import (
	"context"
	"fmt"
	"go-pzn-restful-api/helper"
	"go-pzn-restful-api/model/domain"
	"go-pzn-restful-api/model/web"
	"go-pzn-restful-api/repository"
	"strconv"
)

type CommentServiceImpl struct {
	repository.CommentRepository
	repository.PostRepository
	UserService
}

func (s *CommentServiceImpl) FindByPostId(postId int) []web.CommentResponse {
	comments, err := s.CommentRepository.FindByPostId(postId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	var topLevelComments []domain.Comment
	var childCommentMap = make(map[int][]domain.Comment)

	for _, comment := range comments {
		if comment.CommentFatherId == 0 {
			topLevelComments = append(topLevelComments, comment)
		} else {
			childCommentMap[comment.CommentFatherId] = append(childCommentMap[comment.CommentFatherId], comment)
		}
	}

	var commentsResponse []web.CommentResponse
	for _, comment := range topLevelComments {
		commentResponse := helper.ToCommentResponse(comment)
		if comment.CommentReply {
			commentResponse.CommentChilds = s.buildCommentTree(comment.Id, childCommentMap)
		}
		commentResponse.CommentCount, _ = s.CommentRepository.CountCommentsByFatherId(commentResponse.Id)
		commentsResponse = append(commentsResponse, commentResponse)
	}

	return commentsResponse
}

func (s *CommentServiceImpl) buildCommentTree(commentId int, childCommentMap map[int][]domain.Comment) []web.CommentResponse {
	var childCommentsResponse []web.CommentResponse

	childComments, exists := childCommentMap[commentId]
	if !exists {
		return childCommentsResponse
	}

	for _, childComment := range childComments {
		childCommentResponse := helper.ToCommentResponse(childComment)
		if childComment.CommentReply {
			childCommentResponse.CommentChilds = s.buildCommentTree(childComment.Id, childCommentMap)
		}
		childCommentResponse.CommentCount, _ = s.CommentRepository.CountCommentsByFatherId(childCommentResponse.Id)
		childCommentsResponse = append(childCommentsResponse, childCommentResponse)
	}

	return childCommentsResponse
}

func (s *CommentServiceImpl) FindByUserId(userId int) []web.CommentResponse {
	comments, err := s.CommentRepository.FindByUserId(userId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	commentsResponse := []web.CommentResponse{}
	for _, comment := range comments {
		commentResponse := helper.ToCommentResponse(comment)
		commentsResponse = append(commentsResponse, commentResponse)
	}

	return commentsResponse
}

func (s *CommentServiceImpl) FindByComId(comFarId int) []web.CommentResponse {
	childComments, err := s.CommentRepository.FindByCommentFartherId(comFarId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	var resultComments []web.CommentResponse
	// Lặp qua từng comment con để lấy tất cả các comment con của chúng
	for _, childComment := range childComments {
		commentResponse := helper.ToCommentResponse(childComment)
		resultComments = append(resultComments, commentResponse)
		if childComment.CommentReply {

			// Đệ quy để lấy tất cả các comment con của comment con
			subChildComments := s.FindByComId(childComment.Id)
			// Thêm các comment con vào slice kết quả
			resultComments = append(resultComments, subChildComments...)
		}
	}

	return resultComments
}

func (s *CommentServiceImpl) FindById(commentId int) web.CommentResponse {
	comment, err := s.CommentRepository.FindById(commentId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	comment.CommentCount, _ = s.CommentRepository.CountCommentsByFatherId(comment.Id)
	if comment.CommentReply {
		comments, err := s.CommentRepository.FindByCommentFartherId(comment.Id)
		if err != nil {
			panic(helper.NewNotFoundError(err.Error()))
		}
		comment.CommentChilds = comments
	}
	return helper.ToCommentResponse(comment)
}

func (s *CommentServiceImpl) Create(request web.CommentCreateInput) web.CommentResponse {
	comment := domain.Comment{
		PostId:          request.PostId,
		CommentFatherId: request.CommentFatherId,
		UserId:          request.UserId,
		Content:         request.Content,
	}

	if comment.UserId == 0 {
		panic(helper.NewUnauthorizedError("You're not an user"))
	}

	user := s.UserService.FindById(comment.UserId)
	comment.ProfileImageAlt = user.Avatar
	comment.ProfileImageName = user.LastName

	post, err := s.PostRepository.FindById(comment.PostId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	post.CommentReply = true
	s.PostRepository.Update(post)

	if comment.CommentFatherId != 0 {
		err := s.CommentRepository.CheckCommentExists(comment.PostId, comment.CommentFatherId)
		// if err != nil {
		// 	panic(helper.NewNotFoundError(err.Error()))
		// }
		if err != nil {
			panic(helper.NewNotFoundError(err.Error()))
		}
		commentFather, err := s.CommentRepository.FindById(comment.CommentFatherId)
		if err != nil {
			panic(helper.NewNotFoundError(err.Error()))
		}
		if !commentFather.CommentReply {
			commentFather.CommentReply = true
			s.CommentRepository.Update(commentFather)
		}
	}

	save := s.CommentRepository.Save(comment)

	return helper.ToCommentResponse(save)
}

func (s *CommentServiceImpl) LikeComment(ctx context.Context, userId int, commentId int) (int, bool, error) {
	didUserLike, err := helper.RedisCli.SIsMember(ctx, "user:"+strconv.Itoa(userId)+":liked_comments", strconv.Itoa(commentId)).Result()
	if err != nil {
		return 0, false, err
	}

	findById, err := s.CommentRepository.FindById(commentId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if didUserLike {
		if err := helper.RedisCli.SRem(ctx, "user:"+strconv.Itoa(userId)+":liked_comments", strconv.Itoa(commentId)).Err(); err != nil {
			return 0, false, err
		}
		findById.Likes--
	} else {
		if err := helper.RedisCli.SAdd(ctx, "user:"+strconv.Itoa(userId)+":liked_comments", strconv.Itoa(commentId)).Err(); err != nil {
			return 0, false, err
		}
		findById.Likes++
	}
	commentUpdate := s.CommentRepository.Update(findById)

	return commentUpdate.Likes, didUserLike, nil
}

func (s *CommentServiceImpl) DisLikeComment(ctx context.Context, userId int, commentId int) (int, bool, error) {
	didUserDislike, err := helper.RedisCli.SIsMember(ctx, "user:"+strconv.Itoa(userId)+":dis_likes_comment", strconv.Itoa(commentId)).Result()
	if err != nil {
		return 0, false, err
	}

	findById, err := s.CommentRepository.FindById(commentId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	if didUserDislike {
		if err := helper.RedisCli.SRem(ctx, "user:"+strconv.Itoa(userId)+":dis_likes_comment", strconv.Itoa(commentId)).Err(); err != nil {
			return 0, false, err
		}
		findById.Dislikes--
	} else {
		if err := helper.RedisCli.SAdd(ctx, "user:"+strconv.Itoa(userId)+":dis_likes_comment", strconv.Itoa(commentId)).Err(); err != nil {
			return 0, false, err
		}
		findById.Dislikes++
	}
	commentUpdate := s.CommentRepository.Update(findById)

	return commentUpdate.Dislikes, !didUserDislike, nil
}

func RewardComment(commentId int) {
	fmt.Printf("Rewarding post with ID %d...\n", commentId)
}

// ProcessPosts processes all posts for points calculation and rewards
func (s *CommentServiceImpl) ProcessComments() bool {
	comments, err := s.CommentRepository.FindAll()
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	for _, comment := range comments {
		comment.CommentCount, _ = s.CommentRepository.CountCommentsByFatherId(comment.Id)
		comment.Points = CalculatePoints(comment.Likes, comment.Dislikes, comment.CommentCount)
		if comment.Points > 1000 {
			RewardPost(comment.Id)
		}
		s.CommentRepository.Update(comment)
	}
	return true
}

func NewCommentService(commentRepository repository.CommentRepository, postRepository repository.PostRepository, userService UserService) CommentService {
	return &CommentServiceImpl{CommentRepository: commentRepository, PostRepository: postRepository, UserService: userService}
}
