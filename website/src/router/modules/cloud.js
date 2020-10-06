/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const cloudRouter = {
  path: '/cloud',
  component: Layout,
  redirect: '/cloud/dashboard',
  name: 'cloud',
  meta: {
    title: '容器云',
    icon: 'peoples'
  },
  children: [
    {
      path: 'dashboard',
      component: () => import('@/views/cloud/dashboard/index'),
      name: 'dashboard',
      meta: { title: 'DashBoard' }
    },
    {
      path: 'resources',
      component: () => import('@/views/cloud/resources/index'),
      name: 'resources',
      meta: { title: 'resources' }
    },
    {
      path: 'crd',
      component: () => import('@/views/cloud/crd/index'),
      name: 'crd',
      meta: { title: 'CRD' }
    }

  ]
}
export default cloudRouter
