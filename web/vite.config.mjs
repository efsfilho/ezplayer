// Plugins
import Components from 'unplugin-vue-components/vite'
import Vue from '@vitejs/plugin-vue'
import Vuetify, { transformAssetUrls } from 'vite-plugin-vuetify'
import ViteFonts from 'unplugin-fonts/vite'

// Utilities
// import { defineConfig } from 'vite'
import { defineConfig, loadEnv } from 'vite'
import { fileURLToPath, URL } from 'node:url'

export default defineConfig(({ command, mode }) => {
  // Load env file based on `mode` in the current working directory.
  // Set the third parameter to '' to load all env regardless of the `VITE_` prefix.
  const env = loadEnv(mode, process.cwd(), '')
  return {
    // vite config
    define: {
      __API_PORT__: JSON.stringify(env.API_PORT),
      __API_HOST__: JSON.stringify(env.API_HOST),
    },
    plugins: [
      Vue({
        template: { transformAssetUrls }
      }),
      // https://github.com/vuetifyjs/vuetify-loader/tree/master/packages/vite-plugin#readme
      Vuetify(),
      Components(),
      ViteFonts({
        google: {
          families: [{
            name: 'Roboto',
            styles: 'wght@100;300;400;500;700;900',
          }],
        },
      }),
    ],
    // define: { 'process.env': {} },
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      },
      extensions: [
        '.js',
        '.json',
        '.jsx',
        '.mjs',
        '.ts',
        '.tsx',
        '.vue',
      ],
    },
    server: {
      port: 3000,
    },
  }
})

// // https://vitejs.dev/config/
// export default defineConfig({
//   plugins: [
//     Vue({
//       template: { transformAssetUrls }
//     }),
//     // https://github.com/vuetifyjs/vuetify-loader/tree/master/packages/vite-plugin#readme
//     Vuetify(),
//     Components(),
//     ViteFonts({
//       google: {
//         families: [{
//           name: 'Roboto',
//           styles: 'wght@100;300;400;500;700;900',
//         }],
//       },
//     }),
//   ],
//   define: { 'process.env': {} },
//   resolve: {
//     alias: {
//       '@': fileURLToPath(new URL('./src', import.meta.url))
//     },
//     extensions: [
//       '.js',
//       '.json',
//       '.jsx',
//       '.mjs',
//       '.ts',
//       '.tsx',
//       '.vue',
//     ],
//   },
//   server: {
//     port: 3000,
//   },
// })
