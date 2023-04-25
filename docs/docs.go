// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
	"openapi": "3.0.0",
	"info": {
	  "title": "Resful Api",
	  "description": "Programming Golang",
	  "version": "1.0.11"
	},
	"tags": [
	  {
		"name": "user",
		"description": "Operations about user"
	  },
	  {
		"name": "post",
		"description": "Operations about post"
	  }
	],
	"paths": {
	  "/user/register": {
		"post": {
		  "tags": [
			"user"
		  ],
		  "summary": "Register user",
		  "description": "This can only be done by the logged in user.",
		  "operationId": "createUser",
		  "requestBody": {
			"description": "Created user object",
			"content": {
			  "application/json": {
				"schema": {
				  "$ref": "#/components/schemas/User"
				}
			  }
			}
		  },
		  "responses": {
			"default": {
			  "description": "successful operation",
			  "content": {
				"application/json": {
				  "schema": {
					"$ref": "#/components/schemas/User"
				  }
				}
			  }
			}
		  }
		}
	  },
	  "/user/login": {
		"post": {
		  "tags": [
			"user"
		  ],
		  "summary": "Login user into the system",
		  "description": "",
		  "operationId": "loginUser",
		  "requestBody": {
			"description": "Login user object",
			"content": {
			  "application/json": {
				"schema": {
				  "$ref": "#/components/schemas/Login"
				}
			  }
			}
		  },
		  "responses": {
			"default": {
			  "description": "successful operation",
			  "content": {
				"application/json": {
				  "schema": {
					"$ref": "#/components/schemas/LoginResponse"
				  }
				}
			  }
			}
		  }
		}
	  },
	  "/user/edit": {
		"put": {
		  "tags": [
			"user"
		  ],
		  "summary": "Edit name, password",
		  "description": "Edit user",
		  "security": [
			{
			  "basicAuth": []
			}
		  ],
		  "operationId": "editUser",
		  "requestBody": {
			"description": "Login user object",
			"content": {
			  "application/json": {
				"schema": {
				  "$ref": "#/components/schemas/EditUser"
				}
			  }
			}
		  },
		  "responses": {
			"default": {
			  "description": "successful operation"
			}
		  }
		}
	  },
	  "/user/me": {
		"get": {
		  "tags": [
			"user"
		  ],
		  "summary": "get user info",
		  "description": "User Info",
		  "security": [
			{
			  "basicAuth": []
			}
		  ],
		  "operationId": "userInfo",
		  "responses": {
			"default": {
			  "description": "successful operation"
			}
		  }
		}
	  },
	  "/post/create": {
		"post": {
		  "tags": [
			"post"
		  ],
		  "summary": "Create post",
		  "description": "",
		  "operationId": "createPost",
		  "requestBody": {
			"description": "Created post object",
			"content": {
			  "application/json": {
				"schema": {
				  "$ref": "#/components/schemas/Post"
				}
			  }
			}
		  },
		  "responses": {
			"default": {
			  "description": "successful operation"
			}
		  }
		}
	  },
	  "/post/edit/{id}": {
		"put": {
		  "tags": [
			"post"
		  ],
		  "summary": "Edit post",
		  "description": "",
		  "operationId": "editPost",
		  "parameters": [
			{
			  "name": "id",
			  "in": "path",
			  "description": "Post ID",
			  "required": true,
			  "schema": {
				"type": "integer",
				"format": "int64"
			  }
			}
		  ],
		  "requestBody": {
			"description": "Edit post object",
			"content": {
			  "application/json": {
				"schema": {
				  "$ref": "#/components/schemas/Post"
				}
			  }
			}
		  },
		  "responses": {
			"default": {
			  "description": "successful operation"
			}
		  }
		}
	  },
	  "/post/detail/{id}": {
		"get": {
		  "tags": [
			"post"
		  ],
		  "summary": "Detail post",
		  "description": "",
		  "operationId": "detailPost",
		  "parameters": [
			{
			  "name": "id",
			  "in": "path",
			  "description": "Post ID",
			  "required": true,
			  "schema": {
				"type": "integer",
				"format": "int64"
			  }
			}
		  ],
		  "responses": {
			"default": {
			  "description": "successful operation"
			}
		  }
		}
	  },
	  "/post/list": {
		"get": {
		  "tags": [
			"post"
		  ],
		  "summary": "List post",
		  "description": "",
		  "operationId": "listPost",
		  "parameters": [
			{
			  "name": "page",
			  "in": "query",
			  "description": "",
			  "required": false,
			  "schema": {
				"type": "number"
			  }
			},
			{
			  "name": "size",
			  "in": "query",
			  "description": "",
			  "required": false,
			  "schema": {
				"type": "number"
			  }
			}
		  ],
		  "responses": {
			"default": {
			  "description": "successful operation"
			}
		  }
		}
	  },
	  "/post/delete/{id}": {
		"delete": {
		  "tags": [
			"post"
		  ],
		  "summary": "Delete post",
		  "description": "",
		  "operationId": "deletePost",
		  "parameters": [
			{
			  "name": "id",
			  "in": "path",
			  "description": "Post ID",
			  "required": true,
			  "schema": {
				"type": "integer",
				"format": "int64"
			  }
			}
		  ],
		  "responses": {
			"default": {
			  "description": "successful operation"
			}
		  }
		}
	  }
	},
	"components": {
	  "schemas": {
		"User": {
		  "type": "object",
		  "properties": {
			"name": {
			  "type": "string",
			  "example": "john"
			},
			"email": {
			  "type": "string",
			  "example": "john@email.com"
			},
			"password": {
			  "type": "string",
			  "example": "123456"
			}
		  },
		  "xml": {
			"name": "user"
		  }
		},
		"Login": {
		  "type": "object",
		  "properties": {
			"email": {
			  "type": "string"
			},
			"password": {
			  "type": "string"
			}
		  },
		  "xml": {
			"name": "Login"
		  }
		},
		"EditUser": {
		  "type": "object",
		  "properties": {
			"name": {
			  "type": "string"
			},
			"password": {
			  "type": "string"
			}
		  },
		  "xml": {
			"name": "EditUser"
		  }
		},
		"LoginResponse": {
		  "type": "object",
		  "properties": {
			"accessToken": {
			  "type": "string"
			},
			"ttl": {
			  "type": "number"
			},
			"user": {
			  "type": "object",
			  "properties": {
				"id": {
				  "type": "number",
				  "example": 1
				},
				"name": {
				  "type": "string",
				  "example": "john"
				},
				"email": {
				  "type": "string",
				  "example": "john@email.com"
				}
			  }
			}
		  },
		  "xml": {
			"name": "LoginResponse"
		  }
		},
		"Post": {
		  "type": "object",
		  "properties": {
			"title": {
			  "type": "string"
			},
			"description": {
			  "type": "string"
			},
			"content": {
			  "type": "string"
			}
		  },
		  "xml": {
			"name": "user"
		  }
		}
	  },
	  "securitySchemes": {
		"basicAuth": {
		  "type": "apiKey",
		  "name": "authorization",
		  "in": "header"
		}
	  }
	}
  }`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}