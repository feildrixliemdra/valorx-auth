package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"valorx-auth/internal/config"
	"valorx-auth/internal/payload"
	"valorx-auth/internal/service"
	"valorx-auth/internal/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type IAuthHandler interface {
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
}

type authHandler struct {
	cfg          *config.Config
	UserService  service.IUserService
	GoogleConfig *oauth2.Config
}

func NewAuthHandler(cfg *config.Config, userService service.IUserService) IAuthHandler {
	return &authHandler{
		cfg:         cfg,
		UserService: userService,
		GoogleConfig: &oauth2.Config{
			ClientID:     cfg.Auth.GoogleClientID,
			ClientSecret: cfg.Auth.GoogleClientSecret,
			RedirectURL:  cfg.Auth.GoogleClientCallbackURL,
			Scopes:       []string{"email", "profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (h *authHandler) GoogleLogin(c *gin.Context) {
	state := "state-token"
	url := h.GoogleConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *authHandler) GoogleCallback(c *gin.Context) {

	code := c.Query("code")
	state := c.Query("state")
	if state == "" || code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OAuth response"})
		return
	}

	token, err := h.GoogleConfig.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to exchange token"})
		return
	}

	userData, err := h.fetchGoogleUser(c, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch user token"})
		return
	}

	claims := map[string]interface{}{
		"user_id":  userData.ID,
		"provider": "GOOGLE",
		"email":    userData.Email,
		"exp":      time.Now().Add(time.Hour * 720).Unix(),
	}

	jwtToken, err := util.GenerateToken(claims, h.cfg.JWT.SecretKey)
	if err != nil {
		log.Errorf("Failed to generate token: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate jwt"})

		return
	}
	res := gin.H{
		"token":    jwtToken,
		"provider": claims["provider"],
		"expired":  claims["exp"],
	}

	util.GeneralSuccessResponse(c, "success login", res)
}

func (h *authHandler) fetchGoogleUser(ctx context.Context, token *oauth2.Token) (payload.GoogleUserData, error) {

	var res payload.GoogleUserData
	client := h.GoogleConfig.Client(ctx, token)
	resp, _ := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Errorf("Failed to decode google user data: %v", err)
		return res, err
	}

	return res, err
}
