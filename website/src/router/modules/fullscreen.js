/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const fullScreenRouter = {
  path: '/fullscreen',
  component: Layout,
  redirect: '/fullscreen/dashboard',
  name: 'FullScreen',
  meta: {
    title: '数据展示',
    icon: 'peoples'
  },
  children: [
    {
      path: 'dashboard',
      component: () => import('@/views/fullscreen/dashboard/index'),
      name: 'dashboard',
      meta: { title: 'DashBoard' }
    },
    {
      path: 'resources',
      component: () => import('@/views/fullscreen/iot/index'),
      name: 'resources',
      meta: { title: 'IOT' }
    },
    {
      path: 'crd',
      component: () => import('@/views/fullscreen/bigdata/index'),
      name: 'crd',
      meta: { title: 'BigData' }
    }

  ]
}
export default fullScreenRouter
