package main_page

type Output struct {
	Title title  `json:"title"`
	Video string `json:"video"`
	Info  info   `json:"info"`
}

type title struct {
	FirstPic  string `json:"first_pic"`
	SecondPic string `json:"second_pic"`
	ThirdPic  string `json:"third_pic"`
}

type info struct {
	FirstInfo  string `json:"first_info"`
	SecondInfo string `json:"second_info"`
	ThirdInfo  string `json:"third_info"`
}

// 返回主頁需要的資訊
func Exec() (Output, error) {

	var output Output

	output.Title.FirstPic = "OoBzmPm.jpg"
	output.Title.SecondPic = "p30babL.png"
	output.Title.ThirdPic = "3z1eQBs.jpg"

	output.Video = "https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_1mb.mp4"

	output.Info.FirstInfo = "大拍賣開始喔！趕快來！"
	output.Info.SecondInfo = "年度大減價，全館 3 折起！"
	output.Info.ThirdInfo = "最新鮮食材 / 3C / 家電 / 生活用品"

	return output, nil
}
