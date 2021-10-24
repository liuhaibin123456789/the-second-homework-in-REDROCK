package main

import "fmt"

type Video struct {
	VideoName string//视频名字
	VideoSize string//视频大小
	VideoPraiseNUmber int64//视频点赞数量
	VideoCollectNUmber int64//视频收藏数量
	VideoCoinNumber int64//视频投币数量
}
type VideoContent struct {
	VideoBasicMSG Video//视频基本信息
	VideoIntroduction string//视频介绍
	VideoBulletScreens string//视频弹幕
}
type VideoRecommend struct {
	VideoBasicMSG Video//视频基本信息
	VideoPage string//视频图片
}
type Author struct {
	AuthorName string//作者姓名
	AuthorFansNumber int64//作者粉丝数量
	AuthorIntroduction string//作者介绍
	AuthorIcon string//作者头像
}
type Comment struct {
	CommentMaker string//评论者姓名
	CommentContent string//评论内容
	CommentDialog string//评论对话
}
//创建B站视频详情页得结构体

type VideoIfo struct {
	VideoContent VideoContent
	VideoRecommend VideoRecommend
	Author Author
	Comment Comment
}
//点赞
func (V *VideoIfo) praise()  {
	V.VideoContent.VideoBasicMSG.VideoPraiseNUmber++
}
//收藏
func (V *VideoIfo) collect() {
	V.VideoContent.VideoBasicMSG.VideoCollectNUmber++
}
//投币
func (V *VideoIfo) putCoin()  {
	V.VideoContent.VideoBasicMSG.VideoCoinNumber++
}
//一键三连
func (V *VideoIfo) sanLian()  {
	V.Author.AuthorName="wwwwwwwwww"
	V.Author.AuthorFansNumber+=3
	V.VideoContent.VideoIntroduction="0000000000000______"
	V.VideoContent.VideoBasicMSG.VideoName="MMMMMM"
	V.VideoContent.VideoBasicMSG.VideoSize="120120"
	V.VideoContent.VideoBasicMSG.VideoCoinNumber++
	V.VideoContent.VideoBasicMSG.VideoPraiseNUmber++
	V.VideoContent.VideoBasicMSG.VideoCollectNUmber++

}
func main() {
	var MyVideoIfo = VideoIfo{
		Author: Author{
			"bwj",
			309991,
			"米纳赏~",
			"nil",
		},
		VideoRecommend: VideoRecommend{
			VideoPage: "2333",
			VideoBasicMSG: Video{
				VideoName: "如何煮泡面？",
				VideoSize: "120Mib",
			},
		},
		Comment: Comment{
			CommentMaker:   "lhb",
			CommentContent: "好好好！",
			CommentDialog:  "nil",
		},
		VideoContent: VideoContent{
			VideoBasicMSG: Video{
				VideoName:          "others",
				VideoSize:          "130Mib",
				VideoPraiseNUmber:  int64(0),
				VideoCoinNumber:    int64(0),
				VideoCollectNUmber: int64(0),
			},
			VideoIntroduction:  "this is a good play",
			VideoBulletScreens: "no no no no~~~",
		},
	}
	fmt.Println(MyVideoIfo)
	for i := 0; i < 10; i++ {
		(&MyVideoIfo).sanLian()
		fmt.Println(MyVideoIfo)
	}
}