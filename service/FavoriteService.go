package service

// Favorite
//业务需求：注意看官方user中，有total_favorite(当前用户获赞数量），favorite_count（该用户点赞数量)
//意味着，当前用户点赞的时候：视频作者获赞数量++、自己点赞数量++
func Favorite(videoIdInt64 int64, userIdInt64 int64, actionType int32) (err error) {

	return nil
}
