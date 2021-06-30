// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kube for Edge Management",
    "title": "Kube4EdgeManagement",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/k4e-management/v1",
  "paths": {
    "/device/{device_id}": {
      "get": {
        "security": [
          {
            "agentAuth": []
          }
        ],
        "tags": [
          "devices"
        ],
        "operationId": "GetDeviceConfiguration",
        "parameters": [
          {
            "type": "string",
            "description": "Device to get configuration for",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success.",
            "schema": {
              "$ref": "#/definitions/device-configuration"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden."
          },
          "404": {
            "description": "Error."
          },
          "500": {
            "description": "Error."
          }
        }
      }
    }
  },
  "definitions": {
    "device-configuration": {
      "type": "object",
      "properties": {
        "device_id": {
          "description": "Device identifier",
          "type": "string"
        },
        "workloads": {
          "description": "List of workloads deployed to the device",
          "$ref": "#/definitions/workload-list"
        }
      }
    },
    "workload": {
      "type": "object",
      "properties": {
        "name": {
          "description": "Name of the workload",
          "type": "string"
        }
      }
    },
    "workload-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/workload"
      }
    }
  },
  "securityDefinitions": {
    "agentAuth": {
      "type": "apiKey",
      "name": "X-Secret-Key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Device management",
      "name": "devices"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kube for Edge Management",
    "title": "Kube4EdgeManagement",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/k4e-management/v1",
  "paths": {
    "/device/{device_id}": {
      "get": {
        "security": [
          {
            "agentAuth": []
          }
        ],
        "tags": [
          "devices"
        ],
        "operationId": "GetDeviceConfiguration",
        "parameters": [
          {
            "type": "string",
            "description": "Device to get configuration for",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success.",
            "schema": {
              "$ref": "#/definitions/device-configuration"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden."
          },
          "404": {
            "description": "Error."
          },
          "500": {
            "description": "Error."
          }
        }
      }
    }
  },
  "definitions": {
    "device-configuration": {
      "type": "object",
      "properties": {
        "device_id": {
          "description": "Device identifier",
          "type": "string"
        },
        "workloads": {
          "description": "List of workloads deployed to the device",
          "$ref": "#/definitions/workload-list"
        }
      }
    },
    "workload": {
      "type": "object",
      "properties": {
        "name": {
          "description": "Name of the workload",
          "type": "string"
        }
      }
    },
    "workload-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/workload"
      }
    }
  },
  "securityDefinitions": {
    "agentAuth": {
      "type": "apiKey",
      "name": "X-Secret-Key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Device management",
      "name": "devices"
    }
  ]
}`))
}
