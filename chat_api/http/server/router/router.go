package route

import (
	"chat/ctx"
	"chat/model"
	"chat/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	keyApplicationId = "applicationId"
	keyChatNumber    = "chatNumber"
)

type Router struct {
	serviceContext    ctx.ServiceContext
	chatSystemService *service.ChatSystemService
}

func NewRouter(serviceContext ctx.ServiceContext) *Router {
	return &Router{
		serviceContext:    serviceContext,
		chatSystemService: service.NewChatSystemService(serviceContext),
	}
}

func (r *Router) Install(engine *gin.Engine) {
	prefix := "/api/v1/chat"
	router := engine.Group(prefix)
	router.POST("/application", r.createApplication)
	router.PUT(fmt.Sprintf("/application/:%s", keyApplicationId), r.updateApplication)
	router.GET(fmt.Sprintf("/application/:%s", keyApplicationId), r.getApplication)
	router.POST(fmt.Sprintf("/chat/:%s", keyApplicationId), r.createChat)
	router.GET(fmt.Sprintf("/chat/:%s", keyApplicationId), r.readChats)
	router.POST(fmt.Sprintf("/message/:%s/:%s", keyApplicationId, keyChatNumber), r.createMessage)
	router.GET(fmt.Sprintf("/message/:%s/:%s", keyApplicationId, keyChatNumber), r.getMessages)
	router.GET(fmt.Sprintf("/message/search/:%s/:%s", keyApplicationId, keyChatNumber), r.searchMessages)
}

func (r *Router) createApplication(c *gin.Context) {
	req := new(model.CreateApplicationRequest)
	err := c.BindJSON(req)
	if err != nil {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	res, err := r.chatSystemService.CreateApplication(req)
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)
}

func (r *Router) updateApplication(c *gin.Context) {
	applicationId := c.Param(keyApplicationId)

	if applicationId == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	req := new(model.UpdateApplicationRequest)
	err := c.BindJSON(req)
	if err != nil {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	res, err := r.chatSystemService.UpdateApplication(req, applicationId)
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)
}

func (r *Router) getApplication(c *gin.Context) {
	applicationId := c.Param(keyApplicationId)

	if applicationId == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	res, err := r.chatSystemService.GetApplication(applicationId)
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)
}

func (r *Router) createChat(c *gin.Context) {
	applicationId := c.Param(keyApplicationId)

	if applicationId == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	res, err := r.chatSystemService.CreateChat(applicationId)
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)

}

func (r *Router) readChats(c *gin.Context) {
	applicationId := c.Param(keyApplicationId)

	if applicationId == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	res, err := r.chatSystemService.ReadChats(applicationId)
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)
}

func (r *Router) createMessage(c *gin.Context) {
	applicationId := c.Param(keyApplicationId)

	if applicationId == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	chatNumber := c.Param(keyChatNumber)

	if chatNumber == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	number, err := strconv.ParseInt(chatNumber, 10, 32)

	res, err := r.chatSystemService.CreateMessage(applicationId, int32(number))
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)
}

func (r *Router) getMessages(c *gin.Context) {
	applicationId := c.Param(keyApplicationId)

	if applicationId == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	chatNumber := c.Param(keyChatNumber)

	if chatNumber == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	number, err := strconv.ParseInt(chatNumber, 10, 32)

	if err != nil {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	res, err := r.chatSystemService.GetMessages(applicationId, int32(number))
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)

}

func (r *Router) searchMessages(c *gin.Context) {

	applicationId := c.Param(keyApplicationId)

	if applicationId == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	chatNumber := c.Param(keyChatNumber)

	if chatNumber == "" {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	number, err := strconv.ParseInt(chatNumber, 10, 32)

	if err != nil {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	req := new(model.SearchMessagesRequest)
	err = c.BindJSON(req)
	if err != nil {
		r.newBadRequestResponse(c, model.InvalidRequest)
		return
	}

	res, err := r.chatSystemService.SearchMessages(applicationId, int32(number), req)
	if err != nil {
		r.newBadRequestResponse(c, err.Error())
		return
	}

	r.newSuccessResponse(c, res)

}
