package ginpongo2

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func Pongo2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		templateName, exist := c.Get("template")
		templateNameValue, isString := templateName.(string)

		if exist && isString {
			templateData, exist := c.Get("data")
			var template = pongo2.Must(pongo2.FromCache(templateNameValue))
			err := template.ExecuteWriter(getContext(templateData, exist), c.Writer)
			if err != nil {
				http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func getContext(templateData interface{}, exist bool) pongo2.Context {
	if templateData == nil || !exist {
		return nil
	}
	contextData, isMap := templateData.(map[string]interface{})
	if isMap {
		return contextData
	}
	return nil
}
