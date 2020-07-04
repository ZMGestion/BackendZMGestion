define({ "api": [
  {
    "type": "POST",
    "url": "/ciudades",
    "title": "Listar Ciudades",
    "description": "<p>Devuelve la lista de ciudades de una provincia</p>",
    "group": "Ciudades",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Provincias",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Provincias.IdProvincia",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Provincias.IdCiudad",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Provincias\":{\n        \"IdPais\":\"AR\",\n        \"IdProvincia\":1\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": [\n        {\n            \"IdCiudad\": 1,\n            \"IdProvincia\": 1,\n            \"IdPais\": \"AR\",\n            \"Ciudad\": \"San Miguel de Tucumán\"\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/ciudadesController.go",
    "groupTitle": "Ciudades",
    "name": "PostCiudades"
  },
  {
    "type": "POST",
    "url": "/clientes",
    "title": "Buscar Clienter",
    "description": "<p>Permite buscar clientes</p>",
    "group": "Clientes",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.IdPais",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Documento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Nombres",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Apellidos",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.RazonSocial",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Tipo",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Email",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Telefono",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Clientes\":{\n        \"IdPais\": \"AR\",\n        \"Documento\": \"41144069\",\n        \"Tipo\":\"F\",\n        \"Nombres\":\"Loik\",\n        \"Apellidos\":\"Choua\",\n        \"Email\":\"loikchoua4@gmail.com\",\n        \"Telefono\":\"3815483777\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Clientes\": [\n            {\n                \"IdCliente\": 3,\n                \"IdPais\": \"AR\",\n                \"IdTipoDocumento\": 1,\n                \"Documento\": \"41144069\",\n                \"Tipo\": \"F\",\n                \"FechaNacimiento\": \"1998-05-27\",\n                \"Nombres\": \"Loik\",\n                \"Apellidos\": \"Choua\",\n                \"RazonSocial\": \"\",\n                \"Email\": \"loikchoua4@gmail.com\",\n                \"Telefono\": \"+543815483777\",\n                \"FechaAlta\": \"2020-06-24 15:32:47.000000\",\n                \"FechaBaja\": \"\",\n                \"Estado\": \"A\"\n            }\n        ]\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientes"
  },
  {
    "type": "POST",
    "url": "/clientes/borrar",
    "title": "Borrar Cliente",
    "description": "<p>Permite borrar un cliente</p>",
    "group": "Clientes",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Clientes.IdCliente",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Clientes\":{\n        \"IdCliente\": 3\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Clientes\": {\n            \"IdCliente\": 3,\n            \"IdPais\": \"AR\",\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Tipo\": \"F\",\n            \"FechaNacimiento\": \"\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"RazonSocial\": \"\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"Telefono\": \"+543815483777\",\n            \"FechaAlta\": \"2020-06-24 15:32:47.000000\",\n            \"FechaBaja\": \"\",\n            \"Estado\": \"B\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientesBorrar"
  },
  {
    "type": "POST",
    "url": "/clientes/crear",
    "title": "Crear Cliente",
    "description": "<p>Permite crear un cliente</p>",
    "group": "Clientes",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.IdPais",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Clientes.IdTipoDocumento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Documento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Nombres",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Apellidos",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.RazonSocial",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Tipo",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Email",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Telefono",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.FechaNacimiento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Domicilios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Domicilios.Domicilio",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Domicilios.CodigoPostal",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Domicilios.IdPais",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdProvincia",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdCiudad",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{  \"Clientes\":{\n        \"IdPais\": \"AR\",\n        \"IdTipoDocumento\": 1,\n        \"Documento\": \"41144069\",\n        \"Tipo\":\"F\",\n        \"FechaNacimiento\":\"1998-05-27\",\n        \"Nombres\":\"Loik\",\n        \"Apellidos\":\"Choua\",\n        \"Email\":\"loikchoua4@gmail.com\",\n        \"Telefono\":\"3815483777\"\n    },\n    \"Domicilios\":{\n        \"IdCiudad\":1,\n        \"IdProvincia\":1,\n        \"IdPais\":\"AR\",\n        \"Domicilio\":\"Domicilio\",\n        \"CodigoPostal\":\"4000\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Clientes\": {\n            \"IdCliente\": 3,\n            \"IdPais\": \"AR\",\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Tipo\": \"F\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"RazonSocial\": \"\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"Telefono\": \"3815483777\",\n            \"FechaAlta\": \"2020-06-24 15:32:47.000000\",\n            \"FechaBaja\": \"\",\n            \"Estado\": \"Choua\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientesCrear"
  },
  {
    "type": "POST",
    "url": "/clientes/darAlta",
    "title": "Dar alta Cliente",
    "description": "<p>Permite dar de alta un cliente</p>",
    "group": "Clientes",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Clientes.IdCliente",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Clientes\":{\n        \"IdCliente\": 3\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Clientes\": {\n            \"IdCliente\": 3,\n            \"IdPais\": \"AR\",\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Tipo\": \"F\",\n            \"FechaNacimiento\": \"\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"RazonSocial\": \"\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"Telefono\": \"+543815483777\",\n            \"FechaAlta\": \"2020-06-24 15:32:47.000000\",\n            \"FechaBaja\": \"\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientesDaralta"
  },
  {
    "type": "POST",
    "url": "/clientes/darBaja",
    "title": "Dar baja Cliente",
    "description": "<p>Permite dar de baja un cliente</p>",
    "group": "Clientes",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Clientes.IdCliente",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Clientes\":{\n        \"IdCliente\": 3\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Clientes\": {\n            \"IdCliente\": 3,\n            \"IdPais\": \"AR\",\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Tipo\": \"F\",\n            \"FechaNacimiento\": \"\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"RazonSocial\": \"\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"Telefono\": \"+543815483777\",\n            \"FechaAlta\": \"2020-06-24 15:32:47.000000\",\n            \"FechaBaja\": \"\",\n            \"Estado\": \"B\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientesDarbaja"
  },
  {
    "type": "POST",
    "url": "/clientes/domicilios",
    "title": "Listar Domicilios",
    "description": "<p>Permite listar los domicilios de un cliente</p>",
    "group": "Clientes",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Clientes.IdCliente",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"Clientes\":{\n\t\t\"IdCliente\":1\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Domicilios\": [\n            {\n                \"IdDomicilio\": 14,\n                \"IdCiudad\": 1,\n                \"IdProvincia\": 1,\n                \"IdPais\": \"AR\",\n                \"Domicilio\": \"El Tipal Lote 13\",\n                \"CodigoPostal\": \"El Tipal Lote 13\",\n                \"FechaAlta\": \"2020-06-28 23:04:26.000000\",\n                \"Observaciones\": \"\"\n            },\n            {\n                \"IdDomicilio\": 15,\n                \"IdCiudad\": 1,\n                \"IdProvincia\": 1,\n                \"IdPais\": \"AR\",\n                \"Domicilio\": \"El Tipal Lote 14\",\n                \"CodigoPostal\": \"El Tipal Lote 14\",\n                \"FechaAlta\": \"2020-06-28 23:04:49.000000\",\n                \"Observaciones\": \"\"\n            }\n        ]\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientesDomicilios"
  },
  {
    "type": "POST",
    "url": "/clientes/domicilios/agregar",
    "title": "Agregar Domicilio",
    "description": "<p>Permite crear un domicilio para un cliente</p>",
    "group": "Clientes",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Domicilios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdCiudad",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdProvincia",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdPais",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Domicilios.Domicilio",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Domicilio.CodigoPostal",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Domicilios.Observaciones",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Clientes.IdCliente",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"Clientes\":{\n\t\t\"IdCliente\":1\n\t},\n\t\"Domicilios\":{\n\t\t\"Domicilio\":\"El Tipal Lote 13\",\n\t\t\"IdCiudad\":1,\n\t\t\"IdProvincia\":1,\n\t\t\"IdPais\":\"AR\",\n\t\t\"CodigoPostal\":\"4107\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Domicilios\": {\n            \"IdDomicilio\": 14,\n            \"IdCiudad\": 1,\n            \"IdProvincia\": 1,\n            \"IdPais\": \"AR\",\n            \"Domicilio\": \"El Tipal Lote 13\",\n            \"CodigoPostal\": \"4107\",\n            \"FechaAlta\": \"2020-06-28 23:04:26.000000\",\n            \"Observaciones\": \"\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientesDomiciliosAgregar"
  },
  {
    "type": "POST",
    "url": "/clientes/domicilios/quitar",
    "title": "Quitar Domicilio",
    "description": "<p>Permite borrar un domicilio de un cliente</p>",
    "group": "Clientes",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Domicilios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdDomicilio",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Clientes.IdCliente",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Domicilios\":{\n\t\t \"IdDomicilio\":1\n\t },\n\t \"Clientes\":{\n\t\t \"IdCliente\":1\n\t }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientesDomiciliosQuitar"
  },
  {
    "type": "POST",
    "url": "/clientes/modificar",
    "title": "Modificar Cliente",
    "description": "<p>Permite modificar un cliente</p>",
    "group": "Clientes",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Clientes",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.IdPais",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Clientes.IdTipoDocumento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Documento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Nombres",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Apellidos",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.RazonSocial",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Tipo",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Email",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.Telefono",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Clientes.FechaNacimiento",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Clientes\":{\n        \"IdPais\": \"AR\",\n        \"IdTipoDocumento\": 1,\n        \"Documento\": \"41144069\",\n        \"Tipo\":\"F\",\n        \"FechaNacimiento\":\"1998-05-27\",\n        \"Nombres\":\"Loik\",\n        \"Apellidos\":\"Choua\",\n        \"Email\":\"loikchoua4@gmail.com\",\n        \"Telefono\":\"3815483777\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Clientes\": {\n            \"IdCliente\": 3,\n            \"IdPais\": \"AR\",\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Tipo\": \"F\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"RazonSocial\": \"\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"Telefono\": \"+543815483777\",\n            \"FechaAlta\": \"2020-06-24 15:32:47.000000\",\n            \"FechaBaja\": \"\",\n            \"Estado\": \"Choua\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/clientesController.go",
    "groupTitle": "Clientes",
    "name": "PostClientesModificar"
  },
  {
    "type": "POST",
    "url": "/gruposProducto",
    "title": "Buscar Grupos de Producto",
    "description": "<p>Permite buscar un grupo de producto</p>",
    "group": "GruposProducto",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "GruposProducto",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "GruposProducto.Grupo",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "GruposProducto.Estado",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"GruposProducto\":{\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"GruposProducto\":[\n\t\t\t{\n\t\t\t\t\"IdGrupoProducto\": 2,\n\t\t\t\t\"Grupo\": \"Grupo de prueba 2\",\n\t\t\t\t\"FechaAlta\": \"2020-07-04 16:38:50.000000\",\n\t\t\t\t\"FechaBaja\": \"\",\n\t\t\t\t\"Descripcion\": \"\",\n\t\t\t\t\"Estado\": \"A\"\n\t\t\t},\n\t\t\t{\n\t\t\t\t\"IdGrupoProducto\": 5,\n\t\t\t\t\"Grupo\": \"Grupo1\",\n\t\t\t\t\"FechaAlta\": \"2020-07-04 22:15:38.000000\",\n\t\t\t\t\"FechaBaja\": \"2020-07-04 22:15:44.000000\",\n\t\t\t\t\"Descripcion\": \"\",\n\t\t\t\t\"Estado\": \"A\"\n\t\t\t}\n\t\t]\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/gruposProductoController.go",
    "groupTitle": "GruposProducto",
    "name": "PostGruposproducto"
  },
  {
    "type": "POST",
    "url": "/gruposProducto/borrar",
    "title": "Borrar Grupo de Producto",
    "description": "<p>Permite borrar un grupo de producto</p>",
    "group": "GruposProducto",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "GruposProducto",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "GruposProducto.IdGrupoProducto",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"GruposProducto\":{\n\t\t\"IdGrupoProducto\":3\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/gruposProductoController.go",
    "groupTitle": "GruposProducto",
    "name": "PostGruposproductoBorrar"
  },
  {
    "type": "POST",
    "url": "/gruposProducto/crear",
    "title": "Crear Grupo de Producto",
    "description": "<p>Permite crear un grupo de producto</p>",
    "group": "GruposProducto",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "GruposProducto",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "GruposProducto.Grupo",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "GruposProducto.Descripcion",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"GruposProducto\":{\n        \"Grupo\":\"Grupo1\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"GruposProducto\":{\n\t\t\t\"IdGrupoProducto\": 3,\n\t\t\t\"Grupo\": \"Grupo1\",\n\t\t\t\"FechaAlta\": \"2020-07-04 21:39:47.000000\",\n\t\t\t\"FechaBaja\": \"\",\n\t\t\t\"Descripcion\": \"\",\n\t\t\t\"Estado\": \"A\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/gruposProductoController.go",
    "groupTitle": "GruposProducto",
    "name": "PostGruposproductoCrear"
  },
  {
    "type": "POST",
    "url": "/gruposProducto/darAlta",
    "title": "Dar alta GrupoProducto",
    "description": "<p>Permite dar de alta un grupo de producto</p>",
    "group": "GruposProducto",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "GruposProducto",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "GruposProducto.IdGrupoProducto",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"GruposProducto\": {\n            \"IdGrupoProducto\":4,\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"GruposProducto\":{\n\t\t\t\"IdGrupoProducto\": 5,\n\t\t\t\"Grupo\": \"Grupo1\",\n\t\t\t\"FechaAlta\": \"2020-07-04 22:15:38.000000\",\n\t\t\t\"FechaBaja\": \"2020-07-04 22:15:44.000000\",\n\t\t\t\"Descripcion\": \"\",\n\t\t\t\"Estado\": \"A\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/gruposProductoController.go",
    "groupTitle": "GruposProducto",
    "name": "PostGruposproductoDaralta"
  },
  {
    "type": "POST",
    "url": "/gruposProducto/darBaja",
    "title": "Dar baja GrupoProducto",
    "description": "<p>Permite dar de baja un grupo de producto</p>",
    "group": "GruposProducto",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "GruposProducto",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "GruposProducto.IdGrupoProducto",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"GruposProducto\": {\n            \"IdGrupoProducto\":4,\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"GruposProducto\":{\n\t\t\t\"IdGrupoProducto\": 5,\n\t\t\t\"Grupo\": \"Grupo1\",\n\t\t\t\"FechaAlta\": \"2020-07-04 22:15:38.000000\",\n\t\t\t\"FechaBaja\": \"2020-07-04 22:15:44.000000\",\n\t\t\t\"Descripcion\": \"\",\n\t\t\t\"Estado\": \"B\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/gruposProductoController.go",
    "groupTitle": "GruposProducto",
    "name": "PostGruposproductoDarbaja"
  },
  {
    "type": "POST",
    "url": "/gruposProducto/modificar",
    "title": "Modificar Grupo de Producto",
    "description": "<p>Permite modificar un grupo de producto</p>",
    "group": "GruposProducto",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "GruposProducto",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "GruposProducto.IdGrupoProducto",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "GruposProducto.Grupo",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "GruposProducto.Descripcion",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"GruposProducto\":{\n\t\t\"IdGrupoProducto\":3,\n        \"Grupo\":\"Grupo1Modificado\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"GruposProducto\":{\n\t\t\t\"IdGrupoProducto\": 3,\n\t\t\t\"Grupo\": \"Grupo1Modificado\",\n\t\t\t\"FechaAlta\": \"2020-07-04 21:39:47.000000\",\n\t\t\t\"FechaBaja\": \"\",\n\t\t\t\"Descripcion\": \"\",\n\t\t\t\"Estado\": \"A\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/gruposProductoController.go",
    "groupTitle": "GruposProducto",
    "name": "PostGruposproductoModificar"
  },
  {
    "type": "GET",
    "url": "/paises",
    "title": "Listar Paises",
    "description": "<p>Devuelve una lista de paises</p>",
    "group": "Paises",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": [\n\t\t{\n\t\t\t\"Roles\":{\n\t\t\t\t\"IdRol\": 1,\n\t\t\t\t\"Rol\": \"Administradores\",\n\t\t\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\t\t\"Observaciones\": \"\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"Roles\":{\n\t\t\t\t\"IdRol\": 2,\n\t\t\t\t\"Rol\": \"Vendedores\",\n\t\t\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\t\t\"Observaciones\": \"\"\n\t\t\t}\n\t\t}\n\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/paisesController.go",
    "groupTitle": "Paises",
    "name": "GetPaises"
  },
  {
    "type": "POST",
    "url": "/provincias",
    "title": "Listar Provincias",
    "description": "<p>Devuelve la lista de provincias de un país</p>",
    "group": "Provincias",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Paises",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Paises.IdPais",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Paises\":{\n        \"IdPais\":\"AR\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": [\n        {\n            \"IdPais\": \"\",\n            \"IdProvincia\": 2,\n            \"Provincia\": \"Salta\"\n        },\n        {\n            \"IdPais\": \"\",\n            \"IdProvincia\": 1,\n            \"Provincia\": \"Tucumán\"\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/provinciasController.go",
    "groupTitle": "Provincias",
    "name": "PostProvincias"
  },
  {
    "type": "GET",
    "url": "/roles",
    "title": "Listar Roles",
    "permission": [
      {
        "name": "Administradores"
      }
    ],
    "description": "<p>Devuelve una lista de roles</p>",
    "group": "Roles",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": [\n\t\t{\n\t\t\t\"Roles\":{\n\t\t\t\t\"IdRol\": 1,\n\t\t\t\t\"Rol\": \"Administradores\",\n\t\t\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\t\t\"Observaciones\": \"\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"Roles\":{\n\t\t\t\t\"IdRol\": 2,\n\t\t\t\t\"Rol\": \"Vendedores\",\n\t\t\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\t\t\"Observaciones\": \"\"\n\t\t\t}\n\t\t}\n\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "GetRoles"
  },
  {
    "type": "POST",
    "url": "/roles/listarPermisos",
    "title": "Listar Permisos",
    "permission": [
      {
        "name": "Administradores"
      }
    ],
    "name": "Listar_Permisos",
    "description": "<p>Devuelve una lista de permisos para un determinado rol</p>",
    "group": "Roles",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Roles",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "IdRol",
            "description": ""
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": [\n\t\t{\n            \"Permisos\": {\n                \"IdPermiso\": 1,\n                \"Permiso\": \"Crear rol\"\n            }\n        },\n        {\n            \"Permisos\": {\n                \"IdPermiso\": 2,\n                \"Permiso\": \"Borrar rol\"\n            }\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles"
  },
  {
    "type": "POST",
    "url": "/roles/asignarPermisos",
    "title": "Asignar Permisos",
    "permission": [
      {
        "name": "Administradores"
      }
    ],
    "description": "<p>Permite asignar los permisos a un determinado rol</p>",
    "group": "Roles",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Roles",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Roles.IdRol",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object[]",
            "optional": false,
            "field": "Permisos",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": " {\n\t\"Roles\": {\n\t\t\"IdRol\": 9\n\t},\n\t\"Permisos\": [\n\t\t{\n\t\t\t\"IdPermiso\": 3\n\t\t},\n\t\t{\n\t\t\t\"IdPermiso\": 4\n\t\t}\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "PostRolesAsignarpermisos"
  },
  {
    "type": "POST",
    "url": "/roles/borrar",
    "title": "Borrar Rol",
    "permission": [
      {
        "name": "Administradores"
      }
    ],
    "description": "<p>Borra un rol a partir de su Id</p>",
    "group": "Roles",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Roles",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Roles.IdRol",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": " {\n\t \"Roles\": {\n\t\t \"IdRol\": \"2\"\n\t }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "PostRolesBorrar"
  },
  {
    "type": "POST",
    "url": "/roles/crear",
    "title": "Crear Rol",
    "permission": [
      {
        "name": "Administradores"
      }
    ],
    "description": "<p>Permite crear un rol</p>",
    "group": "Roles",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Roles",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Roles.Rol",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Roles.Descripcion",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": " {\n\t \"Roles\": {\n\t\t \"Rol\": \"Encargados\"\n\t }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": {\n\t\t\"Roles\" : {\n\t\t\t\"IdRol\": 7,\n\t\t\t\"Rol\": \"Encargados\",\n\t\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\t\"Descripcion\": \"\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_EXISTE_NOMBREROL\",\n        \"mensaje\": \"El nombre de rol ya existe.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "PostRolesCrear"
  },
  {
    "type": "POST",
    "url": "/roles/dame",
    "title": "Dame Rol",
    "permission": [
      {
        "name": "Administradores"
      }
    ],
    "description": "<p>Devuelve un rol a partir de un Id</p>",
    "group": "Roles",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Roles",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Roles.IdRol",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": " {\n\t \"Roles\": {\n\t\t \"IdRol\":2\n\t }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": {\n\t\t\"Roles\": {\n            \"IdRol\": 2,\n            \"Rol\": \"Vendedores\",\n            \"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n            \"Descripcion\": \"Este rol es para los vendedores\"\n        }\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_NOEXISTE_ROL\",\n        \"mensaje\": \"No existe el rol.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "PostRolesDame"
  },
  {
    "type": "POST",
    "url": "/roles/modificar",
    "title": "Modificar Rol",
    "permission": [
      {
        "name": "Administradores"
      }
    ],
    "description": "<p>Permite modificar un rol</p>",
    "group": "Roles",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Roles",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Roles.IdRol",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Roles.Rol",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Roles.Descripcion",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": " {\n\t \"Roles\": {\n\t\t \"IdRol\": 7,\n\t\t \"Rol\": \"Los encargados\",\n\t\t \"Descripcion\": \"Nueva descripcion\"\n\t }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": {\n\t\t\"IdRol\": 7,\n\t\t\"Rol\": \"Los encargados\",\n\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\"Descripcion\": \"Nueva descripcion\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_EXISTE_NOMBREROL\",\n        \"mensaje\": \"El nombre de rol ya existe.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "PostRolesModificar"
  },
  {
    "type": "POST",
    "url": "/telas",
    "title": "Buscar Telas",
    "description": "<p>Permite buscar una tela</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Telas.Tela",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Telas.Estado",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Telas\":{\n\t\t\"IdTela\":1,\n        \"Tela\": \"Prueba13\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Telas\": {\n            \"IdTela\": 4,\n            \"Tela\": \"Prueba5\",\n            \"FechaAlta\": \"2020-06-30 23:39:57.000000\",\n            \"FechaBaja\": \"\",\n            \"Observaciones\": \"\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su peticion.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelas"
  },
  {
    "type": "POST",
    "url": "/telas/borar",
    "title": "Borrar Tela",
    "description": "<p>Permite borrar una tela</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.IdTela",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Telas\": {\n            \"IdTela\":2\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelasBorar"
  },
  {
    "type": "POST",
    "url": "/telas/crear",
    "title": "Crear Tela",
    "description": "<p>Permite crear una tela</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.Tela",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.Observaciones",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Precios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Precios.Precio",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Telas\":{\n        \"Tela\": \"Prueba3\"\n    },\n    \"Precios\":{\n        \"Precio\":1.20\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Telas\": {\n            \"IdTela\": 4,\n            \"Tela\": \"Prueba5\",\n            \"FechaAlta\": \"2020-06-30 23:39:57.000000\",\n            \"FechaBaja\": \"\",\n            \"Observaciones\": \"\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_EXISTE_UBICACION\",\n        \"mensaje\": \"La ubicacion ingresada ya existe.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelasCrear"
  },
  {
    "type": "POST",
    "url": "/telas/dame",
    "title": "Dame Tela",
    "description": "<p>Permite instanciar una tela a partir de su Id</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.IdTela",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Telas\":{\n\t\t\"IdTela\":5\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Telas\": {\n            \"IdTela\": 4,\n            \"Tela\": \"Prueba5\",\n            \"FechaAlta\": \"2020-06-30 23:39:57.000000\",\n            \"FechaBaja\": \"\",\n            \"Observaciones\": \"\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su peticion.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelasDame"
  },
  {
    "type": "POST",
    "url": "/telas/darAlta",
    "title": "Dar alta Tela",
    "description": "<p>Permite dar de alta una tela</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.IdTela",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Telas\": {\n            \"IdTela\":3\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"Telas\":{\n\t\t\t\"IdTela\": 1,\n\t\t\t\"Tela\": \"Prueba65\",\n\t\t\t\"FechaAlta\": \"2020-06-29 22:56:45.000000\",\n\t\t\t\"FechaBaja\": \"2020-07-03 12:16:53.000000\",\n\t\t\t\"Observaciones\": \"\",\n\t\t\t\"Estado\": \"A\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelasDaralta"
  },
  {
    "type": "POST",
    "url": "/telas/darBaja",
    "title": "Dar baja Tela",
    "description": "<p>Permite dar de baja una tela</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.IdTela",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Telas\": {\n            \"IdTela\":3\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"Telas\":{\n\t\t\t\"IdTela\": 1,\n\t\t\t\"Tela\": \"Prueba65\",\n\t\t\t\"FechaAlta\": \"2020-06-29 22:56:45.000000\",\n\t\t\t\"FechaBaja\": \"2020-07-03 12:16:53.000000\",\n\t\t\t\"Observaciones\": \"\",\n\t\t\t\"Estado\": \"A\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelasDarbaja"
  },
  {
    "type": "POST",
    "url": "/telas/modificar",
    "title": "Modificar Tela",
    "description": "<p>Permite modificar una tela</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.IdTela",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.Tela",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.Observaciones",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Telas\":{\n\t\t\"IdTela\":1,\n        \"Tela\": \"Prueba13\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Telas\": {\n            \"IdTela\": 4,\n            \"Tela\": \"Prueba5\",\n            \"FechaAlta\": \"2020-06-30 23:39:57.000000\",\n            \"FechaBaja\": \"\",\n            \"Observaciones\": \"\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su peticion.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelasModificar"
  },
  {
    "type": "POST",
    "url": "/telas/precios",
    "title": "Listar Precios Tela",
    "description": "<p>Permite listar el historico de precios de una tela</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.IdTela",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Telas\": {\n            \"IdTela\":3\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"Precios\":[\n\t\t\t{\"IdPrecio\": 14, \"Precio\": 1.2, \"Tipo\": \"\", \"IdReferencia\": 0, \"FechaAlta\": \"2020-07-03 19:57:18.000000\"…},\n\t\t\t{\"IdPrecio\": 16, \"Precio\": 1.21, \"Tipo\": \"\", \"IdReferencia\": 0, \"FechaAlta\": \"2020-07-03 20:15:10.000000\"…},\n\t\t\t{\"IdPrecio\": 17, \"Precio\": 1.22, \"Tipo\": \"\", \"IdReferencia\": 0, \"FechaAlta\": \"2020-07-03 20:19:15.000000\"…},\n\t\t\t{\"IdPrecio\": 18, \"Precio\": 1.23, \"Tipo\": \"\", \"IdReferencia\": 0, \"FechaAlta\": \"2020-07-03 22:29:53.000000\"…}\n\t\t]\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelasPrecios"
  },
  {
    "type": "POST",
    "url": "/telas/precios/modificar",
    "title": "Modificar Precio Tela",
    "description": "<p>Permite modificar el precio de una tela</p>",
    "group": "Telas",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Telas",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Telas.IdTela",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Precios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Precios.IdPrecio",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"Telas\": {\n\t\t\"IdTela\":3\n\t},\n\t\"Precios\":{\n\t\t\"Precio\":1.21\n\t}\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\t\"error\": null,\n\t\"respuesta\":{\n\t\t\"Precios\":{\n\t\t\t\"IdPrecio\": 17,\n\t\t\t\"Precio\": 1.22,\n\t\t\t\"Tipo\": \"\",\n\t\t\t\"IdReferencia\": 0,\n\t\t\t\"FechaAlta\": \"2020-07-03 20:19:15.000000\"\n\t\t},\n\t\t\"Telas\":{\n\t\t\t\"IdTela\": 5,\n\t\t\t\"Tela\": \"Prueba5\",\n\t\t\t\"FechaAlta\": \"2020-07-03 19:57:18.000000\",\n\t\t\t\"FechaBaja\": \"\",\n\t\t\t\"Observaciones\": \"\",\n\t\t\t\"Estado\": \"A\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/telasController.go",
    "groupTitle": "Telas",
    "name": "PostTelasPreciosModificar"
  },
  {
    "type": "GET",
    "url": "/ubicaciones",
    "title": "Listar Ubicaciones",
    "name": "Listar_Ubicaciones",
    "description": "<p>Devuelve una lista con la ubicación y su dirección.</p>",
    "group": "Ubicaciones",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": [\n        {\n            \"Domicilios\": {\n                \"IdDomicilio\": 1,\n                \"IdCiudad\": 1,\n                \"IdProvincia\": 1,\n                \"IdPais\": \"AR\",\n                \"Domicilio\": \"Av. Manuel Belgrano 1456\",\n                \"CodigoPostal\": \"4000\",\n                \"FechaAlta\": \"2020-06-13 13:52:16.000000\",\n                \"Observaciones\": \"Domicilio de la casa central\"\n            },\n            \"Ubicaciones\": {\n                \"IdUbicacion\": 1,\n                \"IdDomicilio\": 1,\n                \"Ubicacion\": \"Casa Central Tucumán\",\n                \"FechaAlta\": \"2020-06-13 13:52:17.000000\",\n                \"Estado\": \"A\"\n            }\n        },\n        {\n            \"Domicilios\": {\n                \"IdDomicilio\": 2,\n                \"IdCiudad\": 1,\n                \"IdProvincia\": 1,\n                \"IdPais\": \"AR\",\n                \"Domicilio\": \"Ildefonso de Muñecas 374\",\n                \"CodigoPostal\": \"4000\",\n                \"FechaAlta\": \"2020-06-13 13:52:16.000000\",\n                \"Observaciones\": \"Domicilio sucursal Muñecas\"\n            },\n            \"Ubicaciones\": {\n                \"IdUbicacion\": 2,\n                \"IdDomicilio\": 2,\n                \"Ubicacion\": \"Sucursal Muñecas\",\n                \"FechaAlta\": \"2020-06-13 13:52:17.000000\",\n                \"Estado\": \"A\"\n            }\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/ubicacionesController.go",
    "groupTitle": "Ubicaciones"
  },
  {
    "type": "POST",
    "url": "/ubicaciones/borar",
    "title": "Borrar Ubicación",
    "description": "<p>Permite borrar una ubicación existente</p>",
    "group": "Ubicaciones",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Ubicaciones",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Ubicaciones.IdUbicacion",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"Ubicaciones\":{\n\t\t\"IdUbicacion\":5\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_SIN_PERMISOS\",\n        \"mensaje\": \"No cuenta con los permisos para ejecutar esta accion.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/ubicacionesController.go",
    "groupTitle": "Ubicaciones",
    "name": "PostUbicacionesBorar"
  },
  {
    "type": "POST",
    "url": "/ubicaciones/crear",
    "title": "Crear Ubicación",
    "description": "<p>Permite crear una ubicación</p>",
    "group": "Ubicaciones",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Ubicaciones",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Ubicaciones.Ubicacion",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.Observaciones",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Domicilios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdCiudad",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdProvincia",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Domicilios.IdPais",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Domicilios.Domicilio",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Domicilios.CodigoPostal",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Ubicaciones\":{\n        \"Ubicacion\": \"Prueba\"\n    },\n    \"Domicilios\":{\n        \"IdCiudad\":1,\n        \"IdProvincia\":1,\n        \"IdPais\":\"AR\",\n        \"Domicilio\":\"Domicilio de prueba\",\n        \"CodigoPostal\":\"PRUEBA\",\n        \"Observaciones\":\"\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Ubicaciones\": {\n            \"IdUbicacion\": 5,\n            \"IdDomicilio\": 19,\n            \"Ubicacion\": \"Prueba\",\n            \"FechaAlta\": \"2020-06-13 13:38:59.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_EXISTE_UBICACION\",\n        \"mensaje\": \"La ubicacion ingresada ya existe.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/ubicacionesController.go",
    "groupTitle": "Ubicaciones",
    "name": "PostUbicacionesCrear"
  },
  {
    "type": "POST",
    "url": "/ubicaciones/darAlta",
    "title": "Dar alta Ubicación",
    "description": "<p>Permite dar de alta una ubicación</p>",
    "group": "Ubicaciones",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Ubicaciones",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Ubicaciones.IdUbicacion",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Ubicaciones\":{\n        \"IdUbicacion\": 8\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Ubicaciones\": {\n            \"IdUbicacion\": 8,\n            \"IdDomicilio\": 8,\n            \"Ubicacion\": \"Modificar prueba\",\n            \"FechaAlta\": \"2020-06-13 15:43:20.000000\",\n            \"FechaBaja\": \"2020-06-13 15:42:10.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_UBICACION_ESTA_ALTA\",\n        \"mensaje\": \"La ubicación no existe o ya está en estado de 'Alta'.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/ubicacionesController.go",
    "groupTitle": "Ubicaciones",
    "name": "PostUbicacionesDaralta"
  },
  {
    "type": "POST",
    "url": "/ubicaciones/darBaja",
    "title": "Dar baja Ubicación",
    "description": "<p>Permite dar de baja una ubicación</p>",
    "group": "Ubicaciones",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Ubicaciones",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Ubicaciones.IdUbicacion",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Ubicaciones\": {\n            \"IdUbicacion\":8,\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Ubicaciones\": {\n            \"IdUbicacion\": 8,\n            \"IdDomicilio\": 8,\n            \"Ubicacion\": \"Modificar prueba\",\n            \"FechaAlta\": \"2020-06-13 14:55:33.000000\",\n            \"FechaBaja\": \"2020-06-13 15:42:10.000000\",\n            \"Estado\": \"B\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_UBICACION_ESTA_BAJA\",\n        \"mensaje\": \"La ubicación no existe o ya está en estado de 'Baja'.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/ubicacionesController.go",
    "groupTitle": "Ubicaciones",
    "name": "PostUbicacionesDarbaja"
  },
  {
    "type": "POST",
    "url": "/ubicaciones/modificar",
    "title": "Modificar Ubicación",
    "description": "<p>Permite modificar una ubicación existente</p>",
    "group": "Ubicaciones",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Ubicaciones",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Ubicaciones.IdUbicacion",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"Ubicaciones\":{\n        \"IdUbicacion\": 8,\n        \"Ubicacion\":\"Modificar prueba\"\n    },\n    \"Domicilios\":{\n        \"IdDomicilio\":8,\n        \"IdCiudad\":1,\n        \"IdProvincia\":1,\n        \"IdPais\":\"AR\",\n        \"Domicilio\":\"Domicilio de prueba\",\n        \"CodigoPostal\":\"PRUEBA\",\n        \"Observaciones\":\"\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_SIN_PERMISOS\",\n        \"mensaje\": \"No cuenta con los permisos para ejecutar esta accion.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/ubicacionesController.go",
    "groupTitle": "Ubicaciones",
    "name": "PostUbicacionesModificar"
  },
  {
    "type": "GET",
    "url": "/usuarios/damePorToken",
    "title": "Dame Usuario por Token",
    "description": "<p>Devuelve un usuario a partir de un Token</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 1,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Nombres\": \"Adam el super admin\",\n            \"EstadoCivil\": \"C\",\n            \"Telefono\": \"+54(381)4321719\",\n            \"Email\": \"zimmermanmueblesgestion@gmail.com\",\n            \"CantidadHijos\": 2,\n            \"Usuario\": \"adam\",\n            \"FechaNacimiento\": \"2020-04-11\",\n            \"FechaInicio\": \"2020-04-11\",\n            \"FechaAlta\": \"2020-04-11 19:50:01.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "GetUsuariosDameportoken"
  },
  {
    "type": "GET",
    "url": "/usuarios/tiposDocumento",
    "title": "Listar Tipos Documento",
    "description": "<p>Devuelve una lista de tipos documento</p>",
    "group": "Usuarios",
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": [\n\t\t{\n\n\t\t\t\"TiposDocumento\":{\n\t\t\t\t\"IdTipoDocumento\": 1,\n\t\t\t\t\"TipoDocumento\": \"DNI\",\n\t\t\t\t\"Descripcion\": \"Documento Nacional de Identidad\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"TiposDocumento\":{\n\t\t\t\t\"IdTipoDocumento\": 2,\n\t\t\t\t\"TipoDocumento\": \"Pasaporte\",\n\t\t\t\t\"Descripcion\": \"Pasaporte\"\n\t\t\t}\n\t\t}\n\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "GetUsuariosTiposdocumento"
  },
  {
    "type": "POST",
    "url": "/usuarios",
    "title": "Buscar Usuarios",
    "description": "<p>Permite buscar un usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "Usuarios.IdRol",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": true,
            "field": "Usuarios.IdUbicacion",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.Documento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.Nombres",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.Apellidos",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.EstadoCivil",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.Telefono",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.Email",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.Usuario",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"Usuarios\":{\n\t\t\"Usuario\":\"nbachs\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": [\n        {\n            \"Roles\": {\n                \"IdRol\": 1,\n                \"Rol\": \"Administradores\"\n            },\n            \"Ubicaciones\": {\n                \"IdUbicacion\": 1,\n                \"Ubicacion\": \"Casa Central Tucumán\"\n            },\n            \"Usuarios\": {\n                \"IdUsuario\": 5,\n                \"IdRol\": 1,\n                \"IdUbicacion\": 1,\n                \"IdTipoDocumento\": 1,\n                \"Documento\": \"39477073\",\n                \"Nombres\": \"Nicolas\",\n                \"Apellidos\": \"Bachs\",\n                \"EstadoCivil\": \"S\",\n                \"Telefono\": \"+543814491954\",\n                \"Email\": \"nicolas.bachs@gmail.com\",\n                \"CantidadHijos\": 0,\n                \"Usuario\": \"nbachs\",\n                \"FechaUltIntento\": \"2020-05-08 00:52:23.000000\",\n                \"FechaNacimiento\": \"1995-12-27\",\n                \"FechaInicio\": \"2019-11-22\",\n                \"FechaAlta\": \"2020-05-08 00:49:18.000000\",\n                \"Estado\": \"A\"\n            }\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuarios"
  },
  {
    "type": "POST",
    "url": "/usuarios/borrar",
    "title": "Borrar Usuario",
    "description": "<p>Permite borrar un usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Usuarios\": {\n            \"IdUsuario\":6,\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosBorrar"
  },
  {
    "type": "POST",
    "url": "/usuarios/cerrarSesion",
    "title": "Cerrar Sesion",
    "description": "<p>Permite a un usuario cerrar la sesion de otro usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdUsuario",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"Usuarios\":{\n        \"IdUsuario\": 1\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_NOEXISTE_USUARIO\",\n        \"mensaje\": \"El usuario indicado no existe.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosCerrarsesion"
  },
  {
    "type": "POST",
    "url": "/usuarios/crear",
    "title": "Crear Usuario",
    "description": "<p>Permite crear un usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdRol",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdUbicacion",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdTipoDocumento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Documento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Nombres",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Apellidos",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.EstadoCivil",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Telefono",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Email",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.CantidadHijos",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Usuario",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Password",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.FechaNacimiento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.FechaInicio",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Usuarios\": {\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n\t\t\t\"IdTipoDocumento\": 1,\n\t\t\t\"Documento\": \"41144069\",\n\t\t\t\"Nombres\": \"Loik\",\n\t\t\t\"Apellidos\" : \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"CantidadHijos\": 0,\n\t\t\t\"Usuario\": \"lchoua\",\n\t\t\t\"Password\":\"LoikCapo\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\"\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 6,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"lchoua\",\n            \"Password\": \"LoikCapo\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-12 18:11:39.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosCrear"
  },
  {
    "type": "POST",
    "url": "/usuarios/dame",
    "title": "Dame Usuario",
    "description": "<p>Devuelve un usuario a partir de un Id</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdUsuario",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": " {\n\t \"Usuarios\": {\n\t\t \"IdUsuario\":1\n\t }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"error\": null,\n    \"respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 1,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Nombres\": \"Adam el super admin\",\n            \"EstadoCivil\": \"C\",\n            \"Telefono\": \"+54(381)4321719\",\n            \"Email\": \"zimmermanmueblesgestion@gmail.com\",\n            \"CantidadHijos\": 2,\n            \"Usuario\": \"adam\",\n            \"FechaNacimiento\": \"2020-04-11\",\n            \"FechaInicio\": \"2020-04-11\",\n            \"FechaAlta\": \"2020-04-11 19:50:01.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosDame"
  },
  {
    "type": "POST",
    "url": "/usuarios/darAlta",
    "title": "Dar alta Usuario",
    "description": "<p>Permite dar de alta un usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Usuarios\": {\n            \"IdUsuario\":6,\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 8,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"lchoua\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-12 20:24:34.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosDaralta"
  },
  {
    "type": "POST",
    "url": "/usuarios/darBaja",
    "title": "Dar baja Usuario",
    "description": "<p>Permite dar de baja un usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Usuarios\": {\n            \"IdUsuario\":6,\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 8,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"lchoua\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-12 20:24:34.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosDarbaja"
  },
  {
    "type": "POST",
    "url": "/usuarios/iniciarSesion",
    "title": "Iniciar Sesion",
    "description": "<p>Permite a un usuario iniciar sesion</p>",
    "group": "Usuarios",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.Email",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "Usuarios.Usuario",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Password",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Usuarios\": {\n\t\t\t\"Usuario\": \"lchoua\",\n\t\t\t\"Password\":\"LoikCapo\",\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 5,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"39477073\",\n            \"Nombres\": \"Nicolas\",\n            \"Apellidos\": \"Bachs\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+543814491954\",\n            \"Email\": \"nicolas.bachs@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"nbachs\",\n            \"Token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk0MTI5MzIsImlhdCI6MTU4OTQwOTMzMiwidXNlciI6Im5iYWNocyJ9.IX4fpwOpjp8-TjPMVKqT1NmZ0gQG0shyrLfLq6ETi-M\",\n            \"FechaNacimiento\": \"1995-12-27\",\n            \"FechaInicio\": \"2019-11-22\",\n            \"FechaAlta\": \"2020-05-08 00:49:18.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_LOGIN_INCORRECTO\",\n        \"mensaje\": \"El nombre de usuario o contrasena ingresados no son correctos.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosIniciarsesion"
  },
  {
    "type": "POST",
    "url": "/usuarios/modificar",
    "title": "Modificar Usuario",
    "description": "<p>Permite modificar un usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdUsuario",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdRol",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdUbicacion",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdTipoDocumento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Documento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Nombres",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Apellidos",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.EstadoCivil",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Telefono",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Email",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.CantidadHijos",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Usuario",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Password",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.FechaNacimiento",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.FechaInicio",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t \"Usuarios\": {\n            \"IdUsuario\":6,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n\t\t\t\"IdTipoDocumento\": 1,\n\t\t\t\"Documento\": \"41144069\",\n\t\t\t\"Nombres\": \"Loik\",\n\t\t\t\"Apellidos\" : \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"CantidadHijos\": 0,\n\t\t\t\"Usuario\": \"lchoua\",\n\t\t\t\"Password\":\"LoikCapo\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\"\n        }\n }",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 6,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"lchoua\",\n            \"Password\": \"LoikCapo\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-12 18:11:39.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_DEFAULT\",\n        \"mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosModificar"
  },
  {
    "type": "POST",
    "url": "/usuarios/modificarPassword",
    "title": "Modificar contraseña",
    "description": "<p>Permite modificar la contraseña de un usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "UsuariosActual",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "UsuariosActual.Password",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "UsuariosNuevo",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "UsuariosNuevo.Password",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"UsuariosActual\":{\n\t\t\"Password\":\"Nueva Pass\"\n\t},\n\t\"UsuariosNuevo\":{\n\t\t\"Password\":\"Hola123\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null,\n    \"respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 9,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"42664256\",\n            \"Nombres\": \"Chloe\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)4451337\",\n            \"Email\": \"chloechoua@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"cchoua\",\n            \"FechaNacimiento\": \"2000-05-17\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-13 20:25:19.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_PASSWORD_INCORRECTA\",\n        \"mensaje\": \"La contraseña ingresada no es correcta.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosModificarpassword"
  },
  {
    "type": "POST",
    "url": "/usuarios/restablecerPassword",
    "title": "Restablecer contraseña",
    "description": "<p>Permite restablecer la contraseña de un usuario</p>",
    "group": "Usuarios",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "Authorization",
            "description": ""
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Object",
            "optional": false,
            "field": "Usuarios",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "Usuarios.IdUsuario",
            "description": ""
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "Usuarios.Password",
            "description": ""
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n\t\"Usuarios\":{\n        \"IdUsuario\": 1,\n        \"Password\":\"Nueva pass\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"error\": null\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_USUARIO_ESTA_BAJA\",\n        \"mensaje\": \"El usuario no existe o ya está dado de baja.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/usuariosController.go",
    "groupTitle": "Usuarios",
    "name": "PostUsuariosRestablecerpassword"
  }
] });
