
## Installation and build


```bash
npm install
npm run build
```
    
For dev mode
```bash
 npm run dev
```

Vite default port is 3000, but if you want to use a different port or address you'll need to set API_PORT and API_HOST variables to your environment. It can be exported or an ```.env``` file can be created at /web.

\* To take effect, both of the variables must be set.
```sh
# .env file example to http://10.0.2.10:5000
API_PORT=5000
API_HOST=10.0.2.10
```