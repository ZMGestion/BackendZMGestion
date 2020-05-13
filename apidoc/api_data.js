define({ "api": [
  {
    "type": "GET",
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
          "content": " {\n    \"Error\": null,\n    \"Respuesta\": {\n\t\t\"Roles\" : {\n\t\t\t\"IdRol\": 7,\n\t\t\t\"Rol\": \"Encargados\",\n\t\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\t\"Descripcion\": \"\"\n\t\t}\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"Error\": {\n        \"Codigo\": \"ERROR_EXISTE_NOMBREROL\",\n        \"Mensaje\": \"El nombre de rol ya existe.\"\n    },\n    \"Respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "GetRolesCrear"
  },
  {
    "type": "GET",
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
          "content": " {\n    \"Error\": null,\n    \"Respuesta\": {\n\t\t\"IdRol\": 7,\n\t\t\"Rol\": \"Los encargados\",\n\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\"Descripcion\": \"Nueva descripcion\"\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"Error\": {\n        \"Codigo\": \"ERROR_EXISTE_NOMBREROL\",\n        \"Mensaje\": \"El nombre de rol ya existe.\"\n    },\n    \"Respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "GetRolesModificar"
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
    "success": {
      "examples": [
        {
          "title": "Success-Response:",
          "content": " {\n    \"Error\": null,\n    \"Respuesta\": [\n\t\t{\n            \"Permisos\": {\n                \"IdPermiso\": 1,\n                \"Permiso\": \"Crear rol\"\n            }\n        },\n        {\n            \"Permisos\": {\n                \"IdPermiso\": 2,\n                \"Permiso\": \"Borrar rol\"\n            }\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": " {\n    \"Error\": null,\n    \"Respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": " {\n    \"Error\": null,\n    \"Respuesta\": null\n}",
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
          "content": " {\n    \"Error\": null,\n    \"Respuesta\": {\n\t\t\"Roles\": {\n            \"IdRol\": 2,\n            \"Rol\": \"Vendedores\",\n            \"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n            \"Descripcion\": \"Este rol es para los vendedores\"\n        }\n\t}\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"Error\": {\n        \"Codigo\": \"ERROR_NOEXISTE_ROL\",\n        \"Mensaje\": \"No existe el rol.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": " {\n    \"Error\": null,\n    \"Respuesta\": [\n\t\t{\n\t\t\t\"Roles\":{\n\t\t\t\t\"IdRol\": 1,\n\t\t\t\t\"Rol\": \"Administradores\",\n\t\t\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\t\t\"Observaciones\": \"\"\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"Roles\":{\n\t\t\t\t\"IdRol\": 2,\n\t\t\t\t\"Rol\": \"Vendedores\",\n\t\t\t\t\"FechaAlta\": \"2020-04-09 15:01:35.000000\",\n\t\t\t\t\"Observaciones\": \"\"\n\t\t\t}\n\t\t}\n\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": " {\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "internal/controllers/rolesController.go",
    "groupTitle": "Roles",
    "name": "PostRolesListar"
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
          "content": "{\n    \"Error\": null,\n    \"Respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 1,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Nombres\": \"Adam el super admin\",\n            \"EstadoCivil\": \"C\",\n            \"Telefono\": \"+54(381)4321719\",\n            \"Email\": \"zimmermanmueblesgestion@gmail.com\",\n            \"CantidadHijos\": 2,\n            \"Usuario\": \"adam\",\n            \"FechaNacimiento\": \"2020-04-11\",\n            \"FechaInicio\": \"2020-04-11\",\n            \"FechaAlta\": \"2020-04-11 19:50:01.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": "{\n    \"Error\": null,\n    \"Respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": "{\n    \"Error\": null,\n    \"Respuesta\": [\n        {\n            \"Roles\": {\n                \"IdRol\": 1,\n                \"Rol\": \"Administradores\"\n            },\n            \"Ubicaciones\": {\n                \"IdUbicacion\": 1,\n                \"Ubicacion\": \"Casa Central Tucumán\"\n            },\n            \"Usuarios\": {\n                \"IdUsuario\": 5,\n                \"IdRol\": 1,\n                \"IdUbicacion\": 1,\n                \"IdTipoDocumento\": 1,\n                \"Documento\": \"39477073\",\n                \"Nombres\": \"Nicolas\",\n                \"Apellidos\": \"Bachs\",\n                \"EstadoCivil\": \"S\",\n                \"Telefono\": \"+543814491954\",\n                \"Email\": \"nicolas.bachs@gmail.com\",\n                \"CantidadHijos\": 0,\n                \"Usuario\": \"nbachs\",\n                \"FechaUltIntento\": \"2020-05-08 00:52:23.000000\",\n                \"FechaNacimiento\": \"1995-12-27\",\n                \"FechaInicio\": \"2019-11-22\",\n                \"FechaAlta\": \"2020-05-08 00:49:18.000000\",\n                \"Estado\": \"A\"\n            }\n        }\n    ]\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": "{\n    \"Error\": null,\n    \"Respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 6,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"lchoua\",\n            \"Password\": \"LoikCapo\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-12 18:11:39.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": " {\n    \"Error\": null,\n    \"Respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 1,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Nombres\": \"Adam el super admin\",\n            \"EstadoCivil\": \"C\",\n            \"Telefono\": \"+54(381)4321719\",\n            \"Email\": \"zimmermanmueblesgestion@gmail.com\",\n            \"CantidadHijos\": 2,\n            \"Usuario\": \"adam\",\n            \"FechaNacimiento\": \"2020-04-11\",\n            \"FechaInicio\": \"2020-04-11\",\n            \"FechaAlta\": \"2020-04-11 19:50:01.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": "{\n    \"Error\": null,\n    \"Respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 8,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"lchoua\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-12 20:24:34.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": "{\n    \"Error\": null,\n    \"Respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 8,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"lchoua\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-12 20:24:34.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": "{\n    \"Error\": null,\n    \"Respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 6,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"41144069\",\n            \"Nombres\": \"Loik\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)5483777\",\n            \"Email\": \"loikchoua4@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"lchoua\",\n            \"Password\": \"LoikCapo\",\n            \"FechaNacimiento\": \"1998-05-27\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-12 18:11:39.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_DEFAULT\",\n        \"Mensaje\": \"Ha ocurrido un error mientras se procesaba su petición.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": "{\n    \"Error\": null,\n    \"Respuesta\": {\n        \"Usuarios\": {\n            \"IdUsuario\": 9,\n            \"IdRol\": 1,\n            \"IdUbicacion\": 1,\n            \"IdTipoDocumento\": 1,\n            \"Documento\": \"42664256\",\n            \"Nombres\": \"Chloe\",\n            \"Apellidos\": \"Choua\",\n            \"EstadoCivil\": \"S\",\n            \"Telefono\": \"+54(381)4451337\",\n            \"Email\": \"chloechoua@gmail.com\",\n            \"CantidadHijos\": 0,\n            \"Usuario\": \"cchoua\",\n            \"FechaNacimiento\": \"2000-05-17\",\n            \"FechaInicio\": \"2020-01-01\",\n            \"FechaAlta\": \"2020-05-13 20:25:19.000000\",\n            \"Estado\": \"A\"\n        }\n    }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_PASSWORD_INCORRECTA\",\n        \"Mensaje\": \"La contraseña ingresada no es correcta.\"\n    },\n    \"Respuesta\": null\n}",
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
          "content": "{\n    \"Error\": null\n    \"Respuesta\": null\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "examples": [
        {
          "title": "Error-Response:",
          "content": "{\n    \"Error\": {\n        \"Codigo\": \"ERROR_USUARIO_ESTA_BAJA\",\n        \"Mensaje\": \"El usuario no existe o ya está dado de baja.\"\n    },\n    \"Respuesta\": null\n}",
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
