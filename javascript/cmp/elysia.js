import { Elysia } from 'elysia';

new Elysia()
    .get('/', () => 'hello')
    .listen(8080)
