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


    /*
    {
      path: 'index',
      component: () => import('@/views/users/rbac/index'),
      name: 'rbac',
      meta: { title: '用户权限' }
    },
    {
      path: 'dynamic-table',
      component: () => import('@/views/table/dynamic-table/index'),
      name: 'DynamicTable',
      meta: { title: 'Dynamic Table' }
    },
    {
      path: 'drag-table',
      component: () => import('@/views/table/drag-table'),
      name: 'DragTable',
      meta: { title: 'Drag Table' }
    },
    {
      path: 'inline-edit-table',
      component: () => import('@/views/table/inline-edit-table'),
      name: 'InlineEditTable',
      meta: { title: 'Inline Edit' }
    }
*/
  ]
}
export default resourcesRouter
