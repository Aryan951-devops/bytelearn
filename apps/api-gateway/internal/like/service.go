package like

import "github.com/google/uuid"

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

func CheckVideoLikeService(
	userID,
	videoID uuid.UUID,
) (bool, error) {

	return IsVideoLiked(userID, videoID)
}

func CheckCommentLikeService(
	userID,
	commentID uuid.UUID,
) (bool, error) {

	return IsCommentLiked(userID, commentID)
}

func GetTotalVideoLikesService(
	videoID uuid.UUID,
) (int64, error) {

	return GetTotalVideoLikes(videoID)
}

func GetTotalCommentLikesService(
	commentID uuid.UUID,
) (int64, error) {

	return GetTotalCommentLikes(commentID)
}
