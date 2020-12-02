package article

// 定义表数据传输类
type Article struct {
	// 文章ID
	Id int64
	// 文章标题
	Title string
	// 文章内容
	Content string
	// 文章创建时间，使用日期格式
	Created string
	// 文章更新时间，使用日期格式
	Updated string
}