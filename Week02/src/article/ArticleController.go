package article

import(
	"encoding/json"
	"net/http"
	"fmt"
)

// A大写表示public
func Articlemethod() {
     http.HandleFunc("/article", ArticleDetailController )
}

func Articleok() {
	 http.HandleFunc("/article", ArticleDetailController )
}

func ArticleDetailController( w http.ResponseWriter,r *http.Request ) {
	// 接受文章ID,此处只是简单的接收ID 不做ID参数检验逻辑处理。
	articleID := r.FormValue("id")
	// 调用SERVICE
	res ,err:= ArticleDetailService(articleID)
	if err != nil {
		fmt.Printf("article not found ,id %s ,\n %+v\n",articleID,err)
	}
	b, _ := json.Marshal(res)
	w.Header().Set("Content-type","application/json;charset=utf-8");
	w.Write(b)
}