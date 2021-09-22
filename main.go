package main
import(
	"github.com/gin-gonic/gin"
)
type myForm struct {
    Colors []string `form:"colors[]"`
}
func main()  {
	r := gin.Default()
	r.LoadHTMLGlob("views/*")
	//r.LoadHTMLFiles("html_Gin/index.html")
	r.GET("/",IndexHandler)
	r.POST("/",formHandler)
	r.Run()
}
func IndexHandler(c *gin.Context)  {
	c.HTML(200,"index.html",nil)
}
func formHandler(c *gin.Context) {
    var fakeForm myForm
    c.ShouldBind(&fakeForm)
    c.JSON(200, gin.H{"color": fakeForm.Colors})
}

