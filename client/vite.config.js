// // import { defineConfig } from 'vite'
// // import reactRefresh from '@vitejs/plugin-react-refresh'
// //
// // // https://vitejs.dev/config/
// // export default defineConfig({
// //   const env = loadEnv(mode, process.cwd(), '')
// //   __HMR_PROTOCOL__: JSON.stringify('wss'),
// //   __HMR_HOSTNAME__: JSON.stringify('custom-hostname'),
// //   __HMR_PORT__: JSON.stringify('custom-port'),
// //   __BASE__: JSON.stringify('custom-base-path')
// // })
//
// import { defineConfig, loadEnv } from 'vite'
//
// export default defineConfig(({ command, mode }) => {
//   // Load env file based on `mode` in the current working directory.
//   // Set the third parameter to '' to load all env regardless of the `VITE_` prefix.
//   // const env = loadEnv(mode, process.cwd(), '')
//   return {
//     // vite config
//     define: {
//       __HMR_PORT__: 3000,
//     },
//   }
// })

// import { defineConfig } from 'vite';
//
// export default ({ mode }) => {
//   return defineConfig({
//     define: {
//       __HMR_PROTOCOL__:JSON.stringify('wss'),
//       __HMR_HOSTNAME__:JSON.stringify('dgsdgsdg.ru'),
//       __HMR_PORT__: JSON.stringify(819),
//     },
//     // server: {
//     //   port: 3000,
//     //   https: true,
//     //   host: "dgsdgsdg.ru",
//     // },
//   });
// }

// import { defineConfig } from 'vite';
//
// export default ({ mode }) => {
//   return defineConfig({
//     // define: {
//     //   __HMR_PROTOCOL__: JSON.stringify('wss'),
//     //   __HMR_HOSTNAME__: JSON.stringify('dgsdgsdg.ru'),
//     //   __HMR_PORT__: JSON.stringify(819),
//     // },
//     server: {
//       hmr:false,
//       // port: 4000,
//       // host: 'fsfsffs.ru',
//       // https: true, // Если вам необходим HTTPS, раскомментируйте эту строку
//       // host: "dgsdgsdg.ru", // Если хотите указать другой хост, раскомментируйте и укажите нужный
//     },
//   });

// }
//

import { defineConfig } from 'vite';

export default ({ mode }) => {
  return defineConfig({
    server: {
      host: 'localhost',
      port: 3000,
    },
  });
}
