# Series Tracker API - Backend

REST API para gestion de series de TV con PostgreSQL y Go.

Deploy en produccion: https://empowering-prosperity-production-d1ee.up.railway.app
Frontend: https://proy1-fullstack-front-end.vercel.app
Repositorio Frontend: https://github.com/Adrian-24914/Proy1_Fullstack_FrontEnd

<img width="1914" height="911" alt="image" src="https://github.com/user-attachments/assets/fd4e14ac-1470-4b81-ab9f-0b5349327713" />



---

Correr Localmente

Requisitos:
- Go 1.21+
- PostgreSQL 14+

Pasos:

1. Clonar repositorio
git clone https://github.com/Adrian-24914/Proy1_Fullstack_BackEnd.git
cd Proy1_Fullstack_BackEnd

2. Instalar dependencias
go mod download

3. Configurar variables de entorno
cp .env.example .env

Editar .env:
PORT=8080
DATABASE_URL=postgresql://postgres:password@localhost:5432/series_tracker?sslmode=disable
ALLOWED_ORIGINS=http://localhost:5500

4. Crear base de datos
createdb series_tracker

5. Correr servidor
go run cmd/server/main.go

El servidor estara en http://localhost:8080

---

Endpoints

GET /series - Listar todas las series
GET /series/:id - Obtener una serie por ID
POST /series - Crear una serie nueva
PUT /series/:id - Editar una serie existente
DELETE /series/:id - Eliminar una serie
GET /health - Health check
POST /upload - Subir imagen (max 1MB)

Filtros:
?search=texto - Buscar en titulo y descripcion
?genre=Crime - Filtrar por genero
?page=1&limit=10 - Paginacion

---

Challenges Implementados

API y Backend:
- Codigos HTTP correctos (20 pts) - 201 al crear, 204 al eliminar, 404 si no existe, 400 en validaciones
- Validacion server-side con JSON descriptivo (20 pts) - Respuestas de error estructuradas
- Paginacion (30 pts) - ?page= y ?limit= con metadata
- Busqueda por nombre (15 pts) - ?search= en titulo y descripcion

Challenges:
- Upload de imagenes (30 pts) - Endpoint /upload con limite de 1MB
- Exportar CSV (20 pts) - Generado desde JavaScript en frontend

Total: 135 puntos

No Implementados:
- OpenAPI/Swagger Spec (20 pts)
- Swagger UI (20 pts)
- Ordenamiento con ?sort= (15 pts)
- Exportar Excel (30 pts)
- Sistema de Rating (30 pts)

---

CORS

Que es CORS?
CORS (Cross-Origin Resource Sharing) es un mecanismo de seguridad del navegador que controla que dominios pueden hacer peticiones a un servidor; configuramos el servidor para permitir peticiones desde el dominio del frontend.

Configuracion:
Access-Control-Allow-Origin: https://proy1-fullstack-front-end.vercel.app
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type

---

Reflexion Tecnica

Go como Backend
Lo usaria de nuevo? Si

Pros:
- Compilacion estatica facilita el deploy (un solo binario)
- Performance excelente sin necesidad de optimizacion
- Manejo de concurrencia nativo perfecto para APIs
- Type safety ayuda a evitar errores en runtime

Contras:
- Verboso comparado con Node.js/Express
- Menos librerias maduras para ciertas tareas

PostgreSQL
Lo usaria de nuevo? Si

Pros:
- Robusto y confiable para produccion
- SERIAL para IDs auto-incrementales funciona perfecto
- Excelente integracion con Railway
- Migraciones simples con SQL puro

Contras:
- Mas pesado que SQLite para desarrollo local
- Requiere servidor corriendo (vs archivo SQLite)

Railway
Lo usaria de nuevo? Absolutamente

Pros:
- Deploy automatico desde GitHub sin configuracion
- PostgreSQL incluido gratis
- Variables de entorno auto-configuradas
- Logs en tiempo real excelentes para debugging

Contras:
- Limites del tier gratuito

---

Estructura del Proyecto

Proy1_Fullstack_BackEnd/
├── cmd/server/main.go
├── internal/
│   ├── handlers/
│   │   ├── series.go
│   │   └── upload.go
│   ├── models/series.go
│   ├── database/
│   │   ├── db.go
│   │   └── series.go
│   └── middleware/cors.go
├── uploads/
├── go.mod
└── README.md

Autor: Adrian-24914
