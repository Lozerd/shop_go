package v1

import "github.com/gin-gonic/gin"

type User struct {
  ID   uint64 `json:"id" example:"100" description:"User identity"`
  Name string `json:"name" example:"Parvez"` 
}

type UsersResponse struct {
  Data []User `json:"users" example:"[{\"id\":100, \"name\":\"Parvez\"}]"`
}

type Error struct {
  Code string `json:"code"`
  Msg  string `json:"msg" skip:"true"`
  // use skip:"true" if you want to skip the field in the documentation spec
}

type ErrorResponse struct {
  ErrorInfo Error `json:"error"`
}

// RequestHeaders represents the model for header params
// @HeaderParameters RequestHeaders
type RequestHeaders struct {
    Authorization  string  `json:"Authorization"`
}

// @Title Get user list of a group.
// @Description Get users related to a specific group.
// @Param  groupID  path  int  true  "Id of a specific group." "120"
// @Success  200  object  UsersResponse  "UsersResponse JSON"
// @Failure  400  object  ErrorResponse  "ErrorResponse JSON"
// @Resource users
// @Route /api/group/{groupID}/users [get]
func Ping(ctx *gin.Context) {
	ctx.String(200, "ok")
}

func Routes(r *gin.RouterGroup) {
	products := r.Group("/products")
	products.GET("/ping", Ping)
}
