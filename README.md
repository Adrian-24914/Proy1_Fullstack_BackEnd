# Series Tracker API

REST API backend para gestión de series con PostgreSQL.

Deploy: https://empowering-prosperity-production-d1ee.up.railway.app
Frontend: https://proy1-fullstack-front-end.vercel.app
Repositorio Frontend: https://github.com/Adrian-24914/Proy1_Fullstack_FrontEnd

Endpoints:
- GET /series — Listar todas las series
- GET /series/:id — Obtener una serie por ID
- POST /series — Crear una serie nueva
- PUT /series/:id — Editar una serie existente
- DELETE /series/:id — Eliminar una serie

CORS

CORS (Cross-Origin Resource Sharing) es un mecanismo que permite que el navegador acepte peticiones desde un dominio distinto al del servidor. Configuramos el servidor para permitir peticiones desde el frontend en Vercel mediante estos headers:

Access-Control-Allow-Origin: https://proy1-fullstack-front-end.vercel.app
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type