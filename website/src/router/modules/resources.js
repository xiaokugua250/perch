/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const resourcesRouter = {
  path: '/resources',
  component: Layout,
  redirect: 'noRedirect',
  name: 'Table',
  meta: {
    title: 'resources',
    icon: 'peoples'
  },
  children: [
    {
      path: '/',
      component: () => import('@/views/resources/resources'),
      name: 'resources',
      meta: { title: 'resources' }
    },
    {
      path: 'blogs',
      component: () => import('@/views/resources/blogs/list'),
      name: 'resources'
      // meta: { title: 'resources' }
    },
    {
      path: 'show',
      component: () => import('@/views/resources/blogs/show'),
      name: 'show',
      meta: { title: 'resources' }
    },
  ]
}
export default resourcesRouter
