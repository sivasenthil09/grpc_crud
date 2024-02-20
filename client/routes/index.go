package routes

import ("github.com/gin-gonic/gin"
"crud_grpc/client/controler"
)

func Approutes(r *gin.Engine) {
	r.POST("/createuser", controler.CreateUser)
	r.GET("/listuser", controler.Listuser)
	r.GET("/get/:name", controler.GetUserByName)
	r.DELETE("/delete/:name", controler.DeleteUser)
}
