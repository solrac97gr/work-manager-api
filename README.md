# Work Manager API

Yendo Api es un desarrollo hecho con golang y mongoDB que te permite gestionar un sistema para solicitar servicios para el hogar. Cuenta con la tecnología de JWT para mayor seguridad.

| Nombre                | Ruta       | Parametros                  | Estado     | Protegida | Método |
| --------------------- | ---------- | --------------------------- | ---------- | --------- | ------ |
| Registrarse           | /register  | ninguno                     | Lanzada    | No        | POST   |
| Ingresar              | /login     | ninguno                     | Lanzada    | No        | POST   |
| Obtener Perfil        | /profile   | id                          | Lanzada    | Si        | GET    |
| Editar perfil         | /profile   | ninguno                     | Lanzada    | Si        | PUT    |
| Crear Trabajo         | /work      | ninguno                     | Lanzada    | Si        | POST   |
| Actualizar Trabajo    | /work      | ninguno                     | Lanzada    | Si        | PUT    |
| Eliminar Trabajo      | /work      | id                          | No lanzada | Si        | DELETE |
| Listar Trabajos       | /work      | category,id,idate,fdate     | No lanzada | Si        | GET    |
| Crear pago            | /pay       | ninguno                     | No lanzada | Si        | POST   |
| Editar pago           | /pay       | id                          | No lanzada | Si        | PUT    |
| Eliminar pago         | /pay       | id                          | No lanzada | Si        | DELETE |
| Listar pagos          | /pay       | idate,fdate,category,amount | No lanzada | Si        | GET    |
| Crear categoria       | /category  | ninguno                     | No lanzada | Si        | POST   |
| Editar categoria      | /category  | id                          | No lanzada | Si        | PUT    |
| Eliminar categoria    | /category  | id                          | No lanzada | Si        | DELETE |
| Listar categoria      | /category  | ninguno                     | No lanzada | Si        | GET    |
| Listar formas de pago | /paymethod | ninguno                     | No lanzada | Si        | GET    |

# Deploy

El api se encuentra alojado en Heroku para realizar un deploy debes tener el Heroku CLI instalado.

`git push heroku master`

Antes de eso asegurate de ejecutar todos los tests.

`go test`

Para ejecutarlo de forma local.(En el directorio donde esta el proyecto)

`go run main.go`

Para obtener los archivos compilados.(En el directorio donde esta el proyecto)

`go build main.go`

# Dependencias

¿Qué hago si agregue una nueva dependencia?

Dirigete al archivo `go.mod` y agrega el nombre del paquete junto con la versión usada.

# Estilos y buenas practicas.

Para este proyecto usamos la guía de recomendaciones y estilos de uber. la podrás encontrar <a href="https://github.com/friendsofgo/uber-go-guide-es">aquí</a>
