define({ "api": [
  {
    "type": "POST",
    "url": "/domicilios/borar",
    "title": "Borrar Domicilio",
    "description": "<p>Permite borrar un domicilio de un cliente</p>",
    "group": "Domicilios",
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
          "content": "{\n\t \"Domicilios\":{\n\t\t \"IdDomicilio\":1\n\t },\n\t \"Cliente\":{\n\t\t \"IdCliente\":1\n\t }\n }",
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
    "filename": "internal/controllers/domiciliosController.go",
    "groupTitle": "Domicilios",
    "name": "PostDomiciliosBorar"
  },
  {
    "type": "POST",
    "url": "/domicilios/crear",
    "title": "Crear Domicilio",
    "description": "<p>Permite crear un domicilio para un cliente</p>",
    "group": "Domicilios",
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
          "content": "{\n    \"error\": {\n        \"codigo\": \"ERROR_NOEXISTE_CLIENTE\",\n        \"mensaje\": \"No existe el cliente ingresado.\"\n    },\n    \"respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n\n}",
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
    "filename": "internal/controllers/domiciliosController.go",
    "groupTitle": "Domicilios",
    "name": "PostDomiciliosCrear"
  },
  {
    "type": "GET",
    "url": "/roles/listar",
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
    "name": "GetRolesListar"
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
    "type": "POST",
    "url": "/usuarios/borar",
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
    "name": "PostUsuariosBorar"
  },
  {
    "type": "POST",
    "url": "/usuarios/buscar",
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
    "name": "PostUsuariosBuscar"
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
