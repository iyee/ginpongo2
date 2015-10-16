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
		templateValue, isString := templateName.(string)

		if exist && isString {
			var template *pongo2.Template
			if gin.Mode() == "debug" {
				template = pongo2.Must(pongo2.FromFile(templateValue))
			} else {
				template = pongo2.Must(pongo2.FromCache(templateValue))
			}

			err := template.ExecuteWriter(getContext(c.Get("data")), c.Writer)
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
