package article

func ArticleDetailService (id string) (code int,err error){
	articleContent , err:= GetDetail(id)
	if articleContent == nil {
		code = 404
	} else {
		code = 200;
	}
	return code , err
}