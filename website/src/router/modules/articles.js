/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const articlesRouter = {
  path: '/articles',
  component: Layout,
  redirect: 'noRedirect',
  name: 'Table',
  meta: {
    title: 'resources',
    icon: 'peoples'
  },
  children: [
    {
      path: 'show',
      component: () => import('@/views/resources/blogs/show'),
      name: 'resources',
      meta: { title: 'resources' }
    },
    {
      path: 'demo',
      component: () => import('@/views/resources/blogs/demo'),
      name: 'resources',
      meta: { title: 'resources' }
    },
    {
      path: 'list',
      component: () => import('@/views/resources/blogs/list'),
      name: 'resources',
      meta: { title: 'resources' }
    },
    {
      path: '/',
      component: () => import('@/views/resources/blogs/index'),
      name: 'resources'
      // meta: { title: 'resources' }
    }

  ]
}
export default articlesRouter
