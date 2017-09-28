// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json",
    "application/yaml"
  ],
  "produces": [
    "application/json",
    "application/yaml"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "An open platform to package and deploy applications into multiple environment such as QingCloud, AWS, Kubernetes etc.",
    "title": "AppHub",
    "version": "0.0.1"
  },
  "basePath": "/v1",
  "paths": {
    "/appruntimes": {
      "get": {
        "description": "Returns a list containing all app runtimes.",
        "tags": [
          "app-runtimes"
        ],
        "summary": "Gets some app runtimes",
        "parameters": [
          {
            "minimum": 0,
            "exclusiveMinimum": true,
            "multipleOf": 10,
            "type": "integer",
            "format": "int32",
            "default": 20,
            "description": "Number of clusters returned",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 1,
            "description": "Page number",
            "name": "pageNumber",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A list of app runtimes",
            "schema": {
              "$ref": "#/definitions/getAppruntimesOKBody"
            }
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Adds a new app runtime to the runtimes list.",
        "tags": [
          "app-runtimes"
        ],
        "summary": "Creates an app runtime",
        "parameters": [
          {
            "description": "The runtime to create.",
            "name": "appruntime",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/AppRuntime"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "App runtime succesfully created."
          },
          "400": {
            "description": "App runtime couldn't have been created."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/appruntimes/{appRuntimeId}": {
      "get": {
        "description": "Returns a single runtime by its id",
        "tags": [
          "app-runtimes"
        ],
        "summary": "Gets an app runtime",
        "parameters": [
          {
            "type": "string",
            "description": "The app runtime's id",
            "name": "appRuntimeId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "An app runtime",
            "schema": {
              "$ref": "#/definitions/AppRuntime"
            }
          },
          "404": {
            "description": "The app runtime does not exist."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a single app runtime identified via its id",
        "tags": [
          "app-runtimes"
        ],
        "summary": "Deletes an app runtime",
        "parameters": [
          {
            "type": "string",
            "description": "The app runtime's id",
            "name": "appRuntimeId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "App runtime successfully deleted."
          },
          "404": {
            "description": "The app runtime does not exist."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/apps": {
      "get": {
        "description": "Returns a list containing all apps.",
        "tags": [
          "apps"
        ],
        "summary": "Gets some apps",
        "parameters": [
          {
            "minimum": 0,
            "exclusiveMinimum": true,
            "multipleOf": 10,
            "type": "integer",
            "format": "int32",
            "default": 20,
            "description": "Number of clusters returned",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 1,
            "description": "Page number",
            "name": "pageNumber",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A list of apps",
            "schema": {
              "$ref": "#/definitions/getAppsOKBody"
            }
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Adds a new app to the app runtimes list.",
        "tags": [
          "apps"
        ],
        "summary": "Creates an app",
        "parameters": [
          {
            "description": "The app to create.",
            "name": "app",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/App"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "App succesfully created."
          },
          "400": {
            "description": "App couldn't have been created."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/apps/{appId}": {
      "get": {
        "description": "Returns a single app by its id",
        "tags": [
          "apps"
        ],
        "summary": "Gets an app",
        "parameters": [
          {
            "type": "string",
            "description": "The app's id",
            "name": "appId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "An app",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "The app does not exist."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a single app identified via its id",
        "tags": [
          "apps"
        ],
        "summary": "Deletes an app",
        "parameters": [
          {
            "type": "string",
            "description": "The app's id",
            "name": "appId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "App successfully deleted."
          },
          "404": {
            "description": "The app does not exist."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/clusters": {
      "get": {
        "description": "Returns a list containing all clusters.",
        "tags": [
          "clusters"
        ],
        "summary": "Gets some clusters",
        "parameters": [
          {
            "minimum": 0,
            "exclusiveMinimum": true,
            "multipleOf": 10,
            "type": "integer",
            "format": "int32",
            "default": 20,
            "description": "Number of clusters returned",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 1,
            "description": "Page number",
            "name": "pageNumber",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A list of Cluster",
            "schema": {
              "$ref": "#/definitions/getClustersOKBody"
            }
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Adds a new cluster to the clusters list.",
        "tags": [
          "clusters"
        ],
        "summary": "Creates a cluster",
        "parameters": [
          {
            "description": "The cluster to create.",
            "name": "cluster",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Cluster"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Cluster succesfully created."
          },
          "400": {
            "description": "Cluster couldn't have been created."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/clusters/{clusterId}": {
      "get": {
        "description": "Returns a single cluster by its id",
        "tags": [
          "clusters"
        ],
        "summary": "Gets a cluster",
        "parameters": [
          {
            "type": "string",
            "description": "The cluster's id",
            "name": "clusterId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A cluster",
            "schema": {
              "$ref": "#/definitions/Cluster"
            }
          },
          "404": {
            "description": "The cluster does not exist."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a single cluster identified via its id",
        "tags": [
          "clusters"
        ],
        "summary": "Deletes a cluster",
        "parameters": [
          {
            "type": "string",
            "description": "The cluster's id",
            "name": "clusterId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Cluster successfully deleted."
          },
          "404": {
            "description": "The cluster does not exist."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/repos": {
      "get": {
        "description": "Returns a list containing all repos.",
        "tags": [
          "repos"
        ],
        "summary": "Gets some repos",
        "parameters": [
          {
            "minimum": 0,
            "exclusiveMinimum": true,
            "multipleOf": 10,
            "type": "integer",
            "format": "int32",
            "default": 20,
            "description": "Number of clusters returned",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 1,
            "description": "Page number",
            "name": "pageNumber",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "A list of repos",
            "schema": {
              "$ref": "#/definitions/getReposOKBody"
            }
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Adds a new repo to the repos list.",
        "tags": [
          "repos"
        ],
        "summary": "Creates a repo",
        "parameters": [
          {
            "description": "The repo to create.",
            "name": "repo",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Repo"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Repo succesfully created."
          },
          "400": {
            "description": "Repo couldn't have been created."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/repos/{repoId}": {
      "get": {
        "description": "Returns a single repo by its id",
        "tags": [
          "repos"
        ],
        "summary": "Gets a repo",
        "parameters": [
          {
            "type": "string",
            "description": "The repo's id",
            "name": "repoId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A repo",
            "schema": {
              "$ref": "#/definitions/Repo"
            }
          },
          "404": {
            "description": "The repo does not exist."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a single repo identified via its id",
        "tags": [
          "repos"
        ],
        "summary": "Deletes a repo",
        "parameters": [
          {
            "type": "string",
            "description": "The repo's id",
            "name": "repoId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Repo successfully deleted."
          },
          "404": {
            "description": "The repo does not exist."
          },
          "500": {
            "description": "An unexpected error occured.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "App": {
      "required": [
        "appId"
      ],
      "properties": {
        "appId": {
          "type": "string",
          "maxLength": 12,
          "minLength": 12,
          "pattern": "app-[a-zA-Z0-9]{11}"
        },
        "createTime": {
          "type": "string",
          "format": "date-time",
          "readOnly": true,
          "example": "2017-09-17T12:36:58.014Z"
        },
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "AppRuntime": {
      "required": [
        "appRuntimeId"
      ],
      "properties": {
        "appRuntimeId": {
          "type": "string",
          "maxLength": 11,
          "minLength": 11,
          "pattern": "rt-[a-zA-Z0-9]{11}"
        },
        "createTime": {
          "type": "string",
          "format": "date-time",
          "readOnly": true,
          "example": "2017-09-17T12:36:58.014Z"
        },
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "AppRuntimes": {
      "properties": {
        "items": {
          "$ref": "#/definitions/appRuntimesItems"
        }
      }
    },
    "Apps": {
      "properties": {
        "items": {
          "$ref": "#/definitions/appsItems"
        }
      }
    },
    "Cluster": {
      "required": [
        "clusterId"
      ],
      "properties": {
        "clusterId": {
          "type": "string",
          "maxLength": 11,
          "minLength": 11,
          "pattern": "cl-[a-zA-Z0-9]{11}"
        },
        "createTime": {
          "type": "string",
          "format": "date-time",
          "readOnly": true,
          "example": "2017-09-17T12:36:58.014Z"
        },
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Clusters": {
      "properties": {
        "items": {
          "$ref": "#/definitions/clustersItems"
        }
      }
    },
    "Error": {
      "properties": {
        "code": {
          "type": "string",
          "enum": [
            "DBERR",
            "NTERR",
            "UNERR"
          ]
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Paging": {
      "properties": {
        "currentPage": {
          "type": "integer"
        },
        "pageSize": {
          "type": "integer"
        },
        "totalItems": {
          "type": "integer"
        },
        "totalPages": {
          "type": "integer"
        }
      }
    },
    "Repo": {
      "required": [
        "repoId"
      ],
      "properties": {
        "createTime": {
          "type": "string",
          "format": "date-time",
          "readOnly": true,
          "example": "2017-09-17T12:36:58.014Z"
        },
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "repoId": {
          "type": "string",
          "maxLength": 13,
          "minLength": 13,
          "pattern": "repo-[a-zA-Z0-9]{11}"
        },
        "url": {
          "type": "string"
        }
      }
    },
    "Repos": {
      "properties": {
        "items": {
          "$ref": "#/definitions/reposItems"
        }
      }
    },
    "appRuntimesItems": {
      "type": "array",
      "uniqueItems": true,
      "items": {
        "$ref": "#/definitions/AppRuntime"
      },
      "x-go-gen-location": "models"
    },
    "appsItems": {
      "type": "array",
      "uniqueItems": true,
      "items": {
        "$ref": "#/definitions/App"
      },
      "x-go-gen-location": "models"
    },
    "clustersItems": {
      "type": "array",
      "uniqueItems": true,
      "items": {
        "$ref": "#/definitions/Cluster"
      },
      "x-go-gen-location": "models"
    },
    "getAppruntimesOKBody": {
      "allOf": [
        {
          "$ref": "#/definitions/AppRuntimes"
        },
        {
          "$ref": "#/definitions/Paging"
        }
      ],
      "x-go-gen-location": "operations"
    },
    "getAppsOKBody": {
      "allOf": [
        {
          "$ref": "#/definitions/Apps"
        },
        {
          "$ref": "#/definitions/Paging"
        }
      ],
      "x-go-gen-location": "operations"
    },
    "getClustersOKBody": {
      "allOf": [
        {
          "$ref": "#/definitions/Clusters"
        },
        {
          "$ref": "#/definitions/Paging"
        }
      ],
      "x-go-gen-location": "operations"
    },
    "getReposOKBody": {
      "allOf": [
        {
          "$ref": "#/definitions/Repos"
        },
        {
          "$ref": "#/definitions/Paging"
        }
      ],
      "x-go-gen-location": "operations"
    },
    "reposItems": {
      "type": "array",
      "uniqueItems": true,
      "items": {
        "$ref": "#/definitions/Repo"
      },
      "x-go-gen-location": "models"
    }
  },
  "parameters": {
    "appId": {
      "type": "string",
      "description": "The app's id",
      "name": "appId",
      "in": "path",
      "required": true
    },
    "appRuntimeId": {
      "type": "string",
      "description": "The app runtime's id",
      "name": "appRuntimeId",
      "in": "path",
      "required": true
    },
    "clusterId": {
      "type": "string",
      "description": "The cluster's id",
      "name": "clusterId",
      "in": "path",
      "required": true
    },
    "pageNumber": {
      "type": "integer",
      "default": 1,
      "description": "Page number",
      "name": "pageNumber",
      "in": "query"
    },
    "pageSize": {
      "minimum": 0,
      "exclusiveMinimum": true,
      "multipleOf": 10,
      "type": "integer",
      "format": "int32",
      "default": 20,
      "description": "Number of clusters returned",
      "name": "pageSize",
      "in": "query"
    },
    "repoId": {
      "type": "string",
      "description": "The repo's id",
      "name": "repoId",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "AppDoesNotExistResponse": {
      "description": "The app does not exist."
    },
    "AppRuntimeDoesNotExistResponse": {
      "description": "The app runtime does not exist."
    },
    "ClusterDoesNotExistResponse": {
      "description": "The cluster does not exist."
    },
    "RepoDoesNotExistResponse": {
      "description": "The repo does not exist."
    },
    "Standard500ErrorResponse": {
      "description": "An unexpected error occured.",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  }
}`))
}
