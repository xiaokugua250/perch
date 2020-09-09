/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const fullScreenRouter = {
  path: '/fullscreen',
  component: Layout,
  redirect: '/fullscreen/dashboard',
  name: 'FullScreen',
  meta: {
    title: '全屏展示',
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
      path: 'iot',
      component: () => import('@/views/fullscreen/iot/index'),
      name: 'iot',
      meta: { title: 'IOT' }
    },
    {
      path: 'bigdata',
      component: () => import('@/views/fullscreen/bigdata/index'),
      name: 'bigdata',
      meta: { title: 'BigData' }
    },


  ]
}
export default fullScreenRouter
