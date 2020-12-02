#本作业说明
此前未写过一行go代码。这是第一次编写。补习了下相关语法。这个项目从httpService启动，开启54100端口。访问方法是http://localhost:54110/article?id=1 访问1有数据。访问2将获取不到数据。
目录中有个sql语句。是创建本例中的表。标准的按照以下目录结构组织。真实模拟了sql.ErrNoRows的场景，想做的比较完整，但由于时间有限，家里有事，有些东西还没补充完整。后面我再完善的。
这种学习方法可以一下子知道许多东西。只是碍于时间不够用。

main.go
src
  article
        Article.go
        ArticleController.go
        ArticleDao.go
        ArticleService.go



#作业问题
Week02 作业题目：
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
思路是：
在使用第三方库的时候发生错误，我们正常是要wrap错误往上抛的。但是如果遇到的是sql.ErrNoRows。此为取不到数据的错误。我认为不应该往上抛。而是除此之外的错误再往上抛。
由于他并不算异常。严格点的做法可以抛该错误，也可以选择返回nil就好。
