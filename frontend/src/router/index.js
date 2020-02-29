import Vue from 'vue'
import Router from 'vue-router'

// Containers
const DefaultContainer = () => import('@/containers/DefaultContainer')

// const Switchlist = () => import('@/views/switching/switchlist')
// const DetectedDevices = () => import('@/views/devices/DetectedDevices')
const targetlist = () => import('@/views/targets/targetlist')
const snippetlist = () => import('@/views/snippets/snippetlist')

// const floweditor = () => import('@/views/flows/floweditor')

// Views - Pages
const Page404 = () => import('@/views/pages/Page404')
const Page500 = () => import('@/views/pages/Page500')
const Login = () => import('@/views/pages/Login')
const Register = () => import('@/views/pages/Register')


Vue.use(Router)

function configRoutes() {
  return [
    {
      path: '/',
      redirect: '/admin/targetlist',
      name: 'Home',
      component: DefaultContainer,
      children: [
        {
          path: 'admin',
          redirect: '/admin/targetlist',
          name: 'admin',
          component: {
            render (c) { return c('router-view') }
          },
          children: [
            {
              path: 'targetlist',
              name: 'targetlist',
              component: targetlist
            },
          ]
        },
        {
          path: '/snippets',
          name: 'snippetlist',
          component: snippetlist,
         
        }
      ]
    },
    
    {
      path: '/pages',
      redirect: '/pages/404',
      name: 'Pages',
      component: {
        render (c) { return c('router-view') }
      },
      children: [
        {
          path: '404',
          name: 'Page404',
          component: Page404
        },
        {
          path: '500',
          name: 'Page500',
          component: Page500
        },
        {
          path: 'login',
          name: 'Login',
          component: Login
        },
        {
          path: 'register',
          name: 'Register',
          component: Register
        }
      ]
    }
  ]
}

export default new Router({
  mode: 'hash', // https://router.vuejs.org/api/#mode
  linkActiveClass: 'open active',
  scrollBehavior: () => ({ y: 0 }),
  routes: configRoutes()
})
