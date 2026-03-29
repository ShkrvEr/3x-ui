package controller

import (
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
)

// XUIController is the main controller for the X-UI panel, managing sub-controllers.
type XUIController struct {
	BaseController

	settingController     *SettingController
	xraySettingController *XraySettingController
}

// NewXUIController creates a new XUIController and initializes its routes.
func NewXUIController(g *gin.RouterGroup) *XUIController {
	a := &XUIController{}
	a.initRouter(g)
	return a
}

// initRouter sets up the main panel routes and initializes sub-controllers.
func (a *XUIController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/panel")
	g.Use(a.checkLogin)

	g.GET("/", a.index)
	g.GET("/inbounds", a.inbounds)
	g.GET("/subeditor", a.subeditor)
	g.POST("/subeditor/load", a.loadSubEditorFile)  
    g.POST("/subeditor/save", a.saveSubEditorFile)
	g.GET("/settings", a.settings)
	g.GET("/xray", a.xraySettings)

	a.settingController = NewSettingController(g)
	a.xraySettingController = NewXraySettingController(g)
}

// index renders the main panel index page.
func (a *XUIController) index(c *gin.Context) {
	html(c, "index.html", "pages.index.title", nil)
}

//moded <!--
// inbounds renders the inbounds management page.
func (a *XUIController) subeditor(c *gin.Context) {  
    // Получаем путь из .env или используем значение по умолчанию  
    defaultPath := os.Getenv("EXTRA_KEYS_FILE_PATH")  
    if defaultPath == "" {  
        defaultPath = "/root/3x-ui/extra-keys.txt"  
    }  
      
    // Передаем путь в шаблон  
    data := map[string]interface{}{  
        "defaultFilePath": defaultPath,  
    }  
      
    html(c, "subeditor.html", "pages.subeditor.title", data)  
}

func (a *XUIController) loadSubEditorFile(c *gin.Context) {  
    path := c.PostForm("path")  
    if path == "" {  
        jsonMsg(c, "Invalid request", nil)  
        return  
    }  
      
    content, err := os.ReadFile(path)  
    if err != nil {  
        jsonMsg(c, fmt.Sprintf("Failed to read file: %v", err), nil)  
        return  
    }  
      
    jsonObj(c, map[string]interface{}{  
        "content": string(content),  
    }, nil)  
}  
  
func (a *XUIController) saveSubEditorFile(c *gin.Context) {  
    path := c.PostForm("path")  
    content := c.PostForm("content")  
      
    if path == "" {  
        jsonMsg(c, "Invalid request", nil)  
        return  
    }  
      
    err := os.WriteFile(path, []byte(content), 0644)  
    if err != nil {  
        jsonMsg(c, fmt.Sprintf("Failed to save file: %v", err), nil)  
        return  
    }  
      
    jsonMsg(c, "File saved successfully", nil)  
}
//--!> moded

func (a *XUIController) inbounds(c *gin.Context) {
	html(c, "inbounds.html", "pages.inbounds.title", nil)
}

// settings renders the settings management page.
func (a *XUIController) settings(c *gin.Context) {
	html(c, "settings.html", "pages.settings.title", nil)
}

// xraySettings renders the Xray settings page.
func (a *XUIController) xraySettings(c *gin.Context) {
	html(c, "xray.html", "pages.xray.title", nil)
}
