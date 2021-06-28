package middleware

import (
	"github.com/gin-gonic/gin"
)

func PageNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == 404 {
			_, _ = c.Writer.Write([]byte(`
				<!DOCTYPE html>
				<html lang="en">
				
				<head>
				  <meta charset="UTF-8" />
				  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
				  <meta http-equiv="X-UA-Compatible" content="ie=edge" />
				  <title>Susu</title>
				  <style>
					body {
					  position: absolute;
					  top: 50%;
					  left: 50%;
					  transform: translate(-50%, -50%);
					  text-align: center
					}
				
					h2 {
					  font-size: 150px;
					  display: block;
					  text-align: center;
					  font-family: Calibri, 'Helvetica Neue', Helvetica, Arial, Verdana, sans-serif;
					  margin-bottom: 20px;
					}
				
					h3 {
					  font-size: 75px;
					  display: block;
					  text-align: center;
					  font-family: Calibri, 'Helvetica Neue', Helvetica, Arial, Verdana, sans-serif;
					}
				  </style>
				</head>
				
				<body>
				  <h2>404</h2>
				  <h3>Page Not Found</h3>
				</body>
				
				</html>`))
		}
	}
}
