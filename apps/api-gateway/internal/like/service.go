package like

import "github.com/google/uuid"

// ToggleVideoLikeService adds or removes a video like.
func ToggleVideoLikeService(
	userID,
	videoID uuid.UUID,
) (bool, error) {

	liked, err := IsVideoLiked(userID, videoID)
	if err != nil {
		return false, err
	}

	if liked {
		err = UnlikeVideo(userID, videoID)
		return false, err
	}

	err = LikeVideo(userID, videoID)
	return true, err
}

// ToggleCommentLikeService adds or removes a comment like.
func ToggleCommentLikeService(
	userID,
	commentID uuid.UUID,
) (bool, error) {

	liked, err := IsCommentLiked(userID, commentID)
	if err != nil {
		return false, err
	}

	if liked {
		err = UnlikeComment(userID, commentID)
		return false, err
	}

	err = LikeComment(userID, commentID)
	return true, err
}

// CheckVideoLikeService verifies if a video is liked by a user.
func CheckVideoLikeService(
	userID,
	videoID uuid.UUID,
) (bool, error) {

	return IsVideoLiked(userID, videoID)
}

// CheckCommentLikeService verifies if a comment is liked by a user.
func CheckCommentLikeService(
	userID,
	commentID uuid.UUID,
) (bool, error) {

	return IsCommentLiked(userID, commentID)
}

// GetTotalVideoLikesService fetches the total like count for a video.
func GetTotalVideoLikesService(
	videoID uuid.UUID,
) (int64, error) {

	return GetTotalVideoLikes(videoID)
}

// GetTotalCommentLikesService fetches the total like count for a comment.
func GetTotalCommentLikesService(
	commentID uuid.UUID,
) (int64, error) {

	return GetTotalCommentLikes(commentID)
}
