import { createApp } from 'vue'
import App from './App.vue'
import installElementPlus from './plugins/element'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import router from './router'

const app = createApp(App).use(router)
installElementPlus(app)
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }

app.mount('#app')